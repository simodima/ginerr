# Gin Err 
Gin err provides a convenient way to transform an internal application error into an HTTP JSON response with the following fields:
- code: the internal error code
- message: the internal error message

Furthermore the HTTP status code can be easily mapped to the internal error by providing a translation map as follows

```go
AutoResponse(map[int]int{
  5001: 500,
  5002: 404,
  5003: 401,
})
```