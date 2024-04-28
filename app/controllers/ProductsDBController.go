package controllers

import (
    "net/http"
    "gorm.io/gorm"
    "strconv"
    "github.com/labstack/echo/v4"

	"zadanie_4_ebiz.com/app/models"
)

const invalidProductID = "nieprawidłowe ID produktu";


type ProductController struct {
    DB *gorm.DB
}

func NewProductController(db *gorm.DB) *ProductController {
    return &ProductController{DB: db}
}

func (controller *ProductController) GetAllDB(c echo.Context) error {
    var products []models.Product
    
	result := controller.DB.Find(&products)
    
	if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, result.Error)
    }
	
    return c.JSON(http.StatusOK, products)
}

func (controller *ProductController) GetOneDB(c echo.Context) error {
    idStr := c.Param("id")
    idNum, err := strconv.Atoi(idStr)

    if err != nil { return c.JSON(http.StatusBadRequest, invalidProductID) }


    var product models.Product

    
    result := controller.DB.First(&product, idNum)
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            return c.JSON(http.StatusNotFound, "Nie znaleziono produktu")
        }
        return c.JSON(http.StatusInternalServerError, "Błąd serwera")
    }

    return c.JSON(http.StatusOK, product)
}

func (controller *ProductController) AddOneDB(c echo.Context) error {
    var newItem models.Product

    err := c.Bind(&newItem) // fill obj based on request bodby
    if err != nil { return c.String(http.StatusBadRequest, "Nie podano prawidłowych parametrów przedmiotu") }

    if newItem.Name == "" || newItem.Icon == "" || newItem.Price == 0 || newItem.Quantity == 0 {
        return c.String(http.StatusBadRequest, "Nie wszystkie informacje o przedmiocie zostały podane")
    }

    result := controller.DB.Create(&newItem)
    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, "Błąd serwera")
    }

    return c.JSON(http.StatusOK, "przedmiot dodany pomyślnie")
}

func (controller *ProductController) DeleteOneDB(c echo.Context) error {
    idStr := c.Param("id")
    idNum, err := strconv.Atoi(idStr)

    if err != nil { return c.JSON(http.StatusBadRequest, invalidProductID) }


    result := controller.DB.Delete(&models.Product{}, idNum)

    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, "Błąd serwera")
    }

    if result.RowsAffected == 0 {
        return c.JSON(http.StatusNotFound, "Nie znaleziono produktu jaki miał zostać usunięty")
    }

    return c.JSON(http.StatusOK, "Produkt został usunięty")
}

func (controller *ProductController) PatchOneDB(c echo.Context) error {
    idStr := c.Param("id")
    idNum, err := strconv.Atoi(idStr)

    if err != nil { return c.JSON(http.StatusBadRequest, invalidProductID) }


    var updateProduct models.Product

    err2 := c.Bind(&updateProduct) // fill obj based on request bodby
    if err2 != nil { return c.String(http.StatusBadRequest, "Nie podano prawidłowych parametrów przedmiotu") }

 
    result := controller.DB.Model(&models.Product{}).Where("id = ?", idNum).Updates(updateProduct)

    if result.Error != nil {
        return c.JSON(http.StatusInternalServerError, invalidProductID)
    }
    
    if result.RowsAffected == 0 {
        return c.JSON(http.StatusNotFound, "Produkt do aktualizacji nie został znaleziony")
    }

    return c.JSON(http.StatusNotFound, "przedmiot został zaktualizowany")
}