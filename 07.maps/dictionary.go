package maps

type DictionaryErr string

// To implement the error interface
func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("the word already exists in the dictionary")
	ErrWordDoesNotExist = DictionaryErr("the word does not exist in the dictionary")
)

type Dictionary map[string]string

func (d Dictionary) Search(name string) (string, error) {
	value, ok := d[name]

	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(name, value string) error {
	_, err := d.Search(name)

	switch err {
	case ErrNotFound:
		d[name] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(name, value string) error {
	_, err := d.Search(name)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[name] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(name string) {
	delete(d, name)
}
