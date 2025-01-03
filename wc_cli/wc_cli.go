package wc_cli


func main() {

	wc_core(os.Args[1:])

        chars, words, lines, bytec := count_chars(data), count_words(data), count_lines(data), len(data)

        output_data(chars, words, lines, bytec, args, filename)
	
}


func output_data(chars int, words int, lines int, bytec int, args []bool, filename string) {

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

