package setdata_questionnaire_store

import setdata_common "github.com/kirigaikabuto/setdata-common"

type Question struct {
	Id     string                 `json:"id"`
	Name   string                 `json:"name"`
	Order  *int                   `json:"order"`
	Fields []setdata_common.Field `json:"fields"`
}

type QuestionUpdate struct {
	Id     string                  `json:"id"`
	Name   *string                 `json:"name"`
	Order  *int                    `json:"order"`
	Fields *[]setdata_common.Field `json:"fields"`
}
