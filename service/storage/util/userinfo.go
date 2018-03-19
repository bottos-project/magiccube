package util

type UserInfo struct {
	Username     string `db:"username"`
	Accountname  string `db:"accountname"`
	Ownerpubkey  string `db:"ownerpubkey"`
	Activepubkey string `db:"activepubkey"`
	User_type    string `db:"user_type"`
	Role_type    string `db:"role_type"`
	Info         string `db:"info"`
}

type UserDBInfo struct {
	Username       string `db:"username"`
	Accountname    string `db:"accountname"`
	Ownerpubkey    string `db:"owner_pub_key"`
	Activepubkey   string `db:"active_pub_key"`
	EncyptedInfo   string `db:"encypted_info"`
	UserType       string `db:"user_type"`
	RoleType       string `db:"role_type"`
	CompanyName    string `db:"company_name"`
	CompanyAddress string `db:"company_address"`
}

type TokenDBInfo struct {
	Username   string `db:"username"`
	Token      string `db:"token"`
	InsertTime int64  `db:"insert_time"`
}

const InserUserInfoSql string = "insert into userinfo(Username, Accountname, Ownerpubkey,Activepubkey,EncyptedInfo,UserType,RoleType,CompanyName,CompanyAddress) values(?,?,?,?,?,?,?,?,?)"
const QueryUserInfoSql string = "select * from userinfo"

const InsertUserTokenSql string = "insert into tokeninfo(Username, Token,InsertTime) values(?,?,?)"

const DefualtAgingTime int64 = 30 * 60
