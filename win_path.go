package unipath

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"text/template"
)

const winHomePath = `{{.DriveLetter}}\Users\{{.User}}`
const winDefaultDrive = `C:`

type WinPath struct {
	driveLetter        string
	isUNCPath          bool
	homePath           string
	defaultDriveLetter string
}

func NewWinPath(options ...func(*WinPath) error) (*WinPath, error) {
	w := &WinPath{}
	w.isUNCPath = false

	var optErr error
	for _, options := range options {
		optErr = options(w)
	}
	if optErr != nil {
		return nil, optErr
	}

	if w.homePath != "" {
		return w, nil
	}

	homeDir := getHomeDir()
	if !detectDriveLetter(homeDir) {
		username := getUsername()
		tmpl := template.New("WinHomePath")
		if tmpl, err := tmpl.Parse(winHomePath); err != nil {
			return nil, err
		} else {
			var cmpTmpl bytes.Buffer
			if err := tmpl.Execute(&cmpTmpl, map[string]string{"User": username, "DriveLetter": winDefaultDrive}); err != nil {
				return nil, err
			}
			homeDir = cmpTmpl.String()
		}
	}
	w.homePath = homeDir
	return w, nil
}

func DriveLetter(s string) func(*WinPath) error {
	return func(w *WinPath) error {
		return w.SetDriveLetter(s)
	}
}
func (w *WinPath) SetDriveLetter(s string) error {
	if detectFPEndingCharacter(s) {
		s = s[:len(s)-1]
	}
	if detectDriveLetter(s) {
		drv, _ := retrieveDriveLetter(s)
		w.driveLetter = drv
	}
	return errors.New("No drive letter detected")
}
func (w *WinPath) GetDriveLetter() string {
	return w.driveLetter
}
func HomePath(s string) func(*WinPath) error {
	return func(w *WinPath) error {
		return w.SetHomePath(s)
	}
}
func (w *WinPath) GetCurrentDirectory() string {
	return getCurrentDir()
}
func (w *WinPath) SetHomePath(s string) error {
	if !detectWinHomePath(s) {
		return errors.New("Not a proper Windows home path")
	}
	if detectFPEndingCharacter(s) {
		s = s[:len(s)-1]
	}
	if detectDriveLetter(s) {
		drv, _ := retrieveDriveLetter(s)
		w.driveLetter = drv
	}
	if isUNCPath(s) {
		w.isUNCPath = true
	}
	w.homePath = s
	return nil
}
func (w *WinPath) GetHomePath() string {
	return w.homePath
}
func ExtractDriveLetter(s string) (string, error) {
	if detectDriveLetter(s) {
		return retrieveDriveLetter(s)
	}
	return "", errors.New("Couldn't detect a drive letter in this path")
}
func (w *WinPath) ConvertFileSeparator(s string) string {
	scrubbedPath := scrubFilePath(s)
	return joinWindowsPath(strings.Split(scrubbedPath, FpReplace))
}
func (w *WinPath) ConvertToAbsPath(s string) (string, error) {
	isPathAbs := checkForWinAbsPath(s)
	if isPathAbs == true {
		var prefix string
		var err error
		if detectDriveLetter(s) {
			if prefix, err = retrieveDriveLetter(s); err != nil {
				return "", err
			}
			w.driveLetter = prefix + WinFPSep
			return processWindowsPrefixedPath(s, prefix, false)
		} else if isUNCPath(s) {
			w.isUNCPath = true
			prefix = "\\\\"
			return processWindowsPrefixedPath(s, prefix, true)
		} else {
			return s, nil
		}
	} else if detectHomeDirChar(s) {
		s = s[1:]
		scrubbedPath := scrubFilePath(s)
		return w.homePath + WinFPSep + joinWindowsPath(strings.Split(scrubbedPath, FpReplace)), nil
	} else {
		scrubbedPath := scrubFilePath(s)
		return w.homePath + WinFPSep + joinWindowsPath(strings.Split(scrubbedPath, FpReplace)), nil
	}
}
func checkForWinAbsPath(s string) bool {
	if len(s) <= 0 {
		return false
	}
	if detectDriveLetter(s) {
		return true
	}
	if isUNCPath(s) {
		return true
	}
	return false
}
func detectWinHomePath(s string) bool {
	return detectDriveLetter(s) || isUNCPath(s)
}
func detectDriveLetter(s string) bool {
	re := regexp.MustCompile(`^[A-Za-z]*:`)
	return re.MatchString(s)
}
func isUNCPath(s string) bool {
	if len(s) > 2 && s[0] == '\\' && s[1] == '\\' {
		return true
	}
	return false
}
func retrieveDriveLetter(s string) (string, error) {
	if !detectDriveLetter(s) {
		return "", errors.New("Couldn't detect a drive letter in this path")
	}
	spl := strings.Split(s, ":")
	if len(spl) > 0 {
		return fmt.Sprintf("%s:", spl[0]), nil
	}
	return "", errors.New("No drive letter found")
}
func processWindowsPrefixedPath(s string, prefix string, isUNCPath bool) (string, error) {
	removedPrefix := strings.Split(s, prefix)
	if len(removedPrefix) >= 2 {
		scrubbedPath := scrubFilePath(removedPrefix[1])
		if !isUNCPath {
			return prefix + WinFPSep + joinWindowsPath(strings.Split(scrubbedPath, FpReplace)), nil
		} else if isUNCPath {
			return prefix + joinWindowsPath(strings.Split(scrubbedPath, FpReplace)), nil
		}
	}
	return "", errors.New("There was an error parsing the drive letter of this path")
}
func joinWindowsPath(s []string) string {
	return strings.Join(s, WinFPSep)
}
