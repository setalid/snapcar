package api

import (
	"net/http"

	"github.com/setalid/snapcar/api/pkg/core"
	"github.com/setalid/snapcar/api/pkg/utils"
	"go.uber.org/zap"
)

func handleGetCarCategories(
	log *zap.Logger,
	carCategoryRepo core.CarCategoryRepo,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		categories, err := carCategoryRepo.All(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Error(err.Error())
		}

		viewCategories := utils.Map(categories, func(c core.CarCategory) map[string]any {
			return map[string]any{
				"name":         c.Name,
				"priceFormula": c.PriceFormula(),
			}
		})

		encode(w, r, http.StatusOK, map[string]any{
			"categories": viewCategories,
		})
	})
}
