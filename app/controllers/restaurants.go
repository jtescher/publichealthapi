package controllers

import (
	"github.com/jtescher/publichealthapi/app/models"
	"github.com/jtescher/publichealthapi/app/routes"
	"github.com/robfig/revel"
	"strings"
)

type Restaurants struct {
	App
}

func (c Restaurants) Index(restaurantName string) revel.Result {
	c.Validation.Required(restaurantName).Message("Restaurant name is required.")
	c.Validation.MinSize(restaurantName, 3).Message("Restaurant name is not long enough.")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(routes.App.Index)
	}

	results, err := c.Txn.Select(models.Restaurant{}, `SELECT * FROM Restaurants WHERE LOWER(Name) like $1`, "%"+strings.ToLower(restaurantName)+"%")
	if err != nil {
		panic(err)
	}

	var restaurants []*models.Restaurant
	for _, result := range results {
		r := result.(*models.Restaurant)
		restaurants = append(restaurants, r)
	}

	return c.Render(restaurantName, restaurants)
}