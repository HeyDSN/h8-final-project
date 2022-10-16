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

type PhotoController struct {
	PhotoSvc services.PhotoSvc
}

var IPhotoSvc interfaces.IPhotoRepo

func (c *PhotoController) PostPhoto(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	photo := models.Photo{}

	// Bind JSON or form data
	if contentType == helpers.GetConstant().AppJSON {
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	// Validating request data
	if photo.Title == "" || photo.PhotoURL == "" {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"title", "photo_url"},
			Message: "title and photo_url are required",
			Extends: nil,
		})
		return
	}
	if !helpers.ValidateURL(photo.PhotoURL) {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"photo_url"},
			Message: "photo_url is not valid URL",
			Extends: nil,
		})
		return
	}

	photo.UserID = helpers.GetUserID(ctx)

	IPhotoSvc = &c.PhotoSvc
	err := IPhotoSvc.PostPhoto(&photo)
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}
	helpers.Response(ctx, http.StatusCreated, gin.H{
		"id":         photo.ID,
		"title":      photo.Title,
		"caption":    photo.Caption,
		"photo_url":  photo.PhotoURL,
		"user_id":    photo.UserID,
		"created_at": photo.CreatedAt,
	}, "OK", nil)
}

func (c *PhotoController) GetPhotos(ctx *gin.Context) {
	IPhotoSvc = &c.PhotoSvc
	data, err := IPhotoSvc.GetPhotos()
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	var photos []models.PhotoResponse
	for _, photo := range data {
		photos = append(photos, models.PhotoResponse{
			ID:        photo.ID,
			Title:     photo.Title,
			Caption:   photo.Caption,
			PhotoURL:  photo.PhotoURL,
			UserID:    photo.UserID,
			CreatedAt: photo.CreatedAt.Format("2006-10-16T17:12:09.331424+07:00"),
			UpdatedAt: photo.UpdatedAt.Format("2006-10-16T17:12:09.331424+07:00"),
			User: struct {
				Email    string `json:"email"`
				Username string `json:"username"`
			}{
				Email:    photo.User.Email,
				Username: photo.User.Username,
			},
		})
	}

	helpers.Response(ctx, http.StatusOK, photos, "OK", nil)
}

func (c *PhotoController) UpdatePhoto(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	photo := models.Photo{}

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
		ctx.ShouldBindJSON(&photo)
	} else {
		ctx.ShouldBind(&photo)
	}

	// Validating request data
	if photo.Title == "" || photo.PhotoURL == "" {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"title", "photo_url"},
			Message: "all title and photo_url are required",
			Extends: nil,
		})
		return
	}
	if !helpers.ValidateURL(photo.PhotoURL) {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"photo_url"},
			Message: "photo_url is not valid URL",
			Extends: nil,
		})
		return
	}

	// Check if photo exist
	IPhotoSvc = &c.PhotoSvc
	photoExist, err := IPhotoSvc.GetPhotoByID(QueryID)
	if err != nil {
		// if photo not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(ctx, http.StatusNotFound, nil, "NOT_FOUND", &models.Error{
				Fields:  nil,
				Message: "photo not found",
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

	// Check if user is owner of photo
	if photoExist.UserID != helpers.GetUserID(ctx) {
		helpers.Response(ctx, http.StatusForbidden, nil, "FORBIDDEN", &models.Error{
			Fields:  nil,
			Message: "you are not owner of this photo",
			Extends: nil,
		})
		return
	}

	data, err := IPhotoSvc.UpdatePhoto(QueryID, &photo)
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
		"title":      data.Title,
		"caption":    data.Caption,
		"photo_url":  data.PhotoURL,
		"user_id":    data.UserID,
		"updated_at": data.UpdatedAt,
	}, "OK", nil)
}

func (c *PhotoController) DeletePhoto(ctx *gin.Context) {
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

	// Check if photo exist
	IPhotoSvc = &c.PhotoSvc
	photoExist, err := IPhotoSvc.GetPhotoByID(QueryID)
	if err != nil {
		// if photo not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(ctx, http.StatusNotFound, nil, "NOT_FOUND", &models.Error{
				Fields:  nil,
				Message: "photo not found",
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

	// Check if user is owner of photo
	if photoExist.UserID != helpers.GetUserID(ctx) {
		helpers.Response(ctx, http.StatusForbidden, nil, "FORBIDDEN", &models.Error{
			Fields:  nil,
			Message: "you are not owner of this photo",
			Extends: nil,
		})
		return
	}

	err = IPhotoSvc.DeletePhoto(QueryID)
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	helpers.Response(ctx, http.StatusOK, gin.H{
		"message": "Your photo has been successfully deleted",
	}, "OK", nil)
}
