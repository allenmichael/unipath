package unipath

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const UnixFPSep = "/"
const WinFPSep = "\\"
const FpReplace = "/"

func scrubFilePath(s string) string {
	sClean := filepath.Clean(s)
	sUnix := replaceUnixSep(sClean)
	return replaceWinSep(sUnix)
}
func replaceUnixSep(s string) string {
	matches := strings.FieldsFunc(s, func(c rune) bool {
		return c == []rune(UnixFPSep)[0]
	})
	return strings.Join(matches, FpReplace)
}
func replaceWinSep(s string) string {
	matches := strings.FieldsFunc(s, func(c rune) bool {
		return c == []rune(WinFPSep)[0]
	})
	return strings.Join(matches, FpReplace)
}
func getHomeDir() string {
	if usr, err := user.Current(); err != nil {
		return ""
	} else {
		return usr.HomeDir
	}
}
func getUsername() string {
	if usr, err := user.Current(); err != nil {
		return ""
	} else {
		username := usr.Username
		res := strings.Split(usr.Username, "\\")
		if len(res) > 1 {
			username = res[len(res)-1]
		}
		return username
	}
}
func detectHomeDirChar(s string) bool {
	if len(s) > 0 && s[0] == '~' {
		return true
	}
	return false
}
func detectFPEndingCharacter(s string) bool {
	lastChar := string(s[len(s)-1])
	if lastChar == WinFPSep || lastChar == UnixFPSep {
		return true
	} else {
		return false
	}
}
func detectDoubleDot(s string) bool {
	if len(s) >= 2 && s[0] == '.' && s[1] == '.' {
		return true
	} else {
		return false
	}
}
func getCurrentDir() string {
	ex, _ := os.Executable()
	exClean := filepath.Clean(ex)
	return filepath.Dir(exClean)
}
