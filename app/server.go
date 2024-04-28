package main

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
	
	"net/http"
	"github.com/labstack/echo/v4"

	"zadanie_4_ebiz.com/app/controllers"
	"zadanie_4_ebiz.com/app/models"
)

func main() {
    db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
    if err != nil {
        panic("Błąd łączenia z bazą danych ❌")
    } 

    db.AutoMigrate(&models.Product{})
	// defer db.Close() // ?


	e := echo.New()
	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	productsFromList := e.Group("/productsList")

	productsFromList.GET("/getAll", controllers.GetAll)
	productsFromList.GET("/getOne/:id", controllers.GetOne)
	productsFromList.POST("/addOne", controllers.AddOne)
	productsFromList.DELETE("/deleteOne/:id", controllers.DeleteOne)
	productsFromList.PATCH("/patchOne/:id", controllers.PatchOne)


	productController := controllers.NewProductController(db)

	productsFromDB := e.Group("/productsDB")

	productsFromDB.GET("/getAll", productController.GetAllDB)
	productsFromDB.GET("/getOne/:id", productController.GetOneDB)
	productsFromDB.POST("/addOne", productController.AddOneDB)
	productsFromDB.DELETE("/deleteOne/:id", productController.DeleteOneDB)
	productsFromDB.PATCH("/patchOne/:id", productController.PatchOneDB)

	e.Logger.Fatal(e.Start(":8080"))
}
