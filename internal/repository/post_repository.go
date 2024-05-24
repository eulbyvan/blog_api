package repository

import (
	"database/sql"

	"github.com/eulbyvan/blog_api/internal/entity"
)

// contract
type PostRepository interface {
	Create(post *entity.Post) (*entity.Post, error)
	Update(post *entity.Post) (*entity.Post, error)
	Delete(id int) error
	GetByID(id int) (*entity.Post, error)
	GetPaged(page, size int) ([]entity.Post, error)
}

// props
type postRepository struct {
	db *sql.DB // Database connection
}

// constructor
func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *entity.Post) (*entity.Post, error) {
	// prepared statement
	stmt, err := r.db.Prepare(`INSERT INTO posts (title, content, status, publish_date) VALUES ($1, $2, $3, $4) RETURNING id`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Title, post.Content, post.Status, post.PublishDate).Scan(&post.ID)
	if err != nil {
		return nil, err
	}

	return post, err
}

func (r *postRepository) Update(post *entity.Post) (*entity.Post, error) {
	// TODO implement me!
	return nil, nil
}

func (r *postRepository) Delete(id int) error {
	// TODO implement me!
	return nil
}

func (r *postRepository) GetByID(id int) (*entity.Post, error) {
	// TODO implement me!
	return nil, nil
}

func (r *postRepository) GetPaged(page, size int) ([]entity.Post, error) {
	// TODO implement me!
	return nil, nil
}
