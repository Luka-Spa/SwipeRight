package cassandra

import "github.com/Luka-Spa/SwipeRight/packages/recommendation/model"

type recommendationRepository struct {
}

func NewRecommendationRepository() *recommendationRepository {
	return &recommendationRepository{}
}

func (*recommendationRepository) CreateRecommendationProfile(user model.RecommendationProfile) error {
	var query = `INSERT INTO recommendation.user_table
							(id,profile_image_url,first_name,
							gender,bio,location,anthem,school,preferences)
							values (?,?,?,?,?,?,?,?,?);`
	err := CassandraWrite(query, user.Id,
		user.ProfileImageUrl, user.Firstname, user.Gender, user.Bio, user.Location, user.Anthem, user.School, user.Preferences)
	return err
}
