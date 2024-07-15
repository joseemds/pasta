package noodles

import (
	"database/sql"

	"github.com/joseemds/pasta/.gen/pasta/public/model"
	. "github.com/joseemds/pasta/.gen/pasta/public/table"
	"go.uber.org/zap"
)

type NoodleService struct{
	DBConn *sql.DB
	Logger *zap.SugaredLogger
}


func NewService( logger *zap.SugaredLogger, db *sql.DB) NoodleService{
	return NoodleService{
		DBConn: db,
		Logger: logger,
	}
}

func (s NoodleService) createNoodle(schema NoodleSchema) (sql.Result, error){
	noodle := model.Noodle {
		Content: schema.Content,
		Language: schema.Language,
		Filename: schema.Filename,
	}

	insertStmt := Noodle.INSERT(Noodle.Content, Noodle.Language, Noodle.Filename).MODEL(noodle)
	res, err := insertStmt.Exec(s.DBConn)
	if err != nil {
		s.Logger.Errorf("Failed to insert noodle, DBError: %w", err.Error())
		return nil, err
	}

	return res, nil
}

func (s NoodleService) createNoodles(schemas []NoodleSchema) error {
	for _, schema := range schemas {
		if _, err := s.createNoodle(schema); err != nil {
			return err
		}
	}
	return nil
}
