package structs

type Address struct {
	ID          string `json:"id"`
	Province    string `json:"province"`
	City        string `json:"city"`
	Description string `json:"description"`
}
