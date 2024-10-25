package repos

import (
	"database/sql"
	"fmt"
	"log"
	"sort"

	"github.com/Kennedy-lsd/StockMarket/data"
)

type StockRepository struct {
	DB *sql.DB
}

func NewStockRepository(db *sql.DB) *StockRepository {
	return &StockRepository{DB: db}
}

func (r *StockRepository) FindAll() ([]data.Stock, error) {
	query := "SELECT * FROM stocks s LEFT JOIN comments c ON s.id = c.stock_id"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stocksMap := make(map[uint]*data.Stock)
	for rows.Next() {
		var stock data.Stock
		var comment data.Comment

		err = rows.Scan(
			&stock.Id, &stock.CompanyName, &stock.CompanySymbol, &stock.Price, &stock.LastDiv,
			&comment.Id, &comment.Title, &comment.CreatedAt, &comment.StockId,
		)
		if err != nil {
			return nil, err
		}

		// Initialize stock comments if not already done
		if _, exists := stocksMap[stock.Id]; !exists {
			stock.Comments = []data.Comment{} // Initialize to an empty slice
			stocksMap[stock.Id] = &stock
		}

		// Handle possible nil comment
		if comment.Id != nil {
			stocksMap[stock.Id].Comments = append(stocksMap[stock.Id].Comments, comment)
		}
	}

	// Convert map to slice
	var stocks []data.Stock
	for _, stock := range stocksMap {
		stocks = append(stocks, *stock)
	}
	sort.Slice(stocks, func(i, j int) bool {
		return stocks[i].Id < stocks[j].Id
	})

	return stocks, nil
}

func (r *StockRepository) Post(stock *data.CreatedStock) error {
	query := `INSERT INTO stocks (company_name, company_symbol, price, last_div) 
		VALUES ($1, $2, $3, $4) RETURNING id`

	err := r.DB.QueryRow(query, &stock.CompanyName, &stock.CompanySymbol, &stock.Price, &stock.LastDiv).Scan(&stock.Id)
	if err != nil {
		log.Printf("Error creating stock: %v", err)
		return err
	}

	return nil
}

func (r *StockRepository) FindById(id int64) (*data.Stock, error) {
	query := `SELECT * FROM stocks WHERE id = $1`
	var stock data.Stock

	err := r.DB.QueryRow(query, id).Scan(&stock.Id, &stock.CompanyName, &stock.CompanySymbol, &stock.Price, &stock.LastDiv)
	if err != nil {
		return nil, err
	}

	commentsQuery := "SELECT * FROM comments WHERE stock_id = $1"
	rows, err := r.DB.Query(commentsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment data.Comment
		err := rows.Scan(&comment.Id, &comment.Title, &comment.CreatedAt, &comment.StockId)
		if err != nil {
			return nil, err
		}

		stock.Comments = append(stock.Comments, comment)
	}
	return &stock, nil
}

func (r *StockRepository) DeleteById(id int64) error {
	query := `DELETE FROM stocks WHERE id = $1`

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
		return fmt.Errorf("no stocks found with the given ID")
	}

	return nil
}
