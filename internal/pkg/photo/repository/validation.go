package repository

import (
	"fmt"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
)

func (r *PhotoRepository) CheckPhoto(row dataStruct.UserPhoto) error {
	var count int64
	err := r.db.Table("user_photos").Where("user_id= ?", row.UserId).Count(&count).Error
	if err != nil {
		return err
	}
	switch count {
	case 5:
		err = fmt.Errorf("too many photos, please delete something")
		return err
	case 0:
		row.Avatar = true
		return nil
	}

	if row.Avatar {
		var oldAvatar dataStruct.UserPhoto
		err = r.db.First(&oldAvatar, "user_id=? AND avatar=?", row.UserId, true).Error
		if err != nil {
			return err
		}
		oldAvatar.Avatar = false
		r.db.Save(&oldAvatar)
		return nil
	}

	return nil
}

func (r *PhotoRepository) CheckDeletedPhoto(row dataStruct.UserPhoto) error {
	var count int64
	err := r.db.Table("user_photos").Where("user_id= ?", row.UserId).Count(&count).Error
	if count == 1 {
		err = fmt.Errorf("you canr delete all your photos")
	}
	if err != nil {
		return err
	}
	if row.Avatar {
		var newAvatar dataStruct.UserPhoto
		err = r.db.Last(&newAvatar, "user_id = ? AND avatar = ?", row.UserId, false).Error
		if err != nil {
			return err
		}
		newAvatar.Avatar = true
		err = r.db.Save(&newAvatar).Error
		if err != nil {
			return err
		}
	}
	return nil
}
