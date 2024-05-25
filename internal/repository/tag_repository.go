package repository

import (
	"database/sql"
	"fmt"

	"github.com/eulbyvan/blog_api/internal/entity"
)

type TagRepository interface {
	Create(tag *entity.Tag) (*entity.Tag, error)
	Update(tag *entity.Tag) (*entity.Tag, error)
	Delete(id int) error
	GetByID(id int) (*entity.Tag, error)
	GetAll() ([]entity.Tag, error)
}

type tagRepository struct {
	db *sql.DB
}

func NewTagRepository(db *sql.DB) TagRepository {
	return &tagRepository{db: db}
}

func (r *tagRepository) Create(tag *entity.Tag) (*entity.Tag, error) {
	err := r.db.QueryRow(`INSERT INTO tags (label) VALUES ($1) RETURNING id`, tag.Label).Scan(&tag.ID)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (r *tagRepository) Update(tag *entity.Tag) (*entity.Tag, error) {
	_, err := r.db.Exec(`UPDATE tags SET label = $1 WHERE id = $2`, tag.Label, tag.ID)
	if err != nil {
		return nil, err
	}
	return tag, nil
}

func (r *tagRepository) Delete(id int) error {
	_, err := r.db.Exec(`DELETE FROM tags WHERE id = $1`, id)
	return err
}

func (r *tagRepository) GetByID(id int) (*entity.Tag, error) {
	var tag entity.Tag
	err := r.db.QueryRow(`SELECT id, label FROM tags WHERE id = $1`, id).Scan(&tag.ID, &tag.Label)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("tag with id %d not found", id)
		}
		return nil, err
	}
	return &tag, nil
}

func (r *tagRepository) GetAll() ([]entity.Tag, error) {
	rows, err := r.db.Query(`SELECT id, label FROM tags`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []entity.Tag
	for rows.Next() {
		var tag entity.Tag
		if err := rows.Scan(&tag.ID, &tag.Label); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}
