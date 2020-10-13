package models

func GetPublishedArticles() (articles []Article, err error) {
	//err = Db.Where("is_publish = ?", true).Joins("User").Find(&articles).Error
	err = Db.Preload("User").Where("is_publish = ?", true).Find(&articles).Error
	//err = Db.Model(&articles).Where("is_publish = ?", true).Association("User").Find(&user).Error
	return articles, err
}

func GetUserArticles(userId uint) (articles []Article, err error) {
	err = Db.Where("user_id = ?", userId).Find(&articles).Error
	return articles, err
}

func GetArticleById(articleID string) (article Article, err error) {
	err = Db.Preload("User").First(&article, articleID).Error
	return article, err
}

func CreateArticle(article *Article) error {
	result := Db.Create(&article)
	return result.Error
}

func EditArticle(article *Article) error {
	result := Db.Save(&article)
	return result.Error
}

func DeleteArticle(article *Article) error {
	result := Db.Delete(&article)
	return result.Error
}
