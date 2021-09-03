# the gRPC server / client

    1. Start the server
    ```bash
    [･‿･] go-examples/07-grpc/ main ✗ go run cmd/server/main.go
    INFO[0000] Starting server...
    INFO[0000] Listen And Serve...
    ```

    2. Run the client
    ```bash
    [･‿･] go-examples/07-grpc/ main ✗ go run cmd/client/main.go
    INFO[0000] sending request: {Val: 4, Name: "hello test name"}
    INFO[0000] received value: {Val: 4, Name: "hello test name"}
    ```

    ---

    3. Example server logs
    ```bash
    [･‿･] go-examples/07-grpc/ main ✗ go run cmd/server/main.go
    INFO[0000] Starting server...
    INFO[0000] Listen And Serve...
    INFO[0046] received: {Val: 4, Name: "hello test name"}
    INFO[0046] sending back: {Val: 4, Name: "hello test name"}
    ```
