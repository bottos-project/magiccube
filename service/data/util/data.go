package util

type DataDBInfo struct {
	Guid           string     `json:"guid"`
	MerkleRootHash string     `json:"merkle_root_hash"`
	Username       string     `json:"username"`
	FileName       string     `json:"file_name"`
	FileSize       uint64     `json:"file_size"`
	FileNumber     uint64     `json:"file_number"`
	FilePolicy     string     `json:"file_policy"`
	Fslice         [][]string `json:"fslice"`
}

const InsertDataSql string = "insert into datainfo(Guid,MerkleRootHash,Username,FileName,FileSize,FileNumber,FilePolicy,Fslice) values(?,?,?,?,?,?,?,?)"
