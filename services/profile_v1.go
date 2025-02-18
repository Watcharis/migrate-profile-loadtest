package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"watcharis/migrate-profile-loadtest/models"
	"watcharis/migrate-profile-loadtest/repositories"
)

type migrateProfileAnnouncementService struct {
	profileRepository repositories.ProfileRepository
	finally           *error
}

func NewMigrateProfileAnnouncementService(profileRepository repositories.ProfileRepository) MigrateProfileAnnouncementService {
	return &migrateProfileAnnouncementService{
		profileRepository: profileRepository,
	}
}

func (s *migrateProfileAnnouncementService) ProcessMigrateProfileAnnouncement(ctx context.Context) error {
	return s.mainProcessInsert(ctx)
}

func (s *migrateProfileAnnouncementService) mainProcessInsert(ctx context.Context) error {

	wg := new(sync.WaitGroup)
	wgErr := new(sync.WaitGroup)

	// channel
	transferCH := make(chan []models.Profile, models.BUFFER_CHANNEL)
	errorCH := make(chan error, models.BUFFER_CHANNEL)

	// new context
	nctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// send data to channel
	wg.Add(1)
	go s.transferData(nctx, wg, transferCH, errorCH)

	// process profile
	for worker := 0; worker < models.GO_WORKER; worker++ {
		wg.Add(1)
		go s.processInsertProfileAnnouncement(nctx, wg, transferCH, errorCH)
	}

	// handle error
	for worker := 0; worker < models.GO_WORKER; worker++ {
		wgErr.Add(1)
		go s.processErrorCH(nctx, wgErr, cancel, errorCH)
	}

	wg.Wait()
	close(errorCH)
	wgErr.Wait()

	log.Println("insert success")

	if s.finally != nil {
		return *s.finally
	}

	return nil
}

func (s *migrateProfileAnnouncementService) transferData(ctx context.Context, wg *sync.WaitGroup, transferCH chan<- []models.Profile, errorCH chan<- error) {
	defer func() {
		wg.Done()
		close(transferCH)
	}()
	// send data to channel
	id := 0
	count := 0
	for {
		if count == models.TOTAL {
			break
		}

		result, err := s.profileRepository.GetProfile(ctx, id, models.LIMIT_SIZE)
		// err = fmt.Errorf("mock error, get profile failed")
		if err != nil {
			log.Println("err get profile :", err)
			errorCH <- err
			return
		}

		id = result[len(result)-1].ID

		count += len(result)

		if count == models.TOTAL {
			log.Println("count :", count)
			result = append(result, result[len(result)-1])
		}

		transferCH <- result
		errorCH <- nil
	}
}

func (s *migrateProfileAnnouncementService) processInsertProfileAnnouncement(ctx context.Context, wg *sync.WaitGroup, transferCH <-chan []models.Profile, errorCH chan<- error) {
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
				// process insert data

				// จำลอง error
				if len(profile) > models.LIMIT_SIZE {
					errorCH <- errors.New("mock error")
					return
				}

				errorCH <- nil
			}
			return
		}
	}
}

func (s *migrateProfileAnnouncementService) processErrorCH(ctx context.Context, wg *sync.WaitGroup, cancel context.CancelFunc, errorCH <-chan error) {
	defer func() {
		wg.Done()
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			for err := range errorCH {
				if err != nil {
					fmt.Println("err profile :", err)
					s.finally = &err
					cancel()
					return
				}
			}
			return
		}
	}
}
