package service

import (
	"blog-service/internal/model"
	"blog-service/pkg/app"
)

type CountArticleRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	State uint8 `form:"state,default=1" binding:"oneof=0 1"`
}

type GetArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type CreateArticleRequest struct {
	Title         string `form:"title" binding:"required,min=3,max=100"`
	Desc          string `form:"desc" binding:"min=3,max=100"`
	Content       string `form:"content" binding:"required,min=1"`
	CoverImageUrl string `form:"cover_image_url" binding:"min=3,max=100"`
	CreatedBy     string `form:"created_by" binding:"required,min=3,max=100"`
	State         uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type UpdateArticleRequest struct {
	ID            uint32 `form:"id" binding:"required,gte=1"`
	Title         string `form:"title" binding:"min=3,max=100"`
	Desc          string `form:"desc" binding:"min=3,max=100"`
	Content       string `form:"content" binding:"min=1"`
	CoverImageUrl string `form:"cover_image_url" binding:"omitempty,min=3,max=100"`
	State         uint8  `form:"state" binding:"oneof=0 1"`
	ModifiedBy    string `form:"modified_by" binding:"required,min=2,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func (svc *Service) CountArticle(param *CountArticleRequest) (int, error) {
	return svc.dao.CountArticle(param.State)
}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*model.Article, error) {
	return svc.dao.GetArticleList(param.State, pager.Page, pager.PageSize)
}

func (svc *Service) GetArticle(param *GetArticleRequest) (*model.Article, error) {
	return svc.dao.GetArticle(param.ID)
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		Model:         &model.Model{CreatedBy: param.CreatedBy},
		State:         param.State,
	}
	return svc.dao.CreateArticle(&article)
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	article := model.Article{
		Title:         param.Title,
		Desc:          param.Desc,
		Content:       param.Content,
		CoverImageUrl: param.CoverImageUrl,
		Model:         &model.Model{ID: param.ID, CreatedBy: param.ModifiedBy},
		State:         param.State,
	}
	return svc.dao.UpdateArticle(&article)
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	return svc.dao.DeleteArticle(param.ID)
}
