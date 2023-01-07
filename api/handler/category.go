package handler

import (
	"app/models"
	"app/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCategory(c *gin.Context) {

	var category models.CreateCategory

	err := c.ShouldBindJSON(&category)
	if err != nil {
		log.Println("error whiling marshal json:", err.Error())
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := storage.InsertCategory(h.db, category)
	if err != nil {
		log.Println("error whiling create category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := storage.GetByIdCategory(h.db, models.CategoryPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling get by id category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) UpdateCategory(c *gin.Context) {
	var category models.UpdateCategory

	err := c.ShouldBindJSON(&category)
	if err != nil {
		log.Println("error while update marshal json: ", err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = storage.UpdateCategory(h.db, category)
	if err != nil {
		log.Println("error while update category: ", err)
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	res, err := storage.GetByIdCategory(h.db, models.CategoryPrimeryKey{Id: category.Id})
	if err != nil {
		log.Println("error whiling update get by id category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(202, res)

}

func (h *Handler) GetByIdCategory(c *gin.Context) {

	id := c.Param("id")

	res, err := storage.GetByIdCategory(h.db, models.CategoryPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling get by id category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) GetListCategory(c *gin.Context) {
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

	res, err := storage.GetListCategory(h.db, models.GetListCategoryRequest{
		Offset: int64(offset),
		Limit:  int64(limit),
	})

	if err != nil {
		log.Println("error whiling get list category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, res)
}

func (h *Handler) DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	err := storage.DeleteCategory(h.db, models.CategoryPrimeryKey{Id: id})
	if err != nil {
		log.Println("error whiling delete  Category:", err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusCreated, "delete Category")
}
