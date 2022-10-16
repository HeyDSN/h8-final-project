package controllers

import (
	"errors"
	"final-project/helpers"
	"final-project/interfaces"
	"final-project/models"
	"final-project/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CommentController struct {
	CommentSvc services.CommentSvc
}

var ICommentSvc interfaces.ICommentRepo

func (c *CommentController) PostComment(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	comment := models.Comment{}

	// Bind JSON or form data
	if contentType == helpers.GetConstant().AppJSON {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	// Validating request data
	if comment.Message == "" || comment.PhotoID == 0 {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"title", "comment_url"},
			Message: "message and photo_id are required",
			Extends: nil,
		})
		return
	}
	comment.UserID = helpers.GetUserID(ctx)

	ICommentSvc = &c.CommentSvc
	err := ICommentSvc.PostComment(&comment)
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}
	helpers.Response(ctx, http.StatusCreated, gin.H{
		"id":         comment.ID,
		"message":    comment.Message,
		"photo_id":   comment.PhotoID,
		"user_id":    comment.UserID,
		"created_at": comment.CreatedAt,
	}, "OK", nil)
}

func (c *CommentController) GetComments(ctx *gin.Context) {
	ICommentSvc = &c.CommentSvc
	data, err := ICommentSvc.GetComments()
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	var comments []models.CommentResponse
	for _, comment := range data {
		comments = append(comments, models.CommentResponse{
			ID:        comment.ID,
			Message:   comment.Message,
			PhotoID:   comment.PhotoID,
			UserID:    comment.UserID,
			UpdatedAt: comment.UpdatedAt.Format("2006-10-16T17:12:09.331424+07:00"),
			CreatedAt: comment.CreatedAt.Format("2006-10-16T17:12:09.331424+07:00"),
			User: struct {
				ID       uint   `json:"id"`
				Email    string `json:"email"`
				Username string `json:"username"`
			}{
				ID:       comment.User.ID,
				Email:    comment.User.Email,
				Username: comment.User.Username,
			},
			Photo: struct {
				ID       uint   `json:"id"`
				Title    string `json:"title"`
				Caption  string `json:"caption"`
				PhotoURL string `json:"photo_url"`
				UserID   uint   `json:"user_id"`
			}{
				ID:       comment.Photo.ID,
				Title:    comment.Photo.Title,
				Caption:  comment.Photo.Caption,
				PhotoURL: comment.Photo.PhotoURL,
				UserID:   comment.Photo.UserID,
			},
		})
	}

	helpers.Response(ctx, http.StatusOK, comments, "OK", nil)
}

func (c *CommentController) UpdateComment(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	comment := models.Comment{}

	QueryIDString := ctx.Param("id")
	QueryID, err := helpers.StringToUint(QueryIDString)
	if err != nil {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"id"},
			Message: "id must be a number",
			Extends: nil,
		})
		return
	}

	// Bind JSON or form data
	if contentType == helpers.GetConstant().AppJSON {
		ctx.ShouldBindJSON(&comment)
	} else {
		ctx.ShouldBind(&comment)
	}

	// Validating request data
	if comment.Message == "" {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"title", "comment_url"},
			Message: "message is required",
			Extends: nil,
		})
		return
	}

	// Check if comment exist
	ICommentSvc = &c.CommentSvc
	commentExist, err := ICommentSvc.GetCommentByID(QueryID)
	if err != nil {
		// if comment not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(ctx, http.StatusNotFound, nil, "NOT_FOUND", &models.Error{
				Fields:  nil,
				Message: "comment not found",
				Extends: nil,
			})
			return
		}
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	// Check if user is owner of comment
	if commentExist.UserID != helpers.GetUserID(ctx) {
		helpers.Response(ctx, http.StatusForbidden, nil, "FORBIDDEN", &models.Error{
			Fields:  nil,
			Message: "you are not owner of this comment",
			Extends: nil,
		})
		return
	}

	data, err := ICommentSvc.UpdateComment(QueryID, &comment)
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	helpers.Response(ctx, http.StatusOK, gin.H{
		"id":         data.ID,
		"message":    data.Message,
		"photo_id":   data.PhotoID,
		"user_id":    data.User.ID,
		"created_at": data.CreatedAt,
		"updated_at": data.UpdatedAt,
	}, "OK", nil)
}

func (c *CommentController) DeleteComment(ctx *gin.Context) {
	QueryIDString := ctx.Param("id")
	QueryID, err := helpers.StringToUint(QueryIDString)
	if err != nil {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"id"},
			Message: "id must be a number",
			Extends: nil,
		})
		return
	}

	// Check if comment exist
	ICommentSvc = &c.CommentSvc
	commentExist, err := ICommentSvc.GetCommentByID(QueryID)
	if err != nil {
		// if comment not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(ctx, http.StatusNotFound, nil, "NOT_FOUND", &models.Error{
				Fields:  nil,
				Message: "comment not found",
				Extends: nil,
			})
			return
		}
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	// Check if user is owner of comment
	if commentExist.UserID != helpers.GetUserID(ctx) {
		helpers.Response(ctx, http.StatusForbidden, nil, "FORBIDDEN", &models.Error{
			Fields:  nil,
			Message: "you are not owner of this comment",
			Extends: nil,
		})
		return
	}

	err = ICommentSvc.DeleteComment(QueryID)
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	helpers.Response(ctx, http.StatusOK, gin.H{
		"message": "Your comment has been successfully deleted",
	}, "OK", nil)
}
