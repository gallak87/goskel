# goskel
skeleton full-stack app with typescript/react + go, oauth, pg and protobuf

## structure
```
cmd/    --> go entry points (client, server)
common/ --> tooling
pkg/    --> go server-side implementation
proto/  --> proto files + generated clients
web/    --> react webapp
    public/     --> web assets
    src/        --> webapp implementation
        api/        --> proto-generated typescript client
        components/ --> web components (ex: login component)
```

## progress
#### WORKING
- protoc generating typescript and go clients
- go cmds using grpc client connects to grpc server
- react typescript webapp with dummy login form
    - talks to back-end via typescript proto client
- proto/gen-proto.ps1 to generate protos
- start-grpc.ps1 to start up both go grpc server + grpcwebproxy locally
- dockerfile + docker-compose with grpcwebproxy + goskel

#### IN PROGRESS
- add dependencies: hydra, pg
- implement server-side logic for user-service
- add makefile and automate stuffz
- revisit protoc generation, make-ify, restructure?

## design
- front-end: typescript + react
- back-end: go + proto + pg
- oauth: hydra
- deploy: docker-compose (stretch: k8s)

### plan
1. basic back-end go app
    1. proto service
    1. user registration
    1. login handler
    1. "features" tbd
    1. dockerize
    1. client/server cmds
    1. tests!!!!
    1. add session management (stretch: mfa?)

1. add pg database
    1. dockerize --> docker-compose
    1. add pg db dep
    1. add go-pg dep
    1. create user-models
    1. add first migration (stretch: create migration framework)

1. add oauth with hydra
    1. introduce
    1. dockerize --> docker-compose
    1. add session management (token creation/expiration)

1. front-end
    1. create-react-app with typescript
    1. add user registration page
    1. add login form
    1. add "features" page behind oauth
    1. dockerize --> docker-compose
    1. add session management
    