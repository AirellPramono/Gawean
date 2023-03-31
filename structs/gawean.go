package structs

type Gawean struct {
	ID              int    `json:"id"`
	ClientID        string `json:"clientid"`
	ClientName      string `json:"clientname"`
	Description     string `json:"description"`
	CreateDate      string `json:"created"`
	UpdateDate      string `json:"updated"`
	CategoryID      int    `json:"category"`
	NextAppointment string `json:"next"`
}
