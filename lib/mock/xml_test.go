package mock

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestXML_ContentType(t *testing.T) {
	tests := []struct {
		name string
		s    XML
		want string
	}{
		{"test_content_type", XML{Content: 1}, ContentTypeXML},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ContentType(); got != tt.want {
				t.Errorf("ContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXML_Serialize(t *testing.T) {
	tests := []struct {
		name    string
		s       XML
		want    []byte
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Serialize()
			if err != tt.wantErr {
				t.Errorf("Serialize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil && tt.want == nil {
				return
			}
			if got, err := ioutil.ReadAll(got); err != nil || !bytes.Equal(got, tt.want) {
				if err != nil {
					t.Error(err)
				} else {
					t.Errorf("Serialize() got = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
