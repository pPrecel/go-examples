# unix socket with the rpc library

    Server logs:

    ```bash
    [･‿･]  go-examples/02-unix-sockets main ✗ go run cmd/server/main.go
    INFO[0000] Remove old socket: /tmp/sample.sock
    INFO[0000] Listen...
    INFO[0000] Serve new unix socket: /tmp/sample.sock
    INFO[0012] Request A: 5, B: 20
    INFO[0012] Result: 25
    ```

    Client logs:
    ```bash
    [･‿･]  go-examples/02-unix-sockets ✓ go run cmd/client/main.go 
    INFO[0000] 5 + 20 = 25
    INFO[0000] Duration: -72.596µs
    INFO[0000] 5 + 20 = 25
    INFO[0000] Duration: -1.242935ms
    ```
