package main

import (
	"os"
	"wc_lib"
)

func main() {

	wc_lib.Output_data(wc_lib.Wc_core(os.Args))

}
