package setdata_questionnaire_store

import (
	"encoding/json"
	"github.com/djumanoff/amqp"
	setdata_common "github.com/kirigaikabuto/setdata-common"
)

type OrderAmqpEndpoints struct {
	ch setdata_common.CommandHandler
}

func NewOrderAmqpEndpoints(ch setdata_common.CommandHandler) OrderAmqpEndpoints {
	return OrderAmqpEndpoints{ch: ch}
}

func (q *OrderAmqpEndpoints) MakeCreateOrderAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &CreateOrderCommand{}
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

func (q *OrderAmqpEndpoints) MakeListOrderAmqpEndpoint() amqp.Handler {
	return func(message amqp.Message) *amqp.Message {
		cmd := &ListOrderCommand{}
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
