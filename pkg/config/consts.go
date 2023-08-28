package config

const (
	VideoTableName   = "Video"
	UserTableName   = "user"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Notes           = "notes"
	ApiServiceName  = "demoapi"
	NoteServiceName = "demonote"
	UserServiceName = "demouser"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	UserServiceAddr = ":9000"
	NoteServiceAddr = ":10000"
	ExportEndpoint  = ":4317"
	ETCDAddress     = "127.0.0.1:2379"
	DefaultLimit    = 10
	
	ErrCode_SuccessCode = 0
	ErrCode_ServiceErrCode = 10001
	ErrCode_ParamErrCode = 10002
	ErrCode_UserAlreadyExistErrCode = 10003 
	ErrCode_AuthorizationFailedErrCode = 10004
	ErrCode_PasswdHashFailedCode = 10005
	ErrCode_UserDoNotExistErrCode = 10003 


)