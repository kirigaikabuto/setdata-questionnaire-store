package setdata_questionnaire_store

import (
	"errors"
	com "github.com/kirigaikabuto/setdata-common"
)

var (
	ErrCreateQuestionUnknown      = com.NewMiddleError(errors.New("could not create question: unknown error"), 500, 1)
	ErrQuestionNotFound           = com.NewMiddleError(errors.New("question not found"), 404, 2)
	ErrNothingToUpdate            = com.NewMiddleError(errors.New("nothing to update"), 400, 3)
	ErrQuestionIdProvided         = com.NewMiddleError(errors.New("question id is not provided"), 400, 4)
	ErrCreateQuestionnaireUnknown = com.NewMiddleError(errors.New("could not create questionnaire: unknown error"), 500, 5)
	ErrQuestionnaireNotFound      = com.NewMiddleError(errors.New("questionnaire not found"), 404, 6)
	ErrCreateOrderUnknown         = com.NewMiddleError(errors.New("could not create order: unknown error"), 500, 7)
	ErrOrderNotFound              = com.NewMiddleError(errors.New("order not found"), 404, 8)
)
