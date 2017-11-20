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

var _dataMrlabelHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\x56\x4b\x73\xdb\x36\x10\xbe\xfb\x57\xec\x20\xce\x84\x8a\x62\x50\xa4\xed\x26\x95\x48\x1f\x3b\xd3\x43\x0f\xbd\xb4\x07\xd7\xe3\x81\xc8\x15\x09\x1b\x04\x18\x00\xd4\x63\x12\xfd\xf7\x0e\x40\xf1\xa1\x87\x3b\xc9\xf4\x22\x02\x8b\x7d\x7e\xfb\xd0\x26\xa5\xad\xc4\xc3\x55\x52\x22\xcb\x1f\xae\x00\x12\x93\x69\x5e\x5b\x30\x3a\x4b\x09\xaf\x58\x81\x86\xda\xad\x25\x0f\x49\xd8\xbe\x9c\x31\x95\xd6\xd6\x66\x1e\x86\xec\x85\x6d\x69\xa1\x54\x21\x90\xd5\xdc\xd0\x4c\x55\x9e\x16\x0a\xbe\x34\xe1\xcb\xd7\x06\xf5\x2e\x8c\x69\x44\x6f\x0f\x17\x5a\x71\x49\x5f\xcc\x89\x6a\xc1\xe5\x2b\x68\x14\x29\x31\x76\x27\xd0\x94\x88\x96\x40\xa9\x71\xf5\x73\xa6\x1a\x1e\x46\x34\x8a\xe8\x5d\x68\x4b\xac\xd0\x84\xa6\x52\xca\x96\x12\x4d\xc7\x71\xd3\x70\x9a\x19\x43\xfe\x57\x44\x83\x99\x41\xe9\xc5\xb8\x7c\x34\xee\x44\x6b\x66\xb3\xf2\x2f\x8e\x1b\xf8\x76\x05\x00\x90\x73\x53\x0b\xb6\x9b\xc3\x52\xa8\xec\x75\xe1\x69\xb5\x32\xdc\x72\x25\xe7\xc0\x96\x46\x89\xc6\x62\x4b\xb7\xaa\x9e\xc3\xac\x3d\x2f\x95\xb5\xaa\xea\xaf\x02\x57\xb6\xbf\x68\x5e\x94\xc3\x6d\xc3\x73\x5b\xce\x21\x9a\xcd\xde\xb7\x84\x12\xdb\xf7\x8e\xb2\x77\x1e\x86\xbd\x8b\x49\xe7\xf7\x15\xc0\xaa\x91\x99\xf3\x04\x6a\xd4\x55\x63\xf1\xcf\x06\x8d\xbb\x9b\x60\x72\xf0\x7f\xa5\x34\x04\x6b\xa6\x81\x43\x0a\xb3\x05\x70\x48\xe0\x6b\xc7\x45\x05\xca\xc2\x96\x70\x03\xd1\x02\xf8\x74\xda\x09\x01\x38\x09\x81\x12\xd2\x4b\xcc\xdc\x0b\x8c\x38\x79\xbe\x85\x14\x38\x4c\x21\x82\x29\xfc\xc1\x6c\x49\x57\x42\x29\x1d\xf8\xa3\x66\x32\x57\x55\x30\x81\x8f\x4e\xe5\x04\xde\xbb\xcf\x58\xdc\x56\xf5\xd8\xd0\x23\xcf\xb7\x4f\xdd\xfb\x31\xf5\x98\xed\x12\x93\x63\xb1\x55\xdd\xbe\xec\x0f\xe8\x75\xef\xbf\x7b\x3f\x6f\xbc\xef\x3d\x74\x12\xb7\xb6\xc3\xad\x87\x6d\x24\x31\x9d\xb6\xca\xf8\x0a\x82\xb1\xa2\x87\x73\x6c\x06\xfc\xae\x03\xf2\x2e\x6b\xf4\xb3\x6f\x52\x32\xa1\xcc\x5a\x1d\x10\xa3\x33\xf2\x09\xfa\x12\xce\x62\x6a\x2c\xb3\x3c\x5b\x09\x9e\xbd\x6a\x5f\xc2\x71\x18\xcd\x3e\xdf\x86\xf7\xb7\x77\x71\x14\x45\x51\xfc\x7c\x8f\x9f\x97\xcb\x2c\xbe\xbd\xbd\x7d\x5e\xd2\x97\xba\x20\x93\x2e\x6a\x8d\xb6\xd1\x72\x88\xf4\xbf\xad\x0e\x18\x8d\x82\x78\x9a\x5c\x8c\x6d\x1a\x5d\xa8\x92\xe3\xe0\x6a\x8d\x42\xb1\xbc\x35\x15\xfd\x88\xad\x69\xd4\x59\xdb\x5f\xb4\x19\xff\x9c\xcd\xf8\x87\x6c\xc6\x63\x9b\x7b\xd7\x32\xd7\x41\xae\xb2\xa6\x42\x69\x27\x54\x23\xcb\x77\x41\x57\x09\x67\xc9\x37\x90\xc2\xe3\xa1\xc8\x7c\x43\x70\x89\x8e\xd6\x4e\xde\x67\x63\x35\x35\xb5\xe0\x36\x20\xff\xc8\x2e\x2d\x97\xfa\xcd\xcb\x1d\x22\xba\xd4\x67\x5c\x22\xa4\x70\x4d\xad\xe6\x55\xe0\x99\x1f\xf9\x53\x9f\x66\x87\x53\xcb\x92\x02\x21\xf0\xfd\xbb\x17\x78\x9c\x3d\x39\xc2\x87\xc5\x87\x53\xca\xbb\x0f\x83\x7a\x80\x4c\x49\xcb\x65\x83\x9d\xb6\xfd\x69\xcb\xd0\xba\x31\xa5\x37\x70\x94\x9d\x4c\x49\xa3\x04\x52\xa1\x8a\x80\x38\xd4\x31\x07\x02\xd3\xf3\x71\x30\x05\x32\x22\x76\x30\x9c\xcf\xa3\x96\x7e\xdc\x6c\x8b\xae\x6a\x97\x2a\xdf\x91\x09\x7d\xc5\x5d\x53\x0f\xf9\xc0\xb5\xcb\x52\x1f\x8c\x03\xc2\x93\xe8\xa6\xe4\x59\xe9\x82\x8d\x3f\x43\xf8\x11\xd0\x64\xac\x46\xf8\x18\x3a\x28\x4e\x38\xbe\x38\x86\x25\xcb\x5e\x4d\xcd\x32\xc7\x73\x02\xce\x10\xe4\x6f\x4a\x17\x68\x2d\x97\x05\x08\x66\x2c\x30\x69\x36\xa8\x87\x76\x03\xb8\xa6\xb5\x32\x36\x20\x61\x23\x73\x35\x7e\x08\x43\xf8\x1b\x81\x09\xa3\xa0\x2d\x51\xb0\x25\x42\xcd\x0a\x04\xab\xa0\x62\xaf\x08\xa6\xd1\xe8\xa9\x1d\x56\xb0\xe1\xb6\xf4\x14\x8d\x95\x5a\x63\x7e\x30\x08\xdc\x00\x97\x1d\x6b\x83\xc0\x0a\xc6\x25\xed\x6d\x6d\xb8\xcc\xd5\x86\x0a\x95\x31\xa7\x86\xb6\x06\x03\xab\x1b\x1c\x79\x34\x9e\x0d\x43\xd2\x4f\x11\x4c\xe0\xce\xe1\x33\x3b\x47\xee\x21\x85\x7b\x87\xed\xaf\xc7\x80\xbd\xad\x76\x3c\x19\x13\x98\x39\x7d\x3f\x3c\x2c\xdf\xd2\xeb\x9a\x83\xf9\x2e\x1c\xbb\x76\x03\x77\x5f\x3a\xce\x71\x02\x2f\x0f\x38\x5f\x9f\xae\x70\x99\x34\x3d\x3e\x7d\x22\x85\x2a\xc8\x27\xf8\x06\xed\x32\x45\xe6\x6f\x4c\xc9\x4f\x40\x04\x5b\xa2\x20\x73\xef\xcf\x7e\x72\xd0\x73\xa9\x9a\xf7\xfe\xdb\xfe\x0e\x3b\x46\x12\xb6\x3b\x5c\xe2\x2a\xdd\xff\x8b\x2f\xb5\xff\xe4\x7c\x0d\xfe\xaf\x3d\x25\xfd\x5a\xa1\x51\x30\xcb\xd7\xb8\x68\x57\x83\xf8\xfe\x97\x7a\xbb\x38\x6c\x05\xed\xc5\xaf\x45\x00\x09\xaf\x0a\xe0\x79\x4a\x86\xa1\xdf\x29\x7b\x4b\x94\x40\x26\x98\x31\x29\xe9\x17\x1d\x12\x7a\x47\xc2\x9c\xaf\xfd\x4e\xd1\x2b\x3d\x19\xf1\x9d\xe6\x6e\x23\x92\x4a\x22\x81\x56\xf8\xa2\x48\xfc\xb6\x48\x12\xb6\x38\x24\x61\xbb\xe1\xfe\x1b\x00\x00\xff\xff\xa0\x8c\xe3\xb3\xe9\x0a\x00\x00")

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

	info := bindataFileInfo{name: "data/mrlabel.html", size: 2793, mode: os.FileMode(436), modTime: time.Unix(1511136094, 0)}
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

