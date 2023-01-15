package dao

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

// 统计个数
func (d *Dao) CountArticle(state uint8) (int, error) {
	tag := model.Article{State: state}
	return tag.Count(d.engine)
}

// 获取标签列表
func (d *Dao) GetArticleList(state uint8, page, pageSize int) ([]*model.Article, error) {
	article := model.Article{State: state}
	pageOffset := app.GetPageOffset(page, pageSize)
	return article.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) GetArticle(id uint32) (*model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.GetArticle(d.engine)

}

func (d *Dao) CreateArticle(article *model.Article) error {
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(article *model.Article) error {
	values := map[string]interface{}{
		"state":       article.State,
		"modified_by": article.ModifiedBy,
	}
	return article.Update(d.engine, values)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

