package logic

import (
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/repository"
)

var recommendationRepository repository.IRecommendationRepository

type IReccomendationLogic interface {
}
type logic struct{}

func NewUserlogic(repository repository.IRecommendationRepository) *logic {
	recommendationRepository = repository
	return &logic{}
}
