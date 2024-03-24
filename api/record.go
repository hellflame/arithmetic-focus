package api

import (
	"github.com/hellflame/arithmetic-focus/model"
	"github.com/labstack/echo/v4"
)

func bindRecord() {
	type SubmitRequest struct {
		Exp    string
		Answer int
	}

	router.POST("/record/submit", func(c echo.Context) error {
		request := SubmitRequest{}
		if e := c.Bind(&request); e != nil {
			return responseSimpleMessage(c, 1, "failed to parse request")
		}
		isCorrect, e := model.SaveRecord(request.Exp, request.Answer)
		if e != nil {
			return responseSimpleMessage(c, 2, "invalid exp: "+e.Error())
		}

		return responseLowerJson(c, isCorrect)
	})

	router.GET("/record/summary", func(c echo.Context) error {
		c.Response().Header().Set("Cache-Control", "no-store")
		return responseLowerJson(c, model.ReadRecordsSummary())
	})
}
