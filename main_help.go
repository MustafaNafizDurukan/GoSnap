package main

import (
	"fmt"

	"github.com/mustafanafizdurukan/GoSnap/pkg/constants"
	"github.com/mustafanafizdurukan/GoSnap/pkg/types"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

func welcome() {
	fmt.Println(`
	 ██████╗  ██████╗ ███████╗███╗   ██╗ █████╗ ██████╗ 
	██╔════╝ ██╔═══██╗██╔════╝████╗  ██║██╔══██╗██╔══██╗
	██║  ███╗██║   ██║███████╗██╔██╗ ██║███████║██████╔╝
	██║   ██║██║   ██║╚════██║██║╚██╗██║██╔══██║██╔═══╝ 
	╚██████╔╝╚██████╔╝███████║██║ ╚████║██║  ██║██║     
	 ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═══╝╚═╝  ╚═╝╚═╝     
	         Take Process Snaphshots Easily!
 `)
}

func initialization() {
	// Open the SQLite database file using GORM
	db, err := gorm.Open(sqlite.Open(constants.FileDBProcesses), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create the schema for the `process_info` table
	err = db.AutoMigrate(&types.ProcessInfo{})
	if err != nil {
		fmt.Println(err)
		return
	}
}
