package model

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	"time"
)

type Signature struct {
	UserID    string    `json:"user_id"`
	Questions []string  `json:"questions"`
	Answers   []string  `json:"answers"`
	Signature string    `json:"signature"`
	Timestamp time.Time `json:"timestamp"`
}

func NewSignature(userID string, questions, answers []string) *Signature {
	signature := generateSignatureString(userID, questions, answers)
	return &Signature{
		UserID:    userID,
		Questions: questions,
		Answers:   answers,
		Signature: signature,
		Timestamp: time.Now(),
	}
}

func generateSignatureString(userID string, questions, answers []string) string {
	data := userID + "|" + strings.Join(questions, "|") + "|" + strings.Join(answers, "|")
	hasher := sha256.New()
	hasher.Write([]byte(data))
	sha := hasher.Sum(nil)
	return hex.EncodeToString(sha)
}
