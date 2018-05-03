# Swagger blog example

Created for medium [story](https://medium.com/@romanyx90/handling-json-null-or-missing-values-with-go-swagger-4d7f37a2a7ca), to show code example.

## Run app

```bash
make start
```

Test with curl

```bash
curl -X PATCH "http://localhost:3332/blog/1" -H "Admin-Id: 1" -H "Content-Type: application/json" -d "{ \"title\": \"Example title\", \"markdown\": \"Example body\" }"
```

## Testing

```
make test
```
