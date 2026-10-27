## chain-2
    
Nested http.Handlers with panic recovery
    
#### client
    
    > curl -v  localhost:8080  -H "Content-Type: application/json" -d '{"name":"johndoe","panic":"false"}'
    
#### server
    
    >  ./chain-2
    server starting on :8080
    + enter Recovery
    + enter Logging: [::1]:54250 - - 26 Oct 26 18:38 -0700 "POST / HTTP/1.1 curl/8.7.1"
     + enter VerifyJson
      + enter Auth
       + enter HeaderSet
        + enter mainHandler
        - exit mainHandler
       - exit HeaderSet
      - exit Auth
     - exit VerifyJson
    - exit Logging
    - exit Recovery: 4.409794292s
    
#### client
    
    > curl -v  localhost:8080  -H "Content-Type: application/json" -d '{"name":"johndoe","panic":"true"}'
    > curl -v  localhost:8080  -H "Content-Type: application/json" -d '{"name":"johndoe","panic":"false"}'
    
#### server
    
    server receives two requests. first causes panic, server recovers and receives second request.
    
    
    + enter Recovery
    + enter Logging: [::1]:54251 - - 26 Oct 26 18:39 -0700 "POST / HTTP/1.1 curl/8.7.1"
     + enter VerifyJson
      + enter Auth
       + enter HeaderSet
        + enter mainHandler
    throwing panic:
    stack:
    goroutine 17 [running]:
    runtime/debug.Stack()
    	/usr/local/go/src/runtime/debug/stack.go:26 +0x64
    main.Recovery.func1.1()
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/chain-2/main.go:38 +0x3c
    panic({0x100c67d00?, 0x100c10150?})
    	/usr/local/go/src/runtime/panic.go:859 +0x120
    main.mainHandler({0x100ca9728, 0x7dee6881e000}, 0x7dee688143c0)
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/chain-2/main.go:131 +0x224
    net/http.HandlerFunc.ServeHTTP(0x100c95080?, {0x100ca9728?, 0x7dee6881e000?}, 0x8?)
    	/usr/local/go/src/net/http/server.go:2358 +0x38
    main.HeaderSet.func1({0x100ca9728, 0x7dee6881e000}, 0x7dee688143c0)
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/chain-2/main.go:117 +0x118
    net/http.HandlerFunc.ServeHTTP(0x7dee688041b0?, {0x100ca9728?, 0x7dee6881e000?}, 0xf?)
    	/usr/local/go/src/net/http/server.go:2358 +0x38
    main.Auth.func1({0x100ca9728, 0x7dee6881e000}, 0x7dee688143c0)
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/chain-2/main.go:108 +0x168
    net/http.HandlerFunc.ServeHTTP(0x100ca9af0?, {0x100ca9728?, 0x7dee6881e000?}, 0x100c10110?)
    	/usr/local/go/src/net/http/server.go:2358 +0x38
    main.VerifyJson.func1({0x100ca9728, 0x7dee6881e000}, 0x7dee68814000)
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/chain-2/main.go:84 +0x2f0
    net/http.HandlerFunc.ServeHTTP(0x100ca74d8?, {0x100ca9728?, 0x7dee6881e000?}, 0x29?)
    	/usr/local/go/src/net/http/server.go:2358 +0x38
    main.Logging.func1({0x100ca9728, 0x7dee6881e000}, 0x7dee68814000)
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/chain-2/main.go:53 +0x16c
    net/http.HandlerFunc.ServeHTTP(0x100ca74d8?, {0x100ca9728?, 0x7dee6881e000?}, 0x11?)
    	/usr/local/go/src/net/http/server.go:2358 +0x38
    main.Recovery.func1({0x100ca9728, 0x7dee6881e000}, 0x7dee68814000)
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/chain-2/main.go:43 +0xbc
    net/http.HandlerFunc.ServeHTTP(0x7dee6879e0c0?, {0x100ca9728?, 0x7dee6881e000?}, 0x89d3728?)
    	/usr/local/go/src/net/http/server.go:2358 +0x38
    net/http.(*ServeMux).ServeHTTP(0x1000000000000?, {0x100ca9728, 0x7dee6881e000}, 0x7dee68814000)
    	/usr/local/go/src/net/http/server.go:2923 +0x194
    net/http.serverHandler.ServeHTTP({0x10?}, {0x100ca9728?, 0x7dee6881e000?}, 0x1?)
    	/usr/local/go/src/net/http/server.go:3436 +0xb0
    net/http.(*conn).serve(0x7dee68808000, {0x100ca9af0, 0x7dee68791080})
    	/usr/local/go/src/net/http/server.go:2157 +0x5a0
    created by net/http.(*Server).Serve in goroutine 1
    	/usr/local/go/src/net/http/server.go:3604 +0x3e8
    
    2026/10/26 18:39:05 http: superfluous response.WriteHeader call from main.Recovery.func1.1 (main.go:40)
    + enter Recovery       <<-- second curl starts here
    + enter Logging: [::1]:54252 - - 26 Oct 26 18:39 -0700 "POST / HTTP/1.1 curl/8.7.1"
     + enter VerifyJson
      + enter Auth
       + enter HeaderSet
        + enter mainHandler
        - exit mainHandler
       - exit HeaderSet
      - exit Auth
     - exit VerifyJson
    - exit Logging
    - exit Recovery: 18.682678334s
    ^C
