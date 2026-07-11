package data

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"time"

	"github.com/eduartepaiva/greenlight/internal/validator"
)

type Scope string

const (
	ScopeActivation     Scope = "activation"
	ScopeAuthentication       = "authentication"
)

type Token struct {
	Plaintext string    `json:"token"`
	Hash      []byte    `json:"-"`
	UserID    int64     `json:"-"`
	Expiry    time.Time `json:"expiry"`
	Scope     Scope     `json:"-"`
}

func generateToken(userID int64, ttl time.Duration, scope Scope) Token {
	token := Token{
		UserID:    userID,
		Expiry:    time.Now().Add(ttl),
		Scope:     scope,
		Plaintext: rand.Text(),
	}

	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]
	return token
}

func ValidateTokenPlaintext(v *validator.Validator, tokenPlaintext string) {
	v.Check(tokenPlaintext != "", "token", "must be provided")
	v.Check(len(tokenPlaintext) == 26, "token", "must be 26 bytes long")
}

type TokenModel struct {
	DB *sql.DB
}

func (m TokenModel) Insert(token Token) error {
	query := `
		INSERT INTO tokens(hash, user_id, expiry, scope) 
		VALUES ($1, $2, $3, $4)`

	args := []any{token.Hash, token.UserID, token.Expiry, token.Scope}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, args...)

	return err
}

func (m TokenModel) New(userID int64, ttl time.Duration, scope Scope) (Token, error) {
	token := generateToken(userID, ttl, scope)

	err := m.Insert(token)

	return token, err
}

func (m TokenModel) DeleteAllForUser(userID int64, scope Scope) error {
	query := "DELETE FROM tokens WHERE user_id = $1 AND scope = $2"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, userID, scope)

	return err
}

func (m TokenModel) GetAllForUser(userID int64) ([]Token, error) {
	query := `
		SELECT tokens.hash, tokens.user_id, tokens.expiry, tokens.scope
		FROM users
		RIGHT JOIN tokens
		ON users.id = tokens.user_id
		WHERE users.id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tokens := []Token{}

	for rows.Next() {
		token := Token{}
		err := rows.Scan(
			&token.Hash,
			&token.UserID,
			&token.Expiry,
			&token.Scope,
		)
		if err != nil {
			return nil, err
		}
		tokens = append(tokens, token)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tokens, nil
}
