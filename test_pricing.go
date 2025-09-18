package main

import (
	"analysis-aws-bill/database"
	"analysis-aws-bill/models"
)

func main() {
	database.InitOtherTables()
	models.InsertEc2PriceToDB()
}
