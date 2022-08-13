package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/nbright/nomadcoin/utils"
)

const (
	dbName       = "blockchain.db"
	dataBucket   = "data"
	blocksBucket = "blocks"
)

var db *bolt.DB

func DB() *bolt.DB {
	if db == nil {
		dbPointer, err := bolt.Open(dbName, 0600, nil)
		db = dbPointer
		utils.HandleErr(err)

		err := db.Update(func(tx *bolt.Tx) error {

		})
		defer db.Close()
	}
	fmt.Println(db)
	return db

}
