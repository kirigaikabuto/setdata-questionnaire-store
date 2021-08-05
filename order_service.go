package setdata_questionnaire_store

type OrderService interface {
	CreateOrder(cmd *CreateOrderCommand) (*Order, error)
	ListOrder(cmd *ListOrderCommand) ([]Order, error)
}

type orderService struct {
	store OrderStore
}

func NewOrderService(store OrderStore) OrderService {
	return &orderService{store: store}
}

func (o *orderService) CreateOrder(cmd *CreateOrderCommand) (*Order, error) {
	return o.store.Create(&Order{
		Id:                   "",
		QuestionnaireName:    cmd.QuestionnaireName,
		QuestionnaireAnswers: cmd.QuestionnaireAnswers,
	})
}

func (o *orderService) ListOrder(cmd *ListOrderCommand) ([]Order, error) {
	return o.store.List(cmd.QuestionnaireName)
}
