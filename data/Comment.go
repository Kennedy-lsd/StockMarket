package data

import "time"

type Comment struct {
	Id        *uint      `json:"id"`
	Title     *string    `json:"title"`
	CreatedAt *time.Time `json:"created_at"`
	StockId   *uint      `json:"stock_id"`
}

type CommentCreate struct {
	Id        *uint      `json:"id"`
	Title     *string    `json:"title"`
	CreatedAt *time.Time `json:"created_at"`
	StockId   *uint      `json:"stock_id"`
}
