## chain-1
    
### valid user
    
#### client
    
    >  curl -v -X POST localhost:8080  -H "Content-Type: application/json" -d '{"name": "johndoe"}'
    * Host localhost:8080 was resolved.
    * IPv6: ::1
    * IPv4: 127.0.0.1
    *   Trying [::1]:8080...
    * Connected to localhost (::1) port 8080
    > POST / HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/8.7.1
    > Accept: */*
    > Content-Type: application/json
    > Content-Length: 19
    >
    * upload completely sent off: 19 bytes
    < HTTP/1.1 200 OK
    < X-Server: Go
    < Date: Sun, 25 Oct 2026 12:46:15 GMT
    < Content-Length: 11
    < Content-Type: text/plain; charset=utf-8
    <
    * Connection #0 to host localhost left intact
    hello chain
    >
    
#### server
    
    >  ./chain-1
    server starting on :8080
    + enter Logging: [::1]:52977 - - 25 Oct 26 05:46 -0700 "POST / HTTP/1.1 curl/8.7.1"
     + enter VerifyJson
      + enter Auth
       + enter HeaderSet
        + enter mainHandler
        - exit mainHandler
       - exit HeaderSet
      - exit Auth
     + exit VerifyJson
    - exit Logging: 249.333µs
    ^C
    
    
### malformed json
    
#### client
    
    >  curl -v -X POST localhost:8080  -H "Content-Type: application/json" -d '{"name": "johndoe"'
    * Host localhost:8080 was resolved.
    * IPv6: ::1
    * IPv4: 127.0.0.1
    *   Trying [::1]:8080...
    * Connected to localhost (::1) port 8080
    > POST / HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/8.7.1
    > Accept: */*
    > Content-Type: application/json
    > Content-Length: 18
    >
    * upload completely sent off: 18 bytes
    < HTTP/1.1 400 Bad Request
    < Content-Type: text/plain; charset=utf-8
    < X-Content-Type-Options: nosniff
    < Date: Sun, 25 Oct 2026 12:47:56 GMT
    < Content-Length: 22
    <
    error: malformed json
    * Connection #0 to host localhost left intact
    
#### server
    
    >  ./chain-1
    server starting on :8080
    + enter Logging: [::1]:52982 - - 25 Oct 26 05:47 -0700 "POST / HTTP/1.1 curl/8.7.1"
     + enter VerifyJson
     + error: malformed json
    - exit Logging: 362.75µs
    ^C
    
    
### invalid user
    
#### client
    
    >  curl -v -X POST localhost:8080  -H "Content-Type: application/json" -d '{"name": "janedoe"}'
    Note: Unnecessary use of -X or --request, POST is already inferred.
    * Host localhost:8080 was resolved.
    * IPv6: ::1
    * IPv4: 127.0.0.1
    *   Trying [::1]:8080...
    * Connected to localhost (::1) port 8080
    > POST / HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/8.7.1
    > Accept: */*
    > Content-Type: application/json
    > Content-Length: 19
    >
    * upload completely sent off: 19 bytes
    < HTTP/1.1 401 Unauthorized
    < Content-Type: text/plain; charset=utf-8
    < X-Content-Type-Options: nosniff
    < Date: Sun, 25 Oct 2026 12:52:29 GMT
    < Content-Length: 18
    <
    unauthorized user
    * Connection #0 to host localhost left intact
    
#### server
    
    >  ./chain-1
    server starting on :8080
    + enter Logging: [::1]:52990 - - 25 Oct 26 05:52 -0700 "POST / HTTP/1.1 curl/8.7.1"
     + enter VerifyJson
      + enter Auth
      + unauthorized user
     + exit VerifyJson
    - exit Logging: 540.583µs
    ^C
