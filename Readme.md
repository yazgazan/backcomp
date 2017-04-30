
## API Proposal

Backcomp (good name pending) is a tool intended to watch for breaking changes in HTTP/JSON apis.
A breaking change being change being defined by:

- A type changing for a field or array element
- A field missing in the response

### Definitions

- Versions are Semver
- Record: a pair of request/response
- Collection: a list of records

### Workflow

For two versions, `n` and `m`, with `n < m`:

- The responses sent back by version `n` are recorded
- The responses sent back by version `m` are compared with the responses stored for version `n`.
  Breaking changes are reported back to the developer.

### Models & File format

Records for a version `n` are saved in a file `records-n.json`.

```go
type Collection struct {
  Version string
  Records []Record
}

type Record struct {
  Request
  Response
}

type Request struct{
  Method string
  URL    string
  Header http.Header
  Body   interface{} // contains unmarshaled JSON
}

type Response struct {
  StatusCode int
  Status     string
  Header     http.Header
  Body       interface{} // contains unmarshaled JSON
}
```

### Assumptions

- For simplicity, we only deal with HTTP/1.1
- We presume no https
- The underlying models used by the tested api doesn't use dynamic types (abstract classes, ...)

### Tooling

The developer is provided with the following tools:

- `import`: convert supported formats (curl, har, http, postman, collection) to a collection
- `hydrate`: hydrate a collection with responses when missing.
  (useful when importing from curl, postman and http)
- `test`: send requests from a collection to a specified endpoint and compare with the stored responses.

### Known limitations

- This tool can (obviously) not take semantics into account:
  if two values have the same type/layout, they are considered compatible.
  (Meaning their content is not taken into consideration, such as two numbers on different scales)
- If a tested response contains an empty array in it's response (for a given request), the underlying model
  if the array's elements cannot be tested. (A warning could be displayed)
