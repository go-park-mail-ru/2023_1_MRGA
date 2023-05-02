package repository

func (repository Repository) UploadFile(filePath string, userID uint) (uint, error) {
	file := &File{
		Path:   filePath,
		UserID: userID,
	}
	result := repository.db.Create(&file)

	// Проверка, что запись создана
	if result.Error == nil && result.RowsAffected > 0 {
		// Получение ID только что созданной записи
		return file.ID, nil
	}

	return 0, result.Error
}

func (repository Repository) GetFile(id uint) (string, error) {
	var file File
	err := repository.db.First(&file, id).Error
	return file.Path, err
}
