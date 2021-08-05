package setdata_questionnaire_store

type OrderStore interface {
	Create(order *Order) (*Order, error)
	List(questionnaireName string) ([]Order, error)
}
