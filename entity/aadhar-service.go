package entity

// aadhar object for REST(CRUD)
type Aadhar struct {
	Id int `json:"id"`

	Aadharnum int `json:"aadharnum"`

	Status int    `json:"status"`
	Reason string `json:"reason"`
}
