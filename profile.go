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

// ProfileService contains methods and other services that help with interacting
// with the ocm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.Options = opts
	return
}

// Perform user authentication, returning a model which includes the users profile
// and a JWT auth token to re-use in subsequent requests.
func (r *ProfileService) Authenticate(ctx context.Context, body ProfileAuthenticateParams, opts ...option.RequestOption) (res *ProfileAuthenticateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "profile/authenticate"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ProfileAuthenticateResponse struct {
	Data     ProfileAuthenticateResponseData     `json:"Data" api:"required"`
	Metadata ProfileAuthenticateResponseMetadata `json:"Metadata" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileAuthenticateResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileAuthenticateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileAuthenticateResponseData struct {
	// JWT Bearer Token to use in subsequent authenticated requests
	AccessToken string `json:"access_token" api:"required"`
	// Full user profile, including non-public fields such as Email Address
	UserProfile ProfileAuthenticateResponseDataUserProfile `json:"UserProfile" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		UserProfile respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileAuthenticateResponseData) RawJSON() string { return r.JSON.raw }
func (r *ProfileAuthenticateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Full user profile, including non-public fields such as Email Address
type ProfileAuthenticateResponseDataUserProfile struct {
	DateCreated      string  `json:"DateCreated" api:"required"`
	ID               float64 `json:"ID" api:"required"`
	IsProfilePublic  bool    `json:"IsProfilePublic" api:"required"`
	Username         string  `json:"Username" api:"required"`
	DateLastLogin    string  `json:"DateLastLogin"`
	EmailAddress     string  `json:"EmailAddress"`
	Latitude         float64 `json:"Latitude"`
	Location         string  `json:"Location"`
	Longitude        float64 `json:"Longitude"`
	Permissions      string  `json:"Permissions"`
	Profile          string  `json:"Profile"`
	ProfileImageURL  string  `json:"ProfileImageURL"`
	ReputationPoints float64 `json:"ReputationPoints"`
	WebsiteURL       string  `json:"WebsiteURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DateCreated      respjson.Field
		ID               respjson.Field
		IsProfilePublic  respjson.Field
		Username         respjson.Field
		DateLastLogin    respjson.Field
		EmailAddress     respjson.Field
		Latitude         respjson.Field
		Location         respjson.Field
		Longitude        respjson.Field
		Permissions      respjson.Field
		Profile          respjson.Field
		ProfileImageURL  respjson.Field
		ReputationPoints respjson.Field
		WebsiteURL       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileAuthenticateResponseDataUserProfile) RawJSON() string { return r.JSON.raw }
func (r *ProfileAuthenticateResponseDataUserProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileAuthenticateResponseMetadata struct {
	StatusCode int64 `json:"StatusCode" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileAuthenticateResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ProfileAuthenticateResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProfileAuthenticateParams struct {
	Emailaddress param.Opt[string] `json:"emailaddress,omitzero"`
	Password     param.Opt[string] `json:"password,omitzero"`
	paramObj
}

func (r ProfileAuthenticateParams) MarshalJSON() (data []byte, err error) {
	type shadow ProfileAuthenticateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProfileAuthenticateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
