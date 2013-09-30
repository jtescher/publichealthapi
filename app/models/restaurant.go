package models

import (
	"fmt"
	"github.com/robfig/revel"
)

type Restaurant struct {
	RestaurantId int
	Name         string
}

func (r *Restaurant) String() string {
	return fmt.Sprintf("Restaurant(%s)", r.Name)
}

func (r *Restaurant) Validate(v *revel.Validation) {
	v.Check(r.Name,
		revel.Required{},
		revel.MaxSize{150},
		revel.MinSize{1},
	)
}
