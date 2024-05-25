package usecase

import (
	"github.com/eulbyvan/blog_api/internal/entity"
	"github.com/eulbyvan/blog_api/internal/repository"
)

type TagUseCase interface {
	CreateTag(tag *entity.Tag) (*entity.Tag, error)
	UpdateTag(tag *entity.Tag) (*entity.Tag, error)
	DeleteTag(id int) error
	GetTagByID(id int) (*entity.Tag, error)
	GetAllTags() ([]entity.Tag, error)
}

type tagUseCase struct {
	tagRepo repository.TagRepository
}

// constructor
func NewTagUseCase(tr repository.TagRepository) TagUseCase {
	return &tagUseCase{tagRepo: tr}
}

func (uc *tagUseCase) CreateTag(tag *entity.Tag) (*entity.Tag, error) {
	return uc.tagRepo.Create(tag)
}

func (uc *tagUseCase) UpdateTag(tag *entity.Tag) (*entity.Tag, error) {
	return uc.tagRepo.Update(tag)
}

func (uc *tagUseCase) DeleteTag(id int) error {
	return uc.tagRepo.Delete(id)
}

func (uc *tagUseCase) GetTagByID(id int) (*entity.Tag, error) {
	return uc.tagRepo.GetByID(id)
}

func (uc *tagUseCase) GetAllTags() ([]entity.Tag, error) {
	return uc.tagRepo.GetAll()
}
