package repository

import "github.com/Luka-Spa/SwipeRight/packages/recommendation/model"

type IRecommendationRepository interface {
	CreateRecommendationProfile(user model.RecommendationProfile) error
}
