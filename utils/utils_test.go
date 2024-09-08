package utils

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestReadCSV(t *testing.T) {
	tempFile := filepath.Join(os.TempDir(), "test.csv")
	testData := `Id,Date,Transaction
0,7/15,+60.5
1,7/28,-10.3
2,8/2,-20.46
3,8/13,+10.0`
	if err := os.WriteFile(tempFile, []byte(testData), 0644); err != nil {
		t.Fatalf("could not write temp file: %v", err)
	}
	defer os.Remove(tempFile)

	type args struct {
		filePath string
		columns  []string
	}
	tests := []struct {
		name    string
		args    args
		want    []map[string]string
		wantErr bool
	}{
		{
			name: "Valid file",
			args: args{
				filePath: tempFile,
				columns:  []string{"Id", "Date", "Transaction"},
			},
			want: []map[string]string{
				{"Id": "0", "Date": "7/15", "Transaction": "+60.5"},
				{"Id": "1", "Date": "7/28", "Transaction": "-10.3"},
				{"Id": "2", "Date": "8/2", "Transaction": "-20.46"},
				{"Id": "3", "Date": "8/13", "Transaction": "+10.0"},
			},
			wantErr: false,
		},
		{
			name: "Invalid columns",
			args: args{
				filePath: tempFile,
				columns:  []string{"Nonexistent"},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Invalid file",
			args: args{
				filePath: "invalid_path.csv",
				columns:  []string{"Id", "Date", "Transaction"},
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadCSV(tt.args.filePath, tt.args.columns)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadCSV() = %v, want %v", got, tt.want)
			}
		})
	}
}
func TestGetFileNames(t *testing.T) {
	tempDir := filepath.Join(os.TempDir(), "test_dir")
	if err := os.Mkdir(tempDir, 0755); err != nil {
		t.Fatalf("could not create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up

	testFiles := []string{"file1.txt", "file2.csv", "file3.log"}
	for _, fileName := range testFiles {
		if err := os.WriteFile(filepath.Join(tempDir, fileName), []byte("test"), 0644); err != nil {
			t.Fatalf("could not write temp file: %v", err)
		}
	}

	type args struct {
		folderPath string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Valid directory",
			args: args{
				folderPath: tempDir,
			},
			want: func() []string {
				var files []string
				for _, file := range testFiles {
					files = append(files, filepath.Join(tempDir, file))
				}
				return files
			}(),
			wantErr: false,
		},
		{
			name: "Invalid directory",
			args: args{
				folderPath: "invalid_dir",
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetFileNames(tt.args.folderPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetFileNames() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got = normalizePaths(got)
			want := normalizePaths(tt.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("GetFileNames() = %v, want %v", got, want)
			}
		})
	}
}

// normalizePaths ensures that file paths use the correct directory separator.
func normalizePaths(paths []string) []string {
	for i, path := range paths {
		paths[i] = filepath.Clean(path)
	}
	return paths
}
