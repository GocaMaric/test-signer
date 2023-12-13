package store

import (
	"database/sql"
	"fmt"
	"log"
	"test-signer/model"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) SaveSignature(signature *model.Signature) error {
	query := `INSERT INTO signatures (user_id, questions, answers, signature, timestamp) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(query, signature.UserID, signature.Questions, signature.Answers, signature.Signature, signature.Timestamp)
	if err != nil {
		log.Printf("Error saving signature: %v", err)
		return fmt.Errorf("error saving signature: %w", err)
	}
	return nil
}

func (s *Store) VerifySignature(userID, signature string) (*model.Signature, error) {
	query := `SELECT user_id, questions, answers, signature, timestamp FROM signatures WHERE user_id = $1 AND signature = $2`
	row := s.db.QueryRow(query, userID, signature)

	var sig model.Signature
	err := row.Scan(&sig.UserID, &sig.Questions, &sig.Answers, &sig.Signature, &sig.Timestamp)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No signature found for user %s", userID)
			return nil, fmt.Errorf("no signature found for user %s", userID)
		}
		log.Printf("Error verifying signature: %v", err)
		return nil, fmt.Errorf("error verifying signature: %w", err)
	}

	return &sig, nil
}
