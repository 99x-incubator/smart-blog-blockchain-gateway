package transactions

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Blogchain-Gateway/model"

	"github.com/stellar/go/keypair"
)

func CreateAccount() model.CAResponse {
	var result model.CAResponse

	pair, err := keypair.Random()
	if err != nil {
		log.Fatal(err)
		result.Error.Message = "Account generation failed at creating seed pair!"
		result.Error.Code = 400

		return result
	}

	address := pair.Address()
	resp, err := http.Get("https://friendbot.stellar.org/?addr=" + address)
	if err != nil {
		log.Fatal(err)
		result.Error.Message = "Account generation failed at submitting account!"
		result.Error.Code = 400

		return result
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		result.Error.Message = "Account generation failed!"
		result.Error.Code = 400

		return result
	}
	fmt.Println(string(body))

	result.Error.Message = "Account successfully generated."
	result.Error.Code = 200

	result.PrivateKey = pair.Seed()
	result.PublicKey = pair.Address()
	log.Println(pair.Seed())
	// SAV76USXIJOBMEQXPANUOQM6F5LIOTLPDIDVRJBFFE2MDJXG24TAPUU7
	log.Println(pair.Address())
	// GCFXHS4GXL6BVUCXBWXGTITROWLVYXQKQLF4YH5O5JT3YZXCYPAFBJZB
	return result
}
