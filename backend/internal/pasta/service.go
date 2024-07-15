package pasta

import (
	"database/sql"
	"fmt"

	"github.com/joseemds/pasta/.gen/pasta/public/model"
	. "github.com/joseemds/pasta/.gen/pasta/public/table"
	"github.com/joseemds/pasta/internal/noodles"
	"go.uber.org/zap"
)

type PastaService struct {
	DBConn *sql.DB
	Logger *zap.SugaredLogger
	NoodleService noodles.NoodleService
}

func NewService(logger *zap.SugaredLogger, db *sql.DB) PastaService {
	noodleService := noodles.NoodleService {DBConn: db, Logger: logger}
	return PastaService {
		DBConn: db,
		Logger: logger,
		NoodleService: noodleService,
	}
}


func (s PastaService) CreatePasta(body CreatePastaRequestBody)  (error){
	fmt.Printf("%+v\n", body)
	pasta := model.Pasta {
		Title: &body.Title,
		Description: &body.Description,
	}

	insertStmt := Pasta.INSERT(Pasta.Description, Pasta.Title).MODEL(pasta).RETURNING(Pasta.ID)

	createdPasta := model.Pasta{}

	err := insertStmt.Query(s.DBConn, &createdPasta)

	if err != nil {
		s.Logger.Errorf("Failed to insert pasta, DBError: %w", err)
		return err
	}

	return s.NoodleService.CreateNoodles(body.Noodles, &createdPasta.ID)
}
