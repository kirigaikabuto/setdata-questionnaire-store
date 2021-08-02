package setdata_questionnaire_store

type QuestionsService interface {
	CreateQuestion(cmd *CreateQuestionCommand) (*Question, error)
	UpdateQuestion(cmd *UpdateQuestionsCommand) (*Question, error)
	GetQuestion(cmd *GetQuestionCommand) (*Question, error)
	DeleteQuestion(cmd *DeleteQuestionCommand) error
	ListQuestions(cmd *ListQuestionsCommand) ([]Question, error)
}

type questionsService struct {
	store QuestionStore
}

func NewQuestionsService(store QuestionStore) QuestionsService {
	return &questionsService{store: store}
}

func (q *questionsService) CreateQuestion(cmd *CreateQuestionCommand) (*Question, error) {
	question := &Question{
		Name:   cmd.Name,
		Order:  cmd.Order,
		Fields: cmd.Fields,
	}
	return q.store.Create(question)
}

func (q *questionsService) UpdateQuestion(cmd *UpdateQuestionsCommand) (*Question, error) {
	return q.store.Update(cmd.QuestionUpdate)
}

func (q *questionsService) GetQuestion(cmd *GetQuestionCommand) (*Question, error) {
	return q.store.Get(cmd.Id)
}

func (q *questionsService) DeleteQuestion(cmd *DeleteQuestionCommand) error {
	return q.store.Delete(cmd.Id)
}

func (q *questionsService) ListQuestions(cmd *ListQuestionsCommand) ([]Question, error) {
	return q.store.List()
}
