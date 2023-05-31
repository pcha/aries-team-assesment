package products

// ID is the object value of the field id of a Product.
type ID struct {
	val int64
}

// Int64 returns the representation of the ID.
func (i ID) Int64() int64 {
	return i.val
}

// ParseID parses a int64 value to a ID.
func ParseID(val int64) ID {
	return ID{val: val}
}
