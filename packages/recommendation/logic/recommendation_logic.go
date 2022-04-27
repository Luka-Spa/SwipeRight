package logic

import (
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/model"
	"github.com/Luka-Spa/SwipeRight/packages/recommendation/repository"
)

var recommendationRepository repository.IRecommendationRepository

type IRecomendationLogic interface {
	CreateRecommendationProfile(recommendationProfile model.RecommendationProfile) error
}
type logic struct{}

func NewUserlogic(repository repository.IRecommendationRepository) *logic {
	recommendationRepository = repository
	return &logic{}
}

func (*logic) CreateRecommendationProfile(recommendationProfile model.RecommendationProfile) error {
	return recommendationRepository.CreateRecommendationProfile(recommendationProfile)
}
