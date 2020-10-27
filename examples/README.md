# /examples

### create docker network
```
$ docker network create sandbox-net
```

### run the docker
```
$ docker-compose -f examples/docker-compose.yml -p kv-sandbox up -d
```

### consul ctl
```
# consul
$ consul kv put app/dev @examples/internal/config.yaml
  Success! Data written to: app/dev

$ consul kv get app/dev
...

# etcd
$ cat examples/internal/config.yaml | etcdctl put app/dev
OK

$ etcdctl get app/dev
...
```

### get kv data
```
$ go run examples/internal/env.go

$ go run examples/internal/consul.go

$ go run examples/internal/etcd.go
```