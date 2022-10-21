package admin

const TABLE_NAME_USER = "user"

type UserModel struct {
  Id   int64  `gorm:"column:id"`
  Name string `gorm:"column:name"`
}

func (UserModel) TableName() string {
  return TABLE_NAME_USER
}
