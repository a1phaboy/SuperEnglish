package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

//将词汇和中文对应插入数据库中
func ImportList(chinese string,english string) int64{
	db, err := sql.Open("sqlite3", "./supEnglish.db")//若数据库没有在这个项目文件下，则需要写绝对路径
	checkErr(err)
	//插入数据
	stmt, err := db.Prepare("INSERT INTO wordlist(chinese, english)  values(?, ?)")
	defer stmt.Close()
	checkErr(err)
	res, err := stmt.Exec(chinese, english)
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	return id
}

//负责分解string，将:和;拆分
func decomposeString(input string){
	arr := strings.Split(input,":")
	//fmt.Println(arr)
	chinese_arr := strings.Split(arr[1],";")
	//fmt.Println(len(chinese_arr))
	for i:=0;i<len(chinese_arr);i++{
		if chinese_arr[i] == ""{
			continue
		}
		_ = ImportList(chinese_arr[i],arr[0])
	}

}

func ImportListProcess(number int){
	var input string
	for i:=0;i<number;i++{
		_,err:=fmt.Scanf("%s",&input)
		if err!=nil {
			fmt.Printf("[x] 输入错误%s\n",err)
			return
		}else{
			decomposeString(input)
			fmt.Printf("[*] 已完成：%d / %d\n",i+1,number)
		}
	}
}

func addWord(){
	var number int
	fmt.Println("输入要录入的词汇数量：")
	fmt.Scanf("%d",&number)
	fmt.Println("输入要录入的词汇,单词和中文用:隔开，一词多义则用;隔开，例如 english:chinese;chinese")
	ImportListProcess(number)
}
