package modelos

import "time"

//Usuario representa um usario representando um usuario
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"id,omitempty"`
	Nick     string    `json:"id,omitempty"`
	Email    string    `json:"id,omitempty"`
	Password string    `json:"id,omitempty"`
	CriadoEm time.Time `json:"id,omitempty"`
}
