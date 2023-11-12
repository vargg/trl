package words

import (
	"trl/db"

	"github.com/huandu/go-sqlbuilder"
)

var countWordsQuery = sqlbuilder.
	Select("count(*)").
	From(db.WordsTableName).
	String()

var insertWordQuery = sqlbuilder.
	InsertInto(db.WordsTableName).
	Cols("word", "translation", "transcription", "created_at", "updated_at").
	Values("?", "?", "?", "?", "?").
	String()

var selectWordQuery = sqlbuilder.
	Select("*").
	From(db.WordsTableName).
	Where("word = ?").
	String()

var updateWordQuery = sqlbuilder.
	Update(db.WordsTableName).
	Set("translation = ?", "transcription = ?", "updated_at = ?").
	Where("word = ?").
	String()

var deleteWordQuery = sqlbuilder.
	DeleteFrom(db.WordsTableName).
	Where("word = ?").
	String()
