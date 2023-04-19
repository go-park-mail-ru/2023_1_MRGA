package recommendation

import dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"

type IRepositoryRec interface {
	GetRecommendation(userId uint, history []uint, reasons []uint, hashtags []uint, filters dataStruct.UserFilter) (users []UserRecommend, err error)
	GetRecommendedUser(userId uint) (user Recommendation, err error)

	GetUserHistory(userId uint) ([]uint, error)

	//GetPhotos(userId uint) ([]uint, error) //photo
	//GetAvatar(userId uint) (uint, error)   //photo

}
