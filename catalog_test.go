package translation

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en": &dictionary{index: enIndex, data: enData},
		"fr": &dictionary{index: frIndex, data: frData},
	}
	fallback := language.MustParse("en")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"%d books available": 0,
}

var enIndex = []uint32{ // 2 elements
	0x00000000, 0x00000033,
} // Size: 32 bytes

const enData string = "" + // Size: 51 bytes
	"\x14\x01\x81\x01\x00=\x01\x13\x02One book available\x00\x16\x02%[1]d boo" +
	"ks available"

var frIndex = []uint32{ // 2 elements
	0x00000000, 0x00000037,
} // Size: 32 bytes

const frData string = "" + // Size: 55 bytes
	"\x14\x01\x81\x01\x00=\x01\x14\x02Un livre disponible\x00\x19\x02%[1]d li" +
	"vres disponibles"

	// Total table size 170 bytes (0KiB); checksum: B02F55E7
