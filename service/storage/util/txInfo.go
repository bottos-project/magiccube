package util

type TxInfo struct {
	TransactionID string `json:"transaction_id"`
	Transaction   struct {
		RefBlockNum    uint64        `json:"ref_block_num"`
		RefBlockPrefix uint64        `json:"ref_block_prefix"`
		Expiration     string        `json:"expiration"`
		Scope          []string      `json:"scope"`
		Signatures     []interface{} `json:"signatures"`
		Messages       []struct {
			Code          string        `json:"code"`
			Type          string        `json:"type"`
			Authorization []interface{} `json:"authorization"`
			Data          struct {
				UserName  string `json:"user_name"`
				BasicInfo struct {
					Info string `json:"info"`
				} `json:"basic_info"`
			} `json:"data"`
			HexData string `json:"hex_data"`
		} `json:"messages"`
		Output []struct {
			Notify       []interface{} `json:"notify"`
			DeferredTrxs []interface{} `json:"deferred_trxs"`
		} `json:"output"`
	} `json:"transaction"`
}

type TxDBInfo struct {
	TransactionID string `json:"transaction_id"`
	From          string `json:"from"`
	To            string `json:"to"`
	Price         uint64 `json:"price"`
	Type          uint64 `json:"type"`
	Date          string `json:"date"`
	BlockId       uint64 `json:"block_id"`
}
type TransferDBInfo struct {
	TransactionID string `json:"tx_id"`
	From          string `json:"from"`
	To            string `json:"to"`
	Price         uint64 `json:"price"`
	TxTime        string `json:"tx_time"`
	BlockNum      uint64 `json:"block_num"`
}

const InserTxInfoSql string = "INSERT INTO txinfo(TransactionID,Price,Type,From,To,Date,BlockId) values(?,?,?,?,?,?,?)"
