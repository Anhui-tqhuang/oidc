# setup

We use [dex](https://github.com/dexidp/dex) as an IdP in this workshop.

documentation: https://dexidp.io/docs/

## setup dex

1. build the dex binary
```sh
# make sure that '$GOPATH/src/github.com/dexidp' exists
$ cd $GOPATH/src/github.com/dexidp
$ git clone https://github.com/dexidp/dex.git
$ cd dex/
$ make
```

2. create a configuration file `newcfg,yaml`, we specify the grpc configuration so we could communicate with dex server with grpc client.
```yaml
issuer: http://127.0.0.1:5556/dex
storage:
  type: sqlite3
  config:
    file: examples/dex.db
web:
  http: 0.0.0.0:5556
# we could register client-app here or using grpc client.
staticClients:
- id: example-app
  redirectURIs:
  - 'http://127.0.0.1:5555/callback'
  name: 'Example App'
  secret: ZXhhbXBsZS1hcHAtc2VjcmV0
connectors:
- type: mockCallback
  id: mock
  name: Example
grpc:
  # Cannot be the same address as an HTTP(S) service.
  addr: 127.0.0.1:5557
  reflection: true

# Let dex keep a list of passwords which can be used to login to dex.
enablePasswordDB: true

# A static list of passwords to login the end user. By identifying here, dex
# won't look in its underlying storage for passwords.
#
# If this option isn't chosen users may be added through the gRPC API.
staticPasswords:
- email: "admin@example.com"
  # bcrypt hash of the string "password"
  hash: "$2a$10$2b2cU8CPhOTaGrs1HRQuAueS7JTT5ZHsHSzYiFPm1leZck7Mc8T4W"
  username: "admin"
  userID: "08a8684b-db88-4b73-90a9-3cd1661f5466"
```

3. start the server
```sh
./bin/dex serve newcfg.yaml
```

4. i wrote a simple grpc client in this repo, which could be used to add new user to the dex server.
```sh
# add a new user, for example
$ go run main.go --username=<username> \
    --password=<password> \
    --email=<email> \
    --address=<dex-server-address>
```

5. write a client app, the [reference](https://dexidp.io/docs/using-dex/), note that we only use `id token` in this app and could ignore the`access_token` and `refresh_token`.

6. test your client app.
