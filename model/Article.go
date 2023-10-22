package model

import (
	"goBlog/utlis/errormsg"
	"gorm.io/gorm"
)

type Article struct {
	Category Category `gorm:"foreignkey:Cid"`
	gorm.Model

	Title   string `gorm:"type:varchar(18);not null" json:"Title"`
	Cid     int    `gorm:"type:int;not null" json:"Cid"` //Cid是为了和Category表进行内连接，查询文章分类。从性能考虑
	Desc    string `gorm:"type:varchar(200)" json:"Desc"`
	Context string `gorm:"type:longtext" json:"Context"`
	Img     string `gorm:"type:varchar(100)" json:"Img"`
}

func GetArticlesInCategory(id int, pageSize int, pageNum int) ([]Article, int) {
	var cateArticles []Article
	err := db.Preload("Category").Limit(pageSize).Offset((pageNum-1)*pageSize).Where("cid =?", id).Find(&cateArticles).Error
	if err != nil {
		return nil, errormsg.ERROR_CATEGORY_NOT_EXIST
	}
	return cateArticles, errormsg.SUCCESS

}
func GetArticle(id int) (Article, int) {
	var artic Article
	err := db.Preload("Category").Where("id = ?", id).First(&artic).Error
	if err != nil {
		return artic, errormsg.ERROR_ARTICLE_NOT_EXIST
	}
	return artic, errormsg.SUCCESS
}

// Find ArticleList
func GetArticleList(pageSize int, pageNum int) ([]Article, int) {
	var articleList []Article
	// 不手动设置分页的话，使用Postman和grom是返回为nil的，但是某些情况下也许是有默认值的，建议手动设置分页参数
	err = db.Preload("Category").Limit(pageSize).Offset((pageNum - 1) * pageSize).Find(&articleList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, errormsg.ERROR
	}
	return articleList, errormsg.SUCCESS
}
func CheckArticle(title string) (code int) {
	var artic Article
	db.Select("id").Where("title = ?", title).First(&artic)
	if artic.ID > 0 {
		return errormsg.ERROR_CATEGORY_USED
	}
	return errormsg.SUCCESS
}
func CreateArticle(data *Article) int {
	err := db.Create(&data).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}

func UpdateArticle(id int, data *Article) int {
	//gorm中用struct更新信息，不会更新值为0的参数
	var artic Article
	var maps = make(map[string]interface{})
	maps["Title"] = data.Title
	maps["Cid"] = data.Cid
	maps["Desc"] = data.Desc
	maps["Context"] = data.Context
	maps["Img"] = data.Img

	err := db.Model(&artic).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errormsg.ERROR
	}
	return errormsg.SUCCESS
}
func DeleteArticle(id int) int {
	var artic Article

	result := db.Where("id = ?", id).Delete(&artic)
	err := result.Error
	if err != nil {
		return errormsg.ERROR
	} else {
		if result.RowsAffected == 0 {
			return errormsg.ERROR_CATEGORY_NOT_EXIST
		}
		return errormsg.SUCCESS
	}
}
