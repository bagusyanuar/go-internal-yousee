package controller

import (
	"github.com/bagusyanuar/go-internal-yousee/common"
	"github.com/bagusyanuar/go-internal-yousee/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type DashboardController struct {
	DashboardService service.DashboardService
	Log              *logrus.Logger
}

func NewDashboardController(dashboardService service.DashboardService, log *logrus.Logger) *DashboardController {
	return &DashboardController{
		DashboardService: dashboardService,
		Log:              log,
	}
}

func (c *DashboardController) GetDashboardStatisticInfo(ctx *fiber.Ctx) error {
	response := c.DashboardService.GetDashboardStatisticInfo(ctx.UserContext())
	if response.Error != nil {
		return common.JSONFromError(ctx, response.Status, response.Error, nil)
	}

	return common.JSONSuccess(ctx, common.ResponseMap{
		Message: "successfully show statistic info",
		Data:    response.Data,
	})
}
