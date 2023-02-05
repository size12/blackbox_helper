package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/theplant/luhn"
)

func main() {
	cfg := GetConfig()
	service := NewService(cfg)

	err := service.AddGoods(goods)
	if err != nil {
		log.Fatal("Failed add new goods to service: ", err)
	}

	orders := GenerateOrders(brands, goods, 10)
	err = service.AddOrders(orders)
	if err != nil {
		log.Fatal("Failed add new orders to service: ", err)
	}

	fmt.Println("Number, Accrual")

	for _, order := range orders {
		fmt.Printf("%-7v %v\n", order.Number, order.WaitAccrual)
	}

}

func GenerateItemName(brands []string, goods []string) string {
	name := brands[rand.Intn(len(brands))] + " " + goods[rand.Intn(len(goods))]
	return name
}

func GenerateOrders(brands []string, goods []string, n int) []Order {
	var orders []Order
	LuhnNext := NewLunhGenerator()

	for i := 0; i < n; i++ {
		items := []OrderItem{
			{
				Description: GenerateItemName(brands, goods),
				Price:       float64(rand.Intn(1000)),
			},
		}

		order := Order{
			Number:      LuhnNext(),
			Goods:       items,
			WaitAccrual: items[0].Price / 10,
		}
		orders = append(orders, order)
	}
	return orders
}

func NewLunhGenerator() func() int {
	number := 0
	return func() int {
		number++
		for !luhn.Valid(number) {
			number++
		}
		return number
	}
}
