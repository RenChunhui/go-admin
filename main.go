package main

import (
	"fmt"

	"github.com/renchunhui/go-admin/pkg/config"
	"github.com/renchunhui/go-admin/pkg/db"
)

func init() {
	err := config.NewConfig()

	if err != nil {
		return
	}

	err = db.Open()

	if err != nil {
		return
	}

	fmt.Printf("database %+v\n", config.Database)
}

func main() {

}
