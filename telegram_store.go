package setdata_questionnaire_store

type TelegramStore interface {
	Create(tel *TelegramBot) (*TelegramBot, error)
	Get(id string) (*TelegramBot, error)
	List() ([]TelegramBot, error)
	Delete(id string) error
}
