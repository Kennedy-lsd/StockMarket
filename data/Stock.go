package data

type Stock struct {
	Id            uint      `json:"id"`
	CompanyName   string    `json:"company_name"`
	CompanySymbol string    `json:"company_symbol"`
	Price         string    `json:"price"`
	LastDiv       float32   `json:"last_div"`
	Comments      []Comment `json:"comments"`
}

type CreatedStock struct {
	Id            uint    `json:"id"`
	CompanyName   string  `json:"company_name"`
	CompanySymbol string  `json:"company_symbol"`
	Price         string  `json:"price"`
	LastDiv       float32 `json:"last_div"`
}
