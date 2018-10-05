package builder

import (
	"github.com/Blogchain-Gateway/transactions"
	"github.com/Blogchain-Gateway/api/apiModel"
	"github.com/Blogchain-Gateway/model"
)

// type InsertData struct{}
// type TDPInsertInterface interface {
// 	InsertDataHash() apiModel.SubmitXDR
// }

type AbstractSubmitXDR struct {
	SubmitXDR apiModel.SubmitXDR
	// Hash          string
	// InsertType    string
	// PreviousTXNID string
	// ProfileId     string
}

func (AP *AbstractSubmitXDR) XDRInsert() model.SubmitXDRResponse {

	object := transactions.ConcreteSubmitXDR{InsertXDR: AP.SubmitXDR}
	// object := stellarExecuter.ConcreteInsertData{Hash: AP.Hash, InsertType: AP.InsertType, PreviousTXNID: AP.PreviousTXNID, ProfileId: AP.ProfileId}

	result := object.SubmitXDR()

	return result
}
