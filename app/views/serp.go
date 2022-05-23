package views

import (
	"sort"
	"ushas/models"
)

// SERPWithIcon : The list of this type struct will be returned as a response of `serp` endpoint.
type SERPWithIcon struct {
	models.SearchPage

	// Linked : Users' behavioral data that probably leaked to third party.
	// For more detail, see `Linked` type.
	Linked []models.LinkedPage `json:"linked"`
}

// NewSERPWithIconView : Get search results for Icon UI
func NewSERPWithIconView(pages []models.SearchPage, pageIDs []int, linkedPages []models.LinkedPageWithPageID) []SERPWithIcon {
	serp := make([]SERPWithIcon, 0, len(pages))

	serpMap := map[int]SERPWithIcon{}
	for _, v := range pages {
		serpMap[v.PageID] = SERPWithIcon{
			SearchPage: models.SearchPage{
				PageID:  v.PageID,
				Title:   v.Title,
				URL:     v.URL,
				Snippet: v.Snippet,
			},
			Linked: []models.LinkedPage{},
		}
	}
	for _, v := range linkedPages {
		tmp := serpMap[v.PageID]
		tmp.Linked = append(tmp.Linked, models.LinkedPage{
			ID:       v.ID,
			Title:    v.Title,
			URL:      v.URL,
			Icon:     v.Icon,
			Category: v.Category,
		})
		serpMap[v.PageID] = tmp
	}

	sort.Ints(pageIDs)
	for _, v := range pageIDs {
		serp = append(serp, serpMap[v])
	}

	return serp
}

// SERPWithRatio : The list of this type struct will be returned as a response of `serp` endpoint.
type SERPWithRatio struct {
	models.SearchPage

	// Total : Total number of linked pages.
	Total int `db:"total" json:"total"`

	// Distribution : Distribution information for each categories.
	Distribution []models.CategoryCount `db:"distribution" json:"distribution"`
}

// NewSERPWithRatioView :
func NewSERPWithRatioView(pages []models.SearchPage, pageIDs []int, linkedPages []models.CategoryCountWithPageID, top int) []SERPWithRatio {
	serp := make([]SERPWithRatio, 0, len(pages))

	serpMap := map[int]SERPWithRatio{}
	for _, v := range pages {
		serpMap[v.PageID] = SERPWithRatio{
			SearchPage: models.SearchPage{
				PageID:  v.PageID,
				Title:   v.Title,
				URL:     v.URL,
				Snippet: v.Snippet,
			},
			Total:        0,
			Distribution: []models.CategoryCount{},
		}
	}

	for _, v := range linkedPages {
		tmp := serpMap[v.PageID]
		tmp.Total += v.Count
		if len(tmp.Distribution) < top {
			tmp.Distribution = append(tmp.Distribution, models.CategoryCount{
				Category: v.Category,
				Count:    v.Count,
			})
		}
		serpMap[v.PageID] = tmp
	}

	sort.Ints(pageIDs)
	for _, v := range pageIDs {
		serp = append(serp, serpMap[v])
	}

	return serp
}
