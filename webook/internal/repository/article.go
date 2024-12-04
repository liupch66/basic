package repository

import (
	"context"

	"basic-go/webook/internal/domain"
	"basic-go/webook/internal/repository/dao"
)

type ArticleRepository interface {
	Create(ctx context.Context, art domain.Article) (int64, error)
	Update(ctx context.Context, art domain.Article) error
}

type CachedArticleRepository struct {
	dao dao.ArticleDAO
}

func NewCachedArticleRepository(dao dao.ArticleDAO) ArticleRepository {
	return &CachedArticleRepository{dao: dao}
}

func (repo *CachedArticleRepository) entityToDomain(ae dao.Article) domain.Article {
	return domain.Article{
		Id:      ae.Id,
		Title:   ae.Title,
		Content: ae.Content,
		Author:  domain.Author{Id: ae.AuthorId},
	}
}

func (repo *CachedArticleRepository) domainToEntity(a domain.Article) (ae dao.Article) {
	return dao.Article{
		Id:       a.Id,
		Title:    a.Title,
		Content:  a.Content,
		AuthorId: a.Author.Id,
	}
}

func (repo *CachedArticleRepository) Create(ctx context.Context, art domain.Article) (int64, error) {
	return repo.dao.Create(ctx, repo.domainToEntity(art))
}

func (repo *CachedArticleRepository) Update(ctx context.Context, art domain.Article) error {
	return repo.dao.Update(ctx, repo.domainToEntity(art))
}
