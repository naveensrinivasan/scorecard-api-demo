# Using ossf scorecard API

- https://api.securityscorecards.dev/
- https://github.com/ossf/scorecard#scorecards-rest-api

## Using the API
The API is available at https://api.securityscorecards.dev/. This API doesnt require any authentication. You can use the API to get the scorecard for a repository. The API is a REST API and it returns JSON.

## Example
```
curl -X GET "https://api.securityscorecards.dev/projects/github.com/ossf/scorecard" -H "accept: application/json" | jq
```

## Demo code
The demo code uses the API to get the scorecard for all the dependencies of a repository.
This code uses parse the `go.mod` file to get the dependencies and then uses the API to get the scorecard for each dependency.

## Running the demo code
```
go run main.go PATH_TO_GO_MOD_FILE_DIR
```
``` bash
go run main.go /Users/naveen/go/src/github.com/naveensrinivasan/cosign

```

The demo code shows all the dependencies that have been fuzzed.

## Example output
``` bash
Projects that are being fuzzed:
github.com/containerd/containerd 10
github.com/google/tink 10
github.com/grpc-ecosystem/grpc-gateway 10
github.com/imdario/mergo 10
github.com/docker/distribution 10
github.com/russross/blackfriday 10
github.com/Microsoft/go-winio 10
github.com/prometheus/prometheus 10
github.com/open-policy-agent/opa 10
github.com/google/flatbuffers 10
github.com/opencontainers/runc 10
github.com/pkg/sftp 10
github.com/go-redis/redis 10
github.com/apache/thrift 10
github.com/distribution/distribution 10
github.com/hashicorp/hcl 10
github.com/golang/snappy 10
github.com/jhump/protoreflect 10
github.com/clbanning/mxj 10
github.com/valyala/fasthttp 10
github.com/godbus/dbus 10
github.com/protocolbuffers/txtpbfmt 10
github.com/miekg/dns 10
github.com/pelletier/go-toml 10
github.com/BurntSushi/toml 10
github.com/coreos/etcd 10
github.com/veraison/go-cose 10
github.com/kevinburke/ssh_config 10
github.com/mattn/go-sqlite3 10
github.com/nats-io/nats-server 10
github.com/klauspost/compress 10
github.com/sigstore/sigstore 10
github.com/apache/beam 10
-----------------
Total number of dependencies : 736
The number of dependencies that are fuzzed : 33
```



