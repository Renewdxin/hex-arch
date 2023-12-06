package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	// connect to db
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("did not connect to the database: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("could not ping the database: %v", err)
		return nil, err
	}

	return &Adapter{db: db}, nil
}

func (da Adapter) CloseDBConnection() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("could not close the database connection: %v", err)
	}
	log.Println("database connection closed")
}

func (da Adapter) AddToHistory(answer int32, operation string) {
	stmt, err := da.db.Prepare("INSERT INTO history (data, answer, opration) VALUES (?,?,?)")
	if err != nil {
		log.Fatalf("could not prepare the statement: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(time.Now(), answer, operation)
	if err != nil {
		return
	}
	log.Println("history record added")

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		// Handle the error
		log.Fatal(err)
	}
	// Print or log the number of rows affected
	fmt.Printf("Rows affected: %d\n", rowsAffected)

	//ataValues := []string{"data1", "data2", "data3"}
	//answerValues := []string{"answer1", "answer2", "answer3"}
	//operationValues := []string{"operation1", "operation2", "operation3"}
	//
	//// Begin a transaction
	//tx, err := da.db.Begin()
	//if err != nil {
	//    // Handle the error
	//    log.Fatal(err)
	//}
	//defer tx.Rollback() // Rollback the transaction if there's an error
	//
	//// Loop through the values and execute the batch insert
	//for i := range dataValues {
	//    _, err := tx.Stmt(stmt).Exec(dataValues[i], answerValues[i], operationValues[i])
	//    if err != nil {
	//        // Handle the error
	//        log.Fatal(err)
	//    }
	//}
	//
	//// Commit the transaction
	//if err := tx.Commit(); err != nil {
	//    // Handle the error
	//    log.Fatal(err)
	//}
}
