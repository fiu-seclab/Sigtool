package main

import (
	"log"

	"database/sql"
	"encoding/hex"

	"sigtool"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID           int    `json:"ID"`
	FileLocation string `json:"FileLocation"`
}

func main() {

	db, err := sql.Open("mysql", "forensicengineui:forensicengineui@tcp(131.94.132.2:3306)/forensicengineui")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	results, err := db.Query("SELECT ID, FileLocation FROM MalshareItems where FileLocation<>'' and PKCS7 is null")
	if err != nil {
		panic(err.Error()) // just exit
	}

	for results.Next() {
		var tag Tag
		// for each row, scan the result into our tag composite object
		err = results.Scan(&tag.ID, &tag.FileLocation)
		if err != nil {
			panic(err.Error()) //exit
		}

		log.Printf(tag.FileLocation)

		buf, err := sigtool.ExtractDigitalSignature(tag.FileLocation)
		if err != nil {
			log.Fatal(err)
		}

		// perform a db.Query insert
		encodedStr := hex.EncodeToString(buf)
		insert, err := db.Query("update MalshareItems set PKCS7='?' where ID=?", encodedStr, tag.ID)

		// if there is an error inserting, handle it
		if err != nil {
			panic(err.Error())
		}

		// be careful deferring Queries if you are using transactions
		insert.Close()

	}
	defer results.Close()

}
