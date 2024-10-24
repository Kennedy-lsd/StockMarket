package repos

import (
	"database/sql"
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
