package setdata_questionnaire_store

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
	"strconv"
	"strings"
)

var questionsQueries = []string{
	`create table if not exists questions(
		id text,
		name text,
		inorder int,
		fields text,
		primary key(id)
	)`,
}

type questionsStore struct {
	db *sql.DB
}

func NewQuestionsPostgresStore(cfg PostgresConfig) (QuestionStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range questionsQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &questionsStore{db: db}
	return store, nil
}

func (q *questionsStore) Create(question *Question) (*Question, error) {
	question.Id = uuid.New().String()
	fields, err := json.Marshal(question.Fields)
	if err != nil {
		return nil, err
	}
	result, err := q.db.Exec("INSERT INTO questions (id, name, inorder, fields) VALUES ($1, $2, $3, $4)",
		question.Id, question.Name, question.Order, string(fields),
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateQuestionUnknown
	}
	return question, nil
}

func (q *questionsStore) Update(question *QuestionUpdate) (*Question, error) {
	query := "update questions set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if question.Name != nil {
		cnt++
		parts = append(parts, "name = $"+strconv.Itoa(cnt))
		values = append(values, question.Name)
	}
	if question.Order != nil {
		cnt++
		parts = append(parts, "inorder = $"+strconv.Itoa(cnt))
		values = append(values, question.Order)
	}
	if question.Fields != nil {
		cnt++
		fields, err := json.Marshal(question.Fields)
		if err != nil {
			return nil, err
		}
		parts = append(parts, "fields = $"+strconv.Itoa(cnt))
		values = append(values, fields)
	}
	if len(parts) <= 0 {
		return nil, ErrNothingToUpdate
	}
	cnt++
	query = query + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, question.Id)
	result, err := q.db.Exec(query, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrQuestionNotFound
	}
	return q.Get(question.Id)
}

func (q *questionsStore) Get(id string) (*Question, error) {
	question := &Question{}
	fields := ""
	err := q.db.QueryRow("select id, name, inorder, fields "+
		"from questions where id = $1 limit 1", id).
		Scan(&question.Id, &question.Name, &question.Order, &fields)
	if err == sql.ErrNoRows {
		return nil, ErrQuestionNotFound
	} else if err != nil {
		return nil, err
	}
	if fields != "" {
		err = json.Unmarshal([]byte(fields), &question.Fields)
		if err != nil {
			return nil, err
		}
	}
	return question, nil
}

func (q *questionsStore) Delete(id string) error {
	result, err := q.db.Exec("delete from questions where id= $1", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrQuestionNotFound
	}
	return nil
}

func (q *questionsStore) List() ([]Question, error) {
	items := []Question{}
	query := "select id, name, inorder, fields from questions order by inorder"
	rows, err := q.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := Question{}
		fields := ""
		err = rows.Scan(&item.Id, &item.Name, &item.Order, &fields)
		if err != nil {
			return nil, err
		}
		if fields != "" {
			err = json.Unmarshal([]byte(fields), &item.Fields)
			if err != nil {
				return nil, err
			}
		}
		items = append(items, item)
	}
	return items, nil
}
