package main

import (
	"testing"
)

func TestEBCDIC277CharMap(t *testing.T) {
	var result []rune
	// ÆØÅæøå in EBCDIC 277
	ebcdic277Chars := []byte {0x7B, 0x7C, 0x5B, 0xC0, 0x6A ,0xD0 }
	for _, element := range ebcdic277Chars[:] {
		runes := map_ebcdic_byte_to_ascii_rune(element, ebcdic_to_ascii_277)
		result = append(result, runes...)
	}

	if string(result[:]) != "ÆØÅæøå" {
		t.Fail()
	}
}
