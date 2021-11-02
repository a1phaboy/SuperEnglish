package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func test(){
	count()
	fmt.Println()
}

func count() {
	var count int = 0
	db, err := sql.Open("sqlite3", "./supEnglish.db")
	checkErr(err)
	sql, err := db.Prepare("SELECT count(DISTINCT english) FROM wordlist")
	checkErr(err)
	err = sql.QueryRow().Scan(&count)
	fmt.Printf("[*] 目前已收录了%d个词汇\n",count)
}
