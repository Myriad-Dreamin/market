package mock

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestProtoMessage_ContentType(t *testing.T) {
	tests := []struct {
		name string
		s    ProtoMessage
		want string
	}{
		{"test_content_type", ProtoMessage{Content: 1}, ContentTypeProtocolBufBinary},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.ContentType(); got != tt.want {
				t.Errorf("ContentType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProtoMessage_Serialize(t *testing.T) {
	tests := []struct {
		name    string
		s       ProtoMessage
		want    []byte
		wantErr error
	}{
		{name: "test_base_type_bytes", s: ProtoMessage{Content:[]byte("1")}, want: nil, wantErr:ErrNotProtoMessage},
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