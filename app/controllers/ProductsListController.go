package controllers

import (
    "net/http"
	"github.com/labstack/echo/v4"
    "strconv"
    "zadanie_4_ebiz.com/app/models"
)

const invalidProductID = "nieprawidłowe ID produktu"

func GetAll(c echo.Context) error {
    return c.JSON(http.StatusOK, models.ItemList)

    // return c.String(http.StatusOK, "Hello, World!")
}

func GetOne(c echo.Context) error {
    idStr := c.Param("id")
    idNum, err := strconv.Atoi(idStr)

    if err != nil { return c.JSON(http.StatusBadRequest, invalidProductID) }

    for _, item := range models.ItemList {
        if item.ID == idNum { return c.JSON(http.StatusOK, item) }
    }

    return c.JSON(http.StatusNotFound, "Nie znaleziono produktu")
}

func AddOne(c echo.Context) error {

    var newItem models.Item

    err := c.Bind(&newItem) // fill obj based on request bodby
    if err != nil { return c.String(http.StatusBadRequest, "Nie podano prawidłowych parametrów przedmiotu") }

    if newItem.Name == "" || newItem.Icon == "" || newItem.ID == 0 || newItem.Price == 0 || newItem.Quantity == 0 {
        return c.String(http.StatusBadRequest, "Nie wszystkie informacje o przedmiocie zostały podane")
    }

    for _, item := range models.ItemList {
        if item.ID == newItem.ID { return c.JSON(http.StatusBadRequest, "przedmiot o podanym ID juz istnieje") }
    }

    models.ItemList = append(models.ItemList, newItem)

    return c.JSON(http.StatusOK, "przedmiot dodany pomyślnie")
}

func DeleteOne(c echo.Context) error {
    idStr := c.Param("id")
    idNum, err := strconv.Atoi(idStr)

    if err != nil { return c.JSON(http.StatusBadRequest, invalidProductID) }


    for i, item := range models.ItemList {
        if item.ID == idNum {
            models.ItemList = append(models.ItemList[:i], models.ItemList[i+1:]...)
    
            return c.JSON(http.StatusOK, "usunięto przedmiot z listy")
        }
    }

    return c.JSON(http.StatusNotFound, "przedmiot do usunięcia nie został znaleziony na liście")
}

func PatchOne(c echo.Context) error {
    idStr := c.Param("id")
    idNum, err := strconv.Atoi(idStr)

    if err != nil { return c.JSON(http.StatusBadRequest, invalidProductID) }

    var updatedItem models.Item

    err2 := c.Bind(&updatedItem) // fill obj based on request bodby
    if err2 != nil { return c.String(http.StatusBadRequest, "Nie podano prawidłowych parametrów przedmiotu") }


    for i, item := range models.ItemList {
        if item.ID == idNum {
            updatedItem.ID = idNum
            models.ItemList[i] = updatedItem
            return c.JSON(http.StatusOK, "przedmiot został zaktualizowany")
        }
    }

    return c.JSON(http.StatusNotFound, "przedmiot do aktualizacji nie został znaleziony")
}

