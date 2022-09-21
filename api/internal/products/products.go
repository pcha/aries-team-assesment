package products

import "time"

// Product is the model used to represent a product.
type Product struct {
	id          ID
	name        Name
	description Description
	createdAt   CreatedAt
}

// ID retunrs the id.
func (p Product) ID() ID {
	return p.id
}

// Name returns the name
func (p Product) Name() Name {
	return p.name
}

// Description returns the name.
func (p Product) Description() Description {
	return p.description
}

// CreatedAt returns the creation time.
func (p Product) CreatedAt() CreatedAt {
	return p.createdAt
}

// Make creates a new Product with the given name and name. It also set the creation time to the current.
func Make(nameVal, descriptionVal string) (Product, error) {
	name, err := ParseName(nameVal)
	if err != nil {
		return Product{}, err
	}
	description, err := ParseDescription(descriptionVal)
	if err != nil {
		return Product{}, err
	}

	return Product{
		name:        name,
		description: description,
		createdAt:   MakeCreatedAt(),
	}, nil
}

// BuildFrom return a Product representation with the given values.
//It shouldn't be used to represent a new Product but to represent an already existent one (for example a result of the DB)
func BuildFrom(id int64, name string, description string, createdAt time.Time) Product {
	return Product{
		id:          ID{val: id},
		name:        Name{val: name},
		description: Description{val: description},
		createdAt:   CreatedAt{val: createdAt},
	}
}
