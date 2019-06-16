package plutus

import (
	"time"
)

// CardTokenType represents a type of token from your charger service
type CardTokenType string

// OneUseToken is a one use token from your charger
var OneUseToken CardTokenType = "one_use"

// RecurrentToken is a saved card to recurrent use in your charger service
var RecurrentToken CardTokenType = "recurrent"

// CardToken represents a basic token generated by the charger service
type CardToken struct {
	ID        string
	Type      CardTokenType
	Value     string
	WithCard  EncodedCardDetails
	CreatedAt time.Time
}

func (token *CardToken) fillID() {
	token.ID = ids.New()
}
