package cli

import (
	"flag"
	"fmt"

	"github.com/oskanberg/egoji"
)

func main() {
	input := flag.String("input", "Hello World", "The input string")
	flag.Parse()
	translater := egoji.NewTranslater()
	output := translater.Translate(*input)
	fmt.Println(output)
}
