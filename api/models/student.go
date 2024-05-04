package models

type Student struct {
	Id         string `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	ExternalId string `json:"external_id"`
	Phone      string `json:"phone"`
	Email      string `json:"mail"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	IsActive   bool   `json:"is_active"`
}

type StudentTimeTable struct {
	Id        string `json:"id"`
	Subject   string `json:"subject"`
	Teacher   string `json:"teacher"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}

type GetStudent struct {
	Id         string             `json:"id"`
	FirstName  string             `json:"first_name,omitempty"`
	LastName   string             `json:"last_name,omitempty"`
	Age        int                `json:"age,omitempty"`
	ExternalId string             `json:"external_id,omitempty"`
	Phone      string             `json:"phone,omitempty"`
	Email      string             `json:"email,omitempty"`
	CreatedAt  string             `json:"created_at"`
	UpdatedAt  string             `json:"updated_at"`
	IsActive   bool               `json:"is_active"`
	TimeTables []StudentTimeTable `json:"time_tables"`
}

type GetAllStudentsRequest struct {
	Search string `json:"search"`
	Page   uint64 `json:"page"`
	Limit  uint64 `json:"limit"`
}

type GetAllStudentsResponse struct {
	Students []GetStudent `json:"students"`
	Count    int64        `json:"count"`
}
