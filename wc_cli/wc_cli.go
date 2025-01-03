package main

import(
	"wc_lib"
)


func main() {

	chars, words, lines, bytec := wc_core(os.Args[1:])

	output_data(chars, words, lines, bytec, args, filename)
	
}

