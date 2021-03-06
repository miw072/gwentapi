// API "gwentapi": Application Contexts
//
// Code generated by goagen v1.1.0-dirty, DO NOT EDIT.
//
// Command:
// $ goagen
// --design=github.com/tri125/gwentapi/design
// --out=$(GOPATH)\src\github.com\tri125\gwentapi
// --version=v1.1.0

package app

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"unicode/utf8"
)

// CardFactionCardContext provides the card cardFaction action context.
type CardFactionCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	FactionID string
	Limit     int
	Offset    int
}

// NewCardFactionCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardFaction action.
func NewCardFactionCardContext(ctx context.Context, r *http.Request, service *goa.Service) (*CardFactionCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CardFactionCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFactionID := req.Params["factionID"]
	if len(paramFactionID) > 0 {
		rawFactionID := paramFactionID[0]
		rctx.FactionID = rawFactionID
	}
	paramLimit := req.Params["limit"]
	if len(paramLimit) == 0 {
		rctx.Limit = 20
	} else {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			rctx.Limit = limit
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
		if rctx.Limit < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`limit`, rctx.Limit, 1, true))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) == 0 {
		rctx.Offset = 0
	} else {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			rctx.Offset = offset
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
		if rctx.Offset < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`offset`, rctx.Offset, 0, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardFactionCardContext) OK(r *GwentapiPagecard) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.pagecard+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardFactionCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardFactionCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// CardLeaderCardContext provides the card cardLeader action context.
type CardLeaderCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Limit  int
	Offset int
}

// NewCardLeaderCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardLeader action.
func NewCardLeaderCardContext(ctx context.Context, r *http.Request, service *goa.Service) (*CardLeaderCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CardLeaderCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLimit := req.Params["limit"]
	if len(paramLimit) == 0 {
		rctx.Limit = 20
	} else {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			rctx.Limit = limit
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
		if rctx.Limit < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`limit`, rctx.Limit, 1, true))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) == 0 {
		rctx.Offset = 0
	} else {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			rctx.Offset = offset
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
		if rctx.Offset < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`offset`, rctx.Offset, 0, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardLeaderCardContext) OK(r *GwentapiPagecard) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.pagecard+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardLeaderCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardLeaderCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// CardVariationCardContext provides the card cardVariation action context.
type CardVariationCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	CardID      string
	Limit       int
	Offset      int
	VariationID string
}

// NewCardVariationCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardVariation action.
func NewCardVariationCardContext(ctx context.Context, r *http.Request, service *goa.Service) (*CardVariationCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CardVariationCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCardID := req.Params["cardID"]
	if len(paramCardID) > 0 {
		rawCardID := paramCardID[0]
		rctx.CardID = rawCardID
	}
	paramLimit := req.Params["limit"]
	if len(paramLimit) == 0 {
		rctx.Limit = 20
	} else {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			rctx.Limit = limit
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
		if rctx.Limit < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`limit`, rctx.Limit, 1, true))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) == 0 {
		rctx.Offset = 0
	} else {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			rctx.Offset = offset
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
		if rctx.Offset < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`offset`, rctx.Offset, 0, true))
		}
	}
	paramVariationID := req.Params["variationID"]
	if len(paramVariationID) > 0 {
		rawVariationID := paramVariationID[0]
		rctx.VariationID = rawVariationID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardVariationCardContext) OK(r *GwentapiVariation) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.variation+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *CardVariationCardContext) OKLink(r *GwentapiVariationLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.variation+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardVariationCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardVariationCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// CardVariationsCardContext provides the card cardVariations action context.
type CardVariationsCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	CardID string
	Limit  int
	Offset int
}

// NewCardVariationsCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller cardVariations action.
func NewCardVariationsCardContext(ctx context.Context, r *http.Request, service *goa.Service) (*CardVariationsCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := CardVariationsCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCardID := req.Params["cardID"]
	if len(paramCardID) > 0 {
		rawCardID := paramCardID[0]
		rctx.CardID = rawCardID
	}
	paramLimit := req.Params["limit"]
	if len(paramLimit) == 0 {
		rctx.Limit = 20
	} else {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			rctx.Limit = limit
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
		if rctx.Limit < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`limit`, rctx.Limit, 1, true))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) == 0 {
		rctx.Offset = 0
	} else {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			rctx.Offset = offset
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
		if rctx.Offset < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`offset`, rctx.Offset, 0, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *CardVariationsCardContext) OK(r GwentapiVariationCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.variation+json; type=collection")
	if r == nil {
		r = GwentapiVariationCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *CardVariationsCardContext) OKLink(r GwentapiVariationLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.variation+json; type=collection")
	if r == nil {
		r = GwentapiVariationLinkCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *CardVariationsCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *CardVariationsCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListCardContext provides the card list action context.
type ListCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Limit  int
	Name   *string
	Offset int
}

// NewListCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller list action.
func NewListCardContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramLimit := req.Params["limit"]
	if len(paramLimit) == 0 {
		rctx.Limit = 20
	} else {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			rctx.Limit = limit
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
		if rctx.Limit < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`limit`, rctx.Limit, 1, true))
		}
	}
	paramName := req.Params["name"]
	if len(paramName) > 0 {
		rawName := paramName[0]
		rctx.Name = &rawName
		if rctx.Name != nil {
			if utf8.RuneCountInString(*rctx.Name) < 3 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`name`, *rctx.Name, utf8.RuneCountInString(*rctx.Name), 3, true))
			}
		}
		if rctx.Name != nil {
			if utf8.RuneCountInString(*rctx.Name) > 10 {
				err = goa.MergeErrors(err, goa.InvalidLengthError(`name`, *rctx.Name, utf8.RuneCountInString(*rctx.Name), 10, false))
			}
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) == 0 {
		rctx.Offset = 0
	} else {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			rctx.Offset = offset
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
		if rctx.Offset < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`offset`, rctx.Offset, 0, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCardContext) OK(r *GwentapiPagecard) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.pagecard+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowCardContext provides the card show action context.
type ShowCardContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	CardID string
	Limit  int
	Offset int
}

// NewShowCardContext parses the incoming request URL and body, performs validations and creates the
// context used by the card controller show action.
func NewShowCardContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowCardContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowCardContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCardID := req.Params["cardID"]
	if len(paramCardID) > 0 {
		rawCardID := paramCardID[0]
		rctx.CardID = rawCardID
	}
	paramLimit := req.Params["limit"]
	if len(paramLimit) == 0 {
		rctx.Limit = 20
	} else {
		rawLimit := paramLimit[0]
		if limit, err2 := strconv.Atoi(rawLimit); err2 == nil {
			rctx.Limit = limit
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("limit", rawLimit, "integer"))
		}
		if rctx.Limit < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`limit`, rctx.Limit, 1, true))
		}
	}
	paramOffset := req.Params["offset"]
	if len(paramOffset) == 0 {
		rctx.Offset = 0
	} else {
		rawOffset := paramOffset[0]
		if offset, err2 := strconv.Atoi(rawOffset); err2 == nil {
			rctx.Offset = offset
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("offset", rawOffset, "integer"))
		}
		if rctx.Offset < 0 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`offset`, rctx.Offset, 0, true))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCardContext) OK(r *GwentapiCard) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.card+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowCardContext) OKLink(r *GwentapiCardLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.card+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCardContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowCardContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListCategoryContext provides the category list action context.
type ListCategoryContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListCategoryContext parses the incoming request URL and body, performs validations and creates the
// context used by the category controller list action.
func NewListCategoryContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListCategoryContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListCategoryContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListCategoryContext) OK(r GwentapiCategoryCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.category+json; type=collection")
	if r == nil {
		r = GwentapiCategoryCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListCategoryContext) OKLink(r GwentapiCategoryLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.category+json; type=collection")
	if r == nil {
		r = GwentapiCategoryLinkCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListCategoryContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListCategoryContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowCategoryContext provides the category show action context.
type ShowCategoryContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	CategoryID string
}

// NewShowCategoryContext parses the incoming request URL and body, performs validations and creates the
// context used by the category controller show action.
func NewShowCategoryContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowCategoryContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowCategoryContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramCategoryID := req.Params["categoryID"]
	if len(paramCategoryID) > 0 {
		rawCategoryID := paramCategoryID[0]
		rctx.CategoryID = rawCategoryID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowCategoryContext) OK(r *GwentapiCategory) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.category+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowCategoryContext) OKLink(r *GwentapiCategoryLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.category+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowCategoryContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowCategoryContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListFactionContext provides the faction list action context.
type ListFactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListFactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the faction controller list action.
func NewListFactionContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListFactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListFactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListFactionContext) OK(r GwentapiFactionCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.faction+json; type=collection")
	if r == nil {
		r = GwentapiFactionCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListFactionContext) OKLink(r GwentapiFactionLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.faction+json; type=collection")
	if r == nil {
		r = GwentapiFactionLinkCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListFactionContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListFactionContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowFactionContext provides the faction show action context.
type ShowFactionContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	FactionID string
}

// NewShowFactionContext parses the incoming request URL and body, performs validations and creates the
// context used by the faction controller show action.
func NewShowFactionContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowFactionContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowFactionContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramFactionID := req.Params["factionID"]
	if len(paramFactionID) > 0 {
		rawFactionID := paramFactionID[0]
		rctx.FactionID = rawFactionID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowFactionContext) OK(r *GwentapiFaction) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.faction+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowFactionContext) OKLink(r *GwentapiFactionLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.faction+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowFactionContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowFactionContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ListGroupContext provides the group list action context.
type ListGroupContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListGroupContext parses the incoming request URL and body, performs validations and creates the
// context used by the group controller list action.
func NewListGroupContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListGroupContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListGroupContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListGroupContext) OK(r GwentapiGroupCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.group+json; type=collection")
	if r == nil {
		r = GwentapiGroupCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListGroupContext) OKLink(r GwentapiGroupLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.group+json; type=collection")
	if r == nil {
		r = GwentapiGroupLinkCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListGroupContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListGroupContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowGroupContext provides the group show action context.
type ShowGroupContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	GroupID string
}

// NewShowGroupContext parses the incoming request URL and body, performs validations and creates the
// context used by the group controller show action.
func NewShowGroupContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowGroupContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowGroupContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramGroupID := req.Params["groupID"]
	if len(paramGroupID) > 0 {
		rawGroupID := paramGroupID[0]
		rctx.GroupID = rawGroupID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowGroupContext) OK(r *GwentapiGroup) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.group+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowGroupContext) OKLink(r *GwentapiGroupLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.group+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowGroupContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowGroupContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowIndexContext provides the index show action context.
type ShowIndexContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewShowIndexContext parses the incoming request URL and body, performs validations and creates the
// context used by the index controller show action.
func NewShowIndexContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowIndexContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowIndexContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowIndexContext) OK(r *GwentapiResource) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.resource+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowIndexContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// ListRarityContext provides the rarity list action context.
type ListRarityContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
}

// NewListRarityContext parses the incoming request URL and body, performs validations and creates the
// context used by the rarity controller list action.
func NewListRarityContext(ctx context.Context, r *http.Request, service *goa.Service) (*ListRarityContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ListRarityContext{Context: ctx, ResponseData: resp, RequestData: req}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ListRarityContext) OK(r GwentapiRarityCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.rarity+json; type=collection")
	if r == nil {
		r = GwentapiRarityCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ListRarityContext) OKLink(r GwentapiRarityLinkCollection) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.rarity+json; type=collection")
	if r == nil {
		r = GwentapiRarityLinkCollection{}
	}
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ListRarityContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ListRarityContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}

// ShowRarityContext provides the rarity show action context.
type ShowRarityContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	RarityID string
}

// NewShowRarityContext parses the incoming request URL and body, performs validations and creates the
// context used by the rarity controller show action.
func NewShowRarityContext(ctx context.Context, r *http.Request, service *goa.Service) (*ShowRarityContext, error) {
	var err error
	resp := goa.ContextResponse(ctx)
	resp.Service = service
	req := goa.ContextRequest(ctx)
	req.Request = r
	rctx := ShowRarityContext{Context: ctx, ResponseData: resp, RequestData: req}
	paramRarityID := req.Params["rarityID"]
	if len(paramRarityID) > 0 {
		rawRarityID := paramRarityID[0]
		rctx.RarityID = rawRarityID
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowRarityContext) OK(r *GwentapiRarity) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.rarity+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// OKLink sends a HTTP response with status code 200.
func (ctx *ShowRarityContext) OKLink(r *GwentapiRarityLink) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.gwentapi.rarity+json")
	return ctx.ResponseData.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowRarityContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}

// InternalServerError sends a HTTP response with status code 500.
func (ctx *ShowRarityContext) InternalServerError() error {
	ctx.ResponseData.WriteHeader(500)
	return nil
}
