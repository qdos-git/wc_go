package wc_lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//  The spec says Python standard library plus argparse only, but can non-ASCII bytes be counted without them?

//  Go standard library is lighter due to static compilation; have to explicitly import more.

//  Why is the world still using wc written in C?
//    Backwards compat - see your point.

func wc_core(args []string) {

	// var args []bool

	// var filename string

	// var data []byte

	// var chars, words, lines, bytec int

	args, filename := handle_args(args)

	data := load_data(filename)

	chars, words, lines, bytec := count_chars(data), count_words(data), count_lines(data), len(data)

	output_data(chars, words, lines, bytec, args, filename)

}

func handle_args(all_args []string) ([]bool, string) {

	join_args := strings.Join(all_args, " ")

	split_args := strings.Split(join_args, " ")

	//  Assume last argument is filename.

	filename := split_args[len(split_args)-1]

	var m_flag, w_flag, l_flag, c_flag, big_l_flag bool = false, false, false, false, false

	if split_args[0][0] == '-' {

		for _, char := range split_args[0] {

			switch char {
p
			case 'm':
				m_flag = true

			case 'w':
				w_flag = true

			case 'l':
				l_flag = true

			case 'c':
				c_flag = true

			case 'L':
				big_l_flag = true

			}
		}

	}

	//  If all false, set flags to default context.

	return []bool{m_flag, w_flag, l_flag, c_flag, big_l_flag}, filename

}

func load_data(filename string) []byte {

	data, _ := ioutil.ReadFile(filename)

	return data

}

func count_chars(data []byte) int {

	//  Loop over text, count chars.

	var ascii_count int

	for _, b := range data {

		if b >= 32 && b <= 126 {

			ascii_count++

		}

	}

	return ascii_count

}

func count_words(data []byte) int {

	var prev_byte_alpha bool

	var word_count int

	for _, b := range data {

		if prev_byte_alpha == true && b == 10 {

			word_count++

		}

		if (b >= 65 && b <= 90) || (b >= 97 && b <= 122) {

			prev_byte_alpha = true

		} else {

			prev_byte_alpha = false

		}

	}

	return word_count

}

func count_lines(data []byte) int {

	var newlines_count int

	for _, b := range data {

		if b == 10 {

			newlines_count++

		}

	}

	return newlines_count

}
