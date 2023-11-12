package words_cli

import (
	"database/sql"
	"fmt"
	"time"
	"trl/db"
	"trl/entities/words/processing"
	"trl/srv/errs"
	"trl/srv/files"
	"trl/srv/logger"

	"github.com/huandu/go-sqlbuilder"
)

type DumpCmd struct {
	Path   string `short:"p" required:"" help:"path to output file"`
	Format string `short:"f" enum:"csv,json" default:"csv" help:"output format"`

	Since time.Time `short:"s" format:"2006-01-02" help:"date since (YYYY-MM-DD)"`
	Until time.Time `short:"u" format:"2006-01-02" help:"date until (YYYY-MM-DD)"`

	Offset int `short:"o" help:"skip the specified number of words"`
	Limit  int `short:"l" help:"maximum number of words"`

	Desc bool `short:"d" default:"false" help:"in descending order"`
}

var emptyDate = time.Time{}

func (this *DumpCmd) Run(dbInstance *db.DB) error {
	this.prepare()

	dumper := processing.DumperFactory(this.Format)
	rows := this.getRows(dbInstance)
	num := dumper.Dump(this.Path, rows)

	logger.Outputln(fmt.Sprintf("%d words dumped", num), logger.Colors.Green)
	return nil
}

func (this *DumpCmd) prepare() {
	if this.Until.IsZero() {
		day := 24 * time.Hour
		this.Until = time.Now().UTC().Truncate(day).Add(day)
	}
	this.Path = files.PrepareFile(this.Path)
}

func (this *DumpCmd) getRows(dbInstance *db.DB) *sql.Rows {
	query := this.makeQuery()
	rows, err := dbInstance.Conn.Query(query, this.Since, this.Until)
	errs.LogFatalIfError(err)
	return rows
}

func (this *DumpCmd) makeQuery() string {
	query := sqlbuilder.
		Select("word", "translation", "transcription").
		From(db.WordsTableName).
		OrderBy("id").
		Asc().
		Where("created_at BETWEEN ? AND ?")

	if this.Desc {
		query.Desc()
	}

	if this.Offset > 0 {
		query.Offset(this.Offset)
	}
	if this.Limit > 0 {
		query.Limit(this.Limit)
	}

	return query.String()
}
