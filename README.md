### INTRODUCTION
Simple JSON-RPC TCP server in Golang which takes MSISDN number as input and returns 
MNO identifier, country dialling code, subscriber number and country identifier.

### INSTALLATION
```go
go get github.com/ttacon/libphonenumber
go run server.go
```
 
or

```
docker run -p 1234:1234 slani/go_rpc_json_s
```

### EXAMPLE
Here is an example of a JSON-RPC request object which should be sent through a TCP socket.
```json
--> {"jsonrpc":"2.0", "method":"Arith.ParseMsisdn", "params": ["+33123456789"], "id":3}
<-- {"id":3,"result":{"Cc":"33","Ndc":"1","Sn":"23456789","Ci":"FR"},"error":null}
```
 
 
 

The application uses [libphonenumber](https://github.com/ttacon/libphonenumber) library for parsing a MSISDN number.
