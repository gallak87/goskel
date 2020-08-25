# goskel
skeleton go app with oauth, pg and protobuf

## design
front-end: typescript + react
back-end: go + proto + pg
auth: hydra
deploy: docker-compose (stretch: k8s)


### plan
1. basic back-end go app
    1. proto
    1. user registration
    1. login
    1. "features" tbd
    1. dockerize
    
1. add oauth with hydra
    1. introduce
    1. dockerize --> docker-compose

1. front-end
    1. React-app-rewired
    1. add user registration page
    1. add login page
    1. add "features" page behind oauth
    