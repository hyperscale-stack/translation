// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package translation

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

type Option func(t *translator)

func WithCatalog(c catalog.Catalog) Option {
	return func(t *translator) {
		t.catalog = c
	}
}

func WithDefaultLocale(locale language.Tag) Option {
	return func(t *translator) {
		t.defaultLocale = locale
	}
}
