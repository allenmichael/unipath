package unipath

import (
	"os/user"
	"runtime"
	"strings"
	"testing"
)

func TestWinHomePathOnUnix(t *testing.T) {
	if runtime.GOOS != "windows" {
		if currentUser, err := user.Current(); err == nil {
			winPath, _ := NewWinPath()
			expectedPath := "C:\\Users\\" + currentUser.Username
			if winPath.GetHomePath() != expectedPath {
				t.Error("Expected path to be ", expectedPath,
					"Path: ", winPath.GetHomePath())
			}
		}
	}
}

// func TestSetWinDriveLetterFail(t *testing.T) {
// 	letter := "D"
// 	winPath, _ := NewWinPath(DriveLetter(letter))
// 	// if winPath.GetHomePath() != path {
// 	// 	t.Error("Expected path to be ", path,
// 	// 		"Path: ", winPath.GetHomePath())
// 	// }
// }
func TestSetWinHomePath(t *testing.T) {
	path := "D:\\Users\\me"
	winPath, _ := NewWinPath(HomePath(path))
	if winPath.GetHomePath() != path {
		t.Error("Expected path to be ", path,
			"Path: ", winPath.GetHomePath())
	}
}
func TestSetWinHomePathEndingFileSepGetAbsPath(t *testing.T) {
	homePath := "D:\\Users\\me\\"
	winPath, _ := NewWinPath(HomePath(homePath))

	path := "my_folder/file.txt"
	if cvtPath, err := winPath.ConvertToAbsPath(path); err == nil {
		if cvtPath != homePath+"my_folder\\file.txt" {
			t.Error("Expected path to be ", homePath+"my_folder\\file.txt",
				"Path: ", cvtPath)
		}
	} else {
		t.Error("An error occurred while parsing the path: ", err)
	}
}

func TestHomeCharWithWinAbsPathOnUnix(t *testing.T) {
	if runtime.GOOS != "windows" {
		path := "~/my_path/to/file.jpg"
		if currentUser, err := user.Current(); err == nil {
			winPath, _ := NewWinPath()
			expectedPath := "C:\\Users\\" + currentUser.Username + "\\my_path\\to\\file.jpg"
			if cvtPath, err := winPath.ConvertToAbsPath(path); err == nil {
				if cvtPath != expectedPath {
					t.Error("Expected path to be ", expectedPath,
						"Path: ", cvtPath)
				}
			} else {
				t.Error("An error occurred while parsing the path: ", err)
			}
		}
	}
}

func TestCurrentDirectory(t *testing.T) {
	winPath, _ := NewWinPath()
	if runtime.GOOS != "windows" {
		if !strings.Contains(winPath.GetCurrentDirectory(), "/var/folders") {
			t.Error("GetCurrentDirectory didn't find temp folder path.")
		}
	}
}

// func TestPathWithDriveLetter(t *testing.T) {
// 	s := `C:\my_path\to_files`
// 	if transformed, err := ConvertToPath(s); err != nil {
// 		t.Error(fmt.Sprintf(`
// 		Something went wrong.
// 		Error: %s
// 		`, err))
// 	} else {
// 		if !(transformed == `C:\my_path\to_files`) {
// 			t.Error("Expected path to be ", s,
// 				"Path: ", transformed)
// 		}
// 	}
// }
// func TestPathWithDriveLetterAndUnixSeperators(t *testing.T) {
// 	s := `C:/my_path/to_files`
// 	if transformed, err := ConvertToPath(s); err != nil {
// 		t.Error(fmt.Sprintf(`
// 		Something went wrong.
// 		Error: %s
// 		`, err))
// 	} else {
// 		if !(transformed == `C:\my_path\to_files`) {
// 			t.Error("Expected path to be ", s,
// 				"Path: ", transformed)
// 		}
// 	}
// }
// func TestPathWithDriveLetterAndWinSeperators(t *testing.T) {
// 	s := `C:\my_path\to_files`
// 	if transformed, err := ConvertToPath(s); err != nil {
// 		t.Error(fmt.Sprintf(`
// 		Something went wrong.
// 		Error: %s
// 		`, err))
// 	} else {
// 		if !(transformed == `C:\my_path\to_files`) {
// 			t.Error("Expected path to be ", s,
// 				"Path: ", transformed)
// 		}
// 	}
// }
// func TestPathWithDriveLetterAndMixedSeperators(t *testing.T) {
// 	s := `C:/my_path\to_files`
// 	if transformed, err := ConvertToPath(s); err != nil {
// 		t.Error(fmt.Sprintf(`
// 		Something went wrong.
// 		Error: %s
// 		`, err))
// 	} else {
// 		if !(transformed == `C:\my_path\to_files`) {
// 			t.Error("Expected path to be ", s,
// 				"Path: ", transformed)
// 		}
// 	}
// }
// func TestPathWithDriveLetterAndMissingAfterDriveWithMixedSeperators(t *testing.T) {
// 	s := `C:my_path\to_files`
// 	if transformed, err := ConvertToPath(s); err != nil {
// 		t.Error(fmt.Sprintf(`
// 		Something went wrong.
// 		Error: %s
// 		`, err))
// 	} else {
// 		if !(transformed == `C:\my_path\to_files`) {
// 			t.Error("Expected path to be ", s,
// 				"Path: ", transformed)
// 		}
// 	}
// }
// func TestPathWithHome(t *testing.T) {
// 	s := `~/my_path/`
// 	if transformed, err := ConvertToPath(s); err != nil {
// 		t.Error(fmt.Sprintf(`
// 		Something went wrong.
// 		Error: %s
// 		`, err))
// 	} else {
// 		if usr, err := user.Current(); err != nil {
// 			log.Fatal(err)
// 		} else {
// 			homeDir := usr.HomeDir
// 			if !strings.Contains(transformed, homeDir) {
// 				t.Error("Expected path to include", homeDir,
// 					"Path: ", transformed)
// 			}
// 		}
// 	}
// }
// func TestPathWithHomeWinSeperators(t *testing.T) {
// 	s := `~\my_path`
// 	if transformed, err := ConvertToPath(s); err != nil {
// 		t.Error(fmt.Sprintf(`
// 		Something went wrong.
// 		Error: %s
// 		`, err))
// 	} else {
// 		if usr, err := user.Current(); err != nil {
// 			log.Fatal(err)
// 		} else {
// 			homeDir := usr.HomeDir
// 			targetDir := homeDir + `\my_path`
// 			if !(transformed == targetDir) {
// 				t.Error("Expected path to match", targetDir,
// 					"Path: ", transformed)
// 			}
// 		}
// 	}
// }
