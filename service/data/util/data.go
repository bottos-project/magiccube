package util

type DataDBInfo struct {
	Filehash           string     `json:"filehash"`
	Username string     `json:"username"`
	Filename       string     `json:"filename"`
	Filesize       uint64     `json:"filesize"`
	Filepolicy       string     `json:"filepolicy"`
	Filenumber       uint64     `json:"filenumber"`
	Simorass     uint64     `json:"simorass"`
	Optype     uint64     `json:"optype"`
	Storeaddr         string `json:"storeaddr"`
}

const InsertDataSql string = "insert into datainfo(Guid,MerkleRootHash,Username,FileName,FileSize,FileNumber,FilePolicy,Fslice) values(?,?,?,?,?,?,?,?)"
