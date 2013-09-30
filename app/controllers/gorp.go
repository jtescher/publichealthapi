package controllers

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/jtescher/publichealthapi/app/models"
	_ "github.com/lib/pq"
	r "github.com/robfig/revel"
	"github.com/robfig/revel/modules/db/app"
)

var (
	Dbm *gorp.DbMap
)

func Init() {
	db.Init()
	Dbm = &gorp.DbMap{Db: db.Db, Dialect: gorp.PostgresDialect{}}

	setColumnSizes := func(t *gorp.TableMap, colSizes map[string]int) {
		for col, size := range colSizes {
			t.ColMap(col).MaxSize = size
		}
	}

	t := Dbm.AddTableWithName(models.Restaurant{}, "Restaurants").SetKeys(true, "RestaurantId")
	setColumnSizes(t, map[string]int{
		"Name": 150,
	})
	Dbm.TraceOn("[gorp]", r.INFO)
	Dbm.CreateTables()

	count, err := Dbm.SelectInt("SELECT COUNT(*) FROM Restaurants")
	if err != nil {
		panic(err)
	}

	if count == 0 {
		restaurants := []*models.Restaurant{
			&models.Restaurant{0, "Barbacco"},
			&models.Restaurant{0, "Bar Agricole"},
			&models.Restaurant{0, "Lolinda"},
		}

		for _, restaurant := range restaurants {
			if err := Dbm.Insert(restaurant); err != nil {
				panic(err)
			}
		}
	}
}

type GorpController struct {
	*r.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() r.Result {
	txn, err := Dbm.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
