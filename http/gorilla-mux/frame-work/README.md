## frame-work

    A basic webserver using gorilla/mux. The setup allows pkg creation and use without having to publish packages initially. _go.work belongs at the root of the project as go.work.


### server

    > app
    2026/11/15 01:55:48 listening @ :8080

### client

    > curl localhost:8080/
    {"message":"hello world basic api"}

### server

    2026/11/15 01:56:08 GET / 2026-11-15 01:56:08.039515 -0700 MST m=+19.565615210
    2026/11/15 01:56:08 auth user
    2026/11/15 01:56:08 HomeHandler

### client

    > curl localhost:8080/users/
    {"message":"list=user1,user2,user3"}

### server

    2026/11/15 01:56:14 GET /users/ 2026-11-15 01:56:14.17707 -0700 MST m=+25.703181626
    2026/11/15 01:56:14 auth user
    2026/11/15 01:56:14 GetAllUsers

### client

    > curl localhost:8080/users/123
    {"message":"user_id=123"}

### server

    2026/11/15 01:56:18 GET /users/123 2026-11-15 01:56:18.555093 -0700 MST m=+30.081212918
    2026/11/15 01:56:18 auth user
    2026/11/15 01:56:18 GetUserById

### client

    > curl localhost:8080/panic/false
    {"message":"panic == false"}

### server

    2026/11/15 01:56:30 GET /panic/false 2026-11-15 01:56:30.108644 -0700 MST m=+41.634785585
    2026/11/15 01:56:30 auth user
    2026/11/15 01:56:30 GetPanic

### client

    > curl localhost:8080/panic/true
    {"message":"internal server error test"}

### server

    2026/11/15 01:56:35 GET /panic/true 2026-11-15 01:56:35.577288 -0700 MST m=+47.103440085
    2026/11/15 01:56:35 auth user
    2026/11/15 01:56:35 GetPanic
    2026/11/15 01:56:35 throwing panic:
    stack:
    goroutine 7 [running]:
    runtime/debug.Stack()
    	/opt/homebrew/Cellar/go/1.27.3/libexec/src/runtime/debug/stack.go:26 +0x64
    github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/middleware.Recovery.func1.1()
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/middleware/middleware.go:14 +0x30
    panic({0x1046b3548?, 0x104659190?})
    	/opt/homebrew/Cellar/go/1.27.3/libexec/src/runtime/panic.go:859 +0x120
    github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/handlers.GetPanic({0x1046f71d8, 0x22d4c0b180f0}, 0x22d4c0a168c0)
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/handlers/handlers.go:42 +0x2ac
    net/http.HandlerFunc.ServeHTTP(0x22d4c0b92210?, {0x1046f71d8?, 0x22d4c0b180f0?}, 0x22d4c0bdd7d8?)
    	/opt/homebrew/Cellar/go/1.27.3/libexec/src/net/http/server.go:2358 +0x38
    github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/middleware.Auth.func1({0x1046f71d8, 0x22d4c0b180f0}, 0x22d4c0a168c0)
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/middleware/middleware.go:32 +0x98
    net/http.HandlerFunc.ServeHTTP(0x22d4c0b92210?, {0x1046f71d8?, 0x22d4c0b180f0?}, 0x22d4c0bdd8a8?)
    	/opt/homebrew/Cellar/go/1.27.3/libexec/src/net/http/server.go:2358 +0x38
    github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/middleware.Logging.func1({0x1046f71d8, 0x22d4c0b180f0}, 0x22d4c0a168c0)
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/middleware/middleware.go:25 +0x118
    net/http.HandlerFunc.ServeHTTP(0x22d4c0bdd998?, {0x1046f71d8?, 0x22d4c0b180f0?}, 0x22d4c0b0c3f0?)
    	/opt/homebrew/Cellar/go/1.27.3/libexec/src/net/http/server.go:2358 +0x38
    github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/middleware.Recovery.func1({0x1046f71d8?, 0x22d4c0b180f0?}, 0x3?)
    	/Users/ric/code/go/src/github.com/rpcox/gobox/http/gorilla-mux/frame-work/pkg/middleware/middleware.go:17 +0x58
    net/http.HandlerFunc.ServeHTTP(0x22d4c0a16780?, {0x1046f71d8?, 0x22d4c0b180f0?}, 0x104158544?)
    	/opt/homebrew/Cellar/go/1.27.3/libexec/src/net/http/server.go:2358 +0x38
    github.com/gorilla/mux.(\*Router).ServeHTTP(0x22d4c0bbc0c0, {0x1046f71d8, 0x22d4c0b180f0}, 0x22d4c0a16640)
    	/Users/ric/code/go/pkg/mod/github.com/gorilla/mux@v1.8.1/mux.go:212 +0x190
    net/http.serverHandler.ServeHTTP({0x22d4c0b100c0?}, {0x1046f71d8?, 0x22d4c0b180f0?}, 0x1?)
    	/opt/homebrew/Cellar/go/1.27.3/libexec/src/net/http/server.go:3436 +0xb0
    net/http.(\*conn).serve(0x22d4c0b0e120, {0x1046f7678, 0x22d4c0b93860})
    	/opt/homebrew/Cellar/go/1.27.3/libexec/src/net/http/server.go:2157 +0x5a0
    created by net/http.(\*Server).Serve in goroutine 1
    	/opt/homebrew/Cellar/go/1.27.3/libexec/src/net/http/server.go:3604 +0x3e8

### client

    > curl localhost:8080/
    {"message":"hello world basic api"}

### server

    2026/11/15 01:57:20 GET / 2026-11-15 01:57:20.290974 -0700 MST m=+91.817209835
    2026/11/15 01:57:20 auth user
    2026/11/15 01:57:20 HomeHandler

