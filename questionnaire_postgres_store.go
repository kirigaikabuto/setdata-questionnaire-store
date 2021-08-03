package setdata_questionnaire_store

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"
)

var questionnaireQueries = []string{
	`create table if not exists questionnaire(
		id text,
		name text,
		questions text,
		primary key(id)
	)`,
}

type questionnaireStore struct {
	db *sql.DB
}

func NewQuestionnairePostgresStore(cfg PostgresConfig) (QuestionnaireStore, error) {
	db, err := getDbConn(getConnString(cfg))
	if err != nil {
		return nil, err
	}
	for _, q := range questionnaireQueries {
		_, err = db.Exec(q)
		if err != nil {
			log.Println(err)
		}
	}
	db.SetMaxOpenConns(10)
	store := &questionnaireStore{db: db}
	return store, nil
}

func (q *questionnaireStore) Create(questionnaire *Questionnaire) (*Questionnaire, error) {
	questionnaire.Id = uuid.New().String()
	questions, err := json.Marshal(questionnaire.Questions)
	if err != nil {
		return nil, err
	}
	result, err := q.db.Exec("INSERT INTO questionnaire (id, name, questions) VALUES ($1, $2, $3)",
		questionnaire.Id, questionnaire.Name, string(questions),
	)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrCreateQuestionnaireUnknown
	}
	return questionnaire, nil
}

func (q *questionnaireStore) Update(questionnaire *QuestionnaireUpdate) (*Questionnaire, error) {
	query := "update questionnaire set "
	parts := []string{}
	values := []interface{}{}
	cnt := 0
	if questionnaire.Name != nil {
		cnt++
		parts = append(parts, "name = $"+strconv.Itoa(cnt))
		values = append(values, questionnaire.Name)
	}
	if questionnaire.Questions != nil {
		cnt++
		questions, err := json.Marshal(questionnaire.Questions)
		if err != nil {
			return nil, err
		}
		parts = append(parts, "questions = $"+strconv.Itoa(cnt))
		values = append(values, questions)
	}
	if len(parts) <= 0 {
		return nil, ErrNothingToUpdate
	}
	cnt++
	query = query + strings.Join(parts, " , ") + " WHERE id = $" + strconv.Itoa(cnt)
	values = append(values, questionnaire.Id)
	result, err := q.db.Exec(query, values...)
	if err != nil {
		return nil, err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if n <= 0 {
		return nil, ErrQuestionnaireNotFound
	}
	return q.Get(questionnaire.Id)
}

func (q *questionnaireStore) Get(id string) (*Questionnaire, error) {
	questionnaire := &Questionnaire{}
	questions := ""
	err := q.db.QueryRow("select id, name, questions "+
		"from questionnaire where id = $1 limit 1", id).
		Scan(&questionnaire.Id, &questionnaire.Name, &questions)
	if err == sql.ErrNoRows {
		return nil, ErrQuestionnaireNotFound
	} else if err != nil {
		return nil, err
	}
	if questions != "" {
		err = json.Unmarshal([]byte(questions), &questionnaire.Questions)
		if err != nil {
			return nil, err
		}
	}
	return questionnaire, nil
}

func (q *questionnaireStore) List() ([]Questionnaire, error) {
	items := []Questionnaire{}
	query := "select id, name, questions from questionnaire"
	rows, err := q.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		item := Questionnaire{}
		questions := ""
		err = rows.Scan(&item.Id, &item.Name, &questions)
		if err != nil {
			return nil, err
		}
		if questions != "" {
			err = json.Unmarshal([]byte(questions), &item.Questions)
			if err != nil {
				return nil, err
			}
		}
		items = append(items, item)
	}
	return items, nil
}

func (q *questionnaireStore) Delete(id string) error {
	result, err := q.db.Exec("delete from questionnaire where id= $1", id)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if n <= 0 {
		return ErrQuestionnaireNotFound
	}
	return nil
}

func (q *questionnaireStore) GetByName(name string) (*Questionnaire, error) {
	questionnaire := &Questionnaire{}
	questions := ""
	err := q.db.QueryRow("select id, name, questions "+
		"from questionnaire where name = $1 limit 1", name).
		Scan(&questionnaire.Id, &questionnaire.Name, &questions)
	if err == sql.ErrNoRows {
		return nil, ErrQuestionnaireNotFound
	} else if err != nil {
		return nil, err
	}
	if questions != "" {
		err = json.Unmarshal([]byte(questions), &questionnaire.Questions)
		if err != nil {
			return nil, err
		}
	}
	return questionnaire, nil
}
