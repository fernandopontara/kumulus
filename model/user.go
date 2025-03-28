package model

type User struct {
	ID             string
	OrganizationID string
	Name           string
	Email          string
	PasswordHash   string
	Role           string // vendedor, atendimento, financeiro, admin
}
