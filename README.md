# goaccept

This is a simple library to assist in content negotiation. No dependencies and easy to use.

```golang
contentType, err := accept.Negotiate(req.Header().Get("Accept"), "application/json", "text/html")
if err != nil {
  switch err.(type) {
  case accept.ErrorNotAcceptable:
    // return a 406 or set the default content type
  case accept.ErrorBadAccept:
    // return a 400; they passed an improperly formatted accept header
  default:
    // this shouldn't happen; return a 500
  }
}
```