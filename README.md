# context-helpers
Generate getters/setters for the things you store in `gorilla/context`.

## Installation

``` shell
go get github.com/codegangsta/context-helpers
```

## Usage

`gorilla/context` is a great golang package for storing request-specific data.
One of the caveats, however, is its relience on a `map[interface{}]interface{}`
for storing the data. We essentially lose any type information we had when we
put the value into `gorilla/context`, and that is no good!

This little command line tool tries to make this a little easier by generating
getters/setters based on the types of values you put into your
`gorilla/context`. For instance, if I had a `*Session` type that I wanted to
map to the context in a type-safe way, I can do this.

``` shell
context-helpers Session
```

And it will generate a `session_helper.go` file with the following content:

```
package main

import (
	"net/http"

	"github.com/gorilla/context"
)

type sessionHelperKey int

const sessionKey sessionHelperKey = 0

func GetSession(r *http.Request) *Session {
	if rv := context.Get(r, sessionKey); rv != nil {
		return rv.(*Session)
	}
	return nil
}

func SetSession(r *http.Request, val *Session) {
	context.Set(r, sessionKey, val)
}
```
