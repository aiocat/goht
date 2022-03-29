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
            goht.RawDom("h1", goht.Encode(title), goht.Props{"id": "title"}),
            goht.RawDom("p", "An Example", goht.Props{"id": "paragraph"}),
            goht.RawDom("img", "", goht.Props{"src": "https://github.com/aiocat/args/blob/main/static/img/logo.png"}),
        ),
        goht.Props{},
    ),
)

println(dom)
```

## API Reference

- [Click here](https://pkg.go.dev/github.com/aiocat/goht) to get api reference for latest release

## Found a Bug / Error?

If you found a bug or an error, please create a new issue at gitlab repository.

## Contributing

If you want to contribute this project:

- Make sure you add comments for your codes.
- Please don't make something useless.

## Author(s)

- [Aiocat](https://github.com/aiocat)

## License

This project is distributed under MIT license.
