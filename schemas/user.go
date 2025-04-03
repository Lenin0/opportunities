package schemas

import "github.com/google/uuid"

type Role string

const (
	Admin    Role = "ADMIN"
	Resident Role = "RESIDENT"
)

type User struct {
	ID               uuid.UUID
	Email            string
	Password         string
	Role             Role
	Resident         *Residents
	AdminCondominium *Condominium
}
