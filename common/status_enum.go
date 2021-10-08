package common

// user.gender 用户性别
const (
	UserGenderUnknown = 0
	UserGenderMale    = 1
	UserGenderFemale  = 2
)

//用户账号状态
const (
	UserForbidden = 0 //禁用
	UserNormal    = 1 //正常
)

//数据状态
const (
	DataStatusDeleted = 0 //删除
	DataStatusNormal  = 1 //正常
)

const (
	SettingTokenExpireTime = 86400 //登录token过期时间
)
