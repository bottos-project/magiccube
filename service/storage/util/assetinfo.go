package util

type AssetDBInfo struct {
	AssetID     string `bson:"asset_id" json:"asset_id"`
	UserName    string `bson:"user_name" json:"user_name"`
	AssetName   string `bson:"asset_name" json:"asset_name"`
	FeatureTag  uint64 `bson:"feature_tag" json:"feature_tag"`
	SamplePath  string `bson:"sample_path" json:"sample_path"`
	SampleHash  string `bson:"sample_hash" json:"sample_hash"`
	StoragePath string `bson:"storage_path" json:"storage_path"`
	StorageHash string `bson:"storage_hash" json:"storage_hash"`
	ExpireTime  uint32 `bson:"expire_time" json:"expire_time"`
	Price       uint64 `bson:"price" json:"price"`
	Description string `bson:"description" json:"description"`
	UploadDate  uint32 `bson:"upload_date" json:"upload_date"`
}
