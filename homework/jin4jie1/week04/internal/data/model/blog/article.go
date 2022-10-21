package blog

const TABLE_NAME_ARTICLE = "article"

type ArticleModel struct {
  Id      int64  `gorm:"column:id"`
  Content string `gorm:"column:content"`
  UserId  int64  `gorm:"column:user_id"`
}

func (ArticleModel) TableName() string {
  return TABLE_NAME_ARTICLE
}
