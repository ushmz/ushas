package controllers

import (
	"net/http"
	"ushas/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/xerrors"
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
	Offset int `json:"offset" query:"offset" validate:"required,numeric"`
}

// ListSERP :
func (sc *SERPController) ListSERP(c echo.Context) error {
	if sc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, xerrors.New("SERPController is nil"))
	}

	p := ListSERPRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Invalid request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
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
	Offset int `json:"offset" query:"offset" validate:"required,numeric"`
	Top    int `json:"top" query:"top" validate:"numeric"`
}

// ListSERPWithIcon :
func (sc *SERPController) ListSERPWithIcon(c echo.Context) error {
	if sc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, xerrors.New("SERPController is nil"))
	}

	p := ListSERPWithIconRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Invalid request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	pages, err := models.ListSERP(p.TaskID, p.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	pageIDs := []int{}
	for _, v := range pages {
		pageIDs = append(pageIDs, v.PageID)
	}

	linkedPages, err := models.GetLinkedPagesByIDs(pageIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}

// ListSERPWithRatioRequest :
type ListSERPWithRatioRequest struct {
	TaskID int `json:"id" param:"id" validate:"required,numeric"`
	Offset int `json:"offset" query:"offset" validate:"required,numeric"`
	Top    int `json:"top" query:"top" validate:"numeric"`
}

// ListSERPWithRatio :
func (sc *SERPController) ListSERPWithRatio(c echo.Context) error {
	if sc == nil {
		return echo.NewHTTPError(http.StatusInternalServerError, xerrors.New("SERPController is nil"))
	}

	p := ListSERPWithRatioRequest{}
	if err := c.Bind(&p); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Invalid request body.", p))
	}

	if err := c.Validate(&p); err != nil {
		if e, ok := err.(*models.InternalError); ok {
			return echo.NewHTTPError(http.StatusBadRequest, e)
		}
		return echo.NewHTTPError(http.StatusBadRequest, models.NewInternalError(err, "Failed to validate.", p))
	}

	pages, err := models.ListSERP(p.TaskID, p.Offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	pageIDs := []int{}
	for _, v := range pages {
		pageIDs = append(pageIDs, v.PageID)
	}

	linkedPages, err := models.GetLinkedPagesByIDs(pageIDs)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}
