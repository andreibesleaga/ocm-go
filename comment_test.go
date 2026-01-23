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

func TestCommentSubmitWithOptionalParams(t *testing.T) {
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
	_, err := client.Comment.Submit(context.TODO(), ocm.CommentSubmitParams{
		ChargePointID:       0,
		CheckinStatusTypeID: ocm.Int(0),
		Comment:             ocm.String("string"),
		CommentTypeID:       ocm.Int(0),
		Rating:              ocm.Int(3),
		RelatedURL:          ocm.String("string"),
		UserName:            ocm.String("string"),
	})
	if err != nil {
		var apierr *ocm.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
