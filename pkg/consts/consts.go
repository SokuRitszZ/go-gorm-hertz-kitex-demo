package consts

const (
	MySQLDefaultDSN = "root:123456@tcp(127.0.0.1:3306)/gorm2?charset=utf8&parseTime=True&loc=Local"
	ETCDAddress = "127.0.0.1:2379"
	TCP = "tcp"
	NoteServiceAddr = ":10000"
	NoteServiceName = "note"
	UserServiceAddr = ":9000"
	UserServiceName = "user"
	ApiServiceAddr = ""
	ApiServiceName = "api"
	ExportEndpoint = ":4317"
	DefaultLimit = 10
	SecretKey = "secret key"
	IdentityKey = "id"
)
