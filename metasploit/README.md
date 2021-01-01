### metasploit

1. run server
```bash
[ಠ_ಠ]  go-examples/shodan main ✗ msfconsole
<...>

msf6 > load msgrpc
[*] MSGRPC Service:  127.0.0.1:55552
[*] MSGRPC Username: msf
[*] MSGRPC Password: <generated_password>
[*] Successfully loaded plugin: msgrpc
```

2. in other teminal

```bash
[･‿･]  go-examples/metasploit main ✗ export METASPLOIT_SERVER_HOST=127.0.0.1:55552
[･‿･]  go-examples/metasploit main ✗ export METASPLOIT_SERVER_USERNAME=msf
[･‿･]  go-examples/metasploit main ✗ export METASPLOIT_SERVER_PASSWORD=<generated_password>

[･‿･]  go-examples/metasploit main ✗ go run cmd/main.go
INFO[0000] Read configuration...
INFO[0000] Connecting...
INFO[0000] List session...
INFO[0000] Server version: {Version:6.0.24-dev-2af64c43e985ceaa73117c338442c672c9e42924 Ruby:2.7.1 x86_64-darwin15 2020-03-31 Api:1.0}
```