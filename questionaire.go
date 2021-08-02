package setdata_questionnaire_store

type Questionnaire struct {
	Id           string   `json:"id"`
	Name         string   `json:"name"`
	QuestionsIds []string `json:"questions_ids"`
}
