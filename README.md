# go-boilerplate

This is a golang boilerplate for Mixin Bot.

run 

```bash
./go-boilerplate -f YOUR_KEYSTORE_FILE help
```

to see the help.

It enables several widely used functionalities by default:

## Getting Started

### 1. adding more APIs

1. define new API at `/handler/serve.go` 
2. inject necessary dependences at `/cmd/httpd/httpd.go` and pass them into the handler
3. implement them under the `/handlers`
4. build and run `./go-boilerplate -f YOUR_KEYSTORE_FILE httpd [port]`

### 2. adding more workers

1. implement new workers at `/workers`
2. add the new instance of the workers to `workers` at `/cmd/worker/worker.go`
3. inject necessary dependences at `/cmd/worker/worker.go` and pass them into the handler
4. build and run `./go-boilerplate -f YOUR_KEYSTORE_FILE worker [health check port]`

### 3. handling messages

1. implement the handlers at `/messages`
2. inject necessary dependences at `/cmd/worker/worker.go` and pass them into the handler `messenger`
3. call the handlers at `/worker/messenger/messenger.go`

### 4. handling snapshots

1. put your code at `/worker/syncer/syncer.go:run` to handle the snapshots you need.
2. inject necessary dependences at `/cmd/worker/worker.go` and pass them into the handler `syncer`
3. the checkpoint of syncer is stored in the table `properties` with key `syncer:snapshot_checkpoint`, change it if you don't want to wait for a long time.

### 5. some handy tips

1. only inject the dependences (store, i18n, etc) you need into the workers or apis

### 6. removing unnecessary functionalities

feel free to remove any functionality you do not need:

1. syncer
  - `/cmd/worker` - the entry command to start workers
  - `/worker/syncer` - a worker to sync snapshots from Mixin Network. 
  - `/core` - database model definitions of snapshots and assets
  - `/store` - database stores of snapshots and assets
2. httpd
  - `/cmd/httpd` - the entry command to start a httpd server
  - `/handler` - the HTTP requset handlers
3. messenger
  - `/cmd/worker` - the entry command to start workers
  - `/worker/messenger` - a worker to handle incoming messages
  - `/message/` - the handlers to handle messages
4. migrater
  - `/cmd/migrate` - the entry command to migrate database
5. the echo command
  - `/cmd/echo` - a simple command to show you how to write a command
