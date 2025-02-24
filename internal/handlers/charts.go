package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chichigami/EMR/internal/components"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *HandlerConfig) HandlerChartsCreate(c *gin.Context) {
	id := c.Param("id")
	patientID, err := ConvertStringToInt32(id)
	if err != nil {
		log.Println(err.Error())
	}
	dbChart, err := h.Config.Database.CreateChart(c, patientID)
	if err != nil {
		log.Println(err)
	}
	//maybe will redirect page instead. need to somehow refresh chart table partial div and return that string
	c.String(http.StatusOK, fmt.Sprintf("/patients/%s/charts/%s", id, dbChart.ID))
}

func (h *HandlerConfig) HandlerChartsGet(c *gin.Context) {
	param := c.Param("uuid")
	//for when i build the charts
	_, err := h.Config.Database.GetChart(c, uuid.MustParse(param))
	if err != nil {
		log.Println(err)
	}
	page := components.Base("Some person's chart", components.DefaultNavbar(), nil, components.ChartFooter(param, "dr. feng"))
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := page.Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

// func (h *HandlerConfig) HandlerChartsUpdate(c *gin.Context) {
// 	//should receive a request every chart completion? or every minute.
// 	HandlerPlaceholder(c)
// }

func (h *HandlerConfig) HandlerChartsDelete(c *gin.Context) {
	param := c.Param("uuid")
	err := h.Config.Database.DeleteChart(c, uuid.MustParse(param))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
}
