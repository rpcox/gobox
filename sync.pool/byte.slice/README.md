## byte.slice
    
The data sample (data/16xkjv) is 16 concatinated copies of the King James Bible. Additional  workers wasn't faster. This implementation uses bufio.Scanner. Using bufio.Reader, it's ~ 15% slower.
    

    >  ./byte.slice -d data/16xkjv
    worker[1] starting
    no scanner errors
    worker[1] exiting
    
      Line Count: 497632
      Pool Count: 6
    Pools in Use: [32 64 128 256 512 1024]
       Get Total: 497632
       Put Total: 497632
    
     Pool Map:
      32 byte bin slice count:    1
      64 byte bin slice count:    1
     128 byte bin slice count:    1
     256 byte bin slice count:    1
     512 byte bin slice count:    1
    1024 byte bin slice count:    1
    
    elapsed: 97.144667ms
    
    
    >  ./byte.slice -d data/16xkjv -workers 2
    worker[1] starting
    worker[2] starting
    no scanner errors
    worker[2] exiting
    worker[1] exiting
    
      Line Count: 497632
      Pool Count: 6
    Pools in Use: [32 64 128 256 512 1024]
       Get Total: 497632
       Put Total: 497632
    
     Pool Map:
      32 byte bin slice count:    1
      64 byte bin slice count:    1
     128 byte bin slice count:    1
     256 byte bin slice count:    1
     512 byte bin slice count:    1
    1024 byte bin slice count:    1
    
    elapsed: 118.637292ms
    
    
    >  ./byte.slice -d data/16xkjv -workers 3
    worker[1] starting
    worker[3] starting
    worker[2] starting
    no scanner errors
    worker[3] exiting
    worker[1] exiting
    worker[2] exiting
    
      Line Count: 497632
      Pool Count: 6
    Pools in Use: [32 64 128 256 512 1024]
       Get Total: 497632
       Put Total: 497632
    
     Pool Map:
      32 byte bin slice count:    1
      64 byte bin slice count:    1
     128 byte bin slice count:    1
     256 byte bin slice count:    2
     512 byte bin slice count:    1
    1024 byte bin slice count:    1
    
    elapsed: 143.144458ms
    
