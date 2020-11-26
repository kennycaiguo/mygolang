package main

import (
	"database/sql"
	_ "github.com/mattn/go-oci8"
	"log"
)

func main() {
	db, err := sql.Open("oci8", "scott/tiger@ORCL")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT DEPTNO, DNAME, LOC FROM DEPT")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var deptno int64
		var dname string
		var loc string
		rows.Scan(&deptno, &dname, &loc)
		log.Printf("DEPTNO = %v, DNAME = %v, LOC = %v \n", deptno, dname, loc)
	}
	rows.Close()
}