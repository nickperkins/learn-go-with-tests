package main

// Dictionary map
type Dictionary map[string]string

const (
	// ErrNotFound error
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	// ErrWordExists error
	ErrWordExists = DictionaryErr("word already exists")
	// ErrWordDoesNotExist error
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// DictionaryErr custom type
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

// Search for a word in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil

}

// Add a word to the dictionary
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil

}

// Update a word in the dictionary
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// Delete a word from the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
