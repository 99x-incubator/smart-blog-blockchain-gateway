package transactions

import (
	"fmt"
	"net/http"

	// "github.com/stellar/go/build"
	"github.com/stellar/go/clients/horizon"

	"github.com/Blogchain-Gateway/model"

	"github.com/Blogchain-Gateway/api/apiModel"
)

type ConcreteSubmitXDR struct {
	InsertXDR apiModel.SubmitXDR
}

func (cd *ConcreteSubmitXDR) SubmitXDR() model.SubmitXDRResponse {

	var response model.SubmitXDRResponse


	// And finally, send it off to Stellar!
	resp, err := horizon.DefaultTestNetClient.SubmitTransaction(cd.InsertXDR.XDR)
	if err != nil {
		// panic(err)
		response.Error.Code = http.StatusNotFound
		response.Error.Message = "Test net client crashed"
		return response
	}

	fmt.Println("Successful Transaction:")
	fmt.Println("Ledger:", resp.Ledger)
	fmt.Println("Hash:", resp.Hash)

	response.Error.Code = http.StatusOK
	response.Error.Message = "Transaction performed in the blockchain."
	response.TXNID = resp.Hash

	return response

}
