package dao

import (
	"gorm.io/gorm"

	"basic-go/webook/internal/repository/dao/article"
)

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &article.Article{}, &article.PublishedArticle{}, &Interact{},
		&UserLikeBiz{}, &Collection{}, &UserCollectionBiz{}, &AsyncSms{}, &CronJob{})
}
