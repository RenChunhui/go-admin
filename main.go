package main

import (
	"fmt"

	"github.com/renchunhui/go-admin/pkg/config"
)

func init() {
	err := config.NewConfig()

	if err != nil {
		return
	}

	fmt.Printf(":%d", config.Http.Port)
}

func main() {

}
