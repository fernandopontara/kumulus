package model

type OwnerType string

const (
	OwnerOrg  OwnerType = "organization"
	OwnerUser OwnerType = "user"
)

type KnowledgeBaseEntry struct {
	ID             string    `db:"id"`
	OrganizationID string    `db:"organization_id"`
	OwnerType      OwnerType `db:"owner_type"` // "organization" ou "user"
	OwnerID        *string   `db:"owner_id"`   // null se for "organization"
	RoleScope      *string   `db:"role_scope"` // "vendedor", "financeiro", etc. ou null
	Content        string    `db:"content"`
	Embedding      []float32 `db:"embedding"`
}
