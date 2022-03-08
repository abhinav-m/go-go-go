package main

import (
	"fmt"
	"unicode/utf8"
)

/*
	Go string is a read only slice of bytes
	The language and standard library treat strings specially
	containers of text encoded in UTF-8.
	Unlike other languages where strings are made of "characters".
	In GO - a character concept is called a "rune", which is an integer
	representing unicode code point
*/

func main() {
	// Strings are containers of texts encoded in
	// in UTF-8 (supports easier internationalization)
	// Instead of characters, we have runes
	// Runes are integer representations of unicode code point
	// Strings are read only slice of bytes
	const s = "สวัสดี" // Hello in thai assigned to s

	// Since strings are equivalent to []byte,
	// this will produce the length of the raw bytes stored within.
	fmt.Println("Len:", len(s))

	// Raw byte values at each index
	// generates hex values of all bytes that
	// constitute code points in s
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()

	// To count how many runes are in a string,
	// we can use the utf8 package.
	// Note that the run-time of RuneCountInString
	// dependes on the size of the string,
	// because it has to decode each UTF-8 rune sequentially.
	// Some Thai characters are represented by multiple UTF-8
	// code points, so the result of this count may be surprising

	fmt.Println("Rune count", utf8.RuneCountInString(s))

	// Range loop handles strings specially
	// decoding each rune along with its offset in the string
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// We can achieve the same iteration
	// by using the utf8.DecodeRuneInString function explicitly.
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		// Get value of rune and width
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

// rune is built in datatype
func examineRune(r rune) {
	// VALUES enclosed in
	// single quotes are rune
	// literals
	//	We can compare a rune value to a rune literal directly.
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
