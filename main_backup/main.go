package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"watcharis/migrate-profile-loadtest/models"
	"watcharis/migrate-profile-loadtest/pkg"
	"watcharis/migrate-profile-loadtest/repositories"
)

func main_BackUp() {
	ctx := context.Background()

	db := pkg.InitDatabase()
	fmt.Printf("db : %+v\n", db)

	profileRepository := repositories.NewProfileRepository(db)

	transferCH := make(chan []models.Profile, 10)
	// errorCH := make(chan error, 10)
	wg := new(sync.WaitGroup)

	id := 0
	count := 0

	for worker := 0; worker < models.GO_WORKER; worker++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()

			for profiles := range transferCH {
				fmt.Println("profiles in channel:", len(profiles))

				if err := profileRepository.InsertProfileAnnouncement(ctx, profiles); err != nil {
					return
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
			close(transferCH)
		}()

		for {
			if count == models.TOTAL {
				break
			}

			result, err := profileRepository.GetProfile(ctx, id, models.LIMIT_SIZE)
			if err != nil {
				log.Panic(err)
			}
			fmt.Println("result :", len(result))

			id = result[len(result)-1].ID

			count += len(result)

			transferCH <- result
		}
	}()
	wg.Wait()
}
