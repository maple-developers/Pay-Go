package controllers

import (
	"context"
	"fmt"

	"github.com/adshao/go-binance/v2"
	"github.com/gofiber/fiber/v2"
)

var (
	apiKey    = "d2yPJveRVdOqbxoWRdlwpCjeIcma543TAtZlkQdLg1ZURgNhSigKCHyxYvwTxxRl"
	secretKey = "wJiawlKoCHm1aIps6uki2ZerncgdBPd8Hm5fcFRJwTnbcTDCIiiq7zuP2qxaOEXB"
)

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"msg"`
}

func HomeStatus(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "GO"})
}
func CreateOrder(c *fiber.Ctx) error {
	//var errorObj []ErrorResponse
	client := binance.NewClient(apiKey, secretKey)
	order, err := client.NewCreateOrderService().Symbol("BTCUSDT").
		Side(binance.SideTypeSell).Type(binance.OrderTypeLimit).
		TimeInForce(binance.TimeInForceTypeGTC).Quantity("5").Price("800").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "created",
		"message": order,
	})
}
func ListTickerPrices(c *fiber.Ctx) error {
	client := binance.NewClient(apiKey, secretKey)
	prices, err := client.NewListPricesService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "created",
		"message": prices,
	})
}
func Account(c *fiber.Ctx) error {
	client := binance.NewClient(apiKey, secretKey)
	res, err := client.NewGetAccountService().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"Time": res,
	})
}
func ListOrders(c *fiber.Ctx) error {
	client := binance.NewClient(apiKey, secretKey)
	orders, err := client.NewListOrdersService().Symbol("BNBETH").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"status":  "created",
		"message": orders,
	})
}

func C2CTradeHistory(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "Done",
		"Message": binanceTest.TestC2CTradeHistory(),
	})
}
