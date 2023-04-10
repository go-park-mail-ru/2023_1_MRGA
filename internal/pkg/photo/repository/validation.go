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
	if count == 5 {
		err = fmt.Errorf("Too many photos, please delete something")
		return err
	}
	if row.Avatar {
		var oldAvatar dataStruct.UserPhoto
		err = r.db.First(&oldAvatar, "userId=? AND avatar=?", row.UserId, true).Error
		if err != nil {
			return err
		}
		oldAvatar.Avatar = false
		r.db.Save(&oldAvatar)
		return nil
	}

	return nil
}
