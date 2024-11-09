package main

//go:generate swag init -g cmd/server.go

import "main/cmd"

func main() {
    cmd.Execute()
}