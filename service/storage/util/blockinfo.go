package util
type Info struct {
	ServerVersion            string `json:"server_version"`
	HeadBlockNum             uint64    `json:"head_block_num"`
	LastIrreversibleBlockNum uint64    `json:"last_irreversible_block_num"`
	HeadBlockID              string `json:"head_block_id"`
	HeadBlockTime            string `json:"head_block_time"`
	HeadBlockProducer        string `json:"head_block_producer"`
	RecentSlots              string `json:"recent_slots"`
	ParticipationRate        string `json:"participation_rate"`
}
type Block struct {
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	TransactionMerkleRoot string        `json:"transaction_merkle_root"`
	Producer              string        `json:"producer"`
	ProducerChanges       []interface{} `json:"producer_changes"`
	ProducerSignature     string        `json:"producer_signature"`
	Cycles                []interface{} `json:"cycles"`
	ID                    string        `json:"id"`
	BlockNum              uint64           `json:"block_num"`
	RefBlockPrefix        uint64           `json:"ref_block_prefix"`
}