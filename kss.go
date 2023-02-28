package main

import (
	kss "kss/app"
)

func main() {
	k := kss.Initialize("kss.fci")
	k.Start()
}
