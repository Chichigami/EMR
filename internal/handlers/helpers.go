package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/chichigami/EMR/internal/components"
	"github.com/gin-gonic/gin"
)

func NullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func NullInt32(i int32) sql.NullInt32 {
	return sql.NullInt32{
		Int32: i,
		Valid: i != 0,
	}
}

func ConvertStringToInt32(s string) (int32, error) {
	d, err := strconv.Atoi(s)
	if err != nil {
		return -1, err
	}
	return int32(d), nil
}

func RenderView(c *gin.Context, title string, navbar, body, footer templ.Component) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	c.Header("X-Frame-Options", "DENY")
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
	page := components.Base(title, navbar, body, footer)
	if err := page.Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

func RenderAppointmentModal(c *gin.Context) {
	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := components.ModalAppointment(c.Param("id")).Render(c, c.Writer); err != nil {
		c.String(http.StatusInternalServerError, "Failed to render page: %v", err)
		return
	}
}

func RenderDeletedView(c *gin.Context) {
	c.String(http.StatusOK, "Deleted")
}
