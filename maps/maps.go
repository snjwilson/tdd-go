package dictionary

type Dictionary	map[string]string
type DictionaryErr string

const(
	ErrUnkownWord = DictionaryErr("unknown word")
	ErrExistingWord = DictionaryErr("existing word")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrUnkownWord
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrUnkownWord:
		d[word] = definition
	case nil:
		return ErrExistingWord
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case ErrUnkownWord:
		return ErrUnkownWord
	default:
		return err
	}
	return nil
}