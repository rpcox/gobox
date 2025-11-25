## timer

I needed a timer to signal a process when it was time to go to work

    > timer -at "2025-04-04 02:42:00" -interval 1m
    first alert time : 2025-04-04 02:42:00 +0000 UTC
        current time : 2025-04-04 02:41:20
    
    * new timer
    pop: time to inform the worker
    alarm time: 2025-04-03T19:42:00.007129-07:00
    pass the word to the worker
    * new timer
    work, work, work
    pop: time to inform the worker
    alarm time: 2025-04-03T19:43:00.007901-07:00
    pass the word to the worker
    * new timer
    work, work, work
    pop: time to inform the worker
    alarm time: 2025-04-03T19:44:00.000758-07:00
    pass the word to the worker
    * new timer
    work, work, work
    ^C
    signal: shutting down
    exiting main
    leaving work
