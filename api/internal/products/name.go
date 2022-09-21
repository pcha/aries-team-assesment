package products

import "errors"

// NameCanNotBeEmptyErr is the error to returns if the name is empty.
var NameCanNotBeEmptyErr = errors.New("name can't be empty")

// Name is the obeject value used to represent the property name of a Product.
type Name struct {
	val string
}

// String resturns the string representation of the Name.
func (n Name) String() string {
	return n.val
}

// ParseName parses the given string to a Name. NameCanNotBeEmptyErr is returned if the string is empty.
func ParseName(val string) (Name, error) {
	if val == "" {
		return Name{}, NameCanNotBeEmptyErr
	}
	return Name{val: val}, nil
}
