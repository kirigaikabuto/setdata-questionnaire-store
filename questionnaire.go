package setdata_questionnaire_store

type Questionnaire struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Questions []string `json:"questions"`
}

type QuestionnaireUpdate struct {
	Id        string    `json:"id"`
	Name      *string   `json:"name"`
	Questions *[]string `json:"questions"`
}
