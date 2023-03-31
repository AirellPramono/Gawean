package structs

type Client struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreateDate  string `json:"created"`
	UpdateDate  string `json:"updated"`
	City        string `json:"city"`
	Province    string `json:"province"`
	AddressID   string `json:"addressid"`
}
