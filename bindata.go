// Code generated by go-bindata.
// sources:
// data/mrlabel.html
// DO NOT EDIT!

package main

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

var _dataMrlabelHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x56\x4d\x73\xdb\x36\x10\xbd\xeb\x57\xec\x20\xce\x84\x34\x6d\x50\xa4\x9c\x49\x6a\x91\x3a\xf4\xd6\x43\xdb\xe9\xd9\xf5\x78\x20\x72\x25\xc2\x26\x01\x06\x00\x1d\x69\x12\xfd\xf7\x0e\x40\xf1\x43\x32\x9d\x71\x26\x3d\xd8\x22\x17\xfb\x81\x7d\x6f\x1f\x88\xa4\x30\x55\xb9\x9a\x25\x05\xb2\x7c\x35\x03\x48\x74\xa6\x78\x6d\x40\xab\x2c\x25\xbc\x62\x5b\xd4\xd4\xec\x0c\x59\x25\x61\xbb\xf2\xc2\xa9\x30\xa6\xd6\xb7\x61\xc8\x1e\xd9\x8e\x6e\xa5\xdc\x96\xc8\x6a\xae\x69\x26\x2b\x67\x0b\x4b\xbe\xd6\xe1\xe3\x97\x06\xd5\x3e\x8c\x69\x44\x17\xc7\x17\x5a\x71\x41\x1f\xf5\x59\xea\x92\x8b\x27\x50\x58\xa6\x44\x9b\x7d\x89\xba\x40\x34\x04\x0a\x85\x9b\x9f\x2b\xd5\xf0\x30\xa2\x51\x44\x6f\x42\x53\x60\x85\x3a\xd4\x95\x94\xa6\x10\xa8\x3b\x8f\xeb\x86\xd3\x4c\x6b\xf2\x4b\x1d\x0d\x65\x86\xa4\x93\x7d\x75\x8f\x33\x80\x4d\x23\x32\xc3\xa5\x80\x1a\x55\xd5\x18\xfc\xa7\x41\x6d\xdf\xb5\xe7\xc3\xb7\x19\x00\xc0\x46\x2a\xf0\x9e\x99\x02\x0e\x29\xcc\x97\xc0\x21\x81\x2f\x9d\x17\x2d\x51\x6c\x4d\x01\xd7\x10\x2d\x81\x07\x41\x17\x04\x60\x23\x4a\x14\x90\x4e\x39\x73\x17\x30\xf2\xe4\xf9\x0e\x52\xe0\x10\x40\x04\x01\xfc\xc9\x4c\x41\x37\xa5\x94\xca\x73\x8f\x8a\x89\x5c\x56\x9e\x0f\x97\x36\xa5\x0f\xef\xed\xcf\x38\xdc\x54\xf5\xb8\xd0\x1d\xcf\x77\xf7\xdd\xfa\xa9\xf5\xd4\x6d\xca\xc9\xba\x98\xaa\x6e\x57\x0e\xb3\xf6\xaf\x5b\xff\xc3\xed\xf3\xda\xed\xbd\x87\x4e\xe0\xce\x74\xb8\xf5\xb0\x8d\x22\x82\xa0\x4d\xc6\x37\xe0\x8d\x13\xad\x5e\x62\x33\xe0\x77\xe1\x91\x77\x59\xa3\x1e\xdc\xdc\x13\x9f\x32\x63\x94\x47\xb4\xca\xc8\x15\xf4\x53\x91\xc5\x54\x1b\x66\x78\xb6\x29\x79\xf6\xa4\xdc\x54\xc4\x61\x34\xff\xb4\x08\x3f\x2e\x6e\xe2\x28\x8a\xa2\xf8\xe1\x23\x7e\x5a\xaf\xb3\x78\xb1\x58\x3c\xac\xe9\x63\xbd\x25\x7e\xd7\xb5\x42\xd3\x28\x31\x74\xfa\xe3\xaa\x03\x46\xa3\x26\xee\xfd\xc9\xde\x82\x68\x62\x4a\x4e\x9b\xab\x15\x96\x92\xe5\x6d\xa9\xe8\x2d\xb5\x82\xa8\xab\x76\x98\xac\x19\xff\x5c\xcd\xf8\x4d\x35\xe3\x71\xcd\x83\x95\xcc\x85\x97\xcb\xac\xa9\x50\x18\x9f\x2a\x64\xf9\xde\xeb\x26\xe1\x05\xf9\x1a\x52\xb8\x3b\x0e\x99\x13\x04\x17\x68\x6d\xed\x61\xf6\xa0\x8d\xa2\xba\x2e\xb9\xf1\xc8\xbf\xa2\xa3\x65\x4a\x6f\x2e\xee\xd8\xd1\x94\xce\xb8\x40\x48\xe1\x82\x1a\xc5\x2b\xcf\x39\xdf\xf1\xfb\x9e\x66\x8b\x53\xeb\x92\x02\x21\xf0\xfd\xbb\x0b\xb8\x9b\xdf\x5b\xc3\x87\xe5\x87\x73\xcb\xbb\x0f\x43\x7a\x80\x4c\x0a\xc3\x45\x83\x5d\xb6\xc3\xb9\x64\x68\xdd\xe8\xc2\x15\x38\x61\x27\x93\x42\xcb\x12\x69\x29\xb7\x1e\xb1\xa8\x63\x0e\x04\x82\x97\xc7\x41\x00\x64\x64\xec\x60\x10\x4d\xf5\xb0\x6e\x8c\x69\x41\x9c\x2f\xbb\xf1\x6c\x6d\xc4\xa7\xc8\xb2\x62\x40\x9e\x8b\x1c\x77\xc3\xb6\xc7\xd1\x41\x3a\x9c\x35\x17\x9e\x29\xb8\xf6\x69\xbb\xe6\xf9\x34\xb3\xc2\x19\xd2\xe0\xf3\xb8\x75\xcb\x80\x4b\xbc\x1c\xa3\xd1\x77\xc5\xfd\xc1\x7e\xae\xed\x04\xe6\x16\xd6\x37\xcb\xfd\x5c\x8e\x63\xa4\x01\x2e\x68\x2d\xb5\xf1\x48\x58\xca\x2d\xb9\x82\x6f\xd0\x7e\x0e\xc9\xed\x2b\xa2\xbc\x02\x52\xb2\x35\x96\xe4\x16\x38\x1c\xfc\x3e\xcf\xe9\x49\xd5\x33\xda\xf1\x36\x82\x3e\x2b\x99\xd6\x6e\x56\xc7\x50\xf6\xa7\xf6\xcb\xaf\xc5\x31\x74\xa2\x80\x25\x4d\xe6\x7b\xe2\xd3\x27\xdc\x37\xf5\x18\x6c\xab\xa1\x1e\x02\x0b\xa1\x33\xd1\xaf\x05\xcf\x0a\x48\xe0\xe6\x33\x84\x97\xf3\xcb\xd0\x22\x39\x5e\x59\xa5\x70\xf3\x1b\x04\xe3\x8d\x8e\x91\x3c\xc5\xf1\x30\xfb\xbf\x18\x9a\xce\x6b\x05\xc8\xdc\x90\x8e\xb7\x78\x0d\x37\x9f\x3b\xcf\xf1\xcc\x4c\xf3\xe5\x34\x60\xc5\xc1\x84\xee\x69\xf9\x05\xd2\xed\x7e\x7a\xda\xa7\x38\x69\xa9\x6e\xff\x0f\x57\x83\x24\x6c\xaf\x5e\x89\xe5\xcb\xdd\x14\x72\xfe\x0c\x3c\x4f\x8f\xaa\x6b\xef\x27\x00\x49\xfb\xba\xfa\x4b\x82\xab\x98\x84\x47\xc3\xf9\x6a\xc5\xb8\x00\xdd\xac\x1f\x31\x33\xe0\x71\x91\x95\x4d\xce\xc5\xd6\x8a\x4a\x4a\xa5\x81\x89\x1c\x16\x39\x28\x14\x39\x2a\xed\x4f\xe7\xf9\xbd\xe1\xa5\x8d\xba\x02\xa9\xa0\x66\xca\x80\xdc\x80\x14\x38\xed\xfd\xb7\x29\x50\x81\xd9\xd7\xa8\xad\xdf\xfa\x18\xac\x43\x6d\x54\x93\x99\x46\xa1\x7e\x6d\xbb\x06\x58\xef\x3f\xb8\x8f\xbd\x93\x30\xe7\xcf\xee\x61\xad\xdc\x0f\xaf\xb6\x0e\x9f\xe1\x9b\x09\xee\xaa\x98\x92\xaf\x3c\x37\xc5\x6d\x3c\x7f\xbf\x84\x02\xf9\xb6\x30\xf6\x19\xc2\x93\xa8\xb3\x4f\x60\x17\x9a\x73\x5d\x97\x6c\x7f\x2b\xa4\x40\xf2\xa3\x90\xf8\xf5\x90\x24\x6c\x39\x4c\xc2\xf6\x52\xfd\x5f\x00\x00\x00\xff\xff\x34\xae\x1b\x44\x5c\x0b\x00\x00")

func dataMrlabelHtmlBytes() ([]byte, error) {
	return bindataRead(
		_dataMrlabelHtml,
		"data/mrlabel.html",
	)
}

func dataMrlabelHtml() (*asset, error) {
	bytes, err := dataMrlabelHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/mrlabel.html", size: 2908, mode: os.FileMode(436), modTime: time.Unix(1509738165, 0)}
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
	"data/mrlabel.html": dataMrlabelHtml,
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
	"data": &bintree{nil, map[string]*bintree{
		"mrlabel.html": &bintree{dataMrlabelHtml, map[string]*bintree{}},
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

