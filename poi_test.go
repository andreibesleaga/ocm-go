// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ocm_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/ocm-go"
	"github.com/stainless-sdks/ocm-go/internal/testutil"
	"github.com/stainless-sdks/ocm-go/option"
)

func TestPoiListWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ocm.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Poi.List(context.TODO(), ocm.PoiListParams{
		Boundingbox:      []any{map[string]any{}},
		Camelcase:        ocm.Bool(true),
		Chargepointid:    ocm.String("chargepointid"),
		Client:           ocm.String("client"),
		Compact:          ocm.Bool(true),
		Connectiontypeid: []any{map[string]any{}},
		Countrycode:      ocm.String("countrycode"),
		Countryid:        []string{"string"},
		Dataproviderid:   []any{map[string]any{}},
		Distance:         ocm.Float(0),
		Distanceunit:     ocm.String("distanceunit"),
		Greaterthanid:    ocm.String("greaterthanid"),
		Includecomments:  ocm.Bool(true),
		Latitude:         ocm.Int(0),
		Levelid:          []any{map[string]any{}},
		Longitude:        ocm.Float(0),
		Maxresults:       ocm.Int(0),
		Modifiedsince:    ocm.String("modifiedsince"),
		Opendata:         ocm.Bool(true),
		Operatorid:       []any{map[string]any{}},
		Output:           ocm.String("output"),
		Polygon:          ocm.String("polygon"),
		Polyline:         ocm.String("polyline"),
		Sortby:           ocm.String("sortby"),
		Statustypeid:     []any{map[string]any{}},
		Usagetypeid:      []any{map[string]any{}},
		Verbose:          ocm.Bool(true),
	})
	if err != nil {
		var apierr *ocm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
