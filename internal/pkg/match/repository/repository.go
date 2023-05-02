package repository

import (
	"time"

	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/pkg/match"
)

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepo(db *gorm.DB) *MatchRepository {
	return &MatchRepository{
		db,
	}
}

func (r *MatchRepository) GetMatches(userId uint) (users []dataStruct.Match, err error) {
	err = r.db.Table("matches").Find(&users, "user_first_id=?", userId).Error
	return
}

func (r *MatchRepository) GetUser(userId uint) (user match.UserRes, err error) {
	err = r.db.Table("users u").Select("u.id, ui.name, p.photo").
		Where("u.id=?", userId).
		Joins("Join user_infos ui on u.id = ui.user_id").
		Joins("join user_photos p on u.id=p.user_id").
		Where("p.avatar =?", true).
		Find(&user).Error
	return
}

func (r *MatchRepository) GetIdReaction(reaction string) (uint, error) {
	var react dataStruct.Reaction
	err := r.db.First(&react, "reaction = ?", reaction).Error
	return react.Id, err
}

func (r *MatchRepository) AddHistoryRow(row dataStruct.UserHistory) error {
	err := r.db.Create(&row).Error
	return err
}

func (r *MatchRepository) GetUserReaction(userId, userToId uint) (dataStruct.UserReaction, error) {
	var react dataStruct.UserReaction
	err := r.db.First(&react, "user_id = ? AND user_from_id=?", userId, userToId).Error
	return react, err
}

func (r *MatchRepository) AddUserReaction(row dataStruct.UserReaction) error {
	err := r.db.Create(&row).Error
	return err
}

func (r *MatchRepository) DeleteUserReaction(rowId uint) error {
	err := r.db.First(&dataStruct.UserReaction{}, "id =?", rowId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.UserReaction{}, "id=?", rowId).Error
	return err
}

func (r *MatchRepository) AddMatchRow(row dataStruct.Match) error {
	err := r.db.Create(&row).Error
	return err
}

func (r *MatchRepository) ChangeStatusMatch(userId, profileId uint) error {
	var matchDB dataStruct.Match
	err := r.db.First(&matchDB, "user_first_id = ? AND user_second_id = ?", userId, profileId).Error
	if err != nil {
		return err
	}
	matchDB.Shown = true
	err = r.db.Save(&matchDB).Error
	return err
}

func (r *MatchRepository) GetChat(userId uint) (match.ChatAnswer, error) {
	var user match.ChatAnswer
	err := r.db.Table("users u").Select(" u.email, ui.name, p.photo").
		Where("u.id =?", userId).
		Joins("Join user_infos ui on u.id = ui.user_id").
		Joins("Join user_photos p on p.user_id=u.id").
		Where("p.avatar = ?", true).
		Find(&user).Error
	return user, err
}

func (r *MatchRepository) GetAge(userId uint) (int, error) {
	var user dataStruct.User
	err := r.db.First(&user, "id=?", userId).Error
	if err != nil {
		return 0, err
	}
	age, err := CalculateAge(user.BirthDay)
	if err != nil {
		return 0, err
	}
	return age, nil
}

func CalculateAge(birthDay string) (int, error) {
	birth, err := time.Parse("2006-01-02", birthDay[:10])
	if err != nil {
		return 0, err
	}
	now := time.Now()
	age := now.Year() - birth.Year()
	if now.Month() < birth.Month() {
		age -= 1
	}
	if now.Month() == birth.Month() && now.Day() < birth.Day() {
		age -= 1
	}
	return age, nil
}

func (r *MatchRepository) DeleteMatch(userId, userMatchId uint) error {
	var match dataStruct.Match
	err := r.db.First(&match, "user_first_id =? AND user_second_id=?", userId, userMatchId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.Match{}, "id=?", match.Id).Error
	if err != nil {
		return err
	}
	var match2 dataStruct.Match
	err = r.db.First(&match2, "user_first_id =? AND user_second_id=?", userMatchId, userId).Error
	if err != nil {
		return err
	}
	err = r.db.Delete(&dataStruct.Match{}, "id=?", match.Id).Error
	return err
}
