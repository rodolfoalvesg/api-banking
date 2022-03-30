package models

import "time"

type Transfers struct {
	Id                     string    `json:"id"`
	Account_origin_id      string    `json:"account_origin_id"`
	Account_destination_id string    `json:"account_destination_id"`
	Amount                 int       `json:"amount"`
	Created_at             time.Time `json:"created_at"`
}
