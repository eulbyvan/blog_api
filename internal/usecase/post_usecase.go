package usecase

import (
	"github.com/eulbyvan/blog_api/internal/entity"
	"github.com/eulbyvan/blog_api/internal/repository"
)

type PostUseCase interface {
	CreatePost(post *entity.Post) (*entity.Post, error)
	UpdatePost(post *entity.Post) (*entity.Post, error)
	DeletePost(id int) error
	GetPostByID(id int) (*entity.Post, error)
	GetPostsPaged(page, size int) ([]entity.Post, error)
	GetPostsByTag(tag string, page, size int) ([]entity.Post, error)
}

type postUseCase struct {
	postRepo repository.PostRepository
}

// constructor
func NewPostUseCase(pr repository.PostRepository) PostUseCase {
	return &postUseCase{postRepo: pr}
}

func (uc *postUseCase) CreatePost(post *entity.Post) (*entity.Post, error) {
	return uc.postRepo.Create(post)
}

func (uc *postUseCase) UpdatePost(post *entity.Post) (*entity.Post, error) {
	return uc.postRepo.Update(post)
}

func (uc *postUseCase) DeletePost(id int) error {
	return uc.postRepo.Delete(id)
}

func (uc *postUseCase) GetPostByID(id int) (*entity.Post, error) {
	return uc.postRepo.GetByID(id)
}

func (uc *postUseCase) GetPostsPaged(page, size int) ([]entity.Post, error) {
	return uc.postRepo.GetPaged(page, size)
}

func (uc *postUseCase) GetPostsByTag(tag string, page, size int) ([]entity.Post, error) {
	return uc.postRepo.GetByTag(tag, page, size)
}
