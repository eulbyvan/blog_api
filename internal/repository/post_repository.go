package repository

import (
	"database/sql"
	"fmt"

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
	// add transaction
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`INSERT INTO posts (title, content, status, publish_date) VALUES ($1, $2, $3, $4) RETURNING id`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(post.Title, post.Content, post.Status, post.PublishDate).Scan(&post.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// insert tags and post-tags relationships
	for i, tag := range post.Tags {
		// check if the tag already exists
		var tagID int
		err = tx.QueryRow(`SELECT id FROM tags WHERE label = $1`, tag.Label).Scan(&tagID)
		if err != nil {
			if err == sql.ErrNoRows {
				// tag doesn't exist, create it
				err = tx.QueryRow(`INSERT INTO tags (label) VALUES ($1) RETURNING id`, tag.Label).Scan(&tagID)
				if err != nil {
					tx.Rollback()
					return nil, err
				}
			} else {
				tx.Rollback()
				return nil, err
			}
		}

		// update the tag ID in the post.Tags slice
		post.Tags[i].ID = tagID

		// insert into post_tags
		_, err = tx.Exec(`INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2)`, post.ID, tagID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) Update(post *entity.Post) (*entity.Post, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return nil, err
	}

	stmt, err := tx.Prepare(`UPDATE posts SET title = $1, content = $2, status = $3, publish_date = $4 WHERE id = $5`)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Title, post.Content, post.Status, post.PublishDate, post.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	// fetch existing tags
	existingTags := make(map[string]int)
	rows, err := tx.Query(`SELECT t.id, t.label FROM tags t JOIN post_tags pt ON t.id = pt.tag_id WHERE pt.post_id = $1`, post.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tag entity.Tag
		err := rows.Scan(&tag.ID, &tag.Label)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
		existingTags[tag.Label] = tag.ID
	}

	// process each tag in the updated post
	for i, tag := range post.Tags {
		var tagID int
		if existingID, exists := existingTags[tag.Label]; exists {
			tagID = existingID
		} else {
			err = tx.QueryRow(`SELECT id FROM tags WHERE label = $1`, tag.Label).Scan(&tagID)
			if err != nil {
				if err == sql.ErrNoRows {
					err = tx.QueryRow(`INSERT INTO tags (label) VALUES ($1) RETURNING id`, tag.Label).Scan(&tagID)
					if err != nil {
						tx.Rollback()
						return nil, err
					}
				} else {
					tx.Rollback()
					return nil, err
				}
			}
		}
		post.Tags[i].ID = tagID

		_, err = tx.Exec(`INSERT INTO post_tags (post_id, tag_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`, post.ID, tagID)
		if err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	// remove tags that are no longer associated with the post
	for label, id := range existingTags {
		found := false
		for _, tag := range post.Tags {
			if tag.Label == label {
				found = true
				break
			}
		}
		if !found {
			_, err = tx.Exec(`DELETE FROM post_tags WHERE post_id = $1 AND tag_id = $2`, post.ID, id)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *postRepository) Delete(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM posts WHERE id = $1`, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *postRepository) GetByID(id int) (*entity.Post, error) {
	// Query to get the post details
	var post entity.Post
	err := r.db.QueryRow(`SELECT id, title, content, status, publish_date FROM posts WHERE id = $1`, id).
		Scan(&post.ID, &post.Title, &post.Content, &post.Status, &post.PublishDate)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("post with id %d not found", id)
		}
		return nil, err
	}

	// Query to get the associated tags
	rows, err := r.db.Query(`SELECT t.id, t.label FROM tags t JOIN post_tags pt ON t.id = pt.tag_id WHERE pt.post_id = $1`, post.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the tags and add them to the post
	for rows.Next() {
		var tag entity.Tag
		if err := rows.Scan(&tag.ID, &tag.Label); err != nil {
			return nil, err
		}
		post.Tags = append(post.Tags, tag)
	}

	// Check for any errors during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return &post, nil
}

func (r *postRepository) GetPaged(page, size int) ([]entity.Post, error) {
	// TODO implement me!
	return nil, nil
}
