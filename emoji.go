package main

import "fmt"

type translater interface {
	translate(string) (string, error)
}

type simpleTranslate struct {
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

func main() {
	var x = newSimpleEmoji()
	test, _ := x.emoji("hello")
	fmt.Println(string(test))
}
