<p align="center">
	<a href="https://travis-ci.com/3timeslazy/cheky">
		<img alt="Build Status" src="https://travis-ci.com/3timeslazy/cheky.svg?branch=master">
	</a>
  	<a href="https://codecov.io/gh/3timeslazy/cheky">
		<img src="https://codecov.io/gh/3timeslazy/cheky/branch/master/graph/badge.svg" />
	</a>
	<a href="https://godoc.org/github.com/3timeslazy/cheky">
		<img src="https://godoc.org/github.com/3timeslazy/cheky?status.svg" alt="GoDoc">
	</a>
</p>

## What is it?
Cheky is a small library for validating request parameters, using the "Errors are Values" concept described by Rob Pike [here](https://blog.golang.org/errors-are-values)

## Why do you want to use it?
if you use such frameworks as Echo or Gin(or even ```net/http``` handlers), then maybe your request parameters validation code looks like this:
```golang
// HandlerFunc handles request with
// 2 query and 1 path parameters.
// Ex.:
// GET /comments/<lang>?offset=10&limit=20
func HandlerFunc(ctx framework.Context) error {
	lang := ctx.PathParam("lang")
	// langs here is a map[string]struct{}
	if _, ok := langs[lang]; !ok {
		return ctx.Error(
			fmt.Errorf("unsupported language '%s'", lang),
		)
	}

	limitStr := ctx.QueryParam("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		return ctx.Error(
			fmt.Errorf("limit should be integer"),
		)
	}

	if limit <= 0 {
		return ctx.Error(
			fmt.Errorf("limit should be greater or equal to 0"),
		)
	}

	offsetStr := ctx.QueryParam("offset")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		return ctx.Error(
			fmt.Errorf("offset should be integer"),
		)
	}

	if offset < 0 {
		return ctx.Error(
			fmt.Errorf("offset should be greater than 0"),
		)
	}

	// ...
}
```
With Cheky you can reduce the amount of boilerplate code.
So the code above looks like:
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
  		return ctx.Error(err)
	}
  
  // ...
}
```

## Install
```sh
go get github.com/3timeslazy/cheky
```
