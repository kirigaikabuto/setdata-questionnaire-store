package setdata_questionnaire_store

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"strconv"
)

var ordersQueries = []string{
	`create table if not exists orders(
		id text,
		questionnaire_name text,
		questionnaire_answers text,
		primary key(id)
	)`,
}

type ordersStore struct {
	db *sql.DB
}

func NewOrderPostgresStore(cfg PostgresConfig) (OrderStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range ordersQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &ordersStore{db: db}
	return store, nil
}

func (o *ordersStore) Create(order *Order) (*Order, error) {
	order.Id = uuid.New().String()
	questions, err := json.Marshal(order.QuestionnaireAnswers)
	if err != nil {
		return nil, err
	}
	result, err := o.db.Exec("INSERT INTO orders (id, questionnaire_name, questionnaire_answers) VALUES ($1, $2, $3)",
		order.Id, order.QuestionnaireName, string(questions),
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateOrderUnknown
	}
	return order, nil
}

func (o *ordersStore) List(questionnaireName string) ([]Order, error) {
	items := []Order{}
	query := "select id, questionnaire_name, questionnaire_answers from orders "
	var values []interface{}
	cnt := 1
	if questionnaireName != "" {
		query = query + "where questionnaire_name = $" + strconv.Itoa(cnt)
		values = append(values, questionnaireName)
		cnt += 1
	}
	rows, err := o.db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		questionAnswers := ""
		item := Order{}
		err = rows.Scan(&item.Id, &item.QuestionnaireName, &questionAnswers)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal([]byte(questionAnswers), &item.QuestionnaireAnswers)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
