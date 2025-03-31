package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Openning struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

type OpenningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedat,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    string    `json:"salary"`
}
