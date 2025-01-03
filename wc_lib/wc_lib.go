package wc_lib

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//  The spec says Python standard library plus argparse only, but can non-ASCII bytes be counted without them?

//  Go standard library is lighter due to static compilation; have to explicitly import more.

//  Why is the world still using wc written in C?
//    Backwards compat - see your point.

func Wc_core(args []string) (int, int, int, int, []bool, string) {

	// var args []bool

	// var filename string

	// var data []byte

	// var chars, words, lines, bytec int

	args_loaded, filename := handle_args(args)

	data := load_data(filename)

	return count_chars(data), count_words(data), count_lines(data), len(data), args_loaded, filename

}

func handle_args(all_args []string) ([]bool, string) {

	//  Join all arguments with a space because... can't remember.

	join_args := strings.Join(all_args, " ")

	//  Maybe some text processing was suppose to go here.

	//  Split them back up with a space!

	split_args := strings.Split(join_args, " ")

	//  Assume last argument is filename.

	filename := split_args[len(split_args)-1]

	m_flag, w_flag, l_flag, c_flag, big_l_flag := false, false, false, false, false

	//  If first character of first string is a dash (no multidash possible here).

	if split_args[0][0] == '-' {

		for _, char := range split_args[0] {

			switch char {

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

		if ( b >= 32 && b <= 126 ) || ( b >= 9 && b <= 13 ) {

			ascii_count++

		}

	}

	return ascii_count

}

func count_words(data []byte) int {

	var prev_byte_alpha bool

	var word_count int

	for _, b := range data {

		if prev_byte_alpha == true && (b == 63 || b == 58 || b == 59 || b == 46 || b == 44 || b == 41 || b == 32 || b == 10) {

			word_count++

		}

		if (b >= 65 && b <= 90) || (b >= 97 && b <= 122) {

			prev_byte_alpha = true

		} else {

			prev_byte_alpha = false

		}

	}

	fmt.Println(word_count)

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

func Output_data(chars int, words int, lines int, bytec int, args []bool, filename string) {

	if !args[0] && !args[1] && !args[2] && !args[3] && !args[4] {

		fmt.Println(words, lines, chars, filename)

	} else {

		if args[0] {

			fmt.Printf("%d ", words)

		}

		if args[1] {

			fmt.Printf("%d ", lines)

		}

		if args[2] {

			fmt.Printf("%d ", bytec)

		}

		if args[3] {

			fmt.Printf("%d ", chars)

		}

		fmt.Println(filename)

	}

}



