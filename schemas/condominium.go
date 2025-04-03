package schemas

import "github.com/google/uuid"

type Condominium struct {
	ID             uuid.UUID
	Name           string
	Phone          string
	Address        string
	ZipCode        string
	ContractStatus bool

	Residents     []Residents
	AdminUserID   *int
	AdminUser     *User
	Events        []Event
	MonthlyFees   []MonthlyFee
	Notifications []Notification
	Areas         []Area
	Reservations  []Reservation
}
