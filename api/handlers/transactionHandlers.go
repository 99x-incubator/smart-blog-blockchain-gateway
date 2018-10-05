package handlers

import (
	"github.com/Blogchain-Gateway/transactions"
	"github.com/Blogchain-Gateway/builder"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Blogchain-Gateway/model"

	"github.com/Blogchain-Gateway/api/apiModel"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	var response model.CAResponse
	response = transactions.CreateAccount()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(response.Error.Code)
	json.NewEncoder(w).Encode(response)
	return

}

func Transaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	TType := (vars["TType"])
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if r.Body == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Please send a request body")
		return
	} else {
		switch TType {
		// case "0":
		// 	var PObj apiModel.InsertProfileStruct
		// 	err := json.NewDecoder(r.Body).Decode(&PObj)
		// 	if err != nil {
		// 		w.WriteHeader(http.StatusBadRequest)
		// 		json.NewEncoder(w).Encode("Error while Decoding the body")
		// 		fmt.Println(err)
		// 		return
		// 	}
		// 	fmt.Println(PObj)
		// 	response := model.InsertProfileResponse{}

		// 	display := &builder.AbstractProfileInsert{InsertProfileStruct: PObj}
		// 	response = display.ProfileInsert()

		// 	w.WriteHeader(response.Error.Code)
		// 	result := apiModel.ProfileSuccess{
		// 		Message:           response.Error.Message,
		// 		ProfileTxn:        response.ProfileTxn,
		// 		PreviousTXNID:     response.PreviousTXNID,
		// 		PreviousProfileID: response.PreviousProfileID,
		// 		Identifiers:       PObj.Identifier,
		// 		Type:              PObj.Type}
		// 	json.NewEncoder(w).Encode(result)
		// case "1":
		// 	var GObj apiModel.InsertGenesisStruct
		// 	err := json.NewDecoder(r.Body).Decode(&GObj)
		// 	if err != nil {
		// 		w.WriteHeader(http.StatusBadRequest)
		// 		json.NewEncoder(w).Encode("Error while Decoding the body")
		// 		fmt.Println(err)
		// 		return
		// 	}
		// 	fmt.Println(GObj)
		// 	result := model.InsertGenesisResponse{}

		// 	display := &builder.AbstractGenesisInsert{InsertGenesisStruct: GObj}
		// 	result = display.GenesisInsert()

		// 	w.WriteHeader(result.Error.Code)
		// 	result2 := apiModel.GenesisSuccess{
		// 		Message:     result.Error.Message,
		// 		ProfileTxn:  result.ProfileTxn,
		// 		GenesisTxn:  result.GenesisTxn,
		// 		Identifiers: GObj.Identifier,
		// 		Type:        GObj.Type}
		// 	json.NewEncoder(w).Encode(result2)

		case "0":
			var TDP apiModel.SubmitXDR
			err := json.NewDecoder(r.Body).Decode(&TDP)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				json.NewEncoder(w).Encode("Error while Decoding the body")
				fmt.Println(err)
				return
			}
			fmt.Println(TDP)
			response := model.SubmitXDRResponse{}

			// display := &builder.AbstractTDPInsert{Hash: TObj.Data, InsertType: TType, PreviousTXNID: TObj.PreviousTXNID[0], ProfileId: TObj.ProfileID[0]}
			display := &builder.AbstractSubmitXDR{SubmitXDR: TDP}
			response = display.XDRInsert()

			w.WriteHeader(response.Error.Code)
			
			json.NewEncoder(w).Encode(response)

		default:
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("Please send a valid Transaction Type")
			return
		}

	}

	return

}


