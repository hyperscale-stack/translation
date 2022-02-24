// Copyright 2022 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package translation

import (
	"context"

	"github.com/hyperscale-stack/locale"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

var _ Translator = (*translator)(nil)

type Translator interface {
	GetSupportedLocale(tag language.Tag) language.Tag
	Translate(ctx context.Context, key message.Reference, args ...interface{}) string
}

type translator struct {
	catalog          catalog.Catalog
	defaultLocale    language.Tag
	supportedLocales []language.Tag
	locales          map[language.Tag]*message.Printer
	matcher          language.Matcher
}

func New(locales []language.Tag, opts ...Option) Translator {
	t := &translator{
		defaultLocale:    defaultLocale,
		supportedLocales: locales,
		locales:          make(map[language.Tag]*message.Printer, len(locales)+1),
	}

	for _, opt := range opts {
		opt(t)
	}

	t.init()

	return t
}

func (t *translator) init() {
	supportedLocalesMap := map[language.Tag]struct{}{
		t.defaultLocale: {},
	}

	supportedLocales := []language.Tag{
		t.defaultLocale,
	}

	for _, tag := range t.supportedLocales {
		if _, ok := supportedLocalesMap[tag]; !ok {
			supportedLocalesMap[tag] = struct{}{}

			supportedLocales = append(supportedLocales, tag)
		}
	}

	t.supportedLocales = supportedLocales

	opts := []message.Option{}

	if t.catalog != nil {
		opts = append(opts, message.Catalog(t.catalog))
	}

	for _, tag := range t.supportedLocales {
		t.locales[tag] = message.NewPrinter(tag, opts...)
	}

	t.matcher = language.NewMatcher(supportedLocales)
}

func (t *translator) GetSupportedLocale(tag language.Tag) language.Tag {
	tag, _, _ = t.matcher.Match(tag)

	return tag
}

func (t *translator) Translate(ctx context.Context, key message.Reference, args ...interface{}) string {
	tag := locale.FromContext(ctx)

	tag = t.GetSupportedLocale(tag)

	printer := t.locales[tag]

	return printer.Sprintf(key, args...)
}
