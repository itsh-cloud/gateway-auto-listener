package main

import (
	"slices"
	"testing"
)

func TestParseNamespaceValidation(t *testing.T) {
	tests := []struct {
		name        string
		validateAll bool
		trusted     string
		prefix      string
		want        []string
		wantErr     bool
	}{
		{name: "prefix mode", prefix: "tenant-"},
		{name: "all namespaces, none trusted", validateAll: true},
		{name: "all namespaces with trusted list", validateAll: true, trusted: " platform, ,monitoring ", want: []string{"platform", "monitoring"}},
		{name: "trusted list without all namespaces", trusted: "platform", wantErr: true},
		{name: "trusted list in prefix mode", trusted: "platform", prefix: "tenant-", wantErr: true},
		{name: "all namespaces and prefix", validateAll: true, prefix: "tenant-", wantErr: true},
		{name: "pattern instead of name", validateAll: true, trusted: "tenant-*", wantErr: true},
		{name: "uppercase name", validateAll: true, trusted: "Platform", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseNamespaceValidation(tt.validateAll, tt.trusted, tt.prefix)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected an error, got %v", got)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
