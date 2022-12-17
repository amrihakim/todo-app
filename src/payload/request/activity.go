package request

import (
	"time"

	"github.com/labstack/echo/v4"
)

type GetActivityListRequest struct {
	ID        uint64     `query:"id" json:"id"`
	Email     string     `query:"email" json:"email"`
	Activity  string     `query:"activity" json:"activity"`
	CreatedAt time.Time  `query:"created_at" json:"created_at"`
	UpdatedAt *time.Time `query:"updated_at" json:"updated_at"`
	DeletedAt *time.Time `query:"deleted_at" json:"deleted_at"`
}

func (g *GetActivityListRequest) GetDataFromHTTPRequest(c echo.Context) error {
	if err := c.Bind(g); err != nil {
		return err
	}

	return nil
}
