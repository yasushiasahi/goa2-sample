// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// viron HTTP client types
//
// Command:
// $ goa gen github.com/tonouchi510/goa2-sample/design

package client

import (
	vironviews "github.com/tonouchi510/goa2-sample/gen/viron/views"
	goa "goa.design/goa"
)

// AuthtypeResponseBody is the type of the "viron" service "authtype" endpoint
// HTTP response body.
type AuthtypeResponseBody []*VironAuthtypeResponse

// VironMenuResponseBody is the type of the "viron" service "viron_menu"
// endpoint HTTP response body.
type VironMenuResponseBody struct {
	Name      *string                     `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Thumbnail *string                     `form:"thumbnail,omitempty" json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	Tags      []string                    `form:"tags,omitempty" json:"tags,omitempty" xml:"tags,omitempty"`
	Color     *string                     `form:"color,omitempty" json:"color,omitempty" xml:"color,omitempty"`
	Theme     *string                     `form:"theme,omitempty" json:"theme,omitempty" xml:"theme,omitempty"`
	Pages     []*VironPageResponseBody    `form:"pages,omitempty" json:"pages,omitempty" xml:"pages,omitempty"`
	Sections  []*VironSectionResponseBody `form:"sections,omitempty" json:"sections,omitempty" xml:"sections,omitempty"`
}

// VironAuthtypeResponse is used to define fields on response body types.
type VironAuthtypeResponse struct {
	// type name
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
	// provider name
	Provider *string `form:"provider,omitempty" json:"provider,omitempty" xml:"provider,omitempty"`
	// url
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
	// request method to submit this auth
	Method *string `form:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
}

// VironPageResponseBody is used to define fields on response body types.
type VironPageResponseBody struct {
	ID         *string                       `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Name       *string                       `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Section    *string                       `form:"section,omitempty" json:"section,omitempty" xml:"section,omitempty"`
	Group      *string                       `form:"group,omitempty" json:"group,omitempty" xml:"group,omitempty"`
	Components []*VironComponentResponseBody `form:"components,omitempty" json:"components,omitempty" xml:"components,omitempty"`
}

// VironComponentResponseBody is used to define fields on response body types.
type VironComponentResponseBody struct {
	Name           *string                   `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Style          *string                   `form:"style,omitempty" json:"style,omitempty" xml:"style,omitempty"`
	Unit           *string                   `form:"unit,omitempty" json:"unit,omitempty" xml:"unit,omitempty"`
	Actions        []string                  `form:"actions,omitempty" json:"actions,omitempty" xml:"actions,omitempty"`
	API            *VironAPIResponseBody     `form:"api,omitempty" json:"api,omitempty" xml:"api,omitempty"`
	Pagination     *bool                     `form:"pagination,omitempty" json:"pagination,omitempty" xml:"pagination,omitempty"`
	Primary        *string                   `form:"primary,omitempty" json:"primary,omitempty" xml:"primary,omitempty"`
	TableLabels    []string                  `form:"table_labels,omitempty" json:"table_labels,omitempty" xml:"table_labels,omitempty"`
	Query          []*VironQueryResponseBody `form:"query,omitempty" json:"query,omitempty" xml:"query,omitempty"`
	AutoRefreshSec *int32                    `form:"auto_refresh_sec,omitempty" json:"auto_refresh_sec,omitempty" xml:"auto_refresh_sec,omitempty"`
}

// VironAPIResponseBody is used to define fields on response body types.
type VironAPIResponseBody struct {
	Method *string `form:"method,omitempty" json:"method,omitempty" xml:"method,omitempty"`
	Path   *string `form:"path,omitempty" json:"path,omitempty" xml:"path,omitempty"`
}

// VironQueryResponseBody is used to define fields on response body types.
type VironQueryResponseBody struct {
	Key  *string `form:"key,omitempty" json:"key,omitempty" xml:"key,omitempty"`
	Type *string `form:"type,omitempty" json:"type,omitempty" xml:"type,omitempty"`
}

// VironSectionResponseBody is used to define fields on response body types.
type VironSectionResponseBody struct {
	ID    *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	Label *string `form:"label,omitempty" json:"label,omitempty" xml:"label,omitempty"`
}

// NewAuthtypeVironAuthtypeCollectionOK builds a "viron" service "authtype"
// endpoint result from a HTTP "OK" response.
func NewAuthtypeVironAuthtypeCollectionOK(body AuthtypeResponseBody) vironviews.VironAuthtypeCollectionView {
	v := make([]*vironviews.VironAuthtypeView, len(body))
	for i, val := range body {
		v[i] = &vironviews.VironAuthtypeView{
			Type:     val.Type,
			Provider: val.Provider,
			URL:      val.URL,
			Method:   val.Method,
		}
	}
	return v
}

// NewVironMenuViewOK builds a "viron" service "viron_menu" endpoint result
// from a HTTP "OK" response.
func NewVironMenuViewOK(body *VironMenuResponseBody) *vironviews.VironMenuView {
	v := &vironviews.VironMenuView{
		Name:      body.Name,
		Thumbnail: body.Thumbnail,
		Color:     body.Color,
		Theme:     body.Theme,
	}
	if body.Tags != nil {
		v.Tags = make([]string, len(body.Tags))
		for i, val := range body.Tags {
			v.Tags[i] = val
		}
	}
	v.Pages = make([]*vironviews.VironPageView, len(body.Pages))
	for i, val := range body.Pages {
		v.Pages[i] = &vironviews.VironPageView{
			ID:      val.ID,
			Name:    val.Name,
			Section: val.Section,
			Group:   val.Group,
		}
		v.Pages[i].Components = make([]*vironviews.VironComponentView, len(val.Components))
		for j, val := range val.Components {
			v.Pages[i].Components[j] = &vironviews.VironComponentView{
				Name:           val.Name,
				Style:          val.Style,
				Unit:           val.Unit,
				Pagination:     val.Pagination,
				Primary:        val.Primary,
				AutoRefreshSec: val.AutoRefreshSec,
			}
			if val.Actions != nil {
				v.Pages[i].Components[j].Actions = make([]string, len(val.Actions))
				for k, val := range val.Actions {
					v.Pages[i].Components[j].Actions[k] = val
				}
			}
			v.Pages[i].Components[j].API = unmarshalVironAPIResponseBodyToVironAPIView(val.API)
			if val.TableLabels != nil {
				v.Pages[i].Components[j].TableLabels = make([]string, len(val.TableLabels))
				for k, val := range val.TableLabels {
					v.Pages[i].Components[j].TableLabels[k] = val
				}
			}
			if val.Query != nil {
				v.Pages[i].Components[j].Query = make([]*vironviews.VironQueryView, len(val.Query))
				for k, val := range val.Query {
					v.Pages[i].Components[j].Query[k] = &vironviews.VironQueryView{
						Key:  val.Key,
						Type: val.Type,
					}
				}
			}
		}
	}
	if body.Sections != nil {
		v.Sections = make([]*vironviews.VironSectionView, len(body.Sections))
		for i, val := range body.Sections {
			v.Sections[i] = &vironviews.VironSectionView{
				ID:    val.ID,
				Label: val.Label,
			}
		}
	}
	return v
}

// ValidateVironAuthtypeResponse runs the validations defined on
// Viron_authtypeResponse
func ValidateVironAuthtypeResponse(body *VironAuthtypeResponse) (err error) {
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	if body.Provider == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("provider", "body"))
	}
	if body.URL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("url", "body"))
	}
	if body.Method == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("method", "body"))
	}
	return
}

// ValidateVironPageResponseBody runs the validations defined on
// Viron_pageResponseBody
func ValidateVironPageResponseBody(body *VironPageResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Section == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("section", "body"))
	}
	if body.Components == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("components", "body"))
	}
	for _, e := range body.Components {
		if e != nil {
			if err2 := ValidateVironComponentResponseBody(e); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ValidateVironComponentResponseBody runs the validations defined on
// Viron_componentResponseBody
func ValidateVironComponentResponseBody(body *VironComponentResponseBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Style == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("style", "body"))
	}
	if body.API == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("api", "body"))
	}
	if body.Style != nil {
		if !(*body.Style == "number" || *body.Style == "table" || *body.Style == "graph-bar" || *body.Style == "graph-scatterplot" || *body.Style == "graph-line" || *body.Style == "graph-horizontal-bar" || *body.Style == "graph-stacked-bar" || *body.Style == "graph-horizontal-stacked-bar" || *body.Style == "graph-stacked-area") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError("body.style", *body.Style, []interface{}{"number", "table", "graph-bar", "graph-scatterplot", "graph-line", "graph-horizontal-bar", "graph-stacked-bar", "graph-horizontal-stacked-bar", "graph-stacked-area"}))
		}
	}
	if body.API != nil {
		if err2 := ValidateVironAPIResponseBody(body.API); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateVironAPIResponseBody runs the validations defined on
// Viron_apiResponseBody
func ValidateVironAPIResponseBody(body *VironAPIResponseBody) (err error) {
	if body.Method == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("method", "body"))
	}
	if body.Path == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("path", "body"))
	}
	return
}

// ValidateVironQueryResponseBody runs the validations defined on
// Viron_queryResponseBody
func ValidateVironQueryResponseBody(body *VironQueryResponseBody) (err error) {
	if body.Key == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("key", "body"))
	}
	if body.Type == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("type", "body"))
	}
	return
}

// ValidateVironSectionResponseBody runs the validations defined on
// Viron_sectionResponseBody
func ValidateVironSectionResponseBody(body *VironSectionResponseBody) (err error) {
	if body.ID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("id", "body"))
	}
	if body.Label == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("label", "body"))
	}
	return
}
