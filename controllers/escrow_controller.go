package controllers

import (
	"context"
	"net/http"
	"time"
	"fund/db"
	"fund/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateEscrow(c *gin.Context) {
	var input struct {
		PayerID string  `json:"payer_id" binding:"required"`
		PayeeID string  `json:"payee_id" binding:"required"`
		Amount  float64 `json:"amount" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	payerID, _ := primitive.ObjectIDFromHex(input.PayerID)
	payeeID, _ := primitive.ObjectIDFromHex(input.PayeeID)

	escrow := models.Escrow{
		PayerID:         payerID,
		PayeeID:         payeeID,
		Amount:          input.Amount,
		Status:          "Pending",
		ApprovedByPayer: false,
		ApprovedByPayee: false,
		CreatedAt:       time.Now().Unix(),
	}

	_, err := db.EscrowCollection.InsertOne(context.TODO(), escrow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create escrow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Escrow created successfully"})
}

func ApproveEscrow(c *gin.Context) {
	escrowID := c.Param("id")
	role := c.Query("role")

	objectID, err := primitive.ObjectIDFromHex(escrowID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid escrow ID"})
		return
	}

	update := bson.M{}
	if role == "payer" {
		update = bson.M{"$set": bson.M{"approved_by_payer": true}}
	} else if role == "payee" {
		update = bson.M{"$set": bson.M{"approved_by_payee": true}}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role"})
		return
	}

	_, err = db.EscrowCollection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to approve escrow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Escrow approved"})
}

func CancelEscrow(c *gin.Context) {
	escrowID := c.Param("id")

	objectID, err := primitive.ObjectIDFromHex(escrowID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid escrow ID"})
		return
	}

	_, err = db.EscrowCollection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, bson.M{"$set": bson.M{"status": "Cancelled"}})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to cancel escrow"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Escrow cancelled"})
}
