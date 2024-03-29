package main

import (
	"flag"
	"fmt"
)

var RULES = flag.String("r", "B3/S23", "Set the rules. Default : B3/S23")
var L = flag.Int("l", 24, "Set the number of lines. Default : 24")
var C = flag.Int("c", 80, "Set the number of columns. Default : 80")
var CYCLE = flag.Int("cycle", 100, "Set the number of cycles. Default : 100")
var MAP = flag.String("m", "", "Set the map a the beginning of the game.")
var FILE = flag.String("f", "", "Set the map a the beginning of the game.")

func main() {
	flag.Parse()
	fmt.Println("The rules are", *RULES, "in a world of", *C, "by", *L)
	err, game := Init(*RULES, *MAP, *FILE, *C, *L, *CYCLE)
	if err != nil {
		fmt.Println("Invalid file.")
		return
	}
	Launch(*C, *L, *CYCLE, game)
}
