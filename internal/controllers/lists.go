package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/indeedhat/todo/internal/store"
	"gorm.io/gorm"
)

type Lists struct {
	db *gorm.DB
}

// NewLists controller
func NewLists(router *gin.Engine, db *gorm.DB) Lists {
	list := Lists{db}

	list.routes(router)

	return list
}

// routes sets up the routes related to lists
func (lst Lists) routes(router *gin.Engine) {
	lists := router.Group("/lists")
	{
		lists.GET("", lst.List())
		lists.POST("", lst.Create())
		lists.PATCH("/:slug", lst.Update())
		lists.DELETE("/:slug", lst.Delete())
	}
}

// List all the TODO lists
func (lst Lists) List() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lists := store.ListLists(lst.db)

		ctx.JSON(http.StatusOK, gin.H{
			"outcome": true,
			"message": nil,
			"lists":   lists,
		})
	}
}

// Create a new TODO list
func (lst Lists) Create() gin.HandlerFunc {
	type formInput struct {
		Slug  string `form:"slug" binding:"required"`
		Title string `form:"title" binding:"required"`
	}

	return func(ctx *gin.Context) {
		var input formInput

		if err := ctx.ShouldBind(&input); err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errBadInput)
			return
		}

		if existing := store.FindListBySlug(lst.db, input.Slug); existing != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, errSlugInUse)
			return
		}

		if list := store.CreateList(lst.db, input.Title, input.Slug); list != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errRequestFailed)
			return
		}

		ctx.JSON(http.StatusOK, errNone)
	}
}

// Update an existing list
func (lst Lists) Update() gin.HandlerFunc {
	type formInput struct {
		Title string `form:"title" binding:"required"`
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

		existing := store.FindListBySlug(lst.db, slug)
		if existing == nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, errNotFound)
			return
		}

		existing.Title = input.Title
		if tx := lst.db.Save(existing); tx.Error != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errRequestFailed)
			return
		}

		ctx.JSON(http.StatusOK, errNone)
	}
}

// Delete a TODO list
func (lst Lists) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var slug = ctx.Param("slug")

		existing := store.FindListBySlug(lst.db, slug)
		if existing == nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, errNotFound)
			return
		}

		err := lst.db.Transaction(func(tx *gorm.DB) error {
			if err := tx.Delete(existing).Error; err != nil {
				return err
			}

			return tx.Delete(&store.Entry{}, "list_id = ?", existing.ID).Error
		})

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, errRequestFailed)
			return
		}

		ctx.JSON(http.StatusOK, errNone)
	}
}
