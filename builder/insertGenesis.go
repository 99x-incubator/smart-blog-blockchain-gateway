package builder

import (
	"github.com/Blogchain-Gateway/transactions"
	"github.com/Blogchain-Gateway/model"
	"github.com/Blogchain-Gateway/api/apiModel"

)

type GenesisInsertInterface interface {
	InsertGenesis() model.InsertGenesisResponse
}

type AbstractGenesisInsert struct {
	InsertGenesisStruct apiModel.InsertGenesisStruct
	// Identifiers string
	// InsertType  string
	// PreviousTXNID string
}

func (AP *AbstractGenesisInsert) GenesisInsert() model.InsertGenesisResponse {

	object1 := transactions.ConcreteGenesis{InsertGenesisStruct: AP.InsertGenesisStruct}

	result := object1.InsertGenesis()

	return result
}
