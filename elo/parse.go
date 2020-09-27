package elo

import "github.com/aleluchesi/PayGo/transferarea"

// Imports within the project

func Parse() {

	Eloparse := new(transferarea.Campos)

	Eloparse.Pan = "1234567890123456"
	Eloparse.MerchantType = 5999
	Eloparse.PointOfService = "051"
	Eloparse.ProcessingCode = "003000"
	Eloparse.TransAmount = "1000"

}
