// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ocm

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/andreibesleaga/ocm-go/internal/apijson"
	"github.com/andreibesleaga/ocm-go/internal/apiquery"
	"github.com/andreibesleaga/ocm-go/internal/requestconfig"
	"github.com/andreibesleaga/ocm-go/option"
	"github.com/andreibesleaga/ocm-go/packages/param"
	"github.com/andreibesleaga/ocm-go/packages/respjson"
)

// PoiService contains methods and other services that help with interacting with
// the ocm API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPoiService] method instead.
type PoiService struct {
	Options []option.RequestOption
}

// NewPoiService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPoiService(opts ...option.RequestOption) (r PoiService) {
	r = PoiService{}
	r.Options = opts
	return
}

// Used to fetch a list of POIs (sites) within a geographic boundary or near a
// specific latitude/longitude. This is the primary method for most applications
// and services to consume data from Open Charge Map.
func (r *PoiService) List(ctx context.Context, query PoiListParams, opts ...option.RequestOption) (res *[]PoiListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "poi"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// A POI (Point of Interest), also referred to as a `Site` or `ChargePoint`, is the
// top-level set of information regarding a geographic site with one or more
// electric vehicle charging equipment present. The term `ChargePointID` is used to
// reference the unique ID for each POI, as called OCM ID. This reference appears
// in various UI elements in the format `OCM-12345` to distinguish the ID number as
// being a reference for a specific POI/site.
//
// Note: If the API is called in verbose mode properties expanded properties are
// included in the results (e.g. UsageType, StatusType, DataProvider, OperatorInfo,
// SubmissionStatus). With the exception of the AddressInfo property, other object
// properties will not be populated in a compact result set and instead only the
// associated reference IDs will be set (e.g. UsageTypeID, DataProviderID etc)
type PoiListResponse struct {
	// Geographic position for site and (nearest) address component information.
	AddressInfo PoiListResponseAddressInfo `json:"AddressInfo"`
	// List of equipment summary information for this site
	Connections []PoiListResponseConnection `json:"Connections"`
	// A Data Provider is the controller of the source data set used to construct the
	// details for this POI. Data has been transformed and interpreted from it's
	// original form. Each Data Provider provides data either by an explicit license or
	// agreement.
	DataProvider PoiListResponseDataProvider `json:"DataProvider"`
	// The reference ID for the Data Provider of this POI
	DataProviderID int64 `json:"DataProviderID"`
	// If present, this is the Data Providers own key for this POI within their source
	// data set
	DataProvidersReference string `json:"DataProvidersReference"`
	// A metric applied during imports to indicate a quality level based on available
	// information detail (5 == best). Largely unused currently.
	DataQualityLevel int64 `json:"DataQualityLevel"`
	// The date and time (UTC, ISO 8601) this POI was added to the Open Charge Map
	// database
	DateCreated time.Time `json:"DateCreated" format:"date-time"`
	// The date and time (UTC, ISO 8601) this POI was last confirmed according to the
	// data provider or a user. See DateLastVerified for a dynamically computed date
	// based on multiple signals.
	DateLastConfirmed time.Time `json:"DateLastConfirmed" format:"date-time"`
	// The date and time (UTC, ISO 8601) this POI or directly related child properties
	// were updated.
	DateLastStatusUpdate time.Time `json:"DateLastStatusUpdate" format:"date-time"`
	// A dynamically computed value, the date and time (UTC, ISO 8601) this POI was
	// last confirmed by a user edit or related user comment
	DateLastVerified time.Time `json:"DateLastVerified" format:"date-time"`
	// The date and time (UTC, ISO 8601) this POI is or was planned for commissioning.
	// In general planned POIs should not be presented to end users until confirmed
	// operational.
	DatePlanned time.Time `json:"DatePlanned" format:"date-time"`
	// General additional factual information for the POI. Users are discouraged from
	// using this field for opinions on site quality etc.
	GeneralComments string `json:"GeneralComments"`
	// The OCM reference ID for this POI (Point of Interest).
	ID int64 `json:"ID"`
	// A dynamically computed value indicating of any recently confirmation activity
	// has taken place for this site (positive check-ins etc)
	IsRecentlyVerified bool `json:"IsRecentlyVerified"`
	// A list of user submitted photos for this site
	MediaItems []PoiListResponseMediaItem `json:"MediaItems"`
	// Optional array of metadata values. Generally used to indicate data attribution
	// but is also intended for future use to indicate surrounding amenties, links or
	// foreign key values into other data sets.
	MetadataValues []any `json:"MetadataValues"`
	// The number of bays or discreet stations available overall at this site. This
	// indicates the limiting for number of simultaneous site users.
	NumberOfPoints int64 `json:"NumberOfPoints"`
	// The reference ID of the equipment network operator or owner
	OperatorID int64 `json:"OperatorID"`
	// An Operator is the public organisation which controls a network of charging
	// points.
	OperatorInfo PoiListResponseOperatorInfo `json:"OperatorInfo"`
	// The network operators own reference for this site (may be a site reference or a
	// single equipment reference)
	OperatorsReference string `json:"OperatorsReference"`
	// If present, this data in this POI supercedes information in another POI.
	// Generally not relevant to consumers.
	ParentChargePointID int64 `json:"ParentChargePointID"`
	// The Status Type of a site or equipment item indicates whether it is generally
	// operational.
	StatusType PoiListResponseStatusType `json:"StatusType"`
	// The overall operational status type reference ID for this POI (i.e. Operational
	// etc). 0 == Unknown
	StatusTypeID int64 `json:"StatusTypeID"`
	// Submission Status object, detailing the POI listing status
	SubmissionStatus PoiListResponseSubmissionStatus `json:"SubmissionStatus"`
	// The reference ID for the submission status type which applied to this POI.
	SubmissionStatusTypeID int64 `json:"SubmissionStatusTypeID"`
	// Free text description of likely usage costs associated with this site. Generally
	// relates to parking charges whether network operates this site as Free
	UsageCost string `json:"UsageCost"`
	// The Usage Type of a site indicates the general restrictions on usage.
	UsageType PoiListResponseUsageType `json:"UsageType"`
	// The reference ID for the site Usage Type, 0 == Unknown
	UsageTypeID int64 `json:"UsageTypeID"`
	// A list of user comments or check-ins for this site
	UserComments []PoiListResponseUserComment `json:"UserComments"`
	// A universally unique identifier used as surrogate key. ID and UUID must be
	// preserved when submitting POI update information.
	Uuid string `json:"UUID" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddressInfo            respjson.Field
		Connections            respjson.Field
		DataProvider           respjson.Field
		DataProviderID         respjson.Field
		DataProvidersReference respjson.Field
		DataQualityLevel       respjson.Field
		DateCreated            respjson.Field
		DateLastConfirmed      respjson.Field
		DateLastStatusUpdate   respjson.Field
		DateLastVerified       respjson.Field
		DatePlanned            respjson.Field
		GeneralComments        respjson.Field
		ID                     respjson.Field
		IsRecentlyVerified     respjson.Field
		MediaItems             respjson.Field
		MetadataValues         respjson.Field
		NumberOfPoints         respjson.Field
		OperatorID             respjson.Field
		OperatorInfo           respjson.Field
		OperatorsReference     respjson.Field
		ParentChargePointID    respjson.Field
		StatusType             respjson.Field
		StatusTypeID           respjson.Field
		SubmissionStatus       respjson.Field
		SubmissionStatusTypeID respjson.Field
		UsageCost              respjson.Field
		UsageType              respjson.Field
		UsageTypeID            respjson.Field
		UserComments           respjson.Field
		Uuid                   respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoiListResponse) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Geographic position for site and (nearest) address component information.
type PoiListResponseAddressInfo struct {
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
func (r PoiListResponseAddressInfo) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseAddressInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details on the equipment type and power capability.
//
// If calling the API in verbose mode related models are also included in the
// result (e.g. ConnectionType, Level, StatusType, CurrentType)
type PoiListResponseConnection struct {
	// EVSE supply max current in Amps
	Amps     int64  `json:"Amps"`
	Comments string `json:"Comments"`
	// The type of end-user connection an EVSE supports.
	ConnectionType   PoiListResponseConnectionConnectionType `json:"ConnectionType"`
	ConnectionTypeID int64                                   `json:"ConnectionTypeID"`
	// Indicates the EVSE power supply type e.g. DC (Direct Current), AC (Single
	// Phase), AC (3 Phase).
	CurrentType PoiListResponseConnectionCurrentType `json:"CurrentType"`
	// The supply type reference ID (e.g. DC etc)
	CurrentTypeID int64 `json:"CurrentTypeID"`
	ID            int64 `json:"ID"`
	// A general category for equipment power capability. Deprecated for general use.
	// Currently computed automatically based on equipment power.
	Level PoiListResponseConnectionLevel `json:"Level"`
	// A general category for power capability. Depreceated in favour of documenting
	// specific equipment power in kW.
	//
	// Deprecated: deprecated
	LevelID int64 `json:"LevelID"`
	// Peak available power in kW
	PowerKw float64 `json:"PowerKW"`
	// Optional summary number of equipment items available with this specification
	Quantity int64 `json:"Quantity"`
	// Optional operators reference for this connection/port
	Reference string `json:"Reference"`
	// The Status Type of a site or equipment item indicates whether it is generally
	// operational.
	StatusType PoiListResponseConnectionStatusType `json:"StatusType"`
	// Status Type reference ID. 0 = Unknown
	StatusTypeID int64 `json:"StatusTypeID"`
	// EVSE supply voltage
	Voltage float64 `json:"Voltage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Amps             respjson.Field
		Comments         respjson.Field
		ConnectionType   respjson.Field
		ConnectionTypeID respjson.Field
		CurrentType      respjson.Field
		CurrentTypeID    respjson.Field
		ID               respjson.Field
		Level            respjson.Field
		LevelID          respjson.Field
		PowerKw          respjson.Field
		Quantity         respjson.Field
		Reference        respjson.Field
		StatusType       respjson.Field
		StatusTypeID     respjson.Field
		Voltage          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoiListResponseConnection) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of end-user connection an EVSE supports.
type PoiListResponseConnectionConnectionType struct {
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
func (r PoiListResponseConnectionConnectionType) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseConnectionConnectionType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the EVSE power supply type e.g. DC (Direct Current), AC (Single
// Phase), AC (3 Phase).
type PoiListResponseConnectionCurrentType struct {
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
func (r PoiListResponseConnectionCurrentType) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseConnectionCurrentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A general category for equipment power capability. Deprecated for general use.
// Currently computed automatically based on equipment power.
type PoiListResponseConnectionLevel struct {
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
func (r PoiListResponseConnectionLevel) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseConnectionLevel) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Status Type of a site or equipment item indicates whether it is generally
// operational.
type PoiListResponseConnectionStatusType struct {
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
func (r PoiListResponseConnectionStatusType) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseConnectionStatusType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A Data Provider is the controller of the source data set used to construct the
// details for this POI. Data has been transformed and interpreted from it's
// original form. Each Data Provider provides data either by an explicit license or
// agreement.
type PoiListResponseDataProvider struct {
	// The reference ID for this Data Provider
	ID int64 `json:"ID" api:"required"`
	// Currently not implemented. Indicates a potential editing restriction.
	IsRestrictedEdit bool `json:"IsRestrictedEdit" api:"required"`
	// General public comments with information about this Data Provider.
	Comments string `json:"Comments"`
	// Status object describing whether this data provider is currently enabled and the
	// type of source (manual entry, imported etc)
	DataProviderStatusType PoiListResponseDataProviderDataProviderStatusType `json:"DataProviderStatusType"`
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
func (r PoiListResponseDataProvider) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseDataProvider) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status object describing whether this data provider is currently enabled and the
// type of source (manual entry, imported etc)
type PoiListResponseDataProviderDataProviderStatusType struct {
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
func (r PoiListResponseDataProviderDataProviderStatusType) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseDataProviderDataProviderStatusType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A user submitted media item related to a specific charge point or site.
// Currently always an image.
type PoiListResponseMediaItem struct {
	ChargePointID      string `json:"ChargePointID"`
	Comment            string `json:"Comment"`
	DateCreated        string `json:"DateCreated"`
	ID                 string `json:"ID"`
	IsEnabled          bool   `json:"IsEnabled"`
	IsExternalResource bool   `json:"IsExternalResource"`
	IsFeaturedItem     bool   `json:"IsFeaturedItem"`
	IsVideo            bool   `json:"IsVideo"`
	ItemThumbnailURL   string `json:"ItemThumbnailURL"`
	ItemURL            string `json:"ItemURL"`
	// Short public summary profile for a specific Open Charge Map user
	User PoiListResponseMediaItemUser `json:"User"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChargePointID      respjson.Field
		Comment            respjson.Field
		DateCreated        respjson.Field
		ID                 respjson.Field
		IsEnabled          respjson.Field
		IsExternalResource respjson.Field
		IsFeaturedItem     respjson.Field
		IsVideo            respjson.Field
		ItemThumbnailURL   respjson.Field
		ItemURL            respjson.Field
		User               respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoiListResponseMediaItem) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseMediaItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Short public summary profile for a specific Open Charge Map user
type PoiListResponseMediaItemUser struct {
	ID               int64  `json:"ID"`
	ProfileImageURL  string `json:"ProfileImageURL"`
	ReputationPoints int64  `json:"ReputationPoints"`
	Username         string `json:"Username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ProfileImageURL  respjson.Field
		ReputationPoints respjson.Field
		Username         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoiListResponseMediaItemUser) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseMediaItemUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An Operator is the public organisation which controls a network of charging
// points.
type PoiListResponseOperatorInfo struct {
	// Id
	ID int64 `json:"ID" api:"required"`
	// Geographic position for site and (nearest) address component information.
	AddressInfo  PoiListResponseOperatorInfoAddressInfo `json:"AddressInfo"`
	BookingURL   string                                 `json:"BookingURL"`
	Comments     string                                 `json:"Comments"`
	ContactEmail string                                 `json:"ContactEmail"`
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
func (r PoiListResponseOperatorInfo) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseOperatorInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Geographic position for site and (nearest) address component information.
type PoiListResponseOperatorInfoAddressInfo struct {
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
func (r PoiListResponseOperatorInfoAddressInfo) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseOperatorInfoAddressInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Status Type of a site or equipment item indicates whether it is generally
// operational.
type PoiListResponseStatusType struct {
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
func (r PoiListResponseStatusType) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseStatusType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Submission Status object, detailing the POI listing status
type PoiListResponseSubmissionStatus struct {
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
func (r PoiListResponseSubmissionStatus) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseSubmissionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Usage Type of a site indicates the general restrictions on usage.
type PoiListResponseUsageType struct {
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
func (r PoiListResponseUsageType) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseUsageType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A user comment or check-in for a specific charging point (POI/Site)
type PoiListResponseUserComment struct {
	ChargePointID int64 `json:"ChargePointID"`
	// Classification for the users comment or experience using a specific charging
	// location.
	CheckinStatusType   PoiListResponseUserCommentCheckinStatusType `json:"CheckinStatusType"`
	CheckinStatusTypeID int64                                       `json:"CheckinStatusTypeID"`
	Comment             string                                      `json:"Comment"`
	// Category for a user comment, e.g. General Comment, Fault Report (Notice To Users
	// And Operator)
	CommentType   PoiListResponseUserCommentCommentType `json:"CommentType"`
	CommentTypeID int64                                 `json:"CommentTypeID"`
	DateCreated   time.Time                             `json:"DateCreated" format:"date-time"`
	ID            string                                `json:"ID"`
	RelatedURL    string                                `json:"RelatedURL"`
	// Short public summary profile for a specific Open Charge Map user
	User     PoiListResponseUserCommentUser `json:"User"`
	UserName string                         `json:"UserName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChargePointID       respjson.Field
		CheckinStatusType   respjson.Field
		CheckinStatusTypeID respjson.Field
		Comment             respjson.Field
		CommentType         respjson.Field
		CommentTypeID       respjson.Field
		DateCreated         respjson.Field
		ID                  respjson.Field
		RelatedURL          respjson.Field
		User                respjson.Field
		UserName            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoiListResponseUserComment) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseUserComment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Classification for the users comment or experience using a specific charging
// location.
type PoiListResponseUserCommentCheckinStatusType struct {
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
func (r PoiListResponseUserCommentCheckinStatusType) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseUserCommentCheckinStatusType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Category for a user comment, e.g. General Comment, Fault Report (Notice To Users
// And Operator)
type PoiListResponseUserCommentCommentType struct {
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
func (r PoiListResponseUserCommentCommentType) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseUserCommentCommentType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Short public summary profile for a specific Open Charge Map user
type PoiListResponseUserCommentUser struct {
	ID               int64  `json:"ID"`
	ProfileImageURL  string `json:"ProfileImageURL"`
	ReputationPoints int64  `json:"ReputationPoints"`
	Username         string `json:"Username"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ProfileImageURL  respjson.Field
		ReputationPoints respjson.Field
		Username         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PoiListResponseUserCommentUser) RawJSON() string { return r.JSON.raw }
func (r *PoiListResponseUserCommentUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PoiListParams struct {
	// Set to true to get a property names in camelCase format.
	Camelcase param.Opt[bool] `query:"camelcase,omitzero" json:"-"`
	// Exact match on a given OCM POI ID (comma separated list)
	Chargepointid param.Opt[string] `query:"chargepointid,omitzero" json:"-"`
	// String to identify your client application. Optional but recommended to
	// distinguish your client from other bots/crawlers
	Client param.Opt[string] `query:"client,omitzero" json:"-"`
	// Set to true to remove reference data objects from output (just returns IDs for
	// common reference data such as DataProvider etc).
	Compact param.Opt[bool] `query:"compact,omitzero" json:"-"`
	// 2-character ISO Country code to filter to one specific country
	Countrycode param.Opt[string] `query:"countrycode,omitzero" json:"-"`
	// Optionally filter results by a max distance from the given latitude/longitude
	Distance param.Opt[float64] `query:"distance,omitzero" json:"-"`
	// `miles` or `km` distance unit
	Distanceunit param.Opt[string] `query:"distanceunit,omitzero" json:"-"`
	// Filter to items with ID greater than given value
	Greaterthanid param.Opt[string] `query:"greaterthanid,omitzero" json:"-"`
	// If true, user comments and media items will be include in result set
	Includecomments param.Opt[bool] `query:"includecomments,omitzero" json:"-"`
	// Latitude for distance calculation and filtering
	Latitude param.Opt[int64] `query:"latitude,omitzero" json:"-"`
	// Longitude for distance calculation and filtering
	Longitude param.Opt[float64] `query:"longitude,omitzero" json:"-"`
	// Limit on max number of results returned
	Maxresults param.Opt[int64] `query:"maxresults,omitzero" json:"-"`
	// Filter to results modified after the given date
	Modifiedsince param.Opt[string] `query:"modifiedsince,omitzero" json:"-"`
	// Use opendata=true for only OCM provided ("Open") data.
	Opendata param.Opt[bool] `query:"opendata,omitzero" json:"-"`
	// Optional output format `json`,`geojson`,`xml`,`csv`, JSON is the default and
	// recommended as the highest fidelity.
	Output param.Opt[string] `query:"output,omitzero" json:"-"`
	// Filter results within a given Polygon. Specify an encoded polyline for the
	// polygon shape. Polygon will be automatically closed from the last point to the
	// first point.
	Polygon param.Opt[string] `query:"polygon,omitzero" json:"-"`
	// Filter results along an encoded polyline, use with distance param to increase
	// search distance along line. Polyline is expanded into a polygon to cover the
	// search distance.
	Polyline param.Opt[string] `query:"polyline,omitzero" json:"-"`
	// Default sort order is based on spatial index but you can optionally sort by
	// `modified_asc` for results in order of modification (oldest to newest), or
	// ` id_asc` for results in order of ID
	Sortby param.Opt[string] `query:"sortby,omitzero" json:"-"`
	// Set to false to get a smaller result set with null items removed.
	Verbose param.Opt[bool] `query:"verbose,omitzero" json:"-"`
	// Filter results to a given bounding box. specify top left and bottom right box
	// corners as: (lat,lng),(lat2,lng2)
	Boundingbox []any `query:"boundingbox,omitzero" json:"-"`
	// Exact match on a given connection type id (comma separated list)
	Connectiontypeid []any `query:"connectiontypeid,omitzero" json:"-"`
	// Exact match on a given numeric country id (comma separated list)
	Countryid []string `query:"countryid,omitzero" json:"-"`
	// Exact match on a given data provider id id (comma separated list).
	Dataproviderid []any `query:"dataproviderid,omitzero" json:"-"`
	// Exact match on a given charging level (1-3) id (comma separated list)
	Levelid []any `query:"levelid,omitzero" json:"-"`
	// Exact match on a given EVSE operator id (comma separated list)
	Operatorid []any `query:"operatorid,omitzero" json:"-"`
	// Exact match on a given status type id (comma separated list)
	Statustypeid []any `query:"statustypeid,omitzero" json:"-"`
	// Exact match on a given usage type id (comma separated list)
	Usagetypeid []any `query:"usagetypeid,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PoiListParams]'s query parameters as `url.Values`.
func (r PoiListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
