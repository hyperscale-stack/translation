// Copyright 2021 Hyperscale. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package translation

import (
	"context"
	"testing"

	"github.com/hyperscale-stack/locale"
	"github.com/stretchr/testify/assert"
	"golang.org/x/text/language"
	"golang.org/x/text/message/catalog"
)

func tagToContext(tag language.Tag) context.Context {
	ctx := context.Background()

	return locale.ToContext(ctx, tag)
}

func TestTranslator(t *testing.T) {
	c := catalog.NewBuilder(catalog.Fallback(language.English))
	c.SetString(language.English, "Hello World", "Hello World")
	c.SetString(language.French, "Hello World", "Bonjour le monde")

	trans := New([]language.Tag{language.English, language.French}, WithCatalog(c))

	for _, item := range []struct {
		locale   language.Tag
		expected language.Tag
	}{
		{
			locale:   language.Norwegian,
			expected: language.English,
		},
		{
			locale:   language.French,
			expected: language.French,
		},
		{
			locale:   language.English,
			expected: language.English,
		},
		{
			locale:   language.Dutch,
			expected: language.English,
		},
	} {
		assert.Equal(t, item.expected, trans.GetSupportedLocale(item.locale))
	}

	assert.Equal(t, "Bonjour le monde", trans.Translate(tagToContext(language.French), "Hello World"))
	assert.Equal(t, "Hello World", trans.Translate(tagToContext(language.English), "Hello World"))
	assert.Equal(t, "Hello World", trans.Translate(tagToContext(language.Italian), "Hello World"))
}

func TestTranslatorWithDefaultLocale(t *testing.T) {
	c := catalog.NewBuilder(catalog.Fallback(language.English))
	c.SetString(language.English, "Hello World", "Hello World")
	c.SetString(language.French, "Hello World", "Bonjour le monde")

	trans := New([]language.Tag{language.English, language.French}, WithCatalog(c), WithDefaultLocale(language.French))

	for _, item := range []struct {
		locale   language.Tag
		expected language.Tag
	}{
		{
			locale:   language.Norwegian,
			expected: language.French,
		},
		{
			locale:   language.French,
			expected: language.French,
		},
		{
			locale:   language.English,
			expected: language.English,
		},
		{
			locale:   language.Dutch,
			expected: language.French,
		},
	} {
		assert.Equal(t, item.expected, trans.GetSupportedLocale(item.locale))
	}

	assert.Equal(t, "Bonjour le monde", trans.Translate(tagToContext(language.French), "Hello World"))
	assert.Equal(t, "Hello World", trans.Translate(tagToContext(language.English), "Hello World"))
	assert.Equal(t, "Bonjour le monde", trans.Translate(tagToContext(language.Italian), "Hello World"))
}

func BenchmarkGetSupportedLocale(b *testing.B) {
	trans := New([]language.Tag{language.English, language.French})

	for i := 0; i < b.N; i++ {
		trans.GetSupportedLocale(language.French)
	}
}

func BenchmarkTranslate(b *testing.B) {
	c := catalog.NewBuilder(catalog.Fallback(language.English))
	c.SetString(language.English, "Hello World", "Hello World")
	c.SetString(language.French, "Hello World", "Bonjour le monde")

	trans := New([]language.Tag{language.English, language.French}, WithCatalog(c))
	ctx := tagToContext(language.French)

	for i := 0; i < b.N; i++ {
		trans.Translate(ctx, "Hello World")
	}
}
