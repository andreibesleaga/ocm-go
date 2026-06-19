// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ocm_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/andreibesleaga/ocm-go"
	"github.com/andreibesleaga/ocm-go/internal/testutil"
	"github.com/andreibesleaga/ocm-go/option"
)

func TestMediaitemNewWithOptionalParams(t *testing.T) {
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
		option.WithBearer("My Bearer"),
	)
	_, err := client.Mediaitem.New(context.TODO(), ocm.MediaitemNewParams{
		ChargePointID:   1234,
		ImageDataBase64: "data:image/jpeg;base64,<BASE64_ENCODED_DATA>",
		Comment:         ocm.String("An example comment"),
	})
	if err != nil {
		var apierr *ocm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
