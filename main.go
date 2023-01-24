package main

import (
	"fmt"
	route2 "github.com/Nailliw/student_simulator_backend/application/route"
)

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}
	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	fmt.Print(stringJson[0])
	//		INSTALAÇÃO E CONFIGURAÇÃO DO KFKA
}
