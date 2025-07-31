package dto

import "time"

type Pagination struct {
	CurrentPage  int `json:"current_page"`
	Limit        int `json:"limit"`
	TotalPages   int `json:"total_pages"`
	TotalRecords int `json:"total_records"`
}

type ResponseUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

type ResponValidatePhone struct {
	Status string `json:"status"`
	Phone  string `json:"phone"`
	Vendor string `json:"vendor"`
}

type RevenueReport struct {
	Total           float64        `json:"total"`
	StatusBreakdown map[string]int `json:"status_breakdown"`
}

type MonthlyRevenue struct {
	Month string  `json:"month"`
	Total float64 `json:"total"`
}

type TopProduct struct {
	Name         string    `json:"name"`
	SellPrice    float64   `json:"sell_price"`
	Profit       float64   `json:"profit"`
	Margin       float64   `json:"margin"`
	TotalRevenue float64   `json:"total_revenue"`
	RevenueDate  time.Time `json:"revenue_date"`
}

type ProfileResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Role    string `json:"role"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}
