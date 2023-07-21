## Albums API

### Running
```
got get .
go run .
```

### GET /albums
```
curl http://localhost:3000/albums
```

### GET /albums/:id
```
curl http://localhost:3000/albums/2
```

### POST /albums

```
curl http://localhost:3000/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'
```