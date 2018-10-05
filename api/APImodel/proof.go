package apiModel

type InsertSuccess struct {
	Message   string
	TxNHash   string
	ProfileID string
	Type      string
}

type GenesisSuccess struct {
	Message     string
	ProfileTxn  string
	GenesisTxn  string
	Identifiers string
	Type        string
}

type TransactionStruct struct {
	TType             string   `json:"TType"`
	ProfileID         []string `json:"ProfileID"`
	PreviousTXNID     []string `json:"PreviousTXNID"`
	Data              string   `json:"Data"`
	Identifiers       []string `json:"Identifiers"`
	Identifier        string
	MergingTXNs       []string
	PreviousProfileID []string `json:"PreviousProfileID"`
	Assets            []string
	Code              string
}

type InsertTDP struct {
	Type          string
	PreviousTXNID string
	ProfileID     string
	Identifier    string `json:"Identifier"`
	DataHash      string
}
type SubmitXDR struct {
	XDR string
}



type InsertGenesisStruct struct {
	Type       string `json:"Type"`
	Identifier string `json:"Identifier"`
}
