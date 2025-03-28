
# 🧠 GPT Seller SaaS – Treinamento de Vendedores com IA

Este projeto é um **SaaS multiempresa** onde empresas podem cadastrar seus vendedores, financeiro, atendimento e outros cargos. Cada usuário interage com um "ChatGPT personalizado" alimentado com:

- Dados da **empresa**
- Dados do **perfil/cargo**
- Dados **individuais** do usuário

As respostas são geradas via **RAG (Retrieval-Augmented Generation)** utilizando a API da OpenAI (GPT-3.5 / GPT-4).

---

## 🧱 Tecnologias e Arquitetura

- **GoLang 1.21+**
- [Gin Gonic](https://github.com/gin-gonic/gin) – roteamento HTTP
- [JWT](https://github.com/golang-jwt/jwt) – autenticação multiempresa
- [pgvector](https://github.com/pgvector/pgvector) + PostgreSQL – embeddings vetoriais
- [OpenAI API](https://platform.openai.com/docs) – GPT-3.5 ou GPT-4
- [golang-migrate](https://github.com/golang-migrate/migrate) – controle de migrações
- [godotenv](https://github.com/joho/godotenv) – variáveis de ambiente

---

## 📂 Estrutura do Projeto

/internal
├── auth/            # JWT
├── config/          # .env loader
├── handler/         # Handlers REST (Gin)
├── middleware/      # CORS, Auth, Logging
├── model/           # Structs de domínio
├── repository/      # DB Access
├── service/         # RAG, Knowledge, Auth
└── openai/          # Cliente OpenAI API

---

## ⚙️ Instalação e Setup

### 1. Clone o projeto

```bash
git clone https://github.com/seuusuario/gpt-seller-saas.git
cd gpt-seller-saas
```

### 2. Crie o arquivo `.env`

```env
OPENAI_API_KEY=sk-xxxxxx
PORT=8080
```

### 3. Instale as dependências

```bash
go mod tidy
```

### 4. Rode as migrações

```bash
migrate -path migrations -database "postgres://user:pass@localhost:5432/dbname?sslmode=disable" up
```

---

## ▶️ Executando

```bash
go run main.go
```

---

## 🔐 Endpoints

### `POST /login`

Autentica o usuário e retorna um JWT.

**Request:**
```json
{
  "email": "vendedor@empresa.com",
  "password": "123456"
}
```

**Response:**
```json
{
  "token": "jwt.aqui..."
}
```

---

### `POST /chat` (autenticado)

**Headers:**
```
Authorization: Bearer <TOKEN>
```

**Request:**
```json
{
  "question": "Como contornar uma objeção de preço?"
}
```

**Response:**
```json
{
  "answer": "Para contornar objeções de preço, você pode..."
}
```

---

### `POST /knowledge` (autenticado)

**Request:**
```json
{
  "content": "Sempre reforçar o valor percebido antes de mencionar o preço.",
  "role_scope": "vendedor"
}
```

---

## 🧠 Como funciona a RAG Engine

A consulta considera 3 camadas de conhecimento:

1. 🏢 **Empresa** (base macro)
2. 👤 **Perfil/Cargo** (ex: "vendedor", "financeiro")
3. 🙋 **Usuário** (base individual)

Esses conteúdos são vetorizados e usados para montar o prompt antes de enviar para o GPT da OpenAI.

---

## 📌 Funcionalidades futuras

- Painel admin (empresas, usuários, permissões)
- Upload de arquivos (PDF, texto) para treinamento
- Histórico e analytics de interações
- Interface web (React)
- Dashboard de uso por empresa e usuário

---

## 👨‍💻 Autor

- Fernando Pontara – Arquitetura, backend Go, integrações IA
- [ChatGPT](https://openai.com/chatgpt) – copiloto de desenvolvimento 🤖

---

## 📄 Licença

MIT License. Livre para usar, melhorar e contribuir!