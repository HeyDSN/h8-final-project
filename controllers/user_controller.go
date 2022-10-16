package controllers

import (
	"errors"
	"final-project/helpers"
	"final-project/interfaces"
	"final-project/models"
	"final-project/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
	"gorm.io/gorm"
)

type UserController struct {
	UserSvc services.UserSvc
}

var IUserSvc interfaces.IUserRepo

func (c *UserController) UserRegistration(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	user := models.User{}

	// Bind JSON or form data
	if contentType == helpers.GetConstant().AppJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	// Validating request data
	if user.Username == "" || user.Email == "" || user.Password == "" {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"username", "email", "password", "age"},
			Message: "all fields are required",
			Extends: nil,
		})
		return
	}
	if !helpers.ValidateEmail(user.Email) {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"email"},
			Message: "invalid email format",
			Extends: nil,
		})
		return
	}
	if user.Age < 8 {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"age"},
			Message: "minimum age is 8",
			Extends: nil,
		})
		return
	}
	if len(user.Password) < 6 {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"password"},
			Message: "minimum password length is 6",
			Extends: nil,
		})
		return
	}

	// Hashing password
	user.Password = helpers.HashPass(user.Password)

	IUserSvc = &c.UserSvc
	err := IUserSvc.CreateUser(&user)
	if err != nil {
		// If error is duplicate key error
		duplicateEntryError := &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			helpers.Response(ctx, http.StatusBadRequest, nil, "DUPLICATE_ENTRY_ERROR", &models.Error{
				Fields:  []string{"username", "email"},
				Message: "username or email already exists",
				Extends: nil,
			})
			return
		}
		// Others error
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}
	helpers.Response(ctx, http.StatusCreated, gin.H{
		"id":       user.ID,
		"username": user.Username,
		"email":    user.Email,
		"age":      user.Age,
	}, "OK", nil)
}

func (c *UserController) UserLogin(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	user := models.User{}

	if contentType == helpers.GetConstant().AppJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	// Validating request data
	if user.Email == "" || user.Password == "" {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"email", "password"},
			Message: "email & password are required",
			Extends: nil,
		})
		return
	}

	// Get password input from user
	inputPassword := user.Password

	IUserSvc = &c.UserSvc
	err := IUserSvc.GetUserByEmail(&user)
	if err != nil {
		helpers.Response(ctx, http.StatusUnauthorized, nil, "UNAUTHORIZED", &models.Error{
			Fields:  []string{"email", "password"},
			Message: "invalid email or password",
			Extends: nil,
		})
		return
	}

	if !helpers.ComparePass(user.Password, inputPassword) {
		helpers.Response(ctx, http.StatusUnauthorized, nil, "UNAUTHORIZED", &models.Error{
			Fields:  []string{"email", "password"},
			Message: "invalid email or password",
			Extends: nil,
		})
		return
	}

	token := helpers.GenerateToken(user.ID, user.Username, user.Email)

	helpers.Response(ctx, http.StatusOK, gin.H{
		"token": token,
	}, "OK", nil)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	contentType := helpers.GetContentType(ctx)
	user := models.User{}

	QueryIDString := ctx.Query("id")
	QueryID, err := helpers.StringToUint(QueryIDString)
	if err != nil {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"id"},
			Message: "id must be a number",
			Extends: nil,
		})
		return
	}

	if contentType == helpers.GetConstant().AppJSON {
		ctx.ShouldBindJSON(&user)
	} else {
		ctx.ShouldBind(&user)
	}

	// Validating request data
	if user.Username == "" || user.Email == "" {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"username", "email"},
			Message: "all fields are required",
			Extends: nil,
		})
		return
	}
	if !helpers.ValidateEmail(user.Email) {
		helpers.Response(ctx, http.StatusBadRequest, nil, "VALIDATION_ERROR", &models.Error{
			Fields:  []string{"email"},
			Message: "invalid email format",
			Extends: nil,
		})
		return
	}
	if QueryID != helpers.GetUserID(ctx) {
		helpers.Response(ctx, http.StatusUnauthorized, nil, "UNAUTHORIZED", &models.Error{
			Fields:  []string{"id"},
			Message: "your user id not match with query params id",
			Extends: nil,
		})
		return
	}

	IUserSvc = &c.UserSvc
	data, err := IUserSvc.UpdateUser(QueryID, &user)
	if err != nil {
		// check if error gorm not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(ctx, http.StatusNotFound, nil, "NOT_FOUND", &models.Error{
				Fields:  []string{"id"},
				Message: "user not found",
				Extends: nil,
			})
			return
		}
		// If error is duplicate key error
		duplicateEntryError := &pgconn.PgError{Code: "23505"}
		if errors.As(err, &duplicateEntryError) {
			helpers.Response(ctx, http.StatusBadRequest, nil, "DUPLICATE_ENTRY_ERROR", &models.Error{
				Fields:  []string{"username", "email"},
				Message: "username or email you select already exists",
				Extends: nil,
			})
			return
		}
		// Others error
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	helpers.Response(ctx, http.StatusOK, gin.H{
		"id":         data.ID,
		"email":      data.Email,
		"username":   data.Username,
		"age":        data.Age,
		"updated_at": data.UpdatedAt,
	}, "OK", nil)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	ID := uint(userData["id"].(float64))

	IUserSvc = &c.UserSvc
	err := IUserSvc.DeleteUser(ID)
	if err != nil {
		// check if error gorm not found
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helpers.Response(ctx, http.StatusNotFound, nil, "NOT_FOUND", &models.Error{
				Fields:  []string{"id"},
				Message: "user not found",
				Extends: nil,
			})
			return
		}
		// Others error
		helpers.Response(ctx, http.StatusInternalServerError, nil, "INTERNAL_SERVER_ERROR", &models.Error{
			Fields:  nil,
			Message: "internal server error",
			Extends: err,
		})
		return
	}

	helpers.Response(ctx, http.StatusOK, gin.H{
		"message": "Your account has been successfully deleted",
	}, "OK", nil)
}
