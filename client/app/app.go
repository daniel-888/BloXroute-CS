package app

import (
	"github.com/daniel-888/BloXroute-CS/models"
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
	"strconv"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

const traceIDkey = "X-Trace-ID"

type App struct {
	client *Client
}

func NewApp(client *Client) *App {
	return &App{
		client: client,
	}
}

func (a *App) Start(commandType models.CommandType, args []string) error {
	rand.Seed(time.Now().Unix())

	for {
		traceID := uuid.New().String()
		ctx := context.WithValue(context.Background(), traceIDkey, traceID)

		command, err := createRandomCommand(commandType, args)
		if err != nil {
			return err
		}

		err = a.client.SendCommand(ctx, command)
		if err != nil {
			log.WithField(traceIDkey, ctx.Value(traceIDkey)).Errorf("Fail send command: %v", err)
		}

		// wait some time to send command again
		secondsWait := rand.Int63n(5)
		time.Sleep(time.Duration(secondsWait) * time.Second)
	}
}

func createRandomCommand(commandType models.CommandType, args []string) (*models.Command, error) {
	switch commandType {
	case models.CommandType_AddItem:
		itemId := rand.Int63()
		randomChar := rand.Intn(25) + 'A' // random character from A to Z

		return &models.Command{
			Type:        commandType,
			ItemID:      itemId,
			ItemPayload: fmt.Sprintf("%c", randomChar),
		}, nil
	case models.CommandType_GetItem:
		var itemId int64
		if len(args) > 1 {
			return nil, errors.New("command argument too much")
		}
		
		if len(args) == 0 {
			itemId = rand.Int63()
		} else {
			n, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return nil, errors.New("command argument invalid")
			}
			itemId = n
		}

		return &models.Command{
			Type:   commandType,
			ItemID: itemId,
		}, nil
	case models.CommandType_RemoveItem:
		var itemId int64
		if len(args) > 1 {
			return nil, errors.New("command argument too much")
		}
		
		if len(args) == 0 {
			itemId = rand.Int63()
		} else {
			n, err := strconv.ParseInt(args[0], 10, 64)
			if err != nil {
				return nil, errors.New("command argument invalid")
			}
			itemId = n
		}

		return &models.Command{
			Type:   commandType,
			ItemID: itemId,
		}, nil
	case models.CommandType_GetAllItems:
		return &models.Command{
			Type: commandType,
		}, nil
	default:
		return nil, errors.New("command type is unknown")
	}
}