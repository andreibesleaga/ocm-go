// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ocm

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/ocm-go/internal/apijson"
	"github.com/stainless-sdks/ocm-go/internal/apiquery"
	"github.com/stainless-sdks/ocm-go/internal/requestconfig"
	"github.com/stainless-sdks/ocm-go/option"
	"github.com/stainless-sdks/ocm-go/packages/respjson"
)

// ReferencedataService contains methods and other services that help with
// interacting with the ocm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReferencedataService] method instead.
type ReferencedataService struct {
	Options []option.RequestOption
}

// NewReferencedataService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReferencedataService(opts ...option.RequestOption) (r ReferencedataService) {
	r = ReferencedataService{}
	r.Options = opts
	return
}

// Returns the core reference data used for looking up IDs such as Connection
// Types, Operators, Countries etc.
//
// This information is useful for UIs such as editing systems or for fetching
// results in the lighter non-verbose mode, then hydrating POI results back into
// complex objects.
func (r *ReferencedataService) Get(ctx context.Context, query ReferencedataGetParams, opts ...option.RequestOption) (res *ReferencedataGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "referencedata"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Country details
type Country struct {
	// The Continentcode Schema
	ContinentCode string `json:"ContinentCode" api:"required"`
	// The Id Schema
	ID int64 `json:"ID" api:"required"`
	// The Isocode Schema
	ISOCode string `json:"ISOCode" api:"required"`
	// The Title Schema
	Title string `json:"Title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ContinentCode respjson.Field
		ID            respjson.Field
		ISOCode       respjson.Field
		Title         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Country) RawJSON() string { return r.JSON.raw }
func (r *Country) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set of core reference data used for other API results and UI
type ReferencedataGetResponse struct {
	ChargerTypes          []ReferencedataGetResponseChargerType          `json:"ChargerTypes"`
	CheckinStatusTypes    []ReferencedataGetResponseCheckinStatusType    `json:"CheckinStatusTypes"`
	ConnectionTypes       []ReferencedataGetResponseConnectionType       `json:"ConnectionTypes"`
	Countries             []Country                                      `json:"Countries"`
	CurrentTypes          []ReferencedataGetResponseCurrentType          `json:"CurrentTypes"`
	DataProviders         []ReferencedataGetResponseDataProvider         `json:"DataProviders"`
	DataTypes             any                                            `json:"DataTypes"`
	MetadataGroups        string                                         `json:"MetadataGroups"`
	Operators             []ReferencedataGetResponseOperator             `json:"Operators"`
	StatusTypes           []ReferencedataGetResponseStatusType           `json:"StatusTypes"`
	SubmissionStatusTypes []ReferencedataGetResponseSubmissionStatusType `json:"SubmissionStatusTypes"`
	UsageTypes            []ReferencedataGetResponseUsageType            `json:"UsageTypes"`
	UserCommentTypes      []ReferencedataGetResponseUserCommentType      `json:"UserCommentTypes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChargerTypes          respjson.Field
		CheckinStatusTypes    respjson.Field
		ConnectionTypes       respjson.Field
		Countries             respjson.Field
		CurrentTypes          respjson.Field
		DataProviders         respjson.Field
		DataTypes             respjson.Field
		MetadataGroups        respjson.Field
		Operators             respjson.Field
		StatusTypes           respjson.Field
		SubmissionStatusTypes respjson.Field
		UsageTypes            respjson.Field
		UserCommentTypes      respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A general category for equipment power capability. Deprecated for general use.
// Currently computed automatically based on equipment power.
type ReferencedataGetResponseChargerType struct {
	Comments string `json:"Comments" api:"required"`
	ID       int64  `json:"ID" api:"required"`
	// If true, this level is considered 'fast' charging, relative to other levels.
	IsFastChargeCapable bool   `json:"IsFastChargeCapable" api:"required"`
	Title               string `json:"Title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Comments            respjson.Field
		ID                  respjson.Field
		IsFastChargeCapable respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseChargerType) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseChargerType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Classification for the users comment or experience using a specific charging
// location.
type ReferencedataGetResponseCheckinStatusType struct {
	ID int64 `json:"ID" api:"required"`
	// If true, checkin or comment was provided by an automated system.
	IsAutomatedCheckin bool `json:"IsAutomatedCheckin" api:"required"`
	// If true, this type of checkin/comment is considered positive.
	IsPositive bool   `json:"IsPositive"`
	Title      string `json:"Title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		IsAutomatedCheckin respjson.Field
		IsPositive         respjson.Field
		Title              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseCheckinStatusType) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseCheckinStatusType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of end-user connection an EVSE supports.
type ReferencedataGetResponseConnectionType struct {
	// Formal (standard) name for this connection type
	FormalName string `json:"FormalName"`
	ID         int64  `json:"ID"`
	// If true, this is an discontinued but used connection type
	IsDiscontinued bool `json:"IsDiscontinued"`
	// If true, this is an obsolete connection type and is unlikely top be present in
	// modern infrastructure
	IsObsolete bool   `json:"IsObsolete"`
	Title      string `json:"Title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FormalName     respjson.Field
		ID             respjson.Field
		IsDiscontinued respjson.Field
		IsObsolete     respjson.Field
		Title          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseConnectionType) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseConnectionType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the EVSE power supply type e.g. DC (Direct Current), AC (Single
// Phase), AC (3 Phase).
type ReferencedataGetResponseCurrentType struct {
	ID    int64  `json:"ID" api:"required"`
	Title string `json:"Title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseCurrentType) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseCurrentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Data Provider is the controller of the source data set used to construct the
// details for this POI. Data has been transformed and interpreted from it's
// original form. Each Data Provider provides data either by an explicit license or
// agreement.
type ReferencedataGetResponseDataProvider struct {
	// The reference ID for this Data Provider
	ID int64 `json:"ID" api:"required"`
	// Currently not implemented. Indicates a potential editing restriction.
	IsRestrictedEdit bool `json:"IsRestrictedEdit" api:"required"`
	// General public comments with information about this Data Provider.
	Comments string `json:"Comments"`
	// Status object describing whether this data provider is currently enabled and the
	// type of source (manual entry, imported etc)
	DataProviderStatusType ReferencedataGetResponseDataProviderDataProviderStatusType `json:"DataProviderStatusType"`
	// Date and time (UTC) the last import was performed for this data provider (if an
	// import).
	DateLastImported time.Time `json:"DateLastImported" format:"date-time"`
	// If false, data may not be imported for this provider.
	IsApprovedImport bool `json:"IsApprovedImport"`
	// If true, data provider uses an Open Data license
	IsOpenDataLicensed bool `json:"IsOpenDataLicensed"`
	// Summary of the licensing which applies for this Data Provider. Each Data
	// Provider has one specific license or agreement. Usage of the data requires
	// acceptance of the given license.
	License string `json:"License"`
	// The Title for this Data Provider
	Title string `json:"Title"`
	// Website URL for this data provider
	WebsiteURL string `json:"WebsiteURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                     respjson.Field
		IsRestrictedEdit       respjson.Field
		Comments               respjson.Field
		DataProviderStatusType respjson.Field
		DateLastImported       respjson.Field
		IsApprovedImport       respjson.Field
		IsOpenDataLicensed     respjson.Field
		License                respjson.Field
		Title                  respjson.Field
		WebsiteURL             respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseDataProvider) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseDataProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status object describing whether this data provider is currently enabled and the
// type of source (manual entry, imported etc)
type ReferencedataGetResponseDataProviderDataProviderStatusType struct {
	// The reference ID for this provider status type
	ID int64 `json:"ID" api:"required"`
	// If false, results from this data provider are not currently enabled
	IsProviderEnabled bool `json:"IsProviderEnabled" api:"required"`
	// The Title of this status type
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		IsProviderEnabled respjson.Field
		Description       respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseDataProviderDataProviderStatusType) RawJSON() string {
	return r.JSON.raw
}
func (r *ReferencedataGetResponseDataProviderDataProviderStatusType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Operator is the public organisation which controls a network of charging
// points.
type ReferencedataGetResponseOperator struct {
	// Id
	ID int64 `json:"ID" api:"required"`
	// Geographic position for site and (nearest) address component information.
	AddressInfo  ReferencedataGetResponseOperatorAddressInfo `json:"AddressInfo"`
	BookingURL   string                                      `json:"BookingURL"`
	Comments     string                                      `json:"Comments"`
	ContactEmail string                                      `json:"ContactEmail"`
	// Used to send automated notification to network operator if a user submits a
	// fault report comment/check-in
	FaultReportEmail string `json:"FaultReportEmail"`
	// If true, this operator represents a private individual
	//
	// Deprecated: deprecated
	IsPrivateIndividual bool `json:"IsPrivateIndividual"`
	// If true, this network restricts community edits for OCM data
	IsRestrictedEdit bool `json:"IsRestrictedEdit"`
	// Primary contact number for network users
	PhonePrimaryContact string `json:"PhonePrimaryContact"`
	// Secondary contact number
	PhoneSecondaryContact string `json:"PhoneSecondaryContact"`
	// Title
	Title string `json:"Title"`
	// Website for more information about this network
	WebsiteURL string `json:"WebsiteURL"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		AddressInfo           respjson.Field
		BookingURL            respjson.Field
		Comments              respjson.Field
		ContactEmail          respjson.Field
		FaultReportEmail      respjson.Field
		IsPrivateIndividual   respjson.Field
		IsRestrictedEdit      respjson.Field
		PhonePrimaryContact   respjson.Field
		PhoneSecondaryContact respjson.Field
		Title                 respjson.Field
		WebsiteURL            respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseOperator) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseOperator) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Geographic position for site and (nearest) address component information.
type ReferencedataGetResponseOperatorAddressInfo struct {
	// The reference ID for the Country
	CountryID int64 `json:"CountryID" api:"required"`
	// ID
	ID int64 `json:"ID" api:"required"`
	// Site latitude coordinate in decimal degrees
	Latitude float64 `json:"Latitude" api:"required"`
	// Site longitude coordinate in decimal degrees
	Longitude float64 `json:"Longitude" api:"required"`
	// Guidance for users to use or find the equipment
	AccessComments string `json:"AccessComments"`
	// First line of nearby street address
	AddressLine1 string `json:"AddressLine1"`
	// Second line of nearby street address
	AddressLine2 string `json:"AddressLine2"`
	// Primary contact email
	ContactEmail string `json:"ContactEmail"`
	// Primary contact number
	ContactTelephone1 string `json:"ContactTelephone1"`
	// Secondary contact number
	ContactTelephone2 string `json:"ContactTelephone2"`
	// Country details
	Country Country `json:"Country"`
	// Distance from search location, if search is around a point
	Distance float64 `json:"Distance"`
	// Unit used for distance, 1= Miles, 2 = KM
	DistanceUnit int64 `json:"DistanceUnit"`
	// Postal code or Zipcode
	Postcode string `json:"Postcode"`
	// Optional website for more information
	RelatedURL string `json:"RelatedURL"`
	// State or Province
	StateOrProvince string `json:"StateOrProvince"`
	// General title for this location to aid user
	Title string `json:"Title"`
	// Town or City
	Town string `json:"Town"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryID         respjson.Field
		ID                respjson.Field
		Latitude          respjson.Field
		Longitude         respjson.Field
		AccessComments    respjson.Field
		AddressLine1      respjson.Field
		AddressLine2      respjson.Field
		ContactEmail      respjson.Field
		ContactTelephone1 respjson.Field
		ContactTelephone2 respjson.Field
		Country           respjson.Field
		Distance          respjson.Field
		DistanceUnit      respjson.Field
		Postcode          respjson.Field
		RelatedURL        respjson.Field
		StateOrProvince   respjson.Field
		Title             respjson.Field
		Town              respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseOperatorAddressInfo) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseOperatorAddressInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Status Type of a site or equipment item indicates whether it is generally
// operational.
type ReferencedataGetResponseStatusType struct {
	ID               int64  `json:"ID" api:"required"`
	IsOperational    bool   `json:"IsOperational" api:"required"`
	IsUserSelectable bool   `json:"IsUserSelectable" api:"required"`
	Title            string `json:"Title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		IsOperational    respjson.Field
		IsUserSelectable respjson.Field
		Title            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseStatusType) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseStatusType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Submission Status object, detailing the POI listing status
type ReferencedataGetResponseSubmissionStatusType struct {
	// Submission Status Type reference ID
	ID int64 `json:"ID" api:"required"`
	// If true, POI listing is live (not draft or de-listed)
	IsLive bool   `json:"IsLive" api:"required"`
	Title  string `json:"Title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsLive      respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseSubmissionStatusType) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseSubmissionStatusType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Usage Type of a site indicates the general restrictions on usage.
type ReferencedataGetResponseUsageType struct {
	ID int64 `json:"ID" api:"required"`
	// If true this usage required a physical access key
	//
	// Deprecated: deprecated
	IsAccessKeyRequired bool `json:"IsAccessKeyRequired" api:"required"`
	// If true, this usage type requires registration or membership with a service.
	IsMembershipRequired bool `json:"IsMembershipRequired" api:"required"`
	// If true, usage requires paying at location
	IsPayAtLocation bool   `json:"IsPayAtLocation" api:"required"`
	Title           string `json:"Title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		IsAccessKeyRequired  respjson.Field
		IsMembershipRequired respjson.Field
		IsPayAtLocation      respjson.Field
		Title                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseUsageType) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseUsageType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Category for a user comment, e.g. General Comment, Fault Report (Notice To Users
// And Operator)
type ReferencedataGetResponseUserCommentType struct {
	ID    int64  `json:"ID"`
	Title string `json:"Title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReferencedataGetResponseUserCommentType) RawJSON() string { return r.JSON.raw }
func (r *ReferencedataGetResponseUserCommentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReferencedataGetParams struct {
	// Optional filter on countryid, exact match on a given numeric country id (comma
	// separated list)
	Countryid []any `query:"countryid,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReferencedataGetParams]'s query parameters as `url.Values`.
func (r ReferencedataGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
