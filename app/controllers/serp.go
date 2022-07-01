package controllers

import (
	"net/http"
	"ushas/models"
	"ushas/views"

	"github.com/labstack/echo/v4"
)

// SERPController : Struct for controll `SERP` resource.
type SERPController struct{}

// NewSERPController : Return pointer to `SERPController`.
func NewSERPController() *SERPController {
	return new(SERPController)
}

// ListSERPRequest :
type ListSERPRequest struct {
	TaskID int `json:"id" param:"id" validate:"required,numeric"`
	Offset int `json:"offset" query:"offset" validate:"numeric"`
}

// ListSERP :
func (sc *SERPController) ListSERP(c echo.Context) error {
	if sc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrNilReceiver)
	}

	p := ListSERPRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.AppError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	pages, err := models.ListSERP(p.TaskID, p.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, newResponse(http.StatusOK, http.StatusText(http.StatusOK), pages))
}

// ListSERPWithIconRequest :
type ListSERPWithIconRequest struct {
	TaskID int `json:"id" param:"id" validate:"required,numeric"`
	Offset int `json:"offset" query:"offset" validate:"numeric"`
	Top    int `json:"top" query:"top" validate:"numeric"`
}

// ListSERPWithIcon : Return search result pages with similarweb icon information
func (sc *SERPController) ListSERPWithIcon(c echo.Context) error {
	if sc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrNilReceiver)
	}

	p := ListSERPWithIconRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.AppError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}
	if p.Top < 1 {
		p.Top = 10
	}

	pages, err := models.ListSERP(p.TaskID, p.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	pageIDs := make([]int, 0, len(pages))
	for _, v := range pages {
		pageIDs = append(pageIDs, v.PageID)
	}

	linkedPages, err := models.GetLinkedPagesByIDs(pageIDs, p.TaskID, p.Top)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	sv := views.NewSERPWithIconView(pages, pageIDs, linkedPages)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		sv,
	))
}

// ListSERPWithRatioRequest :
type ListSERPWithRatioRequest struct {
	TaskID int `json:"id" param:"id" validate:"required,numeric"`
	Offset int `json:"offset" query:"offset" validate:"numeric,min=0,max=10"`
	Top    int `json:"top" query:"top" validate:"numeric"`
}

// ListSERPWithRatio :
func (sc *SERPController) ListSERPWithRatio(c echo.Context) error {
	if sc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, models.ErrNilReceiver)
	}

	p := ListSERPWithRatioRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.AppError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(
			http.StatusBadRequest,
			models.NewAppError(err, http.StatusBadRequest, "Failed bind request body.", p),
		)
	}
	if p.Top < 1 {
		p.Top = 3
	}

	pages, err := models.ListSERP(p.TaskID, p.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	pageIDs := []int{}
	for _, v := range pages {
		pageIDs = append(pageIDs, v.PageID)
	}

	linkedPages, err := models.GetLinkedPagesRatioByIDs(pageIDs, p.TaskID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	sv := views.NewSERPWithRatioView(pages, pageIDs, linkedPages, p.Top)

	return c.JSON(http.StatusOK, newResponse(
		http.StatusOK,
		http.StatusText(http.StatusOK),
		sv,
	))
}
