package orm

type S6APPUserModel struct {
	Id   int
	Name string
	Age  int8
	Sex  int8
}

func (p7this S6APPUserModel) F8TableName() string {
	return "app_user"
}

type S6APPUserModelV2 struct {
	Id  int  `orm:"column_name=user_id"`
	Age int8 `orm:"column_name=user_age"`
}

func (p7this S6APPUserModelV2) F8TableName() string {
	return "app_user_v2"
}

type S6APPUserInfoModel struct {
	UserId int
	Info   string
}

func (p7this S6APPUserInfoModel) F8TableName() string {
	return "app_user_info"
}

type S6APPUserOrder struct {
	UserId  int
	OrderId int
}

func (p7this S6APPUserOrder) F8TableName() string {
	return "app_user_order"
}
