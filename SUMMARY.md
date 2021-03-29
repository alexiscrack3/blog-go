# Summary

## Ready

The following is a list of all the endpoints I was able to implement within 4 hours in Go

- GET /posts -> Get list of posts
- GET /posts/:id -> Get a post by given id
- POST /posts -> Create a post
- PUT /posts/:id -> Update a post by given id
- DELETE /posts/:id -> Delete a post by given id
- GET /users/:id/posts -> Get list of posts by given user id

## TODO

- Use either a database or in-memory database instead of hard-coded list of objects.
- Use Docker to spin up Redis container.
- Use Docker to spin up Redis commander to visualize data from redis.
- Use Docker to spin up Postgres container.
- Use Docker to spin up Visualizer to visualize data from posgres.
- Create a Dockerfile to containerized API service.
- Add authentication to procted API endpoints from malicious users.
- Deploy API to heroku.
