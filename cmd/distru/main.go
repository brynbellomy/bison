package main

import (
	"bison/distru"
	"bison/utils"
	"fmt"
)

var (
	extracts = utils.EnvVarOrPanic("DISTRU_SIGNED_COOKIE_THC")
	hemp     = utils.EnvVarOrPanic("DISTRU_SIGNED_COOKIE_HEMP")
)

func main() {
	clientExtracts := distru.Client{
		AuthToken: extracts,
	}

	// clientHemp := distru.Client{
	// 	AuthToken: hemp,
	// }

	ids, err := clientExtracts.CompanyGetAllIDs()
	if err != nil {
		panic(err)
	}

	fmt.Println(ids)

	// err = clientHemp.CompanyDelete(ids)
	// if err != nil {
	// 	panic(err)
	// }
}
