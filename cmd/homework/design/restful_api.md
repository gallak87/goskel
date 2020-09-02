## What is a restful API?
- Primary features:
  - Client-Server
  - Stateless
  - Uniform design/invocation
  - Cacheable
  - Layered

#### Client-Server
- Model that calls for a client to communicate with a server
- Server exposes API endpoints and handles a client request with a server response
- Client "plugs into" these endpoints, calling them with a request, and handling the response
- Optional: Server-side can generate clients for easy consumption by the user

#### Stateless
- Server should not maintain any state between client requests, aka no session management
- Any required state must be handled at the client side (ex: request login consent, store bearer token, attach to each call)
- Each request must contain all the information needed to fulfill that request

#### Uniform design/invocation
- APIs should be designed with some consistency around resources/structure and verbs
  - Resources: 
    - Bias towards plural nouns `/orders` and `/orders/:id`
    - Consistency in parameterization - if using GET query params `/orders?id=1`, use it 
    everywhere, don't mix and match or if using query paths `/orders/:id` use it everywhere
    - Optional: Consider including "linking" in the response itself to other, related resources (random example: `GET /orders/123` could return order data and include a field with a link to `/orders/123/metadata` which points to a sub-resource of the order `123` and client can just use it directly instead of have to know it or construct it)
  - Verbs: If you are using POST for write calls, don't mix in PUTs, always use POST
- Versioning should be considered from the start
  - Use version prefix `/v1/orders/:id` and when a new response model is developed, bump to `/v2/orders:id`
- Optional: consider formatting for all requests, ex: `/v1/orders/:id?format=json`


#### Cacheable
- Responses that don't change should be cached
- Especially heavy used APIs like trade history, cache as many things as you can, with sane TTLs

#### Layered
- Client doesn't have knowledge of server-side, could be a server, could be a load-balanced proxy