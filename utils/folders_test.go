package utils

import (
	"testing"
)

func TestCreateFile(t *testing.T) {
	type args struct {
		filePath string
	}

	// Get directory

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Testing Create File",
			args:    args{filePath: "./TestDirectory/Testfile.db"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFile(tt.args.filePath); (err != nil) != tt.wantErr {
				t.Errorf("CreateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
