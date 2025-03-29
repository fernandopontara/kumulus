package Service

import (
	"fmt"
	"strings"
)

// ChatService represents the service with dependencies.
type ChatService struct {
	Embedder Embedder
	Repo     Repository
}

func (s *ChatService) MontarPrompt(orgID, userID, role, pergunta string) (string, error) {
	embedding, err := s.Embedder.Generate(pergunta)
	if err != nil {
		return "", err
	}

	contextos, err := s.Repo.BuscarContextos(orgID, userID, role, embedding)
	if err != nil {
		return "", err
	}

	var partes []string
	for _, ctx := range contextos {
		var tag string
		switch {
		case ctx.OwnerType == "organization" && ctx.RoleScope == nil:
			tag = "[Contexto Geral da Empresa]"
		case ctx.OwnerType == "organization" && ctx.RoleScope != nil:
			tag = fmt.Sprintf("[Contexto por Perfil: %s]", *ctx.RoleScope)
		case ctx.OwnerType == "user":
			tag = "[Contexto do Usuário]"
		}
		partes = append(partes, fmt.Sprintf("%s\n%s", tag, ctx.Content))
	}

	return fmt.Sprintf("%s\n\nPergunta:\n%s", strings.Join(partes, "\n\n"), pergunta), nil
}
