package setdata_questionnaire_store

type Question struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Order  string  `json:"order"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Placeholder string `json:"placeholder"`
}
