package transformer

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/model"
)

func ToMetaPagination(pagination *common.Pagination) *model.MetaPagination {
	return &model.MetaPagination{
		Page:      pagination.GetPage(),
		PerPage:   pagination.GetLimit(),
		TotalPage: pagination.TotalPages,
		TotalRows: pagination.TotalRows,
	}
}
