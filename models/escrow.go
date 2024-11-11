package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Escrow struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	PayerID         primitive.ObjectID `bson:"payer_id" json:"payer_id"`
	PayeeID         primitive.ObjectID `bson:"payee_id" json:"payee_id"`
	Amount          float64            `bson:"amount" json:"amount"`
	Status          string             `bson:"status" json:"status"`
	ApprovedByPayer bool               `bson:"approved_by_payer" json:"approved_by_payer"`
	ApprovedByPayee bool               `bson:"approved_by_payee" json:"approved_by_payee"`
	CreatedAt       int64              `bson:"created_at" json:"created_at"`
}
