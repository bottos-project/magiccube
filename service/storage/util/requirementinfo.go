package util

type RequirementDBInfo struct {
	RequirementId   string `bson:"requirement_id" json:"requirement_id"`
	Username        string `bson:"username" json:"username"`
	RequirementName string `bson:"requirement_name" json:"requirement_name"`
	FeatureTag      uint64 `bson:"feature_tag" json:"feature_tag"`
	SamplePath      string `bson:"sample_path" json:"sample_path"`
	SampleHash      string `bson:"sample_hash" json:"sample_hash"`
	ExpireTime      uint32 `bson:"expire_time" json:"expire_time"`
	Price           uint64 `bson:"price" json:"price"`
	Description     string `bson:"description" json:"description"`
	PublishDate     uint32 `bson:"publish_date" json:"publish_date"`
}

const InsertUserRequireSql string = "insert into fileinfo(RequirementId,Username,RequirementName,FeatureTag,SamplePath,SampleHash,ExpireTime,Price,Description,PublishDate) values(?,?,?,?,?,?,?,?,?,?)"
