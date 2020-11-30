package main

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb" // the underscore indicates the package is used
	"log"
)

func main1() {
	fmt.Println("starting app")

	// the user needs to be setup in SQL Server as an SQL Server user.
	// see create login and the create user SQL commands as well as the
	// SQL Server Management Studio documentation to turn on Hybrid Authentication
	// which allows both Windows Authentication and SQL Server Authentication.
	// also need to grant to the user the proper access permissions.
	// also need to enable TCP protocol in SQL Server Configuration Manager.
	//
	// you could also use Windows Authentication if you specify the fully qualified
	// user id which would specify the domain as well as the user id.
	// for instance you could specify "user id=domain\\user;password=userpw;".

	condb, errdb := sql.Open("mssql", "server=117.34.111.35;user id=zsq;password=sa;encrypt=disable")
	if errdb != nil {
		fmt.Println("  Error open db:", errdb.Error())
	}

	defer condb.Close()

	errdb = condb.Ping()
	if errdb != nil {
		log.Fatal(errdb)
	}

	_, errdb = condb.Exec("use  DongFangCloud")
	if errdb != nil {
		fmt.Println("  Error Exec db: using db - ", errdb.Error())
	}

	// Now that we have our database lets read some records and print them.
	var (
		one int
		two int
	)

	_, errdb = condb.Exec("use  DongFangCloud")
	// documentation about a simple query and results loop is at URL
	// http://go-database-sql.org/retrieving.html
	// we use Query() and not Exec() as we expect zero or more rows to
	// be returned. only use Query() if rows may be returned.
	fmt.Println("  Query our table for the three rows we inserted.")
	rows, errdb := condb.Query("use  DongFangCloud; SELECT * FROM t_Order ")
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one, &two)
		if err != nil {
			fmt.Println("  Error Query db: select - ", err.Error())
		} else {
			fmt.Printf("    - one %d and two %d\n", one, two)
		}
	}
	rows.Close()

	errdb = rows.Err()
	if errdb != nil {
		fmt.Println("  Error Query db: processing rows - ", errdb.Error())
	}

	fmt.Println("ending app")
}
