package data

import (
	"fmt"
	"strings"

	"github.com/eduartepaiva/greenlight/internal/validator"
)

type Filters struct {
	Page         int
	PageSize     int
	Sort         string
	SortSafeList []string
}

func (f Filters) sortColumn() string {
	for _, safeValue := range f.SortSafeList {
		if f.Sort == safeValue {
			return strings.TrimPrefix(f.Sort, "-")
		}
	}

	panic("unsafe sort parameter: " + f.Sort)
}

func (f Filters) sortDirection() string {
	if strings.HasPrefix(f.Sort, "-") {
		return "DESC"
	}
	return "ASC"
}

func ValidateFilters(v *validator.Validator, f Filters) {
	v.Check(f.Page > 0 && f.Page <= 10_000_000, "page", "must be between 1 and 10.000.000")
	v.Check(f.PageSize > 0 && f.PageSize <= 100, "page_size", "must be between 1 and 100")
	v.Check(validator.PermittedValue(f.Sort, f.SortSafeList...), "sort", fmt.Sprintf("must be one of the following values %+v", f.SortSafeList))
}
