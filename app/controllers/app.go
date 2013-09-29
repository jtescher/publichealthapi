package controllers

import "github.com/robfig/revel"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Restaurants(restaurantName string) revel.Result {
  c.Validation.Required(restaurantName).Message("Restaurant name is required.")
  c.Validation.MinSize(restaurantName, 3).Message("Restaurant name is not long enough.")
  
  if c.Validation.HasErrors() {
  	c.Validation.Keep()
  	c.FlashParams()
  	return c.Redirect(App.Index)
  }
  
	return c.Render(restaurantName)
}
