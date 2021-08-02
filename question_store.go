package setdata_questionnaire_store

type QuestionStore interface {
	Create(question *Question) (*Question, error)
	Update(question *QuestionUpdate) (*Question, error)
	Get(id string) (*Question, error)
	Delete(id string) error
	List() ([]Question, error)
}
