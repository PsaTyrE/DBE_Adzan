package main

import (
	"github.com/PsaTyrE/dbe_adzan/model"
	"github.com/PsaTyrE/dbe_adzan/route"
)

func main() {
	model.Conn()

	route.Init()
}
