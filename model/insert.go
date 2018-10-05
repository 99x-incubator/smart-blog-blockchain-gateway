package model

type InsertDataResponse struct {
	TXNID     string
	ProfileID string
	TxnType   string
	Error     Error
}
type CAResponse struct {
	PrivateKey     string
	PublicKey string
	Error     Error
}


type SubmitXDRResponse struct {
	TXNID string
	Error     Error
}
type InsertGenesisResponse struct {
	ProfileTxn  string
	GenesisTxn  string
	Identifiers string
	TxnType     string
	Error       Error
}
