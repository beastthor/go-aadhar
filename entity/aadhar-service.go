package entity

// aadhar object for REST(CRUD)
type Aadhar struct {
	demographics_id int `json:"id"`

	Aadharnum int `json:"aadharnum"`

	Status int `json:"status"`
}
