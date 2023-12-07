package paginate

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type Pagination struct {
	Page          int
	Limit         int
	Search        string
	MinSize       int
	MaxSize       int
	MinPrice      int
	MaxPrice      int
	LowerPrice    bool
	HigestPrice   bool
	Recomendation bool
	Status        string

	// Opsional
	PaymentSchemeId int
}

type PaginationTrx struct {
	Page       int
	Limit      int
	Search     string
	ProvinceId int
	Status     string
}

type ItemPages struct {
	Page      int64 `json:"page"`
	Limit     int64 `json:"limit"`
	TotalData int64 `json:"totalData"`
	TotalPage int64 `json:"totalPage"`
}

func Paginate(page, length int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (page - 1) * length
		return db.Offset(offset).Limit(length)
	}
}

func GetParams(c echo.Context) (Pagination, error) {
	params := Pagination{
		Page:            cast.ToInt(c.QueryParam("page")),
		Limit:           cast.ToInt(c.QueryParam("limit")),
		Search:          c.QueryParam("search"),
		MinSize:         cast.ToInt(c.QueryParam("minSize")),
		MaxSize:         cast.ToInt(c.QueryParam("maxSize")),
		MinPrice:        cast.ToInt(c.QueryParam("minPrice")),
		MaxPrice:        cast.ToInt(c.QueryParam("maxPrice")),
		LowerPrice:      cast.ToBool(c.QueryParam("lowerPrice")),
		HigestPrice:     cast.ToBool(c.QueryParam("highestPrice")),
		Recomendation:   cast.ToBool(c.QueryParam("recomendation")),
		PaymentSchemeId: cast.ToInt(c.QueryParam("paymentSchemeId")),
		Status:          c.QueryParam("status"),
	}

	if params.Page == 0 {
		params.Page = 1
	}

	if params.Limit == 0 {
		params.Limit = 10
	}

	if params.Limit > 30 {
		params.Limit = 30
	}

	counter := 0
	if params.LowerPrice {
		counter += 1
	}
	if params.HigestPrice {
		counter += 1
	}
	if params.Recomendation {
		counter += 1
	}

	if counter > 1 {
		params.LowerPrice = false
		params.HigestPrice = false
		params.Recomendation = false
	}

	return params, nil
}
