package setdata_questionnaire_store

type Order struct {
	Id                   string                `json:"id"`
	QuestionnaireName    string                `json:"questionnaire_name"`
	QuestionnaireAnswers []QuestionnaireAnswer `json:"questionnaire_answers"`
}

type QuestionnaireAnswer struct {
	QuestionName string   `json:"question_name"`
	Answers      []string `json:"answers"`
}
