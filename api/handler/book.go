package handler

import (
	"app/models"
	"app/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateBook godoc
// @ID create_book
// @Router /book [POST]
// @Summary Create Book
// @Description Create Book
// @Tags Book
// @Accept json
// @Produce json
// @Param book body models.CreateBook true "CreateBookRequestBody"
// @Success 201 {object} models.CreateBook "GetBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) CreateBook(c *gin.Context) {

	var book models.CreateBook

	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.InsertBook(h.db, book)
	if err != nil {
		log.Println("error whiling create book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := storage.GetByIdBook(h.db, models.BookPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling get by id book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetByIdBook godoc
// @ID get_by_id_book
// @Router /book/{id} [GET]
// @Summary Get By Id Book
// @Description Get By Id Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Book "GetIdBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetByIDBook(c *gin.Context) {

	id := c.Param("id")

	res, err := storage.GetByIdBook(h.db, models.BookPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling get by id book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// GetListBook godoc
// @ID get_list_book
// @Router /book [GET]
// @Summary Get List Book
// @Description Get List Book
// @Tags Book
// @Accept json
// @Produce json
// @Param offset query int false "offset"
// @Param limit query int false "limit"
// @Success 200 {object} models.GetListBookResponse "GetListBook"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) GetListBook(c *gin.Context) {
	var (
		err       error
		offset    int
		limit     int
		offsetStr = c.Query("offset")
		limitStr  = c.Query("limit")
	)

	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			log.Println("error whiling offset:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil {
			log.Println("error whiling limit:", err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
	}

	res, err := storage.GetListBook(h.db, models.GetListBookRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

// UpdateBook godoc
// @ID update_book
// @Router /book/{id} [PUT]
// @Summary Update Book
// @Description Update Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param book body models.UpdateBookSwag true "UpdateBookRequestBody"
// @Success 200 {object} models.UpdateBook "UpdateBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) UpdateBook(c *gin.Context) {
	var book models.UpdateBook

	book.Id = c.Param("id")

	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Println("error while update marshal json: ", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = storage.UpdateBook(h.db, book)
	if err != nil {
		log.Println("error while update book: ", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := storage.GetByIdBook(h.db, models.BookPrimeryKey{Id: book.Id})
	if err != nil {
		log.Println("error whiling update get by id book:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(202, res)

}

// DeleteBook godoc
// @ID delete_book
// @Router /book/{id} [DELETE]
// @Summary Delete Book
// @Description Delete Book
// @Tags Book
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} models.Empty "DeleteBookBody"
// @Response 400 {object} string "Invalid Argumant"
// @Failure 500 {object} string "Server error"
func (h *Handler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := storage.DeleteBook(h.db, models.BookPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  Category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, "delete Category")
}
