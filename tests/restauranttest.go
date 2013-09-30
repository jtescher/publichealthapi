package tests

import "github.com/robfig/revel"

type RestaurantTest struct {
	revel.TestSuite
}

func (t *RestaurantTest) Before() {
	println("Set up")
}

func (t RestaurantTest) TestThatSearchPageWorks() {
	t.Get("/App/Restaurants?restaurantName=Barbac")
	t.AssertOk()
	t.AssertContentType("text/html")
	t.AssertContains("Search: Barbac")
	t.AssertContains("Barbacco")
}

func (t *RestaurantTest) After() {
	println("Tear down")
}
