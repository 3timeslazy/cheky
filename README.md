# Cheky

Cheky is an invisible, small validator of path or query parameters.

## Interface
```golang
.Bool(b *bool)
  .True()
  .False()
    .Not()
.Int(i *int)
  .Gt(i64 int64)
  .Ge(i64 int64)
  .Lt(i64 int64)
  .Le(i64 int64)
  .OneOf(i64 ...int64)
  .NoOne(i64 ...int64)
    .Abs()
.Str(str *string)
  .Gt(i64 int64)
  .Ge(i64 int64)
  .Lt(i64 int64)
  .Le(i64 int64)
  .OneOf(i64 ...int64)
  .NoOne(i64 ...int64)
  .Regexp(re string)
```

## How it should work
```golang
package handlers

import (
  "github.com/3timeslazy/cheky"
  "github.com/author/framework"
)

// Handler handles request with 
// 1 path and 1 query parameters.

func Handler(ctx framework.Context) (err error) {
  var p1 int64
  var q1 bool
  
  var check = cheky.Ctx(ctx)
  // if 'path' is int64 and it -10 < 'path' <= 10 than
  // |'path'| will be stored into p1 variable.
  check.Path("path").Int64(&p).Gt(-10).Le(10).Abs()
  // if 'query' is boolean and it is true than
  // false will be stored into q1 variable.
  check.Query("query").Bool(&q1).True().Not()
  if err = check.Err(); err != nil {
    // do smth
  }
  // code
}
```
