package mydict

import "errors"

var errNotFount = errors.New("Not Found")

type SDictionary map[string]string

func (d SDictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return errNotFount
}

type IDictionary map[string]int
