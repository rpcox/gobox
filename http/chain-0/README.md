## chain-0
    
Simple http.Handler chaining
    
### server side
    
    
    >  chain-0
    server starting on :8080
    + enter Logging: [::1]:52542 - - 24 Oct 26 16:26 -0700 "GET / HTTP/1.1 curl/8.7.1"
     + enter Auth
      + enter HeaderSet
       + enter mainHandler
       - exit mainHandler
      - exit HeaderSet
     - exit Auth
    - exit Logging: 341.584µs
    ^C
    
    
### client side
    
    >  curl -v localhost:8080
    * Host localhost:8080 was resolved.
    * IPv6: ::1
    * IPv4: 127.0.0.1
    *   Trying [::1]:8080...
    * Connected to localhost (::1) port 8080
    > GET / HTTP/1.1
    > Host: localhost:8080
    > User-Agent: curl/8.7.1
    > Accept: */*
    >
    * Request completely sent off
    < HTTP/1.1 200 OK
    < X-Powered-By: Go                      <<-- HeaderSet
    < Date: Sat, 24 Oct 2026 23:26:48 GMT
    < Content-Length: 11
    < Content-Type: text/plain; charset=utf-8
    <
    * Connection #0 to host localhost left intact
    hello chain%
    >
