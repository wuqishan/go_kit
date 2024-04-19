package file

import (
	"github.com/wuqishan/go_kit/base/conv"

	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const (
	// Separator for file system.
	Separator = string(filepath.Separator)

	// DefaultPermOpen is the default perm for file opening.
	DefaultPermOpen = os.FileMode(0666)

	// DefaultPermCopy is the default perm for file/folder copy.
	DefaultPermCopy = os.FileMode(0755)
)

var (
	selfPath = ""
)

func init() {
	// Initialize internal package variable: selfPath.
	selfPath, _ = exec.LookPath(os.Args[0])
	if selfPath != "" {
		selfPath, _ = filepath.Abs(selfPath)
	}
	if selfPath == "" {
		selfPath, _ = filepath.Abs(os.Args[0])
	}
}

// Mkdir creates directories recursively with given `path`.
func Mkdir(path string) (err error) {
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return err
	}
	return nil
}

// Create creates a file with given `path` recursively.
func Create(path string) (*os.File, error) {
	dir := Dir(path)
	if !Exists(dir) {
		if err := Mkdir(dir); err != nil {
			return nil, err
		}
	}
	file, err := os.Create(path)
	return file, err
}

// Open opens file/directory READONLY.
func Open(path string) (*os.File, error) {
	file, err := os.Open(path)
	return file, err
}

// OpenFile opens file/directory with custom `flag` and `perm`.
func OpenFile(path string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := os.OpenFile(path, flag, perm)
	return file, err
}

// OpenWithFlag opens file/directory with default perm and custom `flag`.
func OpenWithFlag(path string, flag int) (*os.File, error) {
	file, err := OpenFile(path, flag, DefaultPermOpen)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// OpenWithFlagPerm opens file/directory with custom `flag` and `perm`.
func OpenWithFlagPerm(path string, flag int, perm os.FileMode) (*os.File, error) {
	file, err := OpenFile(path, flag, perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

// Join joins string array paths with file separator of current system.
func Join(paths ...string) string {
	var s string
	for _, path := range paths {
		if s != "" {
			s += Separator
		}
		s += strings.TrimRight(path, Separator)
	}
	return s
}

// Exists checks whether given `path` exist.
func Exists(path string) bool {
	if stat, err := os.Stat(path); stat != nil && !os.IsNotExist(err) {
		return true
	}
	return false
}

// IsDir checks whether given `path` a directory.
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// Pwd returns absolute path of current working directory.
func Pwd() string {
	path, err := os.Getwd()
	if err != nil {
		return ""
	}
	return path
}

// Chdir changes the current working directory to the named directory.
func Chdir(dir string) (err error) {
	err = os.Chdir(dir)
	return
}

// IsFile checks whether given `path` a file, which means it's not a directory.
func IsFile(path string) bool {
	s, err := Stat(path)
	if err != nil {
		return false
	}
	return !s.IsDir()
}

// Stat returns a FileInfo describing the named file.
func Stat(path string) (os.FileInfo, error) {
	info, err := os.Stat(path)
	return info, err
}

// Move renames (moves) `src` to `dst` path.
func Move(src string, dst string) (err error) {
	err = os.Rename(src, dst)
	return
}

// Rename is alias of Move.
func Rename(src string, dst string) error {
	return Move(src, dst)
}

// DirNames returns sub-file names of given directory `path`.
func DirNames(path string) ([]string, error) {
	f, err := Open(path)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdirnames(-1)
	_ = f.Close()
	if err != nil {
		return nil, err
	}
	return list, nil
}

func Glob(pattern string, onlyNames ...bool) ([]string, error) {
	list, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}
	if len(onlyNames) > 0 && onlyNames[0] && len(list) > 0 {
		array := make([]string, len(list))
		for k, v := range list {
			array[k] = Basename(v)
		}
		return array, nil
	}
	return list, nil
}

// Remove deletes all file/directory with `path` parameter.
func Remove(path string) (err error) {
	// It does nothing if `path` is empty.
	if path == "" {
		return nil
	}
	if err = os.RemoveAll(path); err != nil {
		return err
	}
	return
}

// IsReadable checks whether given `path` is readable.
func IsReadable(path string) bool {
	result := true
	file, err := os.OpenFile(path, os.O_RDONLY, DefaultPermOpen)
	if err != nil {
		result = false
	}
	file.Close()
	return result
}

// IsWritable checks whether given `path` is writable.
func IsWritable(path string) bool {
	result := true
	if IsDir(path) {
		// If it's a directory, create a temporary file to test whether it's writable.
		tmpFile := strings.TrimRight(path, Separator) + Separator + conv.String(time.Now().UnixNano())
		if f, err := Create(tmpFile); err != nil || !Exists(tmpFile) {
			result = false
		} else {
			_ = f.Close()
			_ = Remove(tmpFile)
		}
	} else {
		// If it's a file, check if it can open it.
		file, err := os.OpenFile(path, os.O_WRONLY, DefaultPermOpen)
		if err != nil {
			result = false
		}
		_ = file.Close()
	}
	return result
}

// Chmod is alias of os.Chmod.
func Chmod(path string, mode os.FileMode) (err error) {
	err = os.Chmod(path, mode)
	return
}

// Abs returns an absolute representation of path.
func Abs(path string) string {
	p, _ := filepath.Abs(path)
	return p
}

// RealPath converts the given `path` to its absolute path
func RealPath(path string) string {
	p, err := filepath.Abs(path)
	if err != nil {
		return ""
	}
	if !Exists(p) {
		return ""
	}
	return p
}

// SelfPath returns absolute file path of current running process(binary).
func SelfPath() string {
	return selfPath
}

// SelfName returns file name of current running process(binary).
func SelfName() string {
	return Basename(SelfPath())
}

// SelfDir returns absolute directory path of current running process(binary).
func SelfDir() string {
	return filepath.Dir(SelfPath())
}

// Basename returns the last element of path, which contains file extension.
func Basename(path string) string {
	return filepath.Base(path)
}

// Name returns the last element of path without file extension.
func Name(path string) string {
	base := filepath.Base(path)
	if i := strings.LastIndexByte(base, '.'); i != -1 {
		return base[:i]
	}
	return base
}

// Dir returns all but the last element of path, typically the path's directory.
func Dir(path string) string {
	if path == "." {
		return filepath.Dir(RealPath(path))
	}
	return filepath.Dir(path)
}

// IsEmpty checks whether the given `path` is empty.
func IsEmpty(path string) bool {
	stat, err := Stat(path)
	if err != nil {
		return true
	}
	if stat.IsDir() {
		file, err := os.Open(path)
		if err != nil {
			return true
		}
		defer file.Close()
		names, err := file.Readdirnames(-1)
		if err != nil {
			return true
		}
		return len(names) == 0
	} else {
		return stat.Size() == 0
	}
}

// Ext returns the file name extension used by path.
func Ext(path string) string {
	ext := filepath.Ext(path)
	if p := strings.IndexByte(ext, '?'); p != -1 {
		ext = ext[0:p]
	}
	return ext
}

// ExtName is like function Ext, which returns the file name extension used by path,
func ExtName(path string) string {
	return strings.TrimLeft(Ext(path), ".")
}

// Temp retrieves and returns the temporary directory of current system.
func Temp(names ...string) string {
	path := os.TempDir()
	for _, name := range names {
		path = Join(path, name)
	}
	return path
}
