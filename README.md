# Blog

## Getting started

To get started with the app, first clone the repo and `cd` into the directory:

```bash
$git clone git@github.com:alexiscrack3/blog-go.git
$cd blog-go
```

Then install the needed packages:

```bash
$go get
```

Finally, run the test suite to verify that everything is working correctly:

```bash
$go test -v ./...
```

If the test suite passes, you’ll be ready to run the app in a local server:

```bash
$go run main.go
```

## Endpoints

### GET /posts

```bash
curl --location --request GET 'http://localhost:3000/posts'
```

### GET /posts/:id

```bash
curl --location --request GET 'http://localhost:3000/posts/1'
```

### POST /posts

```bash
curl --location --request POST 'http://localhost:3000/posts' \
--header 'Content-Type: application/json' \
--data-raw '{
"title": "Title",
"body": "lorem ipsum"
}'
```

### PUT /posts

```bash
curl --location --request PUT 'http://localhost:3000/posts/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "title": "New Title",
    "body": "lorem ipsum"
}'
```

### DELETE /posts/:id

```bash
curl --location --request DELETE 'http://localhost:3000/posts/3'
```

### GET /users/:id/posts

```bash
curl --location --request GET 'http://localhost:3000/users/1/posts'
```
