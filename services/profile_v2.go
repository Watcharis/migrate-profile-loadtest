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

type migrateProfileAnnouncementV2Service struct {
	profileRepository repositories.ProfileRepository
}

func NewMigrateProfileAnnouncementV2Service(profileRepository repositories.ProfileRepository) MigrateProfileAnnouncementService {
	return &migrateProfileAnnouncementV2Service{
		profileRepository: profileRepository,
	}
}

func (s *migrateProfileAnnouncementV2Service) ProcessMigrateProfileAnnouncement(ctx context.Context) error {
	return s.mainProcessInsert(ctx)
}

func (s *migrateProfileAnnouncementV2Service) mainProcessInsert(ctx context.Context) error {

	wg := new(sync.WaitGroup)
	wgErr := new(sync.WaitGroup)

	// init control step
	controlStep := new(models.ControlStep)

	// channel
	transferCH := make(chan []models.Profile, models.BUFFER_CHANNEL)
	errorCH := make(chan error, models.BUFFER_CHANNEL)

	// new context
	nctx, cancel := context.WithCancel(ctx)
	defer cancel()

	for worker := 0; worker < models.GO_WORKER; worker++ {
		wgErr.Add(1)
		go s.processErrorCH(nctx, wgErr, cancel, errorCH, controlStep)
	}

	// process profile
	for worker := 0; worker < models.GO_WORKER; worker++ {
		wg.Add(1)
		go s.processInsertProfileAnnouncement(nctx, wg, transferCH, errorCH)
	}

	// send data to channel
	s.transferData(nctx, transferCH, errorCH)

	close(transferCH)
	wg.Wait()
	close(errorCH)
	wgErr.Wait()

	if err := controlStep.Finally; err != nil {
		log.Println("[ERROR] insert failed :", err)
		return err
	}

	log.Println("insert success")
	return nil
}

func (s *migrateProfileAnnouncementV2Service) transferData(ctx context.Context, transferCH chan<- []models.Profile, errorCH chan<- error) {
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

func (s *migrateProfileAnnouncementV2Service) processInsertProfileAnnouncement(ctx context.Context, wg *sync.WaitGroup,
	transferCH <-chan []models.Profile, errorCH chan<- error) {
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
				// จำลอง error
				if len(profile) > models.LIMIT_SIZE {
					errorCH <- errors.New("mock error")
					return
				}

				// process insert data
				if err := s.profileRepository.InsertProfileAnnouncement(ctx, profile); err != nil {
					log.Printf("[ERROR] cannot insert profile_announcement : %+v\n", err)
					errorCH <- err
					return
				}

				errorCH <- nil
			}
			return
		}
	}
}

func (s *migrateProfileAnnouncementV2Service) processErrorCH(ctx context.Context, wg *sync.WaitGroup, cancel context.CancelFunc,
	errorCH <-chan error, controlStep *models.ControlStep) {
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
					// fmt.Println("err profile :", err)
					controlStep.Finally = err
					cancel()
					return
				}
			}
			return
		}
	}
}
