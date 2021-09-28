package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloClient(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "demo1",
			want: "hello client",
		},
		{
			name: "demo2",
			want: "hello client",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(HelloClient))
			defer ts.Close()

			resp, err := http.Get(ts.URL)
			if err != nil {
				t.Errorf("http client err: %v", err)
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("read response error")
			}
			if string(body) != tt.want {
				t.Errorf("HelloClient() = %v, want: %v", string(body), tt.want)
			}
		})
	}
}
