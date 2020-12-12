# rpc-test

Server logs:

```bash
[･‿･]  rpc-test main ✗ go run cmd/server/main.go
INFO[0000] Listen: localhost:1234
INFO[0000] Serve...
INFO[0006] Request A: 5, B: 20
INFO[0006] Result: 25
```

Client logs:
```bash
[･‿･]  rpc-test main ✓ go run cmd/client/main.go 
INFO[0000] 5 + 20 = 25
INFO[0000] Duration: -82.318µs
INFO[0000] 5 + 20 = 25
INFO[0000] Duration: -4.330407ms
```
