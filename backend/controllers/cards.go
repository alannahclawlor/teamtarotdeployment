package controllers

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"main.go/models"
	"main.go/services"
)

func GetThreeCards(ctx *gin.Context) {
	deck, _ := services.FetchTarotCards() //returns a type of []Card
	// if err != nil {
	// 	SendInternalError(ctx, err)
	// 	return
	// } ADD THIS BACK IN WHEN ERROR CONTROLLER EXISTS

	var threeCards []models.Card
	threeCards = append(threeCards, GetRandomCard(deck, threeCards))
	threeCards = append(threeCards, GetRandomCard(deck, threeCards))
	threeCards = append(threeCards, GetRandomCard(deck, threeCards))

	// from here below we convert the three Card into three JSONCard

	var jsonCards []models.JSONCard

	for _, card := range threeCards {
		jsonCards = append(jsonCards, models.JSONCard{
			CardName:       card.CardName,
			Type:           card.Type,
			MeaningUp:      card.MeaningUp,
			MeaningReverse: card.MeaningReverse,
			Description:    card.Description,
			ImageName:      card.ShortName + ".jpg",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"cards": jsonCards})
}

func GetRandomCard(deck []models.Card, currentCards []models.Card) models.Card {
	randomiser := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		randomIndex := randomiser.Intn(len(deck))
		randomCard := deck[randomIndex]

		isDuplicate := false
		for _, card := range currentCards {
			if card.CardName == randomCard.CardName {
				isDuplicate = true
				break
			}
		}

		if !isDuplicate {
			return randomCard
		}
	}
}