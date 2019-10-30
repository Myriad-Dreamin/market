package config

import "testing"

func TestLoad(t *testing.T) {
	type args struct {
		config     *ServerConfig
		configpath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test-load-json", args{
			new(ServerConfig), "server-config.example.json",
		}, false},
		{"test-load-try", args{
			new(ServerConfig), "server-config.example",
		}, false},
		{"test-load-nothing", args{
			new(ServerConfig), "server-config.error",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Load(tt.args.config, tt.args.configpath); (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoadStatic(t *testing.T) {
	type args struct {
		config     interface{}
		configpath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"test-load-json", args{
			new(ServerConfig), "server-config.example.json",
		}, false},
		{"test-load-try", args{
			new(ServerConfig), "server-config.example",
		}, false},
		{"test-load-nothing", args{
			new(ServerConfig), "server-config.error",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := LoadStatic(tt.args.config, tt.args.configpath); (err != nil) != tt.wantErr {
				t.Errorf("LoadStatic() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSave(t *testing.T) {
	type args struct {
		config     *ServerConfig
		configpath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Save(tt.args.config, tt.args.configpath); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}