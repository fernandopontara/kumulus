package repository

import (
	"database/sql"
	"kumulus/internal/model"
)

type KnowledgeRepository struct {
	DB *sql.DB
}

func (r *KnowledgeRepository) BuscarContextos(orgID, userID, role string, embedding []float32) ([]model.KnowledgeBaseEntry, error) {
	query := `
	SELECT * FROM knowledge_base_entries
	WHERE organization_id = $1 AND (
		(owner_type = 'organization' AND role_scope IS NULL) OR
		(owner_type = 'organization' AND role_scope = $2) OR
		(owner_type = 'user' AND owner_id = $3)
	)
	ORDER BY embedding <-> $4
	LIMIT 6;
	`
	rows, err := r.DB.Query(query, orgID, role, userID, embedding)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var entries []model.KnowledgeBaseEntry
	for rows.Next() {
		var e model.KnowledgeBaseEntry
		// faça o Scan dos dados aqui
		entries = append(entries, e)
	}
	return entries, nil
}
