package comments

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/joa23/linear-cli/internal/linear/core"
)

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestResolveCommentThread(t *testing.T) {
	base := core.NewTestBaseClient("test-token", "http://example.test/graphql", &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			body := `{"data":{"commentResolve":{"success":true,"comment":{"id":"comment-123","resolvedAt":"2026-03-10T00:00:00Z"}}}}`
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(body)),
				Header:     make(http.Header),
			}, nil
		}),
	})

	client := NewClient(base)
	if err := client.ResolveCommentThread("comment-123"); err != nil {
		t.Fatalf("ResolveCommentThread() error = %v", err)
	}
}

func TestUnresolveCommentThread(t *testing.T) {
	base := core.NewTestBaseClient("test-token", "http://example.test/graphql", &http.Client{
		Transport: roundTripFunc(func(req *http.Request) (*http.Response, error) {
			body := `{"data":{"commentUnresolve":{"success":true,"comment":{"id":"comment-123","resolvedAt":null}}}}`
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(body)),
				Header:     make(http.Header),
			}, nil
		}),
	})

	client := NewClient(base)
	if err := client.UnresolveCommentThread("comment-123"); err != nil {
		t.Fatalf("UnresolveCommentThread() error = %v", err)
	}
}
