package setdata_questionnaire_store

type CreateQuestionnaireCommand struct {
	Name      string   `json:"name"`
	Questions []string `json:"questions"`
}

func (cmd *CreateQuestionnaireCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionnaireService).CreateQuestionnaire(cmd)
}

type GetQuestionnaireByIdCommand struct {
	Id string `json:"id"`
}

func (cmd *GetQuestionnaireByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionnaireService).GetQuestionnaireById(cmd)
}

type GetQuestionnaireByNameCommand struct {
	Name string `json:"name"`
}

func (cmd *GetQuestionnaireByNameCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionnaireService).GetQuestionnaireByName(cmd)
}

type UpdateQuestionnaireCommand struct {
	*QuestionnaireUpdate
}

func (cmd *UpdateQuestionnaireCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionnaireService).UpdateQuestionnaire(cmd)
}

type DeleteQuestionnaireByIdCommand struct {
	Id string `json:"id"`
}

func (cmd *DeleteQuestionnaireByIdCommand) Exec(svc interface{}) (interface{}, error) {
	return nil, svc.(QuestionnaireService).DeleteQuestionnaire(cmd)
}

type ListQuestionnaireCommand struct {
}

func (cmd *ListQuestionnaireCommand) Exec(svc interface{}) (interface{}, error) {
	return svc.(QuestionnaireService).ListQuestionnaire(cmd)
}
