// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ocm_test

import (
	"context"
	"os"
	"testing"

	"github.com/andreibesleaga/ocm-go"
	"github.com/andreibesleaga/ocm-go/internal/testutil"
	"github.com/andreibesleaga/ocm-go/option"
)

func TestUsage(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	pois, err := client.Poi.List(context.TODO(), ocm.PoiListParams{})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", pois)
}
