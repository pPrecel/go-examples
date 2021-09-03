# first implementation of the rpc library

    Server logs:

    ```bash
    [･‿･]  go-examples/01-rpc main ✗ go run cmd/server/main.go
    INFO[0000] Listen and serve :1234...
    INFO[0004] Request A: 5, B: 20
    INFO[0004] Result: 25
    ```

    Client logs:
    ```bash
    [･‿･]  go-examples/01-rpc ✓ go run cmd/client/main.go 
    INFO[0000] 5 + 20 = 25
    INFO[0000] Duration: -82.318µs
    INFO[0000] 5 + 20 = 25
    INFO[0000] Duration: -4.330407ms
    ```
