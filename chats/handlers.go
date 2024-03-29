package chats

import (
	"context"
	"fmt"
	"strconv"
	"work_in_que/logging"
	"work_in_que/user"

	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	// "strings"
)

type Handlers struct {
	logger       logging.Logger
	chatsService Services
}

func NewInstanceOfchatsHandlers(logger logging.Logger, chatsService Services) *Handlers {
	return &Handlers{logger, chatsService}
}

func (u *Handlers) GetSession(c *gin.Context) (user.Session, bool) {
	i, exists := c.Get("session")
	if !exists {
		return user.Session{}, false
	}
	session, ok := i.(user.Session)
	if !ok {
		return user.Session{}, false
	}
	return session, true
}

func (u *Handlers) GetAll(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logging.CtxDomain, "chats")
	ctx = context.WithValue(ctx, logging.CtxHandlerMethod, "GetAll")
	ctx = context.WithValue(ctx, logging.CtxRequestID, uuid.New().String())

	u.logger.Info(ctx, "Called")

	session, exists := u.GetSession(c)
	if !exists {
		c.JSON(403, gin.H{"message": "error: unauthorized"})
		return
	}

	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "25")
	make := c.DefaultQuery("make", "")
	model := c.DefaultQuery("model", "")
	year := c.DefaultQuery("year", "0")

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	query := ListCarQuery{
		Page:  pageInt,
		Limit: limitInt,
		Make:  make,
		Model: model,
		Year:  yearInt,
	}

	v := validator.New()
	if err := v.Struct(query); err != nil {
		fmt.Print("Validation failed.")
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	chats, err := u.chatsService.GetAll(ctx, session, query)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	if chats == nil {
		chats = []Car{}
	}
	c.JSON(200, gin.H{"message": "chats retrieved", "chats": chats})
	return
}

func (u *Handlers) GetByID(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logging.CtxDomain, "chats")
	ctx = context.WithValue(ctx, logging.CtxHandlerMethod, "GetByID")
	ctx = context.WithValue(ctx, logging.CtxRequestID, uuid.New().String())

	chatsID := c.Param("id")

	session, exists := u.GetSession(c)
	if !exists {
		c.JSON(403, gin.H{"message": "error: unauthorized"})
		return
	}

	car, err := u.chatsService.GetByID(ctx, session, chatsID)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Car retrieved", "car": car})
	return
}

func (u *Handlers) Create(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logging.CtxDomain, "chats")
	ctx = context.WithValue(ctx, logging.CtxHandlerMethod, "Create")
	ctx = context.WithValue(ctx, logging.CtxRequestID, uuid.New().String())

	var body CreateCar
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := body.Valid(); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	session, exists := u.GetSession(c)
	if !exists {
		c.JSON(403, gin.H{"message": "error: unauthorized"})
		return
	}

	v := validator.New()
	if err := v.Struct(body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err := u.chatsService.Create(ctx, session, body)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Created car"})
	return
}

func (u *Handlers) Update(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logging.CtxDomain, "chats")
	ctx = context.WithValue(ctx, logging.CtxHandlerMethod, "Update")
	ctx = context.WithValue(ctx, logging.CtxRequestID, uuid.New().String())

	chatsID := c.Param("id")

	var body UpdateCar
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := body.Valid(); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	session, exists := u.GetSession(c)
	if !exists {
		c.JSON(403, gin.H{"message": "error: unauthorized"})
		return
	}

	v := validator.New()
	if err := v.Struct(body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	err := u.chatsService.Update(ctx, session, chatsID, body)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Updated car"})
	return
}

func (u *Handlers) Delete(c *gin.Context) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, logging.CtxDomain, "chats")
	ctx = context.WithValue(ctx, logging.CtxHandlerMethod, "Delete")
	ctx = context.WithValue(ctx, logging.CtxRequestID, uuid.New().String())

	session, exists := u.GetSession(c)
	if !exists {
		c.JSON(403, gin.H{"message": "error: unauthorized"})
		return
	}

	chatsID := c.Param("id")

	err := u.chatsService.Delete(ctx, session, chatsID)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Deleted car"})
	return
}
