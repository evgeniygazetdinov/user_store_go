package chats

import (
	"context"
	"work_in_que/logging"
	"work_in_que/user"

	// "fmt"
	"time"
	// "common"
	// "github.com/google/uuid"
)

type Services struct {
	logger          logging.Logger
	userRepository  user.Repository
	chatsRepository Repository
}

func NewInstanceOfchatsServices(logger logging.Logger, userRepository user.Repository, chatsRepository Repository) Services {
	return Services{logger, userRepository, chatsRepository}
}

func (c *Services) GetAll(ctx context.Context, session user.Session, query ListCarQuery) ([]Car, error) {
	ctx = context.WithValue(ctx, logging.CtxServiceMethod, "GetAll")

	chats, err := c.chatsRepository.List(session.Email, query)
	if err != nil {
		return []Car{}, err
	}
	return chats, nil
}

func (c *Services) GetByID(ctx context.Context, session user.Session, carID string) (Car, error) {
	ctx = context.WithValue(ctx, logging.CtxServiceMethod, "GetByID")

	car, err := c.chatsRepository.Get(session.Email, carID)
	if err != nil {
		return Car{}, err
	}
	return car, nil
}

func (c *Services) Create(ctx context.Context, session user.Session, body CreateCar) error {
	ctx = context.WithValue(ctx, logging.CtxServiceMethod, "Create")
	car := Car{
		Make:    body.Make,
		Model:   body.Model,
		Year:    body.Year,
		Status:  "",
		Created: time.Now(),
		Email:   session.Email,
	}
	err := c.chatsRepository.Save(car)
	if err != nil {
		return err
	}
	return nil
}

func (c *Services) Update(ctx context.Context, session user.Session, carID string, body UpdateCar) error {
	ctx = context.WithValue(ctx, logging.CtxServiceMethod, "Update")
	// Update car
	err := c.chatsRepository.Update(session.Email, carID, body)
	if err != nil {
		return err
	}
	return nil
}

func (c *Services) Delete(ctx context.Context, session user.Session, carID string) error {
	ctx = context.WithValue(ctx, logging.CtxServiceMethod, "Delete")

	// Delete car
	err := c.chatsRepository.Delete(session.Email, carID)
	if err != nil {
		return err
	}
	return nil
}
