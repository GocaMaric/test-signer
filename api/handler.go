package api

import (
	"encoding/json"
	"net/http"
	"test-signer/model"
	"test-signer/store"
)

func SignHandler(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID    string   `json:"user_id"`
			Questions []string `json:"questions"`
			Answers   []string `json:"answers"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		signature := model.NewSignature(req.UserID, req.Questions, req.Answers)
		if err := s.SaveSignature(signature); err != nil {
			http.Error(w, "Error saving signature", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(map[string]string{"signature": signature.Signature})
	}
}

func VerifyHandler(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			UserID    string `json:"user_id"`
			Signature string `json:"signature"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		sig, err := s.VerifySignature(req.UserID, req.Signature)
		if err != nil {
			http.Error(w, "Error verifying signature", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(sig)
	}
}
