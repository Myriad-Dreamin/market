package mock

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestJSON_ContentType(t *testing.T) {
	tests := []struct {
		name string
		s    JSON
		want string
	}{
		{"test_content_type", JSON{Content: 1}, ContentTypeJSON},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ContentType(); got != tt.want {
				t.Errorf("ContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJSON_Serialize(t *testing.T) {
	tests := []struct {
		name    string
		s       JSON
		want    []byte
		wantErr error
	}{
		{name: "test_base_type_bytes", s: JSON{Content: 1}, want: []byte("1\n"), wantErr: nil},
		{name: "test_base_type_bytes", s: JSON{Content: "1"}, want: []byte("\"1\"\n"), wantErr: nil},
		{name: "test_base_type_bytes", s: JSON{Content: 1.0}, want: []byte("1\n"), wantErr: nil},
		{name: "test_base_type_bytes", s: JSON{Content: true}, want: []byte("true\n"), wantErr: nil},
		{name: "test_base_type_bytes", s: JSON{Content: struct {
			S int `json:"s"`
		}{S: 1}}, want: []byte("{\"s\":1}\n"), wantErr: nil},
		{name: "test_base_type_bytes", s: JSON{Content: 1}, want: []byte("1\n"), wantErr: nil},
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
