package egoji

type translater interface {
	Translate(string) (string, error)
}

type simpleTranslate struct {
}

func (s simpleTranslate) Translate(input string) (string, error) {
	simpleEmoji := newSimpleEmoji()
	//TODO: split
	output, _ := simpleEmoji.emoji(input)
	return string(output), nil
}

type emojier interface {
	emoji(string) (rune, error)
}

type simpleEmoji struct {
	database map[string]rune
}

func newSimpleEmoji() *simpleEmoji {
	database := make(map[string]rune)
	database["hello"] = '👋'
	return &simpleEmoji{database}
}

func (e *simpleEmoji) emoji(s string) (rune, error) {
	return e.database[s], nil
}

func NewSimpleTranslate() translater {
	return simpleTranslate{}
}

// func main() {
// 	var x = newSimpleEmoji()
// 	test, _ := x.emoji("hello")
// 	fmt.Println(string(test))
// }
