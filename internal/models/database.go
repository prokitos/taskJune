package models

import "time"

type HistoryOrder struct {
	Client_name           string    `json:"client_name" example:""`
	Exchange_name         string    `json:"exchange_name" example:""`
	Label                 string    `json:"label" example:""`
	Pair                  string    `json:"pair" example:""`
	Side                  string    `json:"side" example:""`
	Type                  string    `json:"type" example:""`
	Base_qty              float64   `json:"base_qty" example:""`
	Price                 float64   `json:"price" example:""`
	Algorithm_name_placed string    `json:"algorithm_name_placed" example:""`
	Lowest_sell_prc       float64   `json:"lowest_sell_prc" example:""`
	Highest_buy_prc       float64   `json:"highest_buy_prc" example:""`
	Commission_quote_qty  float64   `json:"commission_quote_qty" example:""`
	Time_placed           time.Time `json:"time_placed" example:""`
}

type Client struct {
	Client_name   string `json:"client_name" example:""`
	Exchange_name string `json:"exchange_name" example:""`
	Label         string `json:"label" example:""`
	Pair          string `json:"pair" example:""`
}

type OrderBook struct {
	Id       int64        `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Exchange string       `json:"exchange" example:""`
	Pair     string       `json:"pair" example:""`
	Asks     []DepthOrder `gorm:"foreignKey:Owner;references:Id" json:"asks" example:""`
	Bids     []DepthOrder `gorm:"foreignKey:Owner;references:Id" json:"bids" example:""`
}
type DepthOrder struct {
	Id      int64   `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Price   float64 `json:"price" example:""`
	BaseQty float64 `json:"baseQty" example:""`
	Owner   int64   `json:"owner" example:""`
}
