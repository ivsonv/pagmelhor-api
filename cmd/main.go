package main

import (
	"app/configs"
)

func main() {
	cfg, _ := configs.LoadConfig(".")
	println(cfg.DBDriver)
}
