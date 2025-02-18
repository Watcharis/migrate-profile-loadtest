package services

import (
	"context"
	"fmt"
	"log"
	"sync"
	"watcharis/migrate-profile-loadtest/models"
	"watcharis/migrate-profile-loadtest/pkg"
	"watcharis/migrate-profile-loadtest/repositories"
)

func processInsertProfileAnnouncement(ctx context.Context, wg *sync.WaitGroup, profileRepository repositories.ProfileRepository, transferCH <-chan []models.Profile, errorCH chan<- error) {
	defer func() {
		wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			for profile := range transferCH {

				fmt.Println("profile len:", len(profile))

				if len(profile) > models.LIMIT_SIZE {
					errorCH <- fmt.Errorf("mock error")
					return
				}

				errorCH <- nil
			}
			return
		}
	}
}

func processErrorCH(ctx context.Context, wg *sync.WaitGroup, cancel context.CancelFunc, errorCH <-chan error, errCH *error) {
	defer func() {
		wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			for err := range errorCH {
				fmt.Println("err profile :", err)
				if err != nil {
					*errCH = err
					cancel()
					return
				}
			}
			return
		}
	}
}

func mainProcessInsert(profileRepository repositories.ProfileRepository) (errCH error) {
	wg := new(sync.WaitGroup)
	wgErr := new(sync.WaitGroup)

	// channel
	transferCH := make(chan []models.Profile, models.BUFFER_CHANNEL)

	// var errCH error
	errorCH := make(chan error)

	// new context
	nctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	defer func() error {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			return errCH
		}
		return nil
	}()
	// process profile
	for worker := 0; worker < models.GO_WORKER; worker++ {
		wg.Add(1)
		go processInsertProfileAnnouncement(nctx, wg, profileRepository, transferCH, errorCH)
	}

	// handle error
	wgErr.Add(1)
	go processErrorCH(nctx, wgErr, cancel, errorCH, &errCH)

	id := 0
	count := 0
	for {
		if count == models.TOTAL {
			break
		}

		result, err := profileRepository.GetProfile(nctx, id, models.LIMIT_SIZE)
		if err != nil {
			log.Panic(err)
			// return err
		}

		id = result[len(result)-1].ID
		count += len(result)
		if count == (models.TOTAL - models.LIMIT_SIZE) {
			fmt.Println("count :", count)
			result = append(result, result[len(result)-1])
		}

		transferCH <- result
	}

	close(transferCH)
	wg.Wait()
	close(errorCH)
	wgErr.Wait()

	return nil
}

func main_BackUp_case_failed() {
	// ctx := context.Background()
	db := pkg.InitDatabase()
	fmt.Printf("db : %+v\n", db)

	profileRepository := repositories.NewProfileRepository(db)

	if err := mainProcessInsert(profileRepository); err != nil {
		fmt.Println("[ERROR] insert profile announcement failed : ", err)
	}

	fmt.Println("process success")
}
