package main

import (
	itl "github.com/sagoo-cloud/sagooiot/tools/migration/internal"
)

func main() {
	td := itl.NewTdDestination()

	if err := td.CreateStables(); err != nil {
		panic(err)
	}

	if err := td.CreateTables(); err != nil {
		panic(err)
	}

	if err := td.InsertData(); err != nil {
		panic(err)
	}
}
