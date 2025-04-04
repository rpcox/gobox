## parse

I got caught off guard with time.Parse defaulting to UTC +0

    > parse 
    Parsing time without time zone: 2025-04-03 19:15:00  =>  2025-04-03 19:15:00 +0000 UTC
    Location of parsed time defaults to: UTC
    Parsing time using time zone: 2006-01-02 15:04:05 MST  =>  2025-04-03 19:15:00 +0000 UTC
    Location converts to: UTC
    Parsing time using time zone offset: 2006-01-02 15:04:05 -0700 MST  =>  2025-04-03 19:15:00 -0500 EST
    Location is now: EST
    Parsing time using just the zone offset: 2006-01-02 15:04:05 -0700  =>  2025-04-03 19:15:00 -0500 -0500
    Location is now: EST