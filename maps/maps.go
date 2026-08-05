package maps

import "errors"

type Dictionary map[string]string

func Search(dictionary Dictionary, word string) string {
	return dictionary[word]
}

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}
