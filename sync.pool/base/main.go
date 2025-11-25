// A look into sync.Pool using byte slices. Uses predetermined []byte pools and checks
// to make sure underlying arrays are not reallocated.
package main

import (
	"fmt"
	"os"
	"sort"
	"sync"
)

// PoolManager manages multiple sync.Pools for different []byte sizes.
type PoolManager struct {
	pools map[int]*sync.Pool // Key is the buffer capacity
}

func NewPoolManager() *PoolManager {
	return &PoolManager{
		pools: make(map[int]*sync.Pool),
	}
}

var (
	bufferSize = []int{8, 16, 32, 64, 128}
	debug      = false
)

// Get retrieves a *[]byte from the pool with at least the requested capacity.
func (pm *PoolManager) Get(size int) *[]byte {
	// Determine which pool to draw from based on size. Seeking smallest bucket.
	bucket := 0
	for i, v := range bufferSize {
		if size < v {
			bucket = i
			break
		}
	}

	// Does the pool for this size []byte already exist
	pool, exists := pm.pools[bufferSize[bucket]]
	if debug {
		fmt.Fprintf(os.Stderr, " pool exits: %v\n", exists)
	}

	// If the []byte pool of this size does not yet exist, create it
	if !exists {
		bs := make([]byte, size, bufferSize[bucket])
		if debug {
			fmt.Fprintf(os.Stderr, " ** new: %d bytes, len=%d cap=%d\n", bufferSize[bucket], len(bs), cap(bs))
			fmt.Fprintf(os.Stderr, " ** bs: %p\n", bs)
		}

		pool = &sync.Pool{
			New: func() interface{} {
				return &bs
			},
		}
		pm.pools[bufferSize[bucket]] = pool
	}

	buf := pool.Get().(*[]byte)
	if debug {
		fmt.Fprintf(os.Stderr, " &buf 1: %p\n", *buf)
	}

	// Always reset/zero the buffer before use
	// Operation creates a new slice header that points to the same underlying
	// array as the original myByteSlice, but with a length of 'size'. The
	// capacity of the slice remains unchanged.
	*buf = (*buf)[:size]
	if debug {
		fmt.Fprintf(os.Stderr, " &buf 2: %p\n", *buf)
		fmt.Fprintf(os.Stderr, " buf: %v\n", *buf)
	}

	return buf
}

// Put returns a *[]byte to the appropriate pool.
func (pm *PoolManager) Put(buf *[]byte) {
	capacity := cap(*buf)
	if pool, ok := pm.pools[capacity]; ok {
		pool.Put(buf)
	} else {
		// We should always find the correct pool to return
		fmt.Fprintf(os.Stderr, "error: unable to return buffer to pool\n")
		buf = nil
	}
}

func (pm *PoolManager) PoolCount() int {
	return len(pm.pools)
}

func (pm *PoolManager) PoolsInUse() []int {
	var a []int
	for k, _ := range pm.pools {
		a = append(a, k)
	}

	sort.Ints(a)
	return a
}

func main() {
	pm := NewPoolManager()
	text := []string{"1234567",
		"1234567890123456",
		"ABCDEFG",
		"123456789012345",
		"abcdefg",
		"1234567",
	}

	lineCount := 0
	for _, v := range text {
		fmt.Println("text:", v)
		buf := pm.Get(len(v))
		fmt.Fprintf(os.Stderr, " ** get: len=%d cap=%d\n", len(*buf), cap(*buf))
		// Behavior: (from Go docs)
		// The copy() function copies elements up to the length of the
		// shorter slice between dst and src.
		// * If dst is shorter than src, only len(dst) elements are copied.
		// * If src is shorter than dst, only len(src) elements are copied.
		// * The key point is that dst must be initialized with a sufficient
		//  length before calling copy(). If dst has zero length or is nil,
		//  no elements will be copied. Hence, the 'reset' in Get above
		//      *buf = (*buf)[:size]

		n := copy(*buf, []byte(v)) // fast, but not concurrent.
		fmt.Printf(" &buf: %p\n", *buf)
		if debug {
			fmt.Printf(" copy: %d bytes\n", n)
		}

		fmt.Println(" buf:", *buf)
		fmt.Println(" buf:", string(*buf))
		pm.Put(buf)
		lineCount++
	}

	fmt.Printf("\n  Pool Count: %d\n", pm.PoolCount())
	fmt.Printf("Pools in Use: %d\n", pm.PoolsInUse())
	fmt.Println("  Line Count:", lineCount)
}
