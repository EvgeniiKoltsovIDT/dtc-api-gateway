package domain

type Transaction struct {
	TxID      string `json:"TxID"`
	TxType    string `json:"TxType"`
	Timestamp int    `json:"Timestamp"`
}

type TransactionFilter struct {
	ID     string   `json:"ID"`
	UserID string   `json:"UserID"`
	TxType []string `json:"TxType"`
}
