package setdata_questionnaire_store

type CreateOrderCommand struct {
	QuestionnaireName    string                `json:"questionnaire_name"`
	QuestionnaireAnswers []QuestionnaireAnswer `json:"questionnaire_answers"`
}

func (cmd *CreateOrderCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(OrderService).CreateOrder(cmd)
}

type ListOrderCommand struct {
	QuestionnaireName string `json:"questionnaire_name"`
}

func (cmd *ListOrderCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(OrderService).ListOrder(cmd)
}
