// Code generated by codegen; DO NOT EDIT.

package dns

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/dns/v1"

	"google.golang.org/api/option"
)

type MockPoliciesResult struct {
	Policies []*dns.Policy `json:"policies,omitempty"`
}

func createPolicies() (*client.Services, error) {
	var item dns.Policy
	if err := faker.FakeObject(&item); err != nil {
		return nil, err
	}

	mux := httprouter.New()
	mux.GET("/*filepath", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &MockPoliciesResult{
			Policies: []*dns.Policy{&item},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)
	svc, err := dns.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Dns: svc,
	}, nil
}

func TestPolicies(t *testing.T) {
	client.MockTestHelper(t, Policies(), createPolicies, client.TestOptions{})
}