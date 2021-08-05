package setdata_questionnaire_store

type ChatIdStore interface {
	Create(ch *ChatId) (*ChatId, error)
	List(telegramBotId string) ([]ChatId, error)
}
