package setdata_questionnaire_store

type QuestionnaireService interface {
	CreateQuestionnaire(cmd *CreateQuestionnaireCommand) (*Questionnaire, error)
	UpdateQuestionnaire(cmd *UpdateQuestionnaireCommand) (*Questionnaire, error)
	ListQuestionnaire(cmd *ListQuestionnaireCommand) ([]Questionnaire, error)
	GetQuestionnaireById(cmd *GetQuestionnaireByIdCommand) (*Questionnaire, error)
	GetQuestionnaireByName(cmd *GetQuestionnaireByNameCommand) (*Questionnaire, error)
	DeleteQuestionnaire(cmd *DeleteQuestionnaireByIdCommand) error
}

type questionnaireService struct {
	store QuestionnaireStore
}

func NewQuestionnaireService(store QuestionnaireStore) QuestionnaireService {
	return &questionnaireService{store: store}
}

func (q *questionnaireService) CreateQuestionnaire(cmd *CreateQuestionnaireCommand) (*Questionnaire, error) {
	return q.store.Create(&Questionnaire{
		Name:      cmd.Name,
		Questions: cmd.Questions,
	})
}

func (q *questionnaireService) UpdateQuestionnaire(cmd *UpdateQuestionnaireCommand) (*Questionnaire, error) {
	return q.store.Update(cmd.QuestionnaireUpdate)
}

func (q *questionnaireService) ListQuestionnaire(cmd *ListQuestionnaireCommand) ([]Questionnaire, error) {
	return q.store.List()
}

func (q *questionnaireService) GetQuestionnaireById(cmd *GetQuestionnaireByIdCommand) (*Questionnaire, error) {
	return q.store.Get(cmd.Id)
}

func (q *questionnaireService) GetQuestionnaireByName(cmd *GetQuestionnaireByNameCommand) (*Questionnaire, error) {
	return q.store.GetByName(cmd.Name)
}

func (q *questionnaireService) DeleteQuestionnaire(cmd *DeleteQuestionnaireByIdCommand) error {
	return q.store.Delete(cmd.Id)
}
