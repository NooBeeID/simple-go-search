package main

import (
	"fmt"
	"simple-go-search/config"
	"simple-go-search/db"
	"simple-go-search/server/controllers"
	gormRepo "simple-go-search/server/repositories/gorm"
	"simple-go-search/server/router"
	"simple-go-search/server/services"
)

func main() {
	appDB := db.NewGormDB().GetConnection()
	if appDB != nil {
		fmt.Println("connected")
	}

	// repo
	courseRepo := gormRepo.NewGormCourse(appDB)

	// svc
	courseSvc := services.NewCourseSvc(courseRepo)

	// controller
	courseApp := controllers.NewCourseController(courseSvc)

	port := fmt.Sprintf(":%s", config.GetString(config.APP_PORT, "3001"))
	appRotuer := router.NewRouterGin(port, courseApp)
	appRotuer.Start()

}
