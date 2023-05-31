package products

import "time"

// CreatedAt is the object value for the createdAt field of Product
type CreatedAt struct {
	val time.Time
}

// Time returns a time.Time representation of the CreatedAt
func (c CreatedAt) Time() time.Time {
	return c.val
}

// MakeCreatedAt returns a CreatedAt for the current time
func MakeCreatedAt() CreatedAt {
	return CreatedAt{
		val: time.Now(),
	}
}
