package service

import (
	"context"
	"errors"
	"os"
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"

	"github.com/daniel-888/BloXroute-CS/models"
	"github.com/daniel-888/BloXroute-CS/server/app/repository"
)

const traceIDKey = "X-Trace-ID"

//go:generate mockgen -package=service -source=itemservice.go -destination=itemservice_mock.go
type ItemService interface {
	ProcessItemCommand(ctx context.Context, command *models.Command) error
}

type itemServiceImpl struct {
	repo repository.Repo
}

func New(repo repository.Repo) ItemService {
	return &itemServiceImpl{repo: repo}
}

func (i *itemServiceImpl) ProcessItemCommand(ctx context.Context, command *models.Command) error {
	// open a file
	f, err := os.OpenFile("testlogrus.log", os.O_APPEND | os.O_CREATE | os.O_RDWR, 0666)
	if err != nil {
			fmt.Printf("error opening file: %v", err)
	}
	// defer f.Close()
	log.SetOutput(io.MultiWriter(f, os.Stdout))
	log.SetFormatter(&log.TextFormatter{})

	log.WithField(traceIDKey, ctx.Value(traceIDKey)).
		Info("Start processing command: ", command.String())

	switch command.Type {
	case models.CommandType_AddItem:
		err := i.repo.AddItem(models.Item{
			ID:      command.ItemID,
			Payload: command.ItemPayload,
		})
		if err != nil {
			return err
		}
		log.WithField(traceIDKey, ctx.Value(traceIDKey)).Info("Item was added successfully.")

		return nil
	case models.CommandType_RemoveItem:
		err := i.repo.RemoveItem(command.ItemID)
		if err != nil {
			return err
		}
		log.WithField(traceIDKey, ctx.Value(traceIDKey)).Info("Item was removed successfully.")

		return nil
	case models.CommandType_GetItem:
		item, err := i.repo.GetItem(command.ItemID)
		if err != nil {
			return err
		}
		if item == (models.Item{}) {
			log.WithField(traceIDKey, ctx.Value(traceIDKey)).Info("Item was not found with such id.")
			return nil
		}
		log.WithField(traceIDKey, ctx.Value(traceIDKey)).Info("Item was retrieved successfully.")
		log.WithField(traceIDKey, ctx.Value(traceIDKey)).Info(item)

		return nil
	case models.CommandType_GetAllItems:
		items, err := i.repo.GetAllItems()
		if err != nil {
			return err
		}
		log.WithField(traceIDKey, ctx.Value(traceIDKey)).Info("Get all items")
		log.WithField(traceIDKey, ctx.Value(traceIDKey)).Info(items)

		return nil
	default:
		return errors.New("unknown command type")
	}
}
