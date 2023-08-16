package main

import (
	"fmt"
	"time"

	"bison/leaflink"
	"bison/utils"
)

var (
	brynLeaflinkUserID = 94002
	leaflinkAppToken   = utils.EnvVarOrPanic("LEAFLINK_APP_TOKEN")

	companyBuyoutDate = func() time.Time {
		loc, err := time.LoadLocation("America/Chicago")
		if err != nil {
			panic(err)
		}
		return time.Date(2023, 5, 1, 0, 0, 0, 0, loc)
	}()
)

func main() {
	// Replaces all of the sales representatives on each sales order since the May 1, 2023 company buyout
	// with a single user so that Distru can import these orders without tons of manual labor on our part.

	client := leaflink.Client{
		AppToken: leaflinkAppToken,
	}

	orderSalesReps, err := client.GetAllOrderSalesReps()
	if err != nil {
		panic(err)
	}

	orderSalesReps = utils.Filter(orderSalesReps, func(osr leaflink.OrderSalesRep) bool { return osr.CreatedOn.After(companyBuyoutDate) })
	for _, osr := range orderSalesReps {
		err = client.DeleteOrderSalesRep(osr.ID)
		if err != nil {
			panic(err)
		}
	}

	recentOrders, err := client.GetOrdersSince(companyBuyoutDate)
	if err != nil {
		panic(err)
	}

	recentOrders = utils.Filter(recentOrders, func(o leaflink.Order) bool { return o.CreatedOn.After(companyBuyoutDate) })
	for _, order := range recentOrders {
		fmt.Println("order:", order.Number, order.CreatedOn)
		err = client.SetOrderSalesRep(order.Number, brynLeaflinkUserID)
		if err != nil {
			panic(err)
		}
	}
}
