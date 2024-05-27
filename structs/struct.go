package structs

type Canvasser struct {
	Id       int64  `json:"id"`
	Code     string `json:"code"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

type Item struct {
	Id    int64   `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type Customer struct {
	Id      int64  `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	Address string `json:"address"`
	Email   string `json:"email"`
}

type Stock struct {
	Id          int64 `json:"id"`
	ItemId      int64 `json:"item_id"`
	CanvasserId int64 `json:"canvasser_id"`
	Qty         int64 `json:"qty"`
}

type TrnSales struct {
	Id          int64   `json:"id"`
	CustomerId  int64   `json:"customer_id"`
	CanvasserId int64   `json:"canvasser_id"`
	Code        string  `json:"code"`
	DateSales   string  `json:"date_sales"`
	Description string  `json:"description"`
	SubTotal    float64 `json:"subtotal"`
	Discount    float64 `json:"discount"`
	Total       float64 `json:"total"`
}

type TrnSalesDetail struct {
	Id         int64   `json:"id"`
	TrnSalesId int64   `json:"trn_sales_id"`
	ItemId     int64   `json:"item_id"`
	Qty        int64   `json:"qty"`
	Subtotal   float64 `json:"subtotal"`
	Total      float64 `json:"total"`
}
