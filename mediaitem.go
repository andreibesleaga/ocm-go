// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ocm

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/ocm-go/internal/apijson"
	"github.com/stainless-sdks/ocm-go/internal/requestconfig"
	"github.com/stainless-sdks/ocm-go/option"
	"github.com/stainless-sdks/ocm-go/packages/param"
	"github.com/stainless-sdks/ocm-go/packages/respjson"
)

// MediaitemService contains methods and other services that help with interacting
// with the ocm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaitemService] method instead.
type MediaitemService struct {
	Options []option.RequestOption
}

// NewMediaitemService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMediaitemService(opts ...option.RequestOption) (r MediaitemService) {
	r = MediaitemService{}
	r.Options = opts
	return
}

// Submit a photo for a specific charging location
func (r *MediaitemService) New(ctx context.Context, body MediaitemNewParams, opts ...option.RequestOption) (res *MediaitemNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "mediaitem"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type MediaitemNewResponse struct {
	// status code OK
	Status      string `json:"status,required"`
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaitemNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaitemNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaitemNewParams struct {
	// ID value for the OCM site (POI) this image relates to.
	ChargePointID int64 `json:"chargePointID,required"`
	// BASE64 encoded data
	ImageDataBase64 string `json:"imageDataBase64,required"`
	// Optional description of image or context
	Comment param.Opt[string] `json:"comment,omitzero"`
	paramObj
}

func (r MediaitemNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaitemNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaitemNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
