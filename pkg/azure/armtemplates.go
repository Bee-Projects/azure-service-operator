// Code generated by go-bindata.
// sources:
// pkg/azure/templates/.DS_Store
// pkg/azure/templates/acr.json
// pkg/azure/templates/storageaccount.json
// DO NOT EDIT!

package azure

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

var _pkgAzureTemplatesDs_store = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\xd8\x31\x0a\x02\x31\x10\x85\xe1\x37\x31\x45\xc0\x26\xa5\x65\x1a\x0f\xe0\x0d\xc2\xb2\x9e\xc0\x0b\x58\x78\x05\xfb\x1c\x5d\x96\x79\x60\x60\xd5\x4e\x8c\xcb\xfb\x40\xfe\x05\x37\x2a\x16\x31\x23\x00\x9b\xee\xb7\x13\x90\x01\x24\x78\x71\xc4\x4b\x89\x8f\x95\xd0\x5d\x1b\x5f\x43\x44\x44\x44\xc6\x66\x9e\xb4\xff\xf5\x07\x11\x91\xe1\x2c\xfb\x43\x61\x2b\xdb\xbc\xc6\xe7\x03\x1b\xbb\x35\x99\x2d\x6c\x65\x9b\xd7\x78\x5f\x60\x23\x9b\xd8\xcc\x16\xb6\xb2\xcd\xcb\x4d\xcb\x38\x7c\x18\xdf\xd9\x38\xa1\x18\xa7\x10\x2b\x6c\xfd\xce\x77\x23\xf2\xef\x76\x9e\xbc\xfc\xfe\x9f\xdf\xcf\xff\x22\xb2\x61\x16\xe7\xcb\x3c\x3d\x07\x82\xf5\x0d\x00\xae\xdd\xf5\xa7\x43\x40\xf0\x3f\x0b\x0f\xdd\x5a\x1d\x04\x44\x06\xf3\x08\x00\x00\xff\xff\x6a\x00\x88\x6d\x04\x18\x00\x00")

func pkgAzureTemplatesDs_storeBytes() ([]byte, error) {
	return bindataRead(
		_pkgAzureTemplatesDs_store,
		"pkg/azure/templates/.DS_Store",
	)
}

func pkgAzureTemplatesDs_store() (*asset, error) {
	bytes, err := pkgAzureTemplatesDs_storeBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/azure/templates/.DS_Store", size: 6148, mode: os.FileMode(420), modTime: time.Unix(1545103802, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgAzureTemplatesAcrJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x4d\x6b\xdb\x40\x10\xbd\xfb\x57\x0c\xdb\x42\x5a\x88\xf5\x11\x28\x05\xdf\x42\xe9\xa1\xf4\x83\x52\x3b\xb9\x84\x1c\xc6\xbb\x23\x6b\x1b\xed\xae\xd8\x1d\xa5\xb8\xc5\xff\xbd\xac\x3e\x2c\x2b\xb8\x4d\x6d\x8a\x0e\x42\x3b\xf3\x66\xde\x5b\xbd\xf7\x6b\x06\x20\x5e\x06\x59\x92\x41\xb1\x00\x51\x32\xd7\x61\x91\xa6\xdd\x49\x62\xd0\xe2\x86\x0c\x59\x4e\xf0\x67\xe3\x29\x91\xce\xf4\xb5\x90\x5e\x65\xf9\x9b\x79\x96\xcf\xb3\x3c\x55\x54\x57\x6e\x1b\xfb\x56\x64\xea\x0a\x99\x92\xef\xc1\xd9\x17\xe2\x32\xce\x97\xce\x32\x59\xbe\x25\x1f\xb4\xb3\x71\x4d\x9e\x64\xf1\xe9\xca\x35\x7a\x34\xc4\xe4\x83\x58\x40\x24\x04\x20\x3c\x6d\x74\x60\xbf\xfd\x82\x86\xf6\xa7\x00\x82\xb7\x75\xfc\x16\x81\xbd\xb6\x9b\x16\xdf\x9e\x1b\x62\x54\xc8\x78\xd0\x0b\x20\x14\x05\xe9\x75\xcd\xfd\xd6\x55\x49\x60\xd1\x10\xb8\x02\xb8\x24\x88\xbc\x50\x5b\xf2\x30\xac\x4b\x44\x0f\xde\xb5\xef\xdd\xe5\x94\xcd\x27\x27\xb1\x1f\xf6\xff\x18\x55\xfd\xd0\xbf\xb0\x82\x55\xa9\x03\x48\xb4\xd6\x31\xac\x09\x64\x89\x76\x43\x0a\xb0\x60\xf2\x2d\xc8\x53\x70\x8d\x97\x04\xb1\xcf\x13\x32\xa9\x67\xb4\x2c\x1f\x9a\x7f\x91\xa1\xa8\xc0\xa6\xe2\x5b\xac\x9a\xb6\xbe\x64\xb4\x0a\xbd\x3a\x43\xe8\xf2\xe3\xcd\xd9\x37\x7f\x5d\xeb\xd1\x3e\x27\x93\xbe\xca\xf2\xb7\xf3\x3c\x9b\x67\xf9\x19\xb4\xaf\xbf\x7e\x80\xc7\x6e\xf7\xc9\xf4\x51\x19\x6d\x6f\x02\xf9\xf7\x16\xd7\x15\xa9\x63\xe4\xd7\xce\x55\x7f\xa4\x5e\x60\x15\xe8\x74\xce\x8f\x11\x0e\x5c\x22\x83\xb6\x4a\x4b\x64\x0a\xf0\xa3\x24\x2e\x7b\xbf\xb4\xc4\xa0\x09\xe4\xa3\x63\xa8\x63\xf7\x54\xc4\xac\x17\x22\x06\x77\xc5\x80\xde\xb5\xb5\xbd\x0a\xdb\x05\x54\xdc\x8d\x29\x7e\x75\x71\x98\xde\x8b\xd7\xf7\xa3\xba\x41\xf3\x67\x2d\xbd\x0b\xae\xe0\xe4\xdd\x70\x95\xdf\x7a\x4c\xda\x83\x35\x85\x11\x57\x8d\xc1\x3b\xbe\x69\x48\xe6\x64\x1b\x1e\x9a\xe6\x38\x6e\xf4\xd5\x04\x19\x26\xe1\x78\x4e\xe6\xf2\xa1\x89\xe8\xe1\xf2\xf6\x53\x6a\xef\x6a\xf2\xac\x29\x4c\x87\x1d\x71\xc5\x74\xf0\xd3\x86\xc3\xe9\xfb\x5f\x73\x3f\xdb\xfd\x0e\x00\x00\xff\xff\x2a\x1c\x63\x7a\xbe\x05\x00\x00")

func pkgAzureTemplatesAcrJsonBytes() ([]byte, error) {
	return bindataRead(
		_pkgAzureTemplatesAcrJson,
		"pkg/azure/templates/acr.json",
	)
}

func pkgAzureTemplatesAcrJson() (*asset, error) {
	bytes, err := pkgAzureTemplatesAcrJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/azure/templates/acr.json", size: 1470, mode: os.FileMode(420), modTime: time.Unix(1545090262, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _pkgAzureTemplatesStorageaccountJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xcd\xce\xd3\x30\x10\xbc\xf7\x29\x2c\x83\x54\x90\x9a\xbf\x4a\x08\xd4\x1b\x37\x2e\xd0\x03\x11\x97\xaa\x07\xd7\xd9\xb6\xa6\x89\xd7\xf2\x6e\x90\x0a\xea\xbb\x23\xc7\x21\x3f\xa5\xa0\x7e\xf2\x6d\x76\x66\x76\x3d\xbb\xbf\x16\x42\xc8\xd7\xa4\xcf\xd0\x28\xb9\x11\xf2\xcc\xec\x36\x59\x16\x81\xb4\x51\x56\x9d\xa0\x01\xcb\xa9\xfa\xd9\x7a\x48\x35\x36\x7d\x8d\xb2\x75\x5e\xbc\x4b\xf2\x22\xc9\x8b\xac\x02\x57\xe3\x35\xf0\x4a\x68\x5c\xad\x18\xd2\xef\x84\xf6\x95\x5c\x05\x7b\x8d\x96\xc1\xf2\x37\xf0\x64\xd0\x86\x2e\x45\x9a\x87\x17\xcb\x4e\x79\xd5\x00\x83\x27\xb9\x11\x61\x1e\x21\x64\x8d\x5a\x71\x24\x47\x44\x08\xc9\x57\x07\x41\x4c\xec\x8d\x3d\xc9\x0e\xbe\xad\x22\x9f\x18\xbd\x3a\xc1\x47\xad\xb1\xb5\xfc\x45\x35\xf0\xac\x52\x45\x49\x19\x29\x4f\x49\x2e\xc6\x56\x2f\xb0\x07\xa2\xd2\x80\x7f\xfa\x2b\xad\x73\xe8\x99\x3e\x31\x3b\x2a\xbd\x3a\x1e\x8d\xde\xda\xfa\xfa\x48\x7f\x40\xac\x7b\xf5\xa2\x77\x90\x3f\x94\x37\xea\x50\x43\x97\x66\x84\x3c\x10\xb6\x5e\x77\xd0\xae\xa3\x0f\x4e\x36\x46\x25\x77\xe3\x16\xde\x2c\xff\x4e\x73\xf9\x76\x2f\x57\xf7\xdd\x3f\x1b\xed\x91\xf0\xc8\xe9\xd7\x28\xc8\xe6\x42\x1a\x25\xca\x99\xc9\xfa\xd7\x79\xf1\x21\xc9\xdf\x27\x79\x31\x32\x26\x1b\x9f\x0f\xf3\xa7\x30\x1b\xc1\x79\x74\xe0\xd9\x00\x4d\x62\xb9\x8f\x7b\xee\x33\x96\xa6\x4e\xff\x0f\xfc\x2e\x96\x7f\x10\x83\x5f\x6f\x77\x1b\x26\xac\xc0\x81\xad\x68\x1b\x3e\xb4\xdb\x0f\x30\x5d\xda\xf9\xc4\x0f\x17\x30\x39\xca\x87\xe6\xfd\x05\xce\x45\x01\x1c\xd8\xe1\x1e\xba\xae\x12\x5b\x76\x2d\xc7\x6b\x58\xdc\x7e\x07\x00\x00\xff\xff\xca\x20\x75\x04\xf0\x03\x00\x00")

func pkgAzureTemplatesStorageaccountJsonBytes() ([]byte, error) {
	return bindataRead(
		_pkgAzureTemplatesStorageaccountJson,
		"pkg/azure/templates/storageaccount.json",
	)
}

func pkgAzureTemplatesStorageaccountJson() (*asset, error) {
	bytes, err := pkgAzureTemplatesStorageaccountJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "pkg/azure/templates/storageaccount.json", size: 1008, mode: os.FileMode(420), modTime: time.Unix(1545108859, 0)}
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
	"pkg/azure/templates/.DS_Store": pkgAzureTemplatesDs_store,
	"pkg/azure/templates/acr.json": pkgAzureTemplatesAcrJson,
	"pkg/azure/templates/storageaccount.json": pkgAzureTemplatesStorageaccountJson,
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
	"pkg": &bintree{nil, map[string]*bintree{
		"azure": &bintree{nil, map[string]*bintree{
			"templates": &bintree{nil, map[string]*bintree{
				".DS_Store": &bintree{pkgAzureTemplatesDs_store, map[string]*bintree{}},
				"acr.json": &bintree{pkgAzureTemplatesAcrJson, map[string]*bintree{}},
				"storageaccount.json": &bintree{pkgAzureTemplatesStorageaccountJson, map[string]*bintree{}},
			}},
		}},
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

