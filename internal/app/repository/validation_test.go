package repository_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/constform"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/ds"
	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/repository"
)

func TestCheckAge(t *testing.T) {
	testCases := []struct {
		inp         int
		errRequired bool
	}{
		{
			inp:         -10,
			errRequired: true,
		},
		{
			inp:         10,
			errRequired: true,
		},
		{
			inp:         0,
			errRequired: true,
		},
		{
			inp:         20,
			errRequired: false,
		},
		{
			inp:         567,
			errRequired: true,
		},
	}

	for _, tCase := range testCases {
		err := repository.CheckAge(tCase.inp)

		if tCase.errRequired {
			require.Error(t, err)
			continue
		}

	}
}

func TestRepository_CheckUsername(t *testing.T) {
	repo := repository.NewRepo()
	users := []ds.User{
		{
			UserId:      0,
			Username:    "user1",
			Email:       "email1.com",
			Password:    "123_user1_321",
			Age:         20,
			Sex:         constform.Male,
			City:        "Москва",
			Description: "Im cool",
		},
		{
			UserId:      0,
			Username:    "user2",
			Email:       "email2.com",
			Password:    "123_user2_321",
			Age:         30,
			Sex:         constform.Female,
			City:        "Москва",
			Description: "Im cooler",
		},
	}

	for _, u := range users {
		err := repo.AddUser(&u)
		if err != nil {

			os.Exit(2)
		}
	}

	testCases := []struct {
		inp         ds.User
		errRequired bool
	}{
		{
			inp: ds.User{
				Username:    "user1",
				Email:       "email3.com",
				Password:    "123_user3_321",
				Age:         20,
				Sex:         constform.Male,
				City:        "Москва",
				Description: "Im cool",
			},
			errRequired: true,
		},
		{
			inp: ds.User{
				Username:    "user3",
				Email:       "email3.com",
				Password:    "123_user3_321",
				Age:         20,
				Sex:         constform.Female,
				City:        "Москва",
				Description: "Im cool",
			},
			errRequired: false,
		},
	}

	for _, tCase := range testCases {
		err := repo.CheckUsername(tCase.inp.Username)

		if tCase.errRequired {
			require.Error(t, err)
			continue
		}
	}

}

func TestRepository_CheckEmail(t *testing.T) {
	repo := repository.NewRepo()
	users := []ds.User{
		{
			UserId:      0,
			Username:    "user1",
			Email:       "email1.com",
			Password:    "123_user1_321",
			Age:         20,
			Sex:         constform.Male,
			City:        "Москва",
			Description: "Im cool",
		},
		{
			UserId:      0,
			Username:    "user2",
			Email:       "email2.com",
			Password:    "123_user2_321",
			Age:         30,
			Sex:         constform.Female,
			City:        "Москва",
			Description: "Im cooler",
		},
	}

	for _, u := range users {
		err := repo.AddUser(&u)
		if err != nil {

			os.Exit(2)
		}
	}

	testCases := []struct {
		inp         ds.User
		errRequired bool
	}{
		{
			inp: ds.User{
				Username:    "user3",
				Email:       "email1.com",
				Password:    "123_user3_321",
				Age:         20,
				Sex:         constform.Male,
				City:        "Москва",
				Description: "Im cool",
			},
			errRequired: true,
		},
		{
			inp: ds.User{
				Username:    "user3",
				Email:       "email3.com",
				Password:    "123_user3_321",
				Age:         20,
				Sex:         constform.Female,
				City:        "Москва",
				Description: "Im cool",
			},
			errRequired: false,
		},
	}

	for _, tCase := range testCases {
		err := repo.CheckEmail(tCase.inp.Email)

		if tCase.errRequired {
			require.Error(t, err)
			continue
		}
	}

}
