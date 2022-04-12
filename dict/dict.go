package dict

import "errors"

// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("Aleady Exist")
	errCantUpdate = errors.New("Can't word already exist")
)

// Search for word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to the dictionary
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

// Update a word to the dictionary
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	if err == nil {
		d[word] = def
	} else if err == errNotFound {
		return errCantUpdate
	}
	return nil
}

// Update a word to the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
