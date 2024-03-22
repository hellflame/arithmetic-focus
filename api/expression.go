package api

import (
	"arithmetic/utils"

	"github.com/labstack/echo/v4"
)

func bindExpression() {
	type GenerationRequest struct {
		Max        int
		Difficulty int
		Count      int
		Collisions []string
	}

	difficulties := [][2]float64{}
	for i := 0; i < 5; i++ {
		difficulties = append(difficulties, [2]float64{float64(i+1) * 10 / 100, float64(i) * 3 / 400})
	}

	router.POST("/expression/gen", func(c echo.Context) error {
		request := GenerationRequest{}
		if e := c.Bind(&request); e != nil {
			return responseSimpleMessage(c, 1, "failed to parse request")
		}
		if request.Difficulty < 1 || request.Difficulty > 5 {
			return responseSimpleMessage(c, 2, "no such difficulty")
		}
		if request.Count < 1 {
			return responseSimpleMessage(c, 3, "no enough count")
		}
		if request.Max < 2 {
			return responseSimpleMessage(c, 4, "max is too small")
		}
		ratios := difficulties[request.Difficulty-1]
		result := []string{}
		collisionPool := make(map[string]struct{}, len(request.Collisions))
		for _, c := range request.Collisions {
			collisionPool[c] = struct{}{}
		}

		tryGenerateUniqueExp := func(tries int) string {
			exp := ""
			for i := 0; i < tries; i++ {
				exp = utils.GenerateExpression(request.Max, ratios[0], ratios[1])
				if _, exist := collisionPool[exp]; !exist {
					collisionPool[exp] = struct{}{}
					return exp
				}
			}
			return exp
		}
		for i := 0; i < request.Count; i++ {
			result = append(result, tryGenerateUniqueExp(10))
		}

		return responseLowerJson(c, result)
	})
}
