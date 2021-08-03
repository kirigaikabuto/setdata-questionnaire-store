package setdata_questionnaire_store

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type QuestionnaireAmqpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewQuestionnaireAmqpEndpoints(ch setdata_common.CommandHandler) QuestionnaireAmqpEndpoints {
	return QuestionnaireAmqpEndpoints{ch: ch}
}

func (q *QuestionnaireAmqpEndpoints) MakeCreateQuestionnaireAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateQuestionnaireCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := q.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func (q *QuestionnaireAmqpEndpoints) MakeUpdateQuestionnaireAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &UpdateQuestionnaireCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := q.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func (q *QuestionnaireAmqpEndpoints) MakeListQuestionnaireAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &ListQuestionnaireCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := q.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func (q *QuestionnaireAmqpEndpoints) MakeGetByIdQuestionnaireAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetQuestionnaireByIdCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := q.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func (q *QuestionnaireAmqpEndpoints) MakeGetByNameQuestionnaireAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &GetQuestionnaireByNameCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := q.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}

func (q *QuestionnaireAmqpEndpoints) MakeDeleteQuestionnaireAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &DeleteQuestionnaireByIdCommand{}
		err := json.Unmarshal(message.Body, &cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		response, err := q.ch.ExecCommand(cmd)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			return setdata_common.ErrToAmqpResponse(err)
		}
		return &amqp.Message{Body: jsonResponse}
	}
}
