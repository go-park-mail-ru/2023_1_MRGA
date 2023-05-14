package period_function

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/repository"
)

func RunCronJobs(db *gorm.DB) {
	// 3
	s := gocron.NewScheduler(time.UTC)

	// 4
	s.Every(1).Day().Do(func() {
		repository.CleanCount(db)
		log.Println("count is cleaned")
	})

	// 5
	s.StartBlocking()
}
