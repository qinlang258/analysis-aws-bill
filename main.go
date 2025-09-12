package main

import (
	"analysis-aws-bill/config"
	"analysis-aws-bill/models"
	"context"
)

func initData() {
	models.WriteToMysql(context.Background(), config.CsvName)
}

func main() {
	initData()
}
