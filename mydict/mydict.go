package mydict

import "errors"

// Dictionary type
// alias (like typedef in c)
type Dictionary map[string]string 

var (
	errNotFound = errors.New("Not Found")
	errCantUpdate = errors.New("Can't update non-existing word")
	errWordExists = errors.New("Word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to dictionary
func (d Dictionary) Add(word, def string) (error) {
	_, err := d.Search(word)
	switch err {
	case errNotFound: // doesn't exist yet
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}


// Update a word already exists
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word from dictionary 
func (d Dictionary) Delete(word string) {
	delete(d, word) // return nothing 
}