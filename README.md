<!--
 Copyright (c) 2022 Bi Anlatsana
 
 This software is released under the MIT License.
 https://opensource.org/licenses/MIT
-->

# Goht
DOM String generator for Go.

## Why?
Goht allows you to generate html in server-side part with flexible usage.
## Example
```go
title := "Hello World"

dom := goht.Html(
    goht.Props{"lang": "en"},
    goht.RawDom(
        "body",
        goht.DomList(
            goht.RawDom("h1", title, goht.Props{"id": "title"}),
            goht.RawDom("p", "An Example", goht.Props{"id": "paragraph"}),
        ),
        goht.Props{},
    ),
)

println(dom)
```