package controllers

import (
	"net/http"
	"strconv"

	"Helio/Tugas1.3/models"
	"github.com/gin-gonic/gin"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/mania25/paginator"
)

const Nmax = 9

func ConvertInteger(str string) bool {
	_, errConvert := strconv.Atoi(str)

	return errConvert == nil
}

func GetBook(c *gin.Context) {

	var result = &models.Result{}

	var book []models.Book
	//connect database
	db, errDB := gorm.Open("mysql", "root:@/book")
	db.LogMode(true)
	defer db.Close()

	if errDB != nil {
		panic("failed to connect database")
	}

	BookNumber := c.Query("page[number]")
	BookSize := c.Query("page[size]")
	BookSort := c.Query("sort")
	BookQuery := c.Query("query")
	BookStatus := c.Query("filter[status]")

	switch BookSort {
	case "" : 
		orderby := []int{"id asc"}
	case "-title" :
		orderby := []string{"title desc"}
	case "title" :
		orderby := []string("title asc")
	case "-price" :
		orderby := []int64("price desc")
	case "price" :
		orderby := []int64("price asc")
	case "-status" :
		orderby := []string("status desc")
	case "status" :
		orderby := []string("status asc")
	}

	if ConvertInteger(BookNumber) && ConvertInteger(BookSize) { //validasi integer atau bukan
		pageNumber, _ := strconv.Atoi(BookNumber)
		pageSize, _ := strconv.Atoi(BookSize)

		Paginator := paginator.Paginator{DB: db, OrderBy: orderby, Page: BookNumber, PerPage: BookSize}

		dataPaginator := Paginator.Paginate(&book, "title LIKE ? AND status = ?", "%"+BookQuery+"%", BookStatus)

		result.Meta.Page = pageNumber
		result.Meta.Size = pageSize
		result.Meta.Count = len(book)
		result.Meta.TotalPages = dataPaginator.TotalPages
		result.Meta.TotalData = dataPaginator.TotalRecords

		result.Code = http.StatusOK
		result.Status = "success"
		result.Data = &book
	} else {
		result.Code = http.StatusBadRequest
		result.Status = "failed"
	}
	c.JSON(result.Code, result)
}
