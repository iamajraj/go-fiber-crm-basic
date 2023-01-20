package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/iamajraj/go-fiber-crm-basic/database"
	"github.com/iamajraj/go-fiber-crm-basic/lead"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func setupRoutes(app *fiber.App){
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase(){
	var err error
	database.DBConn, err = gorm.Open("postgres", "host=localhost user=postgres password=admin dbname=fiber_crm port=5432 sslmode=disable TimeZone=Asia/Shanghai")
	if err != nil{
		println(err.Error())
		panic("failed to connect database")
	}
	fmt.Println("connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database Migrated")
}

func main(){
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()
}