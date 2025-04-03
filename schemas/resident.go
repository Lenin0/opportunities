package schemas

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Residents struct {
	gorm.Model
	Name          string
	Phone         string
	Apartment     int
	Block         int
	Plate         *string
	CPF           string
	CondominiumID uuid.UUID

	Condominium   Condominium
	UserID        *int
	User          *User
	Reservations  []Reservation
	Guests        []Guest
	Events        []Event
	Notifications []Notification
	Cards         []Card
	MonthlyFees   []MonthlyFee
}
