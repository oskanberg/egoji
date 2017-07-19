package egoji

type translater interface {
	translate(string) (string, error)
}

type simpleTranslate struct {
}

type emojier interface {
	emoji(string) (rune, error)
}

type simpleEmoji struct {
}

func (*simpleEmoji) emoji(s string) (rune, error) {
	return '👋', nil
}

func NewTranslater() interface{} {
	return nil
}
