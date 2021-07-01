package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/shafiqsaaidin/go-csv-to-mysql/database"
	"github.com/shafiqsaaidin/go-csv-to-mysql/model"
)

func LoadWgCsv() {
	file, err := os.Open("data/wg.csv")
	if err != nil {
		fmt.Println(err)
	}

	reader, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range reader {
		csvData := model.Wireguard{
			UserName:   line[0],
			ServerName: line[1],
		}

		stmt, err := database.DB.Prepare("UPDATE wireguard SET server_name = ? WHERE user_name = ?")
		if err != nil {
			log.Println(err)
		}

		res, err := stmt.Exec(csvData.ServerName, csvData.UserName)
		if err != nil {
			log.Println(err)
		}

		result, err := res.RowsAffected()
		if err != nil {
			log.Println(err)
		}

		fmt.Println(result)

		defer stmt.Close()

	}
}
