package pasta

import (
	"database/sql"


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
	pasta := model.Pasta {
		Title: &body.Title,
		Description: &body.Description,
	}

	insertStmt := Pasta.INSERT(Pasta.Description, Pasta.Title).MODEL(pasta).RETURNING(Pasta.ID)

	res, err := insertStmt.Exec(s.DBConn)

	if err != nil {
		s.Logger.Errorf("Failed to inser pasta, DBError: %w", err.Error())
		return err
	}

	pastaId, err := res.LastInsertId()

	if err != nil{
		s.Logger.Errorf("Failed to get id from created pasta, DBError: %w", err.Error())
	}

	return s.NoodleService.CreateNoodles(body.Noodles, int32(pastaId))
}
