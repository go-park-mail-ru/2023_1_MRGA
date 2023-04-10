package repository

func (repo Repository) UploadFile(filePath string) (uint, error) {
	file := &File{
		Path: filePath,
	}
	result := repo.db.Create(&file)

	// Проверка, что запись создана
	if result.Error == nil && result.RowsAffected > 0 {
		// Получение ID только что созданной записи
		return file.ID, nil
	}

	return 0, result.Error
}

func (repo Repository) GetFile(id uint) (string, error) {
	var file File
	err := repo.db.First(&file, id).Error
	return file.Path, err
}
