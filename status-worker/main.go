package main

import "github.com/Arturlima/status-worker/configs"

func main() {
	configs.NewProvider().Initialize()
}
