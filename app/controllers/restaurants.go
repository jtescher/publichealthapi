package controllers

import (
	"github.com/jtescher/publichealthapi/app/models"
	"github.com/robfig/revel"
	"github.com/robfig/revel/cache"
	"strings"
	"time"
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
		return c.Redirect(App.Index)
	}

	var restaurants []*models.Restaurant
	if err := cache.Get("restaurants_by_name_"+strings.ToLower(restaurantName), &restaurants); err != nil {
		restaurants = c.loadRestaurants(restaurantName)
		go cache.Set("restaurants_by_name_"+strings.ToLower(restaurantName), restaurants, 1*time.Minute)
	}

	return c.Render(restaurantName, restaurants)
}

func (c Restaurants) loadRestaurants(restaurantName string) []*models.Restaurant {
	results, err := c.Txn.Select(models.Restaurant{}, `SELECT * FROM Restaurants WHERE LOWER(Name) like $1`, "%"+strings.ToLower(restaurantName)+"%")
	if err != nil {
		panic(err)
	}

	var restaurants []*models.Restaurant
	for _, result := range results {
		r := result.(*models.Restaurant)
		restaurants = append(restaurants, r)
	}

	return restaurants
}
