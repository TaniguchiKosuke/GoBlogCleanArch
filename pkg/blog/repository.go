package blog

import (
	"GoBlogClean/domain"
)

type ArticleRepository interface {
	PostArticle(article *domain.Article) (*domain.Article, error)
	GetArticleByID(articleID int) (*domain.Article, error)
	GetArticles() ([]*domain.Article, error)
	// UpdateArticle(article *Article) (*Article, error)
	// DeleteArticle(article *Article) (*Article, error)
	// SearchArticle(text string) ([]*Article, error)
}
