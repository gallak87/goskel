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

## tl;dr how to run
1. docker
    - build everying and run via docker-compose (~2-4 minutes) from root directory:
    - `docker-compose up --build`

## how to run
1. container traffic description:
    - goskel: go app serving grpc on :9090
    - goskel-web: react webapp serving on :3000
    - grpcwebproxy: proxy taking http requests from goskel-web on :8080, and proxying them to goskel on :9090
1. pre-reqs
    - clone the repo (generated clients committed to source)
    - docker for desktop: works out of the box
    - running locally:
        - install all deps mentioned above and added to PATH
        - install nodejs and yarn
1. docker
    - build everying and run via docker-compose (~2-4 minutes) from root directory:
        - `docker-compose up --build`
    - build individually and run afterwards
        - goskel: `docker build -t goskel:latest .`
        - goskel-web: `pushd web; docker build -t goskel-web:latest .; popd`
        - run all: `docker-compose up`
1. locally (WINDOWS)
    1. Download or build from source, put these in a folder and add to PATH variable. (TODO: links)
        - protoc-gen-go.exe
        - protoc-gen-grpc-gateway.exe
        - protoc-gen-swagger.exe
        - protoc-gen-grpc-web.exe
        - protoc.exe
        - grpcwebproxy.exe
    1. run:
        - `pushd proto; ./gen-proto.ps1; popd`
        - `./start-grpc.ps1`
        - in new tab/terminal: `cd web; yarn start`
1. clean up (regain 7-8GB)
    - `docker container prune`
    - `docker image prune -a`


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
- add protoc release links above
- add dependencies: hydra, pg
- implement server-side logic for user-service (with db + go-pg)
- add makefile and automate stuffz
- implement front-end session management and routes
- implement hydra login/consent/logout routes

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
    