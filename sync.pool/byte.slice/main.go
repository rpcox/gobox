// see https://github.com/rpcox/gobox/tree/main/sync.pool/base
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"sync"
	"time"
)

// A slice for read-only purposes
var bufferSizes = []int{8, 16, 32, 64, 128, 256, 512, 1024, 2048}

type SyncPoolManager struct {
	pmu   sync.Mutex
	pools map[int]*sync.Pool // Key is the buffer capacity
}

type MetricsManager struct {
	mmu      sync.Mutex
	GetTotal int64
	PutTotal int64
	NewTotal map[int]int
}

// PoolManager manages multiple sync.Pools for different []byte sizes.
type PoolManager struct {
	spm SyncPoolManager
	mm  MetricsManager
}

// Create a new PoolManager
func NewPoolManager() *PoolManager {
	return &PoolManager{
		SyncPoolManager{pools: make(map[int]*sync.Pool)},
		MetricsManager{GetTotal: 0, PutTotal: 0, NewTotal: make(map[int]int)},
	}
}

// Get retrieves a *[]byte from the pool with at least the requested capacity.
func (pm *PoolManager) Get(length int) *[]byte {
	// Determine which pool to draw from based on size. Seeking smallest bucket.
	capacity := 0
	for _, size := range bufferSizes {
		if length < size {
			capacity = size
			break
		}
	}

	// Does the pool for this size []byte already exist
	pool, exists := pm.spm.pools[capacity]
	// If the []byte pool of this size does not yet exist, create it
	if !exists {
		bs := make([]byte, capacity)

		pool = &sync.Pool{
			New: func() interface{} {
				return &bs
			},
		}
		pm.spm.pmu.Lock()
		pm.spm.pools[capacity] = pool
		pm.spm.pmu.Unlock()
	}

	buf := pool.Get().(*[]byte)

	// Always reset/zero the buffer before use
	reset := make([]byte, capacity)
	*buf = reset

	pm.mm.mmu.Lock()
	pm.mm.GetTotal++
	if !exists {
		pm.mm.NewTotal[capacity]++
	}
	pm.mm.mmu.Unlock()

	return buf
}

// Put returns a *[]byte to the appropriate pool.
func (pm *PoolManager) Put(buf *[]byte) {
	pm.mm.mmu.Lock()
	defer pm.mm.mmu.Unlock()
	pm.spm.pmu.Lock()
	defer pm.spm.pmu.Unlock()

	capacity := cap(*buf)
	if pool, ok := pm.spm.pools[capacity]; ok {
		pool.Put(buf)
		pm.mm.PutTotal++
	} else {
		// We should always find the correct pool to return
		fmt.Fprintf(os.Stderr, "error: unable to return buffer to pool\n")
		buf = nil
	}
}

func (pm *PoolManager) PoolCount() int {
	return len(pm.spm.pools)
}

func (pm *PoolManager) PoolsInUse() []int {
	pm.mm.mmu.Lock()
	defer pm.mm.mmu.Unlock()
	var a []int
	for k, _ := range pm.spm.pools {
		a = append(a, k)
	}

	sort.Ints(a)
	return a
}

func (pm *PoolManager) GetTotal() int64 {
	pm.mm.mmu.Lock()
	defer pm.mm.mmu.Unlock()
	return pm.mm.GetTotal
}

func (pm *PoolManager) PutTotal() int64 {
	pm.mm.mmu.Lock()
	defer pm.mm.mmu.Unlock()
	return pm.mm.PutTotal
}

func (pm *PoolManager) PoolMap() {
	keys := pm.PoolsInUse()
	for _, v := range keys {
		fmt.Fprintf(os.Stdout, "%4d byte bin slice count: %4d\n", v, pm.mm.NewTotal[v])
	}
}

func (pm *PoolManager) Worker(data <-chan *[]byte, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintf(os.Stderr, "worker[%d] starting\n", id)

	for bp := range data {
		lenB := len(*bp)
		buf := pm.Get(lenB)
		n := copy(*buf, *bp)
		if n < lenB {
			fmt.Fprintf(os.Stderr, "error: missing %d bytes on copy\n", lenB-n)
		}
		pm.Put(buf)

		//
		// Do work on the buf data
		//

	}
	fmt.Fprintf(os.Stderr, "worker[%d] exiting\n", id)
}

func main() {
	_workers := flag.Int("workers", 1, "Specify the number of workers")
	_file := flag.String("d", "", "Identify the data file")
	flag.Parse()

	//fmt.Println("pid:", os.Getpid())  // in case I do something dumb and need a pid real quick
	pm := NewPoolManager()

	fh, err := os.Open(*_file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	var wg sync.WaitGroup
	data := make(chan *[]byte, 100)

	for i := 1; i <= *_workers; i++ {
		wg.Add(1)
		go pm.Worker(data, i, &wg)
	}

	lineCount := 0
	scanner := bufio.NewScanner(fh)
	start := time.Now()
	for scanner.Scan() { // Note: default size of scanner buffer is 4096 bytes
		B := scanner.Bytes()
		data <- &B
		lineCount++
	}

	if scanner.Err() != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", scanner.Err())
	} else {
		fmt.Fprintf(os.Stdout, "no scanner errors\n")
	}

	close(data)
	wg.Wait()
	fmt.Fprintln(os.Stdout, "\n  Line Count:", lineCount)
	fmt.Fprintf(os.Stdout, "  Pool Count: %d\n", pm.PoolCount())
	fmt.Fprintf(os.Stdout, "Pools in Use: %d\n", pm.PoolsInUse())
	fmt.Fprintf(os.Stdout, "   Get Total: %d\n", pm.GetTotal())
	fmt.Fprintf(os.Stdout, "   Put Total: %d\n\n", pm.PutTotal())
	fmt.Fprintln(os.Stdout, " Pool Map:")
	pm.PoolMap()
	fmt.Fprintf(os.Stderr, "\nelapsed: %v\n", time.Since(start))
}
