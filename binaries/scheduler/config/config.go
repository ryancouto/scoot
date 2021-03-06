// Code generated by go-bindata.
// sources:
// config/config.go
// config/local.local
// config/local.memory
// DO NOT EDIT!

package config

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configConfigGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func configConfigGoBytes() ([]byte, error) {
	return bindataRead(
		_configConfigGo,
		"config/config.go",
	)
}

func configConfigGo() (*asset, error) {
	bytes, err := configConfigGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/config.go", size: 0, mode: os.FileMode(420), modTime: time.Unix(1501868697, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configLocalLocal = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\x41\x4b\x3b\x31\x10\x47\xef\xfb\x29\x42\xce\xe5\xff\x2f\x15\x2f\x7b\x6d\x4f\x62\x71\x69\x17\x3c\x4f\xb3\x93\x18\x3a\xbb\xb3\xcc\x4c\xc4\x22\xfd\xee\x12\x6d\x55\xa8\xc7\xcc\x7b\x79\xfc\xde\x1b\xe7\xfc\x9a\x8a\x1a\x8a\x6f\x5d\x7d\x3a\xe7\xfb\xd3\x8c\xbe\x75\x9e\x38\x00\xf9\xc6\xb9\xf3\xa2\x7a\xcf\x2c\x47\x14\xbd\xf5\x64\x0e\x7e\xf1\x75\xea\x98\x28\x4f\xa9\x43\xc9\x3c\x54\xb6\xba\x5f\x8e\xfa\xd3\xd8\x87\x17\x1c\x0a\xa1\xac\x79\x8a\x39\xdd\xb6\xd4\xc0\x30\x16\xba\x06\xb7\xf0\xb6\x43\x93\x8c\xda\xa1\xf4\xa0\x47\xef\x5a\xb7\xba\xc0\x0d\x1e\x4a\xda\xf2\x80\xf5\x18\x81\x14\x2f\x60\x87\x81\x5f\x51\x1e\xf8\xa0\x4f\xd3\xde\x40\xac\xcc\xd5\x31\x29\xf8\xfd\x37\x42\x21\xab\xc9\x3e\x8f\xc8\xc5\xaa\xe0\xef\x96\xe3\xaf\xb5\x90\xe0\x91\xff\x58\x19\x33\xe1\x75\xe1\x26\x0b\x06\x63\x39\x55\xf0\x4f\x03\xb3\x0d\x60\xf0\xbf\x3a\x0a\x09\x88\xd3\x67\xb1\x39\x37\x1f\x01\x00\x00\xff\xff\x50\xb0\x50\xfa\x6f\x01\x00\x00")

func configLocalLocalBytes() ([]byte, error) {
	return bindataRead(
		_configLocalLocal,
		"config/local.local",
	)
}

func configLocalLocal() (*asset, error) {
	bytes, err := configLocalLocalBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/local.local", size: 367, mode: os.FileMode(420), modTime: time.Unix(1497560569, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configLocalMemory = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8f\xc1\x4a\x03\x41\x10\x44\xef\xfb\x15\x4d\x9f\x3d\x44\xbd\xcd\x75\x73\x12\x83\x92\xec\x0f\x74\x92\x9a\x75\x71\x66\x3b\xcc\x74\x8b\x41\xf6\xdf\x65\x70\x45\x50\x8f\x5d\xf5\x78\x54\x7f\x74\x44\xdc\x27\xaf\x86\xc2\x81\xda\x49\xc4\xc3\xf5\x02\x0e\xc4\x19\x59\xcb\x95\x6f\xbe\xd2\x5e\x7d\x36\x0e\x74\xbb\xe9\x88\x96\x16\xf2\xe1\xf4\x82\xb3\x27\x94\x5e\xe7\x38\x8d\x7f\x0d\xd5\xc4\x10\x3d\x7d\x3b\x76\xf2\xbe\x87\x95\x09\xf5\x19\x65\x90\xfa\xca\x14\xe8\x6e\x2d\xb7\x38\xfa\xb8\xd3\x33\x5a\x18\x25\x55\xac\xc5\x1e\x27\x7d\x43\x79\xd0\x63\x7d\x9a\x0f\x26\xc5\xfc\xf2\x9b\xd9\x22\x8a\x27\x6b\xce\x61\xca\x50\xb7\x46\xf0\xfd\x26\xf3\xcf\x5c\x19\xe5\x51\xff\x99\xb9\x3e\xda\xc0\x6e\xe9\x3e\x03\x00\x00\xff\xff\xdf\xe0\x48\x63\x15\x01\x00\x00")

func configLocalMemoryBytes() ([]byte, error) {
	return bindataRead(
		_configLocalMemory,
		"config/local.memory",
	)
}

func configLocalMemory() (*asset, error) {
	bytes, err := configLocalMemoryBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/local.memory", size: 277, mode: os.FileMode(420), modTime: time.Unix(1497560569, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"config/config.go":    configConfigGo,
	"config/local.local":  configLocalLocal,
	"config/local.memory": configLocalMemory,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"config": &bintree{nil, map[string]*bintree{
		"config.go":    &bintree{configConfigGo, map[string]*bintree{}},
		"local.local":  &bintree{configLocalLocal, map[string]*bintree{}},
		"local.memory": &bintree{configLocalMemory, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
