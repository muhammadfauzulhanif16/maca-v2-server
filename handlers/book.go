package handlers

import (
	"maca/formatters"
	"maca/inputs"
	"maca/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type bookH struct {
	s services.BookS
}

func NewBookH(s services.BookS) *bookH {
	return &bookH{s}
}

func (h *bookH) Create(c *gin.Context) {
	var input inputs.CreateBook
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})

		return
	}

	data, err := h.s.Create(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, formatters.BookFormat(data))
}

func (h *bookH) ReadAll(c *gin.Context) {
	datas, err := h.s.ReadAll(c.Query("is_completed"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, formatters.BooksFormat(datas))
}

func (h *bookH) ReadSearch(c *gin.Context) {
	var input inputs.SearchBook
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	datas, err := h.s.ReadSearch(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, formatters.BooksFormat(datas))
}

func (h *bookH) UpdateIsCompleted(c *gin.Context) {
	data, err := h.s.UpdateIsCompleted(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, formatters.BookFormat(data))
}

func (h *bookH) Delete(c *gin.Context) {
	data, err := h.s.Delete(c.Params.ByName("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, formatters.BookFormat(data))
}

func (h *bookH) DeleteAll(c *gin.Context) {
	datas, err := h.s.DeleteAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, formatters.BooksFormat(datas))
}
