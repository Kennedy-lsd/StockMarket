package repos

import (
	"database/sql"
	"log"
	"sort"

	"github.com/Kennedy-lsd/StockMarket/data"
)

type CommentRepository struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) *CommentRepository {
	return &CommentRepository{DB: db}
}

func (r *CommentRepository) FindAll() ([]data.Comment, error) {
	query := `SELECT * FROM comments`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []data.Comment

	for rows.Next() {
		var comment data.Comment

		err := rows.Scan(&comment.Id, &comment.Title, &comment.CreatedAt, &comment.StockId)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	sort.Slice(comments, func(i, j int) bool {
		return *comments[i].Id < *comments[j].Id
	})

	return comments, nil
}

func (r *CommentRepository) Post(comment *data.CommentCreate) error {
	query := `INSERT INTO comments (title, stock_id) 
		VALUES ($1, $2) RETURNING id, created_at`

	err := r.DB.QueryRow(query, &comment.Title, &comment.StockId).Scan(&comment.Id, &comment.CreatedAt)
	if err != nil {
		log.Printf("Error creating comment: %v", err)
		return err
	}

	return nil
}
