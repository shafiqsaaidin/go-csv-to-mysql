package data

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/shafiqsaaidin/go-csv-to-mysql/database"
	"github.com/shafiqsaaidin/go-csv-to-mysql/model"
)

func LoadCsv() {
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
			UserName:   line[0],
			UserUuid:   line[1],
			ServerName: line[2],
			Protocol:   line[3],
		}

		stmt, err := database.DB.Prepare("UPDATE v2ray SET user_uuid = ?, protocol = ?, server_name = ? WHERE user_name = ?")
		if err != nil {
			log.Println(err)
		}

		res, err := stmt.Exec(csvData.UserUuid, csvData.Protocol, csvData.ServerName, csvData.UserName)
		if err != nil {
			log.Println(err)
		}

		result, err := res.RowsAffected()
		if err != nil {
			log.Println(err)
		}

		fmt.Println(result)

		// defer stmt.Close()

		// for stmt.Next() {
		// 	v2ray := model.V2ray{}
		// 	err := stmt.Scan(v2ray.UserName)
		// 	if err != nil {
		// 		log.Println(csvData.UserName + " not exist")
		// 	}
		// }

		// UPDATE AND INSERT
		// stmt, err := db.Prepare("UPDATE wireguard SET public_key = ?, preshared_key = ?, server_name = ?, ip_addr = ? WHERE user_name = ?")
		// if err != nil {
		// 	panic(err.Error())
		// }

		// res, err := stmt.Exec(csv_user.public_key, csv_user.preshared_key, "wg03", csv_user.ip_addr, csv_user.user_name)
		// if err != nil {
		// 	panic(err.Error())
		// }

		// a, err := res.RowsAffected()
		// if err != nil {
		// 	panic(err.Error())
		// }

		// fmt.Println(a)

	}
}
