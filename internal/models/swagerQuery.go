package models

import "time"

// структуры для получения из запросов или для свагера.

type ClientAndHistory struct {
	Client  Client       `json:"client"`
	History HistoryOrder `json:"history"`
}

type ClientAndHistorySwag struct {
	Client  ClientSwag       `json:"client" example:ClientSwag`
	History HistoryOrderSwag `json:"history" example:HistoryOrderSwag`
}

type ArrayDepthOrder struct {
	Array []DepthOrder
}

type ArrayDepthOrderSwag struct {
	Array []DepthOrderSwag `json:"array" example:DepthOrderSwag`
}

type DepthOrderSwag struct {
	Price   float64 `json:"price" example:"120"`
	BaseQty float64 `json:"baseQty" example:"95"`
}

type ClientSwag struct {
	Client_name   string `json:"client_name" example:"Artem"`
	Exchange_name string `json:"exchange_name" example:""`
	Label         string `json:"label" example:""`
	Pair          string `json:"pair" example:""`
}
type ClientSwag2 struct {
	Client_name   string `json:"client_name" example:"Artem"`
	Exchange_name string `json:"exchange_name" example:""`
	Label         string `json:"label" example:""`
	Pair          string `json:"pair" example:"rub-eur"`
}

type HistoryOrderSwag struct {
	Client_name           string    `json:"client_name" example:"Artem"`
	Exchange_name         string    `json:"exchange_name" example:""`
	Label                 string    `json:"label" example:""`
	Pair                  string    `json:"pair" example:"rub-eur"`
	Side                  string    `json:"side" example:""`
	Type                  string    `json:"type" example:""`
	Base_qty              float64   `json:"base_qty" example:"10"`
	Price                 float64   `json:"price" example:"112.7"`
	Algorithm_name_placed string    `json:"algorithm_name_placed" example:""`
	Lowest_sell_prc       float64   `json:"lowest_sell_prc" example:"100"`
	Highest_buy_prc       float64   `json:"highest_buy_prc" example:"150"`
	Commission_quote_qty  float64   `json:"commission_quote_qty" example:"12"`
	Time_placed           time.Time `json:"time_placed" example:"2017-07-21T17:32:28Z"`
}
