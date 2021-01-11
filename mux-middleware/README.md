# the negroni tool as a mux middleware

1. Start server
```bash
[･‿･]  go-examples/mux-middlewar main ✗ go run cmd/server/main.go     23:03:52
INFO[0000] Starting server...
INFO[0000] Listen And Serve...
```

2. Use curl or other tool to rich server
```bash
[･‿･]  go-examples/mux-middlewar main ✗ curl "localhost:8080/hello?pass=1234&user=admin"
looks ok :
```

---

3. Example server logs
```bash
[･‿･]  go-examples/mux-middlewar main ✗ go run cmd/server/main.go     23:03:52
INFO[0000] Starting server...
INFO[0000] Listen And Serve...
INFO[0021] Hello from AuthHandler
INFO[0021] URL values                                    pass=1234 user=admin
INFO[0021] Hello from HelloHandler
INFO[0021] Request received for user: admin
INFO[0021] Hello from HelloHandle
```