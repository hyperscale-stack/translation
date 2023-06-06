package main

import (
	"context"

	"github.com/hyperscale-stack/translation"
	_ "github.com/hyperscale-stack/translation/example/translations"
	"golang.org/x/text/language"
)

// nolint: forbidigo
func main() {
	trans := translation.New([]language.Tag{
		language.English,
		language.French,
	})

	totalBookCount := 1

	ctx := context.Background()

	println(trans.Translate(ctx, "%d books available", totalBookCount))

	totalBookCount += 1

	println(trans.Translate(ctx, "%d books available", totalBookCount))
}
