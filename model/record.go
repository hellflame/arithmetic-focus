package model

import "github.com/hellflame/arithmetic-focus/utils"

type Record struct {
	Expression string
	Occur      int
	Correct    int
}

type RecordSummary struct {
	ExpCounts int
	Occurs    int
	Corrects  int
}

func SaveRecord(exp string, answer int) (bool, error) {
	right, e := utils.ParseExpression(exp)
	if e != nil {
		return false, e
	}

	match := Record{Expression: exp}
	if e := _DB.Get(&match, "select expression, occur, correct from record where expression = $1", exp); e != nil {
		if _, e := _DB.Exec("insert into record (expression, occur, correct) values ($1, $2, $2)", exp, 0); e != nil {
			return false, e
		}
	}

	match.Occur += 1
	isCorrect := right == answer
	if isCorrect {
		match.Correct += 1
	}

	_, e = _DB.Exec("update record set occur = $1, correct = $2 where expression = $3", match.Occur, match.Correct, match.Expression)
	return isCorrect, e
}

func removeRecord(exp string) {
	_DB.Exec("delete from record where expression = $1", exp)
}

func ReadRecordsSummary() RecordSummary {
	var count, occurs, corrects int
	_DB.Get(&count, "select count(*) from record")
	rows, _ := _DB.Query("select occur, correct from record")
	for rows.Next() {
		occur := 0
		correct := 0
		rows.Scan(&occur, &correct)
		occurs += occur
		corrects += correct
	}

	return RecordSummary{ExpCounts: count, Corrects: corrects, Occurs: occurs}
}
