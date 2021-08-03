package setdata_questionnaire_store

type QuestionnaireStore interface {
	Create(questionnaire *Questionnaire) (*Questionnaire, error)
	Update(questionnaire *QuestionnaireUpdate) (*Questionnaire, error)
	Get(id string) (*Questionnaire, error)
	List() ([]Questionnaire, error)
	Delete(id string) error
	GetByName(name string) (*Questionnaire, error)
}