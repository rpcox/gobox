## sync.Pool: base

A look into sync.Pool using byte slices. Uses bucketed []byte pools and checks to make sure underlying arrays are not reallocated.
    
    >  ./base 
    text: 1234567
     ** get: len=7 cap=8
     &buf: 0x14000092028
     buf: [49 50 51 52 53 54 55]
     buf: 1234567
    text: 1234567890123456
     ** get: len=16 cap=32
     &buf: 0x140000b6020
     buf: [49 50 51 52 53 54 55 56 57 48 49 50 51 52 53 54]
     buf: 1234567890123456
    text: ABCDEFG
     ** get: len=7 cap=8
     &buf: 0x14000092028
     buf: [65 66 67 68 69 70 71]
     buf: ABCDEFG
    text: 123456789012345
     ** get: len=15 cap=16
     &buf: 0x14000092060
     buf: [49 50 51 52 53 54 55 56 57 48 49 50 51 52 53]
     buf: 123456789012345
    text: abcdefg
     ** get: len=7 cap=8
     &buf: 0x14000092028
     buf: [97 98 99 100 101 102 103]
     buf: abcdefg
    text: 1234567
     ** get: len=7 cap=8
     &buf: 0x14000092028
     buf: [49 50 51 52 53 54 55]
     buf: 1234567
    
      Pool Count: 3
    Pools in Use: [8 16 32]
      Line Count: 6
    
