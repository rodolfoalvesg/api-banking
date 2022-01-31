package controllers

import "github.com/rodolfoalvesg/api-banking/api/src/db"

type Controller struct {
	db.Database
}

func NewController(db db.Database) *Controller {
	return &Controller{}
}
