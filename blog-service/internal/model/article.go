package model

import "github.com/jinzhu/gorm"

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a *Article) TableName() string {
	return "blog_article"
}

// 统计数量
func (a Article) Count(db *gorm.DB) (int, error) {
	var count int
	db = db.Where("state = ? ", a.State)
	if err := db.Model(&a).Where("is_del = ?", 0).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// 查询列表
func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	db = db.Where("state = ? ", a.State)
	if err = db.Where("is_del = ?", 0).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *Article) GetArticle(db *gorm.DB) (*Article, error) {
	var article *Article = &Article{}
	var err error
	if err = db.Where("id = ? AND is_del = ?", a.ID,0).Take(&article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	if err := db.Model(a).Where("id = ? AND is_del = ?", a.ID, 0).Update(values).Error; err != nil {
		return err
	}
	return nil
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND is_del = ?", a.Model.ID, 0).Delete(&a).Error
}
