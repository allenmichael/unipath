package unipath

import (
	"path/filepath"
	"strings"
)

type PathResolver interface {
	ConvertToAbsPath(s string) (string, error)
	ConvertFileSeparator(s string) string
}

type GenericPath struct {
}

// ConvertToPath - finds and replaces file seperators of / and \ values.
// Returns a new string with the transformed file path.
// func ConvertToPath(s string) (string, error) {
// 	if len(s) > 0 {
// 		isPathAbs := checkForWinAbsPath(s)
// 		isWinPath := isLikelyWindowsPath(s)
// 		var cleanPath string

// 		if isPathAbs == true && isWinPath == true {
// 			var prefix string
// 			var err error
// 			if checkForDriveLetter(s) {
// 				if prefix, err = retrieveDriveLetter(s); err != nil {
// 					return "", err
// 				}
// 				return processWindowsPrefixedPath(s, prefix, false)
// 			}
// 			if isUNCPath(s) {
// 				prefix = "\\\\"
// 				return processWindowsPrefixedPath(s, prefix, true)
// 			}
// 		}

// 		homeDir := ""
// 		if detectHomeDirChar(s) {
// 			homeDir := getHomeDir()
// 			if homeDir == "" {
// 				if isWinPath {
// 					homeDir = winDefaultDrive
// 				} else {
// 					homeDir = unixFPSep
// 				}
// 			}
// 			s = s[1:]
// 		}

// 		cleanPath = scrubFilePath(s)
// 		finalPath := strings.Split(cleanPath, fpReplace)

// 		if isPathAbs == true && isWinPath == false {
// 			return joinFilePath(append([]string{unixFPSep}, finalPath...)), nil
// 		}
// 		return joinFilePath(append([]string{homeDir}, finalPath...)), nil
// 	}

// 	return "", errors.New("No file path found")
// }

func joinFilePath(s []string) string {
	return filepath.Join(s...)
}
func joinFilePathWithSeperator(s []string, pathSep string) string {
	return strings.Join(s, pathSep)
}

// func checkForAbsPath(s string) bool {
// 	if len(s) > 0 && s[0] == '/' || checkForDriveLetter(s) {
// 		return true
// 	}
// 	if isUNCPath(s) {
// 		return true
// 	}
// 	return false
// }

// func isLikelyWindowsPath(s string) bool {
// 	isWindowsPath := checkForDriveLetter(s)
// 	if isWindowsPath == true {
// 		return true
// 	}
// 	if isUNCPath(s) {
// 		return true
// 	}
// 	winSeps := strings.Split(s, winFPSep)
// 	unixSeps := strings.Split(s, unixFPSep)
// 	return len(unixSeps) < len(winSeps)
// }

// func resolveDoubleDot(s string) string {
// 	if len(s) >= 2 && s[0] == '.' && s[1] == '.' {
// 		home := getHomeDir()
// 		if home == "" {
// 			return ""
// 		}
// 		return home + s[1:]
// 	}
// 	return s
// }
