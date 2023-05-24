package repository

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	if err != nil {
		t.Fatalf("cant open db: %s", err)
	}

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
