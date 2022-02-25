# Hyperscale Translation [![Last release](https://img.shields.io/github/release/hyperscale-stack/translation.svg)](https://github.com/hyperscale-stack/translation/releases/latest) [![Documentation](https://godoc.org/github.com/hyperscale-stack/translation?status.svg)](https://godoc.org/github.com/hyperscale-stack/translation)

[![Go Report Card](https://goreportcard.com/badge/github.com/hyperscale-stack/translation)](https://goreportcard.com/report/github.com/hyperscale-stack/translation)

| Branch | Status                                                                                                                                                                               | Coverage                                                                                                                                                         |
| ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| master | [![Build Status](https://github.com/hyperscale-stack/translation/workflows/Go/badge.svg?branch=master)](https://github.com/hyperscale-stack/translation/actions?query=workflow%3AGo) | [![Coveralls](https://img.shields.io/coveralls/hyperscale-stack/translation/master.svg)](https://coveralls.io/github/hyperscale-stack/translation?branch=master) |

The Hyperscale translation library provides a simple translation manager over `x/text` package.

## Example

```go
package main

import (
    "fmt"

    "github.com/hyperscale-stack/locale"
    "github.com/hyperscale-stack/translation"
    "golang.org/x/text/language"
    "golang.org/x/text/message/catalog"
)

func main() {
    // plase use golang.org/x/text/cmd/gotext for build catalog
    c := catalog.NewBuilder(catalog.Fallback(language.English))
    c.SetString(language.English, "Hello World", "Hello World")
    c.SetString(language.French, "Hello World", "Bonjour le monde")

    // the default language is always language.English, for override use translation.WithDefaultLocale()
    trans := translation.New(
        []language.Tag{
            language.English,
            language.French,
        },
        translation.WithCatalog(c),
    )

    ctx := context.Background()

    ctx = locale.ToContext(ctx, language.French)

    trans.Translate(ctx, "Hello World") // return Bonjour le monde
}

```

## License

Hyperscale Translation is licensed under [the MIT license](LICENSE.md).
