package products

import "errors"

// DescriptionCanNotBeEmptyErr is the error to return if the description is empty
var DescriptionCanNotBeEmptyErr = errors.New("description can't be empty")

// Description is an object value for the field description of Product
type Description struct {
	val string
}

// String returns the string representation of the Description.
func (d Description) String() string {
	return d.val
}

// ParseDescription parses a string to a description. DescriptionCanNotBeEmptyErr is returned if the string is empty.
func ParseDescription(val string) (Description, error) {
	if val == "" {
		return Description{}, DescriptionCanNotBeEmptyErr
	}

	return Description{val: val}, nil
}
