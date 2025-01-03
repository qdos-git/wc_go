
package wc_lib

import (
	"testing"
	"fmt"
)

func TestWcCore_1(t *testing.T) {

	// chars, words, lines, bytes, args_loaded, filename := Wc_core("empty_file")

	chars, words, lines, bytes, _, _ := Wc_core([]string{"../test_data/empty_file"})

	if (chars != 0 || words != 0 || lines != 0 || bytes != 0) {

		t.Fatalf("screwed_1")

	}

}

func TestWcCore_2(t *testing.T) {

	chars, words, lines, bytes, _, _ := Wc_core([]string{"../test_data/empty_file", "-cml"})

	if (chars != 0 || words != 0 || lines != 0 || bytes != 0) {

		t.Fatalf("screwed_2")

	}

}

func TestWcCore_3(t *testing.T) {

	chars, words, lines, bytes, _, _ := Wc_core([]string{"../test_data/output.bin"})

	fmt.Println(chars, words, bytes)

	if (chars != 0 || words != 0 || lines != 0 || bytes != 26) {

		t.Fatalf("screwed_3")

	}

}


func TestWcCore_4(t *testing.T) {

	chars, words, lines, bytes, _, _ := Wc_core([]string{"../test_data/lorem_ipsum_count.tx"})

	fmt.Println(bytes, chars, lines, words)

	if (chars != 3738 || words != 549 || lines != 9 || bytes != 3738) {

		t.Fatalf("screwed_4")

	}

}
