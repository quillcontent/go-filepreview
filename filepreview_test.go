package filepreview

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestGeneratingPreview(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `"job_id": "5f30e1c1-3fcf-4150-a5d6-25f19b4024f6", "metadata_url": "https://s3.amazonaws.com/demo.filepreviews.io/xxx/metadata.json", "preview_url": "https://s3.amazonaws.com/demo.filepreviews.io/xxx/Test Word document_320>_1.png"}`)
	}))

	svc, err := url.Parse(ts.URL)
	f := "http://foo.com/file_location"

	c := ServiceConfig{}
	c.endPoint = *svc
	c.apiKey = "xxxx123456"

	res, err := GeneratePreview(f, c)

	if err != nil {
		t.Errorf("Error: %v", err)
	}

	if res == nil {
		t.Errorf("%s", "Response should not be nil")
	}

	log.Printf("%s", res)
}
