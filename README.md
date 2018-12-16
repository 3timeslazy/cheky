## Doc
https://godoc.org/github.com/3timeslazy/cheky


## Install
```sh
go get github.com/3timeslazy/cheky
```

# Interface
```golang

// HandlerFunc handles request with
// 2 query and 1 path parameters.
// Ex.:
// GET /comments/<lang>?offset=10&limit=20
func HandlerFunc(ctx framework.Context) (err error) {
  var check = cheky.Ctx(ctx)
  
  var lang string
  var offset, limit int
  
  check.Path("lang").Str(&lang).OneOf("en", "fr", "ru")
  check.Query("offset").Int(&offset).Ge(0)
  check.Query("limit").Int(&limit).Gt(0)
  if err = check.Err(); err != nil {
    return ctx.JSON(400, err)
  }
  
  // do smth...
}
```
