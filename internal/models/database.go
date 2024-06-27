package models

import "time"

type HistoryOrder struct {
	Client_name           string
	Exchange_name         string
	Label                 string
	Pair                  string
	Side                  string
	Type                  string
	Base_qty              float64
	Price                 float64
	Algorithm_name_placed string
	Lowest_sell_prc       float64
	Highest_buy_prc       float64
	Commission_quote_qty  float64
	Time_placed           time.Time
}

type Client struct {
	Client_name   string
	Exchange_name string
	Label         string
	Pair          string
}

type OrderBook struct {
	Id       int64
	Exchange string
	Pair     string
	Asks     []DepthOrder
	Bids     []DepthOrder
}

type DepthOrder struct {
	Price   float64
	BaseQty float64
}
