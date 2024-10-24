package repos

import (
	"database/sql"
	"fmt"
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

func (r *CommentRepository) FindById(id int64) (*data.Comment, error) {
	query := `SELECT * FROM comments WHERE id = $1`
	var comment data.Comment
	err := r.DB.QueryRow(query, id).Scan(&comment.Id, &comment.Title, &comment.CreatedAt, &comment.StockId)
	if err != nil {
		log.Printf("Error : %v", err)
		return nil, err
	}
	return &comment, nil
}

func (r *CommentRepository) DeleteById(id int64) error {
	query := `DELETE FROM comments WHERE id = $1`

	result, err := r.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error executing DELETE: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Error checking affected rows: %v", err)
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no comment found with the given ID")
	}

	return nil
}
