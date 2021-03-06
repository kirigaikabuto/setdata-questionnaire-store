package setdata_questionnaire_store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"strconv"
)

type TelegramService interface {
	CreateTelegramBot(cmd *CreateTelegramBotCommand) (*TelegramBot, error)
	GetTelegramBot(cmd *GetTelegramBotCommand) (*TelegramBot, error)
	ListTelegramBot(cmd *ListTelegramBotCommand) ([]TelegramBot, error)
	DeleteTelegramBot(cmd *DeleteTelegramBotCommand) error
	SendMessage(cmd *SendMessageCommand) error
}

type telegramService struct {
	store             TelegramStore
	chatIdStore       ChatIdStore
	defaultTelegramId string
}

func NewTelegramService(id string, s TelegramStore, c ChatIdStore) TelegramService {
	return &telegramService{defaultTelegramId: id, store: s, chatIdStore: c}
}

func (t *telegramService) CreateTelegramBot(cmd *CreateTelegramBotCommand) (*TelegramBot, error) {
	telegramBot := &TelegramBot{Id: uuid.New().String()}
	telegramBot.Name = cmd.Name
	telegramBot.AccessToken = cmd.AccessToken
	return t.store.Create(telegramBot)
}

func (t *telegramService) GetTelegramBot(cmd *GetTelegramBotCommand) (*TelegramBot, error) {
	return t.store.Get(cmd.Id)
}

func (t *telegramService) ListTelegramBot(cmd *ListTelegramBotCommand) ([]TelegramBot, error) {
	return t.store.List()
}

func (t *telegramService) DeleteTelegramBot(cmd *DeleteTelegramBotCommand) error {
	return t.store.Delete(cmd.Id)
}

func (t *telegramService) SendMessage(cmd *SendMessageCommand) error {
	telegramBotIDs := []TelegramBot{}
	if cmd.TelegramBoId == "" && t.defaultTelegramId == "" {
		telegrams, err := t.ListTelegramBot(&ListTelegramBotCommand{})
		if err != nil {
			return err
		}
		for _, v := range telegrams {
			telegramBotIDs = append(telegramBotIDs, v)
		}
	}
	if cmd.TelegramBoId == "" {
		defaultTelegram, err := t.GetTelegramBot(&GetTelegramBotCommand{Id: t.defaultTelegramId})
		if err != nil {
			return err
		}
		telegramBotIDs = append(telegramBotIDs, *defaultTelegram)
	}

	for _, v := range telegramBotIDs {
		difference := []string{}
		currentChatIds := GetTelegramChatIds(v.AccessToken)
		dbChatIds, err := t.chatIdStore.List(v.Id)
		if err != nil {
			return err
		}
		for _, c := range currentChatIds {
			isNotExist := true
			for _, v := range dbChatIds {
				if c == v.Value {
					isNotExist = false
					break
				}
			}
			if isNotExist {
				difference = append(difference, c)
			}
		}
		for _, d := range difference {
			_, err := t.chatIdStore.Create(&ChatId{
				Id:            uuid.New().String(),
				TelegramBotId: v.Id,
				Value:         d,
			})
			if err != nil {
				return err
			}
		}
		dbChatIds, err = t.chatIdStore.List(v.Id)
		if err != nil {
			return err
		}
		for _, chID := range dbChatIds {
			client := http.Client{}
			baseUrl := "https://api.telegram.org/bot%s/%s"
			sendMessageURl := fmt.Sprintf(baseUrl, v.AccessToken, "sendMessage")
			jsonData, err := json.Marshal(SendTelegramMessage{
				ChatId:    chID.Value,
				ParseMode: cmd.ParseMode,
				Text:      cmd.Message,
			})
			_, err = client.Post(sendMessageURl, "application/json", bytes.NewReader(jsonData))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func GetTelegramChatIds(accessToken string) []string {
	baseUrl := "https://api.telegram.org/bot%s/%s"
	getUpdatesUrl := fmt.Sprintf(baseUrl, accessToken, "getUpdates")
	client := http.Client{}
	response, err := client.Get(getUpdatesUrl)
	if err != nil {
		panic(err)
	}
	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	res := &GetUpdates{}
	fmt.Println(string(dataBytes))
	err = json.Unmarshal(dataBytes, &res)
	if err != nil {
		panic(err)
	}
	unique := []string{}
	for _, v := range res.Result {
		if v.Message.Chat.Type == "group" {
			isNotExist := true
			for _, k := range unique {
				if strconv.Itoa(v.Message.Chat.Id) == k {
					isNotExist = false
					break
				}
			}
			if isNotExist {
				unique = append(unique, strconv.Itoa(v.Message.Chat.Id))
			}
		}
	}
	fmt.Println(unique)
	return unique
}
