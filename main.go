package main

import (
	"restapi/route"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	request.HandleRequest()
}
