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

// CommentService contains methods and other services that help with interacting
// with the ocm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCommentService] method instead.
type CommentService struct {
	Options []option.RequestOption
}

// NewCommentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCommentService(opts ...option.RequestOption) (r CommentService) {
	r = CommentService{}
	r.Options = opts
	return
}

// Submit a user comment or checkin for a specific charging location
func (r *CommentService) Submit(ctx context.Context, body CommentSubmitParams, opts ...option.RequestOption) (res *CommentSubmitResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "comment"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type CommentSubmitResponse struct {
	Description string `json:"description" api:"required"`
	Status      string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CommentSubmitResponse) RawJSON() string { return r.JSON.raw }
func (r *CommentSubmitResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CommentSubmitParams struct {
	// This must be a valid POI ID
	ChargePointID int64 `json:"chargePointID" api:"required"`
	// Optional valid CheckStatusTypeID to indicate overall catgeory and
	// success/failure to use equipment e.g. 10 = Charged Successfully.
	CheckinStatusTypeID param.Opt[int64] `json:"checkinStatusTypeID,omitzero"`
	// This is an optional comment to describe the charging experience, may include
	// guidance for future users.
	Comment param.Opt[string] `json:"comment,omitzero"`
	// This must be a valid Comment Type ID as per UserCommentTypes found in Core
	// Reference Data. If left as null then General Comment will be used.
	CommentTypeID param.Opt[int64] `json:"commentTypeID,omitzero"`
	// Optional integer rating between 1 = Worst, 5 = Best.
	Rating param.Opt[int64] `json:"rating,omitzero"`
	// Optional website URL for related information
	RelatedURL param.Opt[string] `json:"relatedURL,omitzero"`
	// This is an optional name to associate with the submission, for authenticated
	// users their profile username is used.
	UserName param.Opt[string] `json:"userName,omitzero"`
	paramObj
}

func (r CommentSubmitParams) MarshalJSON() (data []byte, err error) {
	type shadow CommentSubmitParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CommentSubmitParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
