package main

import (
	"github.com/shafiqsaaidin/go-csv-to-mysql/data"
	"github.com/shafiqsaaidin/go-csv-to-mysql/database"
)

func main() {
	// open db connection
	database.Connect()
	data.LoadV2rayCsv()
}
