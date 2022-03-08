package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/indeedhat/todo/internal/store"
	"gorm.io/gorm"
)

type Entries struct {
	db *gorm.DB
}

// NewEntries controller
func NewEntries(router *gin.Engine, db *gorm.DB) Entries {
	entries := Entries{db}

	entries.routes(router)

	return entries
}

// routes sets up all the routes related to entries
func (ent Entries) routes(router *gin.Engine) {
	lists := router.Group("/lists/:slug/entries")
	{
		lists.GET("", ent.List())
		lists.POST("", ent.Create())
		lists.PATCH("/:id", ent.Update())
		lists.DELETE("/:id", ent.Delete())
		lists.POST("/:id/toggle", ent.Toggle())
	}
}

// List teh entries on a list
func (ent Entries) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			slug = ctx.Param("slug")
			list = store.FindListBySlug(ent.db, slug)
		)

		if list == nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, errNotFound)
			return
		}

		entries := store.ListEntries(ent.db, list.ID)

		ctx.JSON(http.StatusOK, gin.H{
			"outcome": true,
			"message": nil,
            "list": list,
			"entries": entries,
		})
	}
}

// Create a new entry on a liste
func (ent Entries) Create() gin.HandlerFunc {
	type formInput struct {
		Text string `form:"text" binding:"required"`
	}

	return func(ctx *gin.Context) {
		var (
			input formInput
			slug  = ctx.Param("slug")
		)

		if err := ctx.ShouldBind(&input); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errBadInput)
			return
		}

		list := store.FindListBySlug(ent.db, slug)
		if list != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errSlugInUse)
			return
		}

		if list := store.CreateEntry(ent.db, list.ID, input.Text); list != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errRequestFailed)
			return
		}

		ctx.JSON(http.StatusOK, errNone)
	}
}

// Update the text on an entry
func (ent Entries) Update() gin.HandlerFunc {
	type formInput struct {
		Text string `form:"text" binding:"required"`
	}

	return func(ctx *gin.Context) {
		var (
			input    formInput
			idString = ctx.Param("id")
		)

		if err := ctx.ShouldBind(&input); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errBadInput)
			return
		}

		id, err := strconv.Atoi(idString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errBadInput)
			return
		}

		existing := store.FindEntriy(ent.db, uint(id))
		if existing == nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, errNotFound)
			return
		}

		existing.Text = input.Text
		if tx := ent.db.Save(existing); tx.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errRequestFailed)
			return
		}

		ctx.JSON(http.StatusOK, errNone)
	}
}

// Toggle the done state on an entry
func (ent Entries) Toggle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var idString = ctx.Param("id")

		id, err := strconv.Atoi(idString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errBadInput)
			return
		}

		existing := store.FindEntriy(ent.db, uint(id))
		if existing == nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, errNotFound)
			return
		}

		existing.Done = !existing.Done
		if tx := ent.db.Save(existing); tx.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errRequestFailed)
			return
		}

		ctx.JSON(http.StatusOK, errNone)
	}
}

// Delete an entry
func (ent Entries) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var idString = ctx.Param("id")

		id, err := strconv.Atoi(idString)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errBadInput)
			return
		}

		existing := store.FindEntriy(ent.db, uint(id))
		if existing == nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, errNotFound)
			return
		}

		if err := ent.db.Delete(existing).Error; err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errRequestFailed)
			return
		}

		ctx.JSON(http.StatusOK, errNone)
	}
}
