package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/shafiqsaaidin/go-csv-to-mysql/database"
	"github.com/shafiqsaaidin/go-csv-to-mysql/model"
)

func LoadV2rayCsv() {
	file, err := os.Open("data/v2ray.csv")
	if err != nil {
		fmt.Println(err)
	}

	reader, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range reader {
		csvData := model.V2ray{
			UserName:       line[0],
			UserUuid:       line[1],
			ServerName:     line[2],
			ServerProtocol: line[3],
		}

		stmt, err := database.DB.Prepare("UPDATE v2ray SET user_uuid = ?, server_name = ?, server_protocol = ? WHERE user_name = ?")
		if err != nil {
			log.Println(err)
		}

		res, err := stmt.Exec(csvData.UserUuid, csvData.ServerName, csvData.ServerProtocol, csvData.UserName)
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
