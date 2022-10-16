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

type SocialMediaController struct {
	SocialMediaSvc services.SocialMediaSvc
}

var ISocialMediaSvc interfaces.ISocialMediaRepo

func (c *SocialMediaController) AddSocialMedia(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	socialMedia := models.SocialMedia{}

	// Bind JSON or form data
	if contentType == helpers.GetConstant().AppJSON {
		ctx.ShouldBindJSON(&socialMedia)
	} else {
		ctx.ShouldBind(&socialMedia)
	}

	// Validating request data
	if socialMedia.Name == "" || socialMedia.SocialMediaUrl == "" {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"name", "social_media_url"},
			Message: "name and social_media_url are required",
			Extends: nil,
		})
		return
	}
	if !helpers.ValidateURL(socialMedia.SocialMediaUrl) {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"social_media_url"},
			Message: "social_media_url is not valid URL",
			Extends: nil,
		})
		return
	}

	socialMedia.UserID = helpers.GetUserID(ctx)

	ISocialMediaSvc = &c.SocialMediaSvc
	err := ISocialMediaSvc.AddSocialMedia(&socialMedia)
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	helpers.Response(ctx, http.StatusCreated, gin.H{
		"id":               socialMedia.ID,
		"name":             socialMedia.Name,
		"social_media_url": socialMedia.SocialMediaUrl,
		"user_id":          socialMedia.UserID,
		"created_at":       socialMedia.CreatedAt,
	}, "SUCCESS", nil)
}

func (c *SocialMediaController) GetSocialMedias(ctx *gin.Context) {
	ISocialMediaSvc = &c.SocialMediaSvc
	data, err := ISocialMediaSvc.GetSocialMedias()
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	var socialMedias []models.SocialMediaResponse
	for _, socialMedia := range data {
		socialMedias = append(socialMedias, models.SocialMediaResponse{
			ID:             socialMedia.ID,
			Name:           socialMedia.Name,
			SocialMediaUrl: socialMedia.SocialMediaUrl,
			UserID:         socialMedia.User.ID,
			CreatedAt:      socialMedia.CreatedAt.Format("2006-10-16T17:12:09.331424+07:00"),
			UpdatedAt:      socialMedia.UpdatedAt.Format("2006-10-16T17:12:09.331424+07:00"),
			User: struct {
				ID       uint   `json:"id"`
				Username string `json:"username"`
			}{
				ID:       socialMedia.User.ID,
				Username: socialMedia.User.Username,
			},
		})
	}

	helpers.Response(ctx, http.StatusOK, gin.H{
		"social_medias": socialMedias,
	}, "SUCCESS", nil)
}

func (c *SocialMediaController) UpdateSocialMedia(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	socialMedia := models.SocialMedia{}

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
		ctx.ShouldBindJSON(&socialMedia)
	} else {
		ctx.ShouldBind(&socialMedia)
	}

	// Validating request data
	if socialMedia.Name == "" || socialMedia.SocialMediaUrl == "" {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"name", "social_media_url"},
			Message: "name and social_media_url are required",
			Extends: nil,
		})
		return
	}
	if !helpers.ValidateURL(socialMedia.SocialMediaUrl) {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"social_media_url"},
			Message: "social_media_url is not valid URL",
			Extends: nil,
		})
		return
	}

	// Check if social media exists
	ISocialMediaSvc = &c.SocialMediaSvc
	socialMediaData, err := ISocialMediaSvc.GetSocialMediaByID(QueryID)
	if err != nil {
		// if social media not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(ctx, http.StatusNotFound, nil, "NOT_FOUND", &models.Error{
				Fields:  nil,
				Message: "social media not found",
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

	// Check if user is the owner of the social media
	if socialMediaData.User.ID != helpers.GetUserID(ctx) {
		helpers.Response(ctx, http.StatusForbidden, nil, "FORBIDDEN", &models.Error{
			Fields:  nil,
			Message: "you are not the owner of this social media",
			Extends: nil,
		})
		return
	}

	// Update social media
	data, err := ISocialMediaSvc.UpdateSocialMedia(QueryID, &socialMedia)
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	helpers.Response(ctx, http.StatusOK, gin.H{
		"id":               data.ID,
		"name":             data.Name,
		"social_media_url": data.SocialMediaUrl,
		"user_id":          data.User.ID,
		"updated_at":       data.UpdatedAt,
	}, "SUCCESS", nil)
}

func (c *SocialMediaController) DeleteSocialMedia(ctx *gin.Context) {
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

	// Check if social media exists
	ISocialMediaSvc = &c.SocialMediaSvc
	socialMediaData, err := ISocialMediaSvc.GetSocialMediaByID(QueryID)
	if err != nil {
		// if social media not exist
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(ctx, http.StatusNotFound, nil, "NOT_FOUND", &models.Error{
				Fields:  nil,
				Message: "social media not found",
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

	// Check if user is the owner of the social media
	if socialMediaData.User.ID != helpers.GetUserID(ctx) {
		helpers.Response(ctx, http.StatusForbidden, nil, "FORBIDDEN", &models.Error{
			Fields:  nil,
			Message: "you are not the owner of this social media",
			Extends: nil,
		})
		return
	}

	// Delete social media
	err = ISocialMediaSvc.DeleteSocialMedia(QueryID)
	if err != nil {
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	helpers.Response(ctx, http.StatusOK, gin.H{
		"message": "Your social media has been successfully deleted",
	}, "SUCCESS", nil)
}
