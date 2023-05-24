package period_function

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
	"gorm.io/gorm"

	"github.com/go-park-mail-ru/2023_1_MRGA.git/internal/app/repository"
)

func RunCronJobs(db *gorm.DB) {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.Every(1).Day().Do(func() {
		err := repository.CleanCount(db)
		if err != nil {
			log.Println("error while cleaning count:", err)
			return
		} else {
			log.Println("count is cleaned")
		}

	})
	if err != nil {
		log.Println(err)
		return
	}
	s.StartBlocking()
}
