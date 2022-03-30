package db

import (
	"database/sql"
	"log"
	"math/rand"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func createTables(db *sql.DB) {
	log.Println("Creating tables...")

	statement, _ := db.Prepare(tblIDB_Funds())
	statement.Exec()

	statement, _ = db.Prepare(tblIDB_Investments())
	statement.Exec()

	statement, _ = db.Prepare(tblIDB_Investors())
	statement.Exec()

	statement, _ = db.Prepare(tblIDB_Investments_CF())
	statement.Exec()

	statement, _ = db.Prepare(tblIDB_Investments_CF_IDs())
	statement.Exec()

	statement, _ = db.Prepare(tblusers())
	statement.Exec()

	log.Println("Adding table rows...")
	tblIDB_Funds_Rows(db)
	log.Println("tblIDB_Funds_Rows")
	tblIDB_Investments_Rows(db)
	log.Println("tblIDB_Investments_Rows")
	tblIDB_Investments_CF_Rows(db)
	log.Println("tblIDB_Investments_CF_Rows")
	tblIDB_Investments_Rows(db)
	log.Println("tblIDB_Investments_Rows")
	tblIDB_Investors_Rows(db)
	log.Println("tblIDB_Investors_Rows")
	tblIDB_Investments_CF_IDs_Rows(db)
	log.Println("tblIDB_Investments_CF_IDs_Rows")
	tblusers_Rows(db)
	log.Println("tblusers_Rows")

}

func tblIDB_Funds() (table string) {
	log.Println("tblIDB_Funds")
	table = `
		CREATE TABLE tblIDB_Funds (
		"Fund_ID" TEXT,
		"Feeder" TEXT,
		"Fund_Type" TEXT
	  );`
	return
}

func tblIDB_Funds_Rows(db *sql.DB) {
	insert := `INSERT INTO tblIDB_Funds(
        Fund_ID, Feeder, Fund_Type) 
        VALUES (?, ?, ?)`
	statement, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec("Fund1", "NQ", "Normal")
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec("Fund2", "Q", "Normal")
	if err != nil {
		log.Fatalln(err.Error())
	}

}

func tblIDB_Investments() (table string) {
	log.Println("tblIDB_Investments")
	table = `
		CREATE TABLE tblIDB_Investments (
		"InvID" INTEGER,		
        "VID" INTEGER,
        "CID" INTEGER,
        "SID" INTEGER,
        "Inv_Type" TEXT,
        "Feeder" TEXT,
        "Gross_Capital" NUMERIC,
		"Net_Capital" NUMERIC,
		"Actual_Capital" NUMERIC,
        "Date_Inv" DATE,
        "Date_Eliminate" DATE
	  );`
	return
}

func tblIDB_Investments_Rows(db *sql.DB) {
	insert := `INSERT INTO tblIDB_Investments(
        InvID, VID, CID, SID, Inv_Type, Feeder, Gross_Capital, Net_Capital, Actual_Capital, Date_Inv, Date_Eliminate) 
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	now := time.Now()

	for i := 0; i < 10; i++ {
		amount := (rand.Intn(10-1) + 1) * 10000
		Inv_Type := "NQ-A"
		Feeder := "NQ"
		if i > 5 {
			Inv_Type = "Q-A"
			Feeder = "Q"
		}
		_, err = statement.Exec(i, i, i, i, Inv_Type, Feeder, amount, amount, amount, now.Format("2006-01-02"), now.Format("2006-01-02"))
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func tblIDB_Investors() (table string) {
	log.Println("tblIDB_Investors")
	table = `
		CREATE TABLE tblIDB_Investors (
		"SID" INTEGER,		
        "VID" INTEGER,
        "Account_Name" TEXT
	  );`
	return
}

func tblIDB_Investors_Rows(db *sql.DB) {
	insert := `INSERT INTO tblIDB_Investors(
        SID, VID, Account_Name) 
        VALUES (?, ?, ?)`
	statement, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	for i := 0; i < 10; i++ {
		_, err = statement.Exec(i, i, "Investor"+strconv.Itoa(i))
		if err != nil {
			log.Fatalln(err.Error())
		}
	}
}

func tblIDB_Investments_CF_IDs() (table string) {
	log.Println("tblIDB_Investments_CF_IDs")
	table = `
		CREATE TABLE tblIDB_Investments_CF_IDs (
        "ID" INTEGER,		
        "CFID" TEXT,
        "Code_Name" TEXT,
        "Code_Type" TEXT
	  );`
	return
}

func tblIDB_Investments_CF_IDs_Rows(db *sql.DB) {
	insert := `INSERT INTO tblIDB_Investments_CF_IDs(ID, CFID, Code_Name, Code_Type) VALUES (?, ?, ?, ?)`
	statement, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(1, "G", "Gross Distribution", "Distro")
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(1, "S", "Special Distribution", "Distro")
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(1, "Y", "Composite Tax Distro", "Distro")
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(1, "Q", "Tax Distribution", "Distro")
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(1, "M", "Management Fee", "Expense")
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(1, "V", "Servicing Fee", "Expense")
	if err != nil {
		log.Fatalln(err.Error())
	}

}

func tblusers() (table string) {
	log.Println("tblusers")
	table = `
		CREATE TABLE tblusers (
        "ids" INTEGER,		
        "username" TEXT,
        "password" TEXT
	  );`
	return
}

func tblusers_Rows(db *sql.DB) {
	insert := `INSERT INTO tblusers(ids, username, password) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_, err = statement.Exec(1, "CBetori", "password")
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func tblIDB_Investments_CF() (table string) {
	log.Println("tblIDB_Investments_CF")
	table = `
		CREATE TABLE tblIDB_Investments_CF (
        "ID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"InvID" INTEGER,		
        "CID" INTEGER,
        "Scenario" TEXT,
        "CFID" TEXT,
        "CF_Amount" NUMBER,
        "CF_Date" TEXT,
        "Time_Stamp" TEXT
	  );`
	return
}

func tblIDB_Investments_CF_Rows(db *sql.DB) {
	insert := `INSERT INTO tblIDB_Investments_CF(
        InvID, CID , Scenario , CFID , CF_Amount , CF_Date , Time_Stamp ) 
        VALUES (?, ?, ?, ?, ?, ?, ?)`
	statement, err := db.Prepare(insert)
	if err != nil {
		log.Fatalln(err.Error())
	}
	now := time.Now()
	for i := 0; i < 10; i++ {
		date := now.AddDate(0, i, 0)
		CFID := []string{"G", "S", "Q", "Y", "M", "V"}
		for x := 0; x < 6; x++ {
			amount := (rand.Intn(10-1) + 1) * 1000
			_, err = statement.Exec(i, i, "Actual", CFID[x], amount, date.Format("2006-01-02"), date.Format("2006-01-02"))
			if err != nil {
				log.Fatalln(err.Error())
			}
		}
	}
}
