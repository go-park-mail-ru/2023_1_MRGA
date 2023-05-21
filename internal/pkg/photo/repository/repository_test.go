package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	dataStruct "github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/data_struct"
)

func TestNewUserRepo(t *testing.T) {
	db, _, err := sqlmock.New()
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	pDb, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: false,
		DriverName:           "postgres",
		Conn:                 db,
	}))
	defer db.Close()
	userRepo := NewPhotoRepo(pDb)
	if userRepo != nil {
		return
	}
}

func TestPhotoRepo_GetAvatar(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	pDb, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: false,
		DriverName:           "postgres",
		Conn:                 db,
	}))
	photoRepo := NewPhotoRepo(pDb)

	userId := uint(1)
	avatarDB := uint(1)

	avatar := sqlmock.NewRows([]string{"photo"}).AddRow(avatarDB)

	query := `SELECT photo FROM user_photos p WHERE user_id = $1 AND avatar = $2`
	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(userId, true).
		WillReturnRows(avatar)

	avatarRepo, err := photoRepo.GetAvatar(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.Equal(t, avatarRepo, avatarDB)
}

func TestPhotoRepo_GetAvatar_GetError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	pDb, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: false,
		DriverName:           "postgres",
		Conn:                 db,
	}))
	photoRepo := NewPhotoRepo(pDb)

	userId := uint(1)

	query := `SELECT photo FROM user_photos p WHERE user_id = $1 AND avatar = $2`
	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(userId, true).
		WillReturnError(gorm.ErrRecordNotFound)

	_, err = photoRepo.GetAvatar(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
}

func TestPhotoRepo_GetPhotos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	pDb, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: false,
		DriverName:           "postgres",
		Conn:                 db,
	}))
	photoRepo := NewPhotoRepo(pDb)

	userId := uint(1)
	avatarDB := []uint{1, 2, 3}
	avatar := sqlmock.NewRows([]string{"photo"}).AddRow(avatarDB[0]).AddRow(avatarDB[1]).AddRow(avatarDB[2])
	query := `SELECT up.photo FROM user_photos up WHERE user_id = $1 AND avatar = $2`
	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(userId, false).
		WillReturnRows(avatar)

	avatarRepo, err := photoRepo.GetPhotos(userId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
	require.Equal(t, avatarRepo, avatarDB)
}

func TestPhotoRepo_GetPhotos_GetError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	pDb, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: false,
		DriverName:           "postgres",
		Conn:                 db,
	}))
	photoRepo := NewPhotoRepo(pDb)

	userId := uint(1)

	query := `SELECT up.photo FROM user_photos up WHERE user_id = $1 AND avatar = $2`
	mock.
		ExpectQuery(regexp.QuoteMeta(query)).
		WithArgs(userId, false).
		WillReturnError(gorm.ErrRecordNotFound)

	_, err = photoRepo.GetPhotos(userId)
	if err == nil {
		t.Errorf("unexpected err: %s", err)
		return
	}

	require.EqualError(t, err, gorm.ErrRecordNotFound.Error())
}

func TestPhotoRepository_ChangePhoto(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	pDb, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: false,
		DriverName:           "postgres",
		Conn:                 db,
	}))
	photoRepo := NewPhotoRepo(pDb)

	userId := uint(1)
	photoId := uint(1)
	newPhotoId := uint(2)
	photoDB := dataStruct.UserPhoto{
		Id:     uint(1),
		Photo:  photoId,
		Avatar: true,
		UserId: userId,
	}

	photo := sqlmock.NewRows([]string{"id", "photo", "user_id", "avatar"}).AddRow(photoDB.Id, photoDB.Photo, photoDB.UserId, photoDB.Avatar)

	query1 := `SELECT * FROM "user_photos" WHERE user_id = $1 AND photo = $2 ORDER BY "user_photos"."id" LIMIT 1`
	query2 := `UPDATE "user_photos"" SET "user_id" = $1, "photo" = $2, "avatar" = $3 WHERE "id"=$4 `
	mock.
		ExpectQuery(regexp.QuoteMeta(query1)).
		WithArgs(userId, photoId).
		WillReturnRows(photo)
	mock.ExpectBegin()
	mock.
		ExpectExec(regexp.QuoteMeta(query2)).
		WithArgs(userId, newPhotoId, true, photoId)
	mock.ExpectCommit()
	//mock.ExpectRollback()

	err = photoRepo.ChangePhoto(photoId, userId, newPhotoId)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}

func TestPhotoRepository_SavePhoto(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("cant create mock: %s", err)
	}
	defer db.Close()
	pDb, err := gorm.Open(postgres.New(postgres.Config{
		PreferSimpleProtocol: false,
		DriverName:           "postgres",
		Conn:                 db,
	}))
	photoRepo := NewPhotoRepo(pDb)

	userId := uint(1)
	photoId := uint(1)
	newPhotoId := uint(2)
	photoDB := dataStruct.UserPhoto{
		Photo:  photoId,
		Avatar: true,
		UserId: userId,
	}

	//photo := sqlmock.NewRows([]string{"id", "photo", "user_id", "avatar"}).AddRow(photoDB.Id, photoDB.Photo, photoDB.UserId, photoDB.Avatar)

	//query1 := `SELECT * FROM "user_photos" WHERE user_id = $1 AND photo = $2 ORDER BY "user_photos"."id" LIMIT 1`
	query2 := `INSERT INTO "user_photos" ("user_id", "photo", "avatar") VALUES( $1, $2, $3) RETURNING "id"`
	//mock.
	//	ExpectQuery(regexp.QuoteMeta(query1)).
	//	WithArgs(userId, photoId).
	//	WillReturnRows(photo)
	mock.ExpectBegin()
	mock.
		ExpectQuery(regexp.QuoteMeta(query2)).
		WithArgs(userId, newPhotoId, true).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(uint(0)))
	mock.ExpectCommit()

	err = photoRepo.SavePhoto(photoDB)
	if err != nil {
		t.Errorf("unexpected err: %s", err)
		return
	}
}
