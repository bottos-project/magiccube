package util

type FileDBInfo struct {
	FileHash          string `json:"file_hash"`
	Username          string `json:"username"`
	FileName          string `json:"file_name"`
	FileSize          uint64 `json:"file_size"`
	FileNumber        uint64 `json:"file_number"`
	FilePolicy        string `json:"file_policy"`
	AuthorizedStorage string `json:"authorized_storage"`
}

const InsertUserFileSql string = "insert into fileinfo(FileHash,Username,FileName,FileSize,FileNumber,FilePolicy,AuthorizedStorage) values(?,?,?,?,?,?,?)"
