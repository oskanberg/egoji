package main

import (
	"flag"
	"fmt"

	"github.com/oskanberg/egoji"
)

func main() {
	input := flag.String("input", "Hello World", "The input string")
	flag.Parse()
	translater := egoji.NewSimpleTranslate()
	output, _ := translater.Translate(*input)
	fmt.Println(output)
}
