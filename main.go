package main

import (
	"context"
	"fmt"
	"log"
	"watcharis/migrate-profile-loadtest/pkg"
	"watcharis/migrate-profile-loadtest/repositories"
	"watcharis/migrate-profile-loadtest/services"
)

func main() {
	ctx := context.Background()
	db := pkg.InitDatabase()
	fmt.Printf("db : %+v\n", db)

	profileRepository := repositories.NewProfileRepository(db)

	profileService := services.NewMigrateProfileAnnouncementV2Service(profileRepository)

	if err := profileService.ProcessMigrateProfileAnnouncement(ctx); err != nil {
		fmt.Println("[ERROR] insert profile announcement failed : ", err)
	}
	log.Println("process success")
}
