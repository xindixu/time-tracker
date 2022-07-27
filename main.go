/*
Copyright © 2022 Xindi Xu <xindixu0@gmail.com>
*/
package main

import (
	"github.com/xindixu/todo-time-tracker/cmd"
	"github.com/xindixu/todo-time-tracker/db"
)

func main() {
	db.InitDB()
	cmd.Execute()
}
