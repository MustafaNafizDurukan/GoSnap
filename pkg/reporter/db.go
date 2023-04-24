package reporter

import (
	"fmt"

	"github.com/mustafanafizdurukan/GoSnap/pkg/constants"
	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// saveDB saves taken processes to db file
func saveDB(processes ...*types.ProcessInfo) {
	// Open the SQLite database file using GORM
	db, err := gorm.Open(sqlite.Open(constants.FileDBProcesses), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = db.Create(&processes).Error
	if err != nil {
		return
	}
}
