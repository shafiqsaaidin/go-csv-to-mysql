package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Wireguard struct {
	// user_id   int
	user_name     string
	public_key    string
	preshared_key string
	ip_addr       string
}

func main() {
	db, err := sql.Open("mysql", "root:@/vpnje")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	file, err := os.Open("wg03.csv")
	if err != nil {
		fmt.Println(err)
	}

	reader, err := csv.NewReader(file).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range reader {
		csv_user := Wireguard{
			user_name:     line[0],
			public_key:    line[1],
			preshared_key: line[2],
			ip_addr:       line[3],
		}

		// fmt.Println(user.public_key)
		// var result Wireguard

		// sqlStmt := `SELECT user_name FROM wireguard WHERE user_name = ?`
		// err := db.QueryRow(sqlStmt, csv_user.user_name).Scan(&result.user_name)
		// if err != nil {
		// 	if err != sql.ErrNoRows {
		// 		log.Print(err)
		// 	}

		// 	fmt.Println(csv_user.user_name + " not exist")
		// }

		// UPDATE AND INSERT
		stmt, err := db.Prepare("UPDATE wireguard SET public_key = ?, preshared_key = ?, server_name = ?, ip_addr = ? WHERE user_name = ?")
		if err != nil {
			panic(err.Error())
		}

		res, err := stmt.Exec(csv_user.public_key, csv_user.preshared_key, "wg03", csv_user.ip_addr, csv_user.user_name)
		if err != nil {
			panic(err.Error())
		}

		a, err := res.RowsAffected()
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(a)

	}
	defer db.Close()
}
