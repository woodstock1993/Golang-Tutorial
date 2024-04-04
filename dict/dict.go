package dict

import "errors"

var NotFound = errors.New("Not Found")
var CantUpdate = errors.New("Can't update")
var WordExists = errors.New("That word already exists")

// Dictionary type
type Dictionary map[string]string

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", NotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case NotFound:
		d[word] = def
	case nil:
		return WordExists
	}
	return nil
}

// Update Dictioniary
func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		d[key] = value
	case NotFound:
		return CantUpdate
	}
	return nil
}

// Delete Dictionary
func (d Dictionary) Delete(key string){
	delete(d, key)
} 