package utils

import (
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestAll(t *testing.T) {
	// TEST CREATE FILE -- THIS will also create a folder

	wg = sync.WaitGroup{}
	wg.Add(1)
	TestCreateFile(t)

	wg.Wait()
	// TEST DELETE THE FOLDER
	TestRemoveFolder(t)
}

func TestCreateFile(t *testing.T) {
	t.Log("Testing Create File")
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
	wg.Done()
}

func TestCreateFolder(t *testing.T) {
	t.Log("Testing Create Folder")
	type args struct {
		folderPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Create Folder test",
			args:    args{folderPath: "./TestDirectory"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFolder(tt.args.folderPath); (err != nil) != tt.wantErr {
				t.Errorf("CreateFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	wg.Done()
}

func TestRemoveFolder(t *testing.T) {
	t.Log("Testing Remove Folder")

	type args struct {
		folderPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "Remove Folder test",
			args:    args{folderPath: "./TestDirectory"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveFolder(tt.args.folderPath); (err != nil) != tt.wantErr {
				t.Errorf("RemoveFolder() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
