/*The functions of the Sqlite database that provided here are not ready yet，they should never be exposed to users。They are 
in the Bottos's service layer,which delivering database service pluggable to provide users with queries.And we plan to support
it in a future point release.At present, we only support mongodb to provide users with queries.*/
package sqlite

import (
	"github.com/code/bottos/service/storage/util"
)

func (r *SqliteRepository) CallGetUserRequirementList(username string) ([]*util.RequirementDBInfo, error) {
	var reqs = []*util.RequirementDBInfo{}
	dbtag := new(util.RequirementDBInfo)
	dbtag.RequirementId = "idtest"
	dbtag.RequirementName="requirename"
	dbtag.FeatureTag=111
	dbtag.SamplePath="pathtest"
	dbtag.SampleHash="hashtest"
	dbtag.ExpireTime=222
	dbtag.Price=333
	dbtag.Description = "destest"
	dbtag.PublishDate =444
	reqs = append(reqs, dbtag)
	return reqs, nil
}
