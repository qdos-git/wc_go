
package wc_lib

import (
	"testing"
	"fmt"
)

func TestWcCore_1(t *testing.T) {

	// chars, words, lines, bytes, args_loaded, filename := Wc_core("empty_file")

	chars, words, lines, bytes, _, _ := Wc_core([]string{"empty_file"})

	fmt.Println(chars, words)

	if (chars != 0 && words != 0 && lines != 0 && bytes != 0) {

		t.Fatalf("screwed_1")

	}

}
