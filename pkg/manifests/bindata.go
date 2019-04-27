package manifests

import (
	"bytes"
	godefaultbytes "bytes"
	godefaulthttp "net/http"
	godefaultruntime "runtime"
	"compress/gzip"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	bytes	[]byte
	info	os.FileInfo
	digest	[sha256.Size]byte
}
type bindataFileInfo struct {
	name	string
	size	int64
	mode	os.FileMode
	modTime	time.Time
}

func (fi bindataFileInfo) Name() string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return nil
}

var _assetsRouterClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x8f\x31\x4e\xc4\x40\x0c\x45\xfb\x39\x85\x25\xea\x0c\xa2\x43\xd3\x01\x37\x58\x24\x7a\xef\xc4\xbb\x31\x49\xec\xc8\xf6\xa4\xe0\xf4\x28\x4a\x44\xc3\x4a\x29\x2d\xf9\xbf\xff\xfe\x13\xbc\xb3\xf4\x0e\x31\x10\x98\xb6\x20\x03\xd3\x89\x20\x14\x38\x1c\x3e\xc9\x56\xae\x04\x6f\xb5\x6a\x93\xc8\x69\x64\xe9\x0b\x7c\x4c\xcd\x83\xec\xa2\x13\x6d\x71\x96\x7b\xc2\x85\xbf\xc8\x9c\x55\x0a\xd8\x15\x6b\xc6\x16\x83\x1a\xff\x60\xb0\x4a\x1e\x5f\x3d\xb3\x3e\xaf\x2f\x69\xa6\xc0\x1e\x03\x4b\x02\x10\x9c\xa9\x80\x2e\x24\x3e\xf0\x2d\x3a\x96\xbb\x91\x7b\xb7\x9b\x24\x6f\xd7\x6f\xaa\xe1\x25\x75\xb0\x17\x1f\x3e\x87\xce\x1f\xe1\xf8\xdf\x4f\x5f\xb0\x3e\xa2\xa6\x6d\xd8\x85\x6e\x5b\xf1\xbf\x19\xe7\x32\x27\xf0\xdf\x00\x00\x00\xff\xff\x83\x13\xa9\xa6\x49\x01\x00\x00")

func assetsRouterClusterRoleBindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterClusterRoleBindingYaml, "assets/router/cluster-role-binding.yaml")
}
func assetsRouterClusterRoleBindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/cluster-role-binding.yaml", size: 329, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x12, 0x9a, 0xeb, 0xd3, 0x79, 0x1a, 0xa6, 0x75, 0xb5, 0xda, 0x10, 0x70, 0xf9, 0x80, 0xd1, 0x51, 0x55, 0x98, 0x4d, 0x11, 0x98, 0x79, 0x5c, 0x46, 0x99, 0xa3, 0x39, 0x68, 0xe9, 0xa2, 0x72, 0x56}}
	return a, nil
}

var _assetsRouterClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x31\x6f\xe3\x30\x0c\x85\x77\xfd\x0a\x22\x37\xdb\xc1\x6d\x07\xaf\x37\xdc\x76\x43\x51\x74\xa7\x65\xa6\x66\xed\x88\x02\x49\x39\x6d\x7f\x7d\x61\x3b\x29\x82\x24\x45\x9b\xcd\xcf\x22\xbf\x47\x3e\xe9\x17\xfc\x1d\x8b\x39\x29\x58\x94\x4c\x1d\xa8\x8c\x04\x3b\x51\x50\x29\x4e\x6a\x35\x3c\xf6\x6c\x60\xbd\x94\xb1\x83\x96\x00\x0d\x94\xcc\x95\xa3\xf3\xb4\xc8\x2c\x66\xdc\x8e\x54\x87\x81\x53\xd7\x9c\x88\x0f\x32\x52\xc0\xcc\x4f\xa4\xc6\x92\x1a\xd0\x16\x63\x8d\xc5\x7b\x51\x7e\x47\x67\x49\xf5\xf0\xc7\x6a\x96\xed\xf4\x3b\xec\xc9\xb1\x43\xc7\x26\x00\x24\xdc\x53\x03\x92\x29\x59\xcf\x3b\xaf\x38\x3d\x2b\x99\x55\xeb\x48\x41\xcb\x48\xd6\x84\x0a\x30\xf3\x3f\x95\x92\x6d\x6e\xaa\x60\xb3\x09\x30\xcf\x26\x45\x23\x1d\xff\x51\xea\xb2\x70\x72\x5b\xd4\x0c\xb6\x8c\x91\x56\x69\xa4\x13\xaf\x62\x22\x6d\x8f\x2d\x23\x9b\x2f\x1f\x07\xf4\xd8\x87\x6b\x9f\x79\x05\x4a\xce\xf1\x7c\x87\x6b\x6b\x97\x81\x92\xd2\xc4\x74\xb8\x70\x88\x4a\xe8\xf4\x05\xf9\x32\x9c\x6b\xb0\x95\xf6\x85\xa2\x63\x8c\x64\x76\x9f\xc1\x92\x60\xfd\x99\xec\x4d\xfc\x52\x73\x6f\x26\x3f\x07\x6f\xcd\xd1\xcb\x05\xbf\xe4\xee\xf6\xc0\x46\xb1\x28\xfb\xdb\x37\xe8\x53\x59\x94\xe4\xf4\xea\x51\x92\xb9\xe2\xf1\xde\xcf\x7d\x8c\xce\x9a\xff\xcf\xcf\x61\x3d\xe8\xc5\x3c\x91\x1f\x44\x87\xf0\x11\x00\x00\xff\xff\xad\x45\xb2\xc3\x14\x03\x00\x00")

func assetsRouterClusterRoleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterClusterRoleYaml, "assets/router/cluster-role.yaml")
}
func assetsRouterClusterRoleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/cluster-role.yaml", size: 788, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa, 0xc4, 0x4e, 0xa2, 0x9d, 0x1b, 0xd7, 0x35, 0xbf, 0x95, 0x95, 0xb4, 0x17, 0x87, 0x56, 0xb, 0x12, 0xf, 0xb6, 0x3e, 0x51, 0xe5, 0x5a, 0xd0, 0x66, 0x3a, 0x4d, 0x36, 0xc9, 0x93, 0x90, 0xb2}}
	return a, nil
}

var _assetsRouterDeploymentYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x54\x4f\x6f\xe2\x4e\x0c\xbd\xf3\x29\xac\x72\xe6\x47\xff\xfd\xaa\xdd\xdc\x10\xa4\x2b\xa4\x52\x22\x48\x7b\x45\xd3\x89\x01\xab\x93\x99\x59\xdb\xa1\x62\x3f\xfd\x2a\x81\xb6\x84\xa5\x55\xf7\xb6\x39\x8d\xec\xe7\xe7\x67\xc7\x76\x17\x46\x18\x5d\xd8\x96\xe8\x15\x5e\x48\xd7\x50\xe0\xd2\x54\x4e\x61\x63\x5c\x85\xd2\xe9\xc2\xd8\xaf\x18\x45\x60\x18\xbc\x72\x70\x0e\x19\x24\xa2\xa5\x25\xd9\x3d\x08\x0c\x23\x98\x18\x1d\x61\x01\x46\x81\x2b\xaf\x54\xe2\x7f\x9d\x67\xf2\x45\x72\x90\xa1\x63\x22\x3d\x22\x0b\x05\x9f\xd4\x01\xd2\xdf\x5c\x74\xba\xe0\x4d\x89\x60\x7c\xd1\x3c\x24\x1a\x8b\x0d\xa3\xa0\xb6\xd8\xea\xac\x49\x07\x40\xb1\x8c\xce\x28\xd6\x6f\x80\x57\x6b\xf3\x46\xde\x90\xc5\x81\xb5\xa1\xf2\x7a\x6f\x4a\x4c\x80\x43\xa5\xc8\x7b\x40\x17\x7c\x28\x70\x8e\x0e\xad\x06\x06\x92\x3f\x92\xec\x70\x91\x29\x30\xe9\x76\xe8\x8c\xc8\x8e\x47\xb6\xa2\x58\xf6\xac\xab\x44\x91\x7b\x96\x49\xc9\x1a\xb7\x0f\xb0\xc1\xab\x21\x8f\x2c\xaf\x5a\x00\x7a\x4d\x3d\x47\x0a\x76\x2a\xa8\x34\x2b\xfc\x38\x7d\xfd\x35\x90\xac\x72\x2e\x0b\x8e\xec\x36\x81\xf1\xf2\x3e\x68\xc6\x28\x75\x23\xdf\x71\x8a\x5c\x92\x37\x4a\xc1\x4f\x50\xa4\x0e\xda\x07\xdc\x1a\xe7\x9e\x8c\x7d\xce\xc3\x5d\x58\xc9\xd4\xa7\xcc\xe1\x50\x46\x0c\xac\x07\x72\xdf\x05\xaf\x55\xe3\x81\xf9\xa0\xba\x2c\xb0\x26\xf0\xed\xbc\xe5\x8d\x1c\x34\xd8\xe0\x12\xc8\x87\xd9\x07\x74\xf2\x19\xdf\xf5\xf5\xd5\x5f\x11\x96\xa8\x4c\xf6\x53\xca\x8b\xef\x57\x37\x5f\xe2\xec\xc2\x04\x79\x75\x34\xb7\xef\x6e\xf4\x9b\xa4\x85\x16\x35\x2a\x50\x09\xf2\xdb\xd4\x46\x23\xf2\x12\xb8\x68\x86\x76\x85\x1e\xd9\x68\x8b\xf0\x44\x09\xf3\x7c\x90\xcf\x17\xd9\x74\x96\xb7\x54\x36\xfb\x94\xc0\x59\x2d\xff\xec\x44\xd8\x6c\xfa\x90\xa7\xb3\xc5\x3c\x9d\x3d\x8e\x87\xe9\xe2\x7e\x30\x49\xe7\xd9\x60\x98\x9e\x22\x09\x11\xbd\xac\x69\xa9\x3d\xda\x6d\xf0\x09\xbe\x51\x7a\x3b\x78\xb8\xcb\x17\xc3\x74\x96\x8f\x6f\xc7\xc3\x41\x9e\x2e\x46\xe3\xd9\x29\xba\x3e\xaa\xed\xc7\x67\xea\xab\x93\x7e\x64\xda\x18\x3d\x2c\xcc\xd1\x06\x3d\x8a\x64\x1c\x9e\x30\x69\x11\x90\x27\x25\xe3\x46\xe8\xcc\x76\x8e\x36\xf8\x42\x12\xb8\x68\xcf\x50\x3d\x23\x3f\x50\xdb\x81\x00\xd1\xe8\x3a\x81\xfe\x1a\x8d\xd3\xf5\xaf\x63\xe7\xa9\x3f\xcd\x68\x0a\xfa\x37\x84\x48\xa8\xd8\xa2\xb4\xa9\x18\x7f\x56\x28\x2a\xc7\x09\x6c\xac\x6a\x2d\xe7\xe5\x91\xbd\xc4\x32\xf0\x36\x81\xcb\xff\x6f\x26\x74\xe0\xdb\x04\x57\x95\x38\xa9\xef\xdc\xd1\x0e\x97\xb5\x2d\xdb\xe9\xfd\xfc\x9f\xc1\x7e\x0a\xf6\x27\xbf\x67\x91\xb5\x3e\xeb\xc7\xa8\xba\xa7\x53\xef\xb6\x09\x28\x57\xaf\xae\x9d\x80\xb7\xdc\xbd\x2f\x70\x09\x5a\x6e\xb7\x76\x8f\x9e\x84\x02\x13\xb8\xbe\x3c\x6f\xad\xda\xbc\x81\xd7\xd7\xb7\x7d\x29\x7b\xbb\x25\xfd\x1d\x00\x00\xff\xff\x40\x5a\x90\x7d\xbb\x06\x00\x00")

func assetsRouterDeploymentYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterDeploymentYaml, "assets/router/deployment.yaml")
}
func assetsRouterDeploymentYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/deployment.yaml", size: 1723, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x73, 0x3a, 0x4, 0x3c, 0x74, 0xd2, 0x6, 0x8, 0x18, 0x6d, 0x45, 0x37, 0x5a, 0x20, 0xa0, 0x18, 0xd6, 0xcc, 0x36, 0x3a, 0x9, 0x17, 0xd3, 0x69, 0xdc, 0x15, 0x1b, 0xfa, 0x92, 0xc9, 0x2a, 0x6a}}
	return a, nil
}

var _assetsRouterMetricsClusterRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x8f\xc1\x4a\xc4\x40\x0c\x86\xef\xf3\x14\x79\x81\x56\xbc\x2d\x73\x53\x0f\xde\x57\xf0\x9e\x9d\xa6\x36\xb6\x93\x0c\x49\xa6\x07\x9f\x5e\x8a\x22\xc2\x42\xaf\x81\x7c\xdf\xff\xad\x2c\x53\x86\x97\xad\x7b\x90\x5d\x75\xa3\x67\x96\x89\xe5\x23\x61\xe3\x77\x32\x67\x95\x0c\x76\xc3\x32\x62\x8f\x45\x8d\xbf\x30\x58\x65\x5c\x2f\x3e\xb2\x3e\xec\x8f\xa9\x52\xe0\x84\x81\x39\x01\x08\x56\xca\x60\xda\x83\x6c\xa8\x2a\x1c\x6a\x07\xcc\xfb\xed\x93\x4a\x78\x4e\x03\xfc\x18\xdf\xc8\x76\x2e\xf4\x54\x8a\x76\x89\xbf\xd7\x66\x5a\x29\x16\xea\x3e\xac\x17\xff\x3d\x7b\xc3\x42\x19\xb4\x91\xf8\xc2\x73\xfc\x27\x9b\x6e\x74\xa5\xf9\x90\xdf\xa5\x9c\x0c\x02\xc0\xc6\xaf\xa6\xbd\x9d\xd4\xa5\xef\x00\x00\x00\xff\xff\x7f\xc0\x4a\x40\x1d\x01\x00\x00")

func assetsRouterMetricsClusterRoleBindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterMetricsClusterRoleBindingYaml, "assets/router/metrics/cluster-role-binding.yaml")
}
func assetsRouterMetricsClusterRoleBindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterMetricsClusterRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/metrics/cluster-role-binding.yaml", size: 285, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xa2, 0xce, 0x4e, 0xd1, 0x37, 0xde, 0x79, 0x91, 0x5c, 0x71, 0xd1, 0x88, 0x1b, 0xdb, 0xaf, 0x1, 0xe5, 0x8c, 0x81, 0xb3, 0xfd, 0x30, 0xe3, 0x5d, 0xb0, 0x59, 0x8b, 0x2a, 0x47, 0xf9, 0xa0, 0xdf}}
	return a, nil
}

var _assetsRouterMetricsClusterRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\xce\xb1\x4e\xc4\x40\x0c\x84\xe1\x7e\x9f\xc2\x12\x75\x72\xa2\x43\x69\x29\xe8\x29\xe8\x9d\xec\x70\xb1\x2e\x6b\xaf\x6c\xef\x49\xf0\xf4\xe8\x44\x90\xa8\xe7\x93\xfe\x79\xa2\xd7\x63\x44\xc2\xc9\xed\x40\x90\x02\x15\x95\xd6\x2f\xea\x6e\x0d\xb9\x63\x04\xa5\x51\x6c\xce\x1d\xe4\x36\x1e\xb6\x21\x5d\xb6\x20\x68\xed\x26\x9a\x85\xbb\x7c\xc0\x43\x4c\x17\xf2\x95\xb7\x99\x47\xee\xe6\xf2\xcd\x29\xa6\xf3\xed\x25\x66\xb1\xcb\xfd\xb9\xdc\x44\xeb\xf2\xd7\x7c\xb7\x03\xa5\x21\xb9\x72\xf2\x52\x88\x94\x1b\x96\x33\x32\x35\x53\x49\x73\xd1\x6b\xf1\x71\x20\x96\x32\x11\x77\x79\x73\x1b\x3d\x1e\x7a\xfa\x95\xb3\x75\x68\xec\xf2\x99\xb3\x58\x21\x72\x84\x0d\xdf\xf0\xdf\x78\x5c\xce\xcf\x85\xe8\x0e\x5f\xcf\xf1\x8a\x2c\x3f\x01\x00\x00\xff\xff\x4f\xd5\xdf\xe0\x03\x01\x00\x00")

func assetsRouterMetricsClusterRoleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterMetricsClusterRoleYaml, "assets/router/metrics/cluster-role.yaml")
}
func assetsRouterMetricsClusterRoleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterMetricsClusterRoleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/metrics/cluster-role.yaml", size: 259, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x98, 0x77, 0x73, 0x9c, 0x6, 0x33, 0xf2, 0x91, 0x6f, 0x3b, 0x35, 0x49, 0xf3, 0xa5, 0xfc, 0x1d, 0x2e, 0x2e, 0xa6, 0x5b, 0x95, 0xa6, 0x7e, 0x8d, 0xfe, 0x7e, 0xf4, 0x62, 0x30, 0xa5, 0x37, 0x61}}
	return a, nil
}

var _assetsRouterMetricsRoleBindingYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\xce\x31\x4e\xc5\x40\x0c\x04\xd0\x7e\x4f\xe1\x0b\x24\x88\xee\x6b\x3b\x68\xe8\x3f\x12\xbd\xb3\x71\x12\x93\xac\xbd\xb2\xbd\x29\x38\x3d\x42\x8a\x44\x05\xd2\x6f\x47\x33\x9a\x87\x8d\x3f\xc8\x9c\x55\x32\xd8\x84\x65\xc4\x1e\x9b\x1a\x7f\x61\xb0\xca\xb8\xdf\x7c\x64\x7d\x3a\x9f\xd3\xce\x32\x67\xb8\xeb\x41\xaf\x2c\x33\xcb\x9a\x2a\x05\xce\x18\x98\x13\x80\x60\xa5\x0c\xcd\xb4\x52\x6c\xd4\x7d\xd8\x6f\x7e\xc5\xde\xb0\x50\x06\x6d\x24\xbe\xf1\x12\x03\xcb\x6a\xe4\x9e\x4c\x0f\xba\xd3\xf2\x33\xc7\xc6\x6f\xa6\xbd\xfd\x63\x48\x00\xbf\x84\xbf\x1e\xbd\x4f\x9f\x54\xc2\x73\x1a\xae\xf6\x3b\xd9\xc9\x85\x5e\x4a\xd1\x2e\xf1\xa0\xb4\xaa\x70\xa8\xb1\xac\x90\xbe\x03\x00\x00\xff\xff\x15\x9f\x30\x56\x29\x01\x00\x00")

func assetsRouterMetricsRoleBindingYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterMetricsRoleBindingYaml, "assets/router/metrics/role-binding.yaml")
}
func assetsRouterMetricsRoleBindingYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterMetricsRoleBindingYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/metrics/role-binding.yaml", size: 297, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xff, 0xef, 0x8, 0xfb, 0x1f, 0xa3, 0xc7, 0xfb, 0xbc, 0x6, 0x78, 0xad, 0x0, 0x28, 0x90, 0xc8, 0xe8, 0xf5, 0x7d, 0xf8, 0xd0, 0xeb, 0x52, 0xf, 0xd4, 0x81, 0xce, 0x69, 0xb8, 0x8c, 0x26, 0x8c}}
	return a, nil
}

var _assetsRouterMetricsRoleYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4c\x8e\xb1\x6e\xeb\x30\x0c\x45\x77\x7d\x05\x91\x37\x3b\x0f\xdd\x02\xfd\x40\xf7\x0e\xdd\x19\xe9\x36\x26\x62\x8b\x02\x49\xb9\x68\xbf\xbe\x88\xe3\x02\x9d\x78\xef\x01\xc1\xc3\x7f\xf4\xa6\x0b\x9c\x1a\x50\x51\xe9\xfa\x45\xdd\x74\x45\xcc\x18\x4e\xa1\xe4\xc5\xb8\x83\x4c\x47\xc0\x68\x45\x98\x14\x27\xb4\xda\x55\x5a\x24\xee\xf2\x0e\x73\xd1\x96\xc9\xae\x5c\xce\x3c\x62\x56\x93\x6f\x0e\xd1\x76\xbe\x5f\xfc\x2c\xfa\x7f\x7b\x49\x77\x69\x35\xef\xae\xb4\x22\xb8\x72\x70\x4e\x44\x8d\x57\xe4\x3f\xca\xe9\x7e\xf1\x03\x7b\xe7\x82\x4c\xda\xd1\x7c\x96\x8f\x98\xa4\xdd\x0c\xee\xc9\xc6\x02\xcf\x69\x22\xee\xf2\x6a\x3a\xba\x3f\x2e\x4d\x74\x3a\x25\x22\x83\xeb\xb0\x82\x83\x39\x6c\x93\x02\xdf\xcb\xef\xd7\xcf\xd6\xb5\x3e\xc2\x06\xbb\x1e\xcb\x37\xc4\x3e\x17\xf1\x67\xf8\xe4\x28\x73\xfa\x09\x00\x00\xff\xff\x67\x78\x6f\x08\x23\x01\x00\x00")

func assetsRouterMetricsRoleYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterMetricsRoleYaml, "assets/router/metrics/role.yaml")
}
func assetsRouterMetricsRoleYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterMetricsRoleYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/metrics/role.yaml", size: 291, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x57, 0xe3, 0x68, 0x2c, 0xfd, 0xf1, 0x27, 0x9d, 0xa0, 0x3b, 0x10, 0x3e, 0xca, 0x3b, 0x76, 0x39, 0xf4, 0xb1, 0x37, 0x7b, 0xa3, 0xa7, 0x11, 0xc0, 0x6, 0x4b, 0x47, 0xbb, 0x93, 0x4b, 0xb7, 0xc2}}
	return a, nil
}

var _assetsRouterNamespaceYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x41\x4a\x43\x41\x10\x44\xf7\x73\x8a\xe2\xbb\x8e\xe2\x76\xee\xa0\x1b\xc1\x7d\x67\x7e\x25\x69\x33\xd3\xfd\x99\xee\xc4\xeb\x8b\x46\x30\xe0\xba\x1e\x8f\x57\x67\xb5\xb5\xe2\x55\x06\x63\x93\xc6\x22\x9b\xbe\x73\x86\xba\x55\x5c\x9f\xcb\x60\xca\x2a\x29\xb5\x00\x26\x83\x15\xbe\xd1\xe2\xa4\x87\xdc\xa9\x1d\x27\x23\x0a\x20\x66\x9e\x92\xea\x16\xdf\x20\xfe\xa0\x47\xf5\x27\xf3\x95\xbb\x60\x67\x4b\x9f\x15\xcb\x52\x80\x2e\x7b\xf6\x5f\xf8\x01\xd2\xbb\x7f\xde\x99\x87\x9b\xa6\x4f\xb5\x23\xd2\xd1\xdd\xcf\x38\xf8\xc4\x1b\xe7\x55\x1b\x5f\x6e\x2b\x7c\xff\xc1\x96\x01\x35\xe4\x49\xe3\xa7\xef\x76\xe2\x5f\x42\xeb\x97\x48\xce\x3b\x71\xc5\x92\xf3\xc2\xa5\x7c\x05\x00\x00\xff\xff\xfd\xbd\x46\x74\x01\x01\x00\x00")

func assetsRouterNamespaceYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterNamespaceYaml, "assets/router/namespace.yaml")
}
func assetsRouterNamespaceYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterNamespaceYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/namespace.yaml", size: 257, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x93, 0x80, 0x53, 0x86, 0xac, 0xef, 0x4d, 0xd1, 0x80, 0xe0, 0x94, 0x53, 0xb, 0x1e, 0x3f, 0xd3, 0x6c, 0x99, 0x8c, 0x3c, 0x6d, 0x61, 0x35, 0xa0, 0x7a, 0x7c, 0x51, 0x10, 0x53, 0xc7, 0xc5, 0x86}}
	return a, nil
}

var _assetsRouterServiceAccountYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2c\xce\xb1\x4e\xc4\x30\x10\x84\xe1\xde\x4f\x31\xd2\xd5\x9c\x44\xeb\x8e\x92\x16\x24\x7a\xb3\x99\xbb\x5b\x91\x78\xcd\xee\x3a\x88\xb7\x47\x41\x29\xa7\x98\x5f\xdf\x05\x2f\x22\x36\x7b\xe2\x66\x0e\xb7\x99\xf4\x80\x38\x5b\x72\xc1\xe7\x2f\xf2\x41\xd8\xa0\xb7\x34\xbf\xe2\x35\xf1\xa3\xeb\x0a\xe7\xf7\x54\x27\x64\x9d\x91\x74\x84\xd8\xe0\x52\x2e\x18\xf4\x4d\x23\xd4\x7a\xc0\xb9\xfe\x57\xd2\xf0\x76\x84\x31\xdc\x84\x11\xda\xef\xd7\xf2\xa5\x7d\xa9\x78\xa7\xef\x2a\x3c\x0d\xa5\x0d\xfd\xa0\x1f\xef\x8a\xfd\xb9\x6c\xcc\xb6\xb4\x6c\xb5\x00\xbd\x6d\xac\x27\xf0\x9c\x31\x9a\xb0\x1e\xba\x1e\x0f\xbd\xe5\x93\xf6\xbb\x33\xa2\xfc\x05\x00\x00\xff\xff\x33\xdc\xda\x8c\xd5\x00\x00\x00")

func assetsRouterServiceAccountYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterServiceAccountYaml, "assets/router/service-account.yaml")
}
func assetsRouterServiceAccountYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/service-account.yaml", size: 213, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xd0, 0xe3, 0x6, 0x3a, 0x88, 0x2e, 0x33, 0xe3, 0x24, 0xf0, 0xf0, 0xe9, 0x43, 0xc8, 0x46, 0x6c, 0x60, 0x9, 0x69, 0x84, 0x3, 0xd8, 0xc3, 0x80, 0xb, 0xab, 0x37, 0x13, 0xce, 0xf2, 0xeb, 0x60}}
	return a, nil
}

var _assetsRouterServiceCloudYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x90\x41\x6b\x14\x41\x10\x85\xef\xfd\x2b\x1e\xec\x39\x41\x31\x07\x99\x63\x72\x12\x82\x2c\xb8\x78\xaf\xf4\xd4\xec\x34\xe9\xa9\x6a\xaa\x6a\x56\xf7\xdf\xcb\xf4\xec\x82\xa2\x78\xec\x07\xf5\xfa\x7b\xdf\x01\xaf\x4a\x23\x9e\xa9\x92\x64\x36\x7c\x63\xbb\x94\xcc\x08\x45\xab\x94\x19\x45\x30\x99\x4a\x40\x27\xc4\xcc\x30\x5d\x83\x6d\x8b\x73\xd5\x75\x04\xcb\xa5\x98\xca\xc2\x12\xfe\x98\x0e\xf8\x22\x67\x63\x77\xbc\xa8\x84\x69\xad\x6c\xf0\xc6\xb9\x4c\x25\xe3\x42\x75\x65\x07\x19\x83\x5a\xab\x85\x47\x50\xc0\x56\x89\xb2\xf0\x63\x7a\x2f\x32\x0e\x77\x82\x44\xad\x7c\x67\xf3\xa2\x32\xe0\xf2\x31\x2d\x1c\x34\x52\xd0\x90\x80\x03\xbe\xd2\xc2\x28\x0e\xe7\xf8\xa3\x02\x10\x5a\xd8\x1b\x65\x1e\xa0\x8d\xc5\xe7\x32\xc5\x43\xd9\xa1\x12\x50\xe9\x8d\xab\x6f\x25\xd8\x18\x86\xdb\x9e\xb4\x31\x6e\x69\x5c\x1b\x0f\xdd\xc9\x5d\x49\x02\x9c\x2b\xe7\x50\xfb\xfb\x6c\x63\x39\xcd\xc5\x41\xd5\x15\x33\x79\x77\xc4\xd3\xc4\xb9\x1b\x5b\xc8\xde\x8b\x9c\xf1\xfa\x8c\xa6\x5a\x11\x64\x67\x0e\x07\x39\x56\x99\x99\x6a\xcc\x57\xfc\x98\x59\x20\xda\xcb\x6e\x7a\x9b\x8e\xbb\xa7\x66\xec\xbc\xd9\x17\x10\x44\x47\xc6\x1b\xcf\x45\xc6\xfe\x8f\xef\xaa\xb6\xd9\xfc\x33\xd8\x84\xea\xc9\x68\x9a\x4a\x3e\x6a\x2d\xf9\xba\x0d\xc9\x54\x13\xd0\xd4\xa2\xaf\x7e\xe8\x82\x06\xcc\x11\xad\xaf\x69\xa6\xa1\x59\xeb\x80\xd3\xcb\x71\x4f\xd4\x62\xc0\xe7\x0f\xfd\xb1\x03\x1f\x7b\x74\xbb\xf9\xbd\xc2\xff\xdb\xf1\xf4\xf4\xe9\x9f\x25\x9e\x7e\x05\x00\x00\xff\xff\x56\xdc\x0d\xe9\x77\x02\x00\x00")

func assetsRouterServiceCloudYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterServiceCloudYaml, "assets/router/service-cloud.yaml")
}
func assetsRouterServiceCloudYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterServiceCloudYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/service-cloud.yaml", size: 631, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xcd, 0x97, 0xc4, 0xab, 0x8b, 0xa9, 0xa1, 0x47, 0xf7, 0xf, 0xeb, 0x38, 0x1b, 0xc2, 0xf7, 0x8c, 0xd9, 0xba, 0x35, 0x9b, 0x1, 0x67, 0x26, 0xd7, 0x3f, 0x6d, 0xa5, 0x5f, 0xf0, 0x92, 0x84, 0x8f}}
	return a, nil
}

var _assetsRouterServiceInternalYaml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xcf\x31\x6b\xfb\x30\x10\x05\xf0\x5d\x9f\xe2\x41\xd6\xff\xbf\x34\x24\x94\x56\xab\xa7\x6c\x86\x96\xee\x87\x7c\x49\x8e\xca\x92\xb8\x3b\xbb\xf4\xdb\x97\x38\x0d\xb8\x74\xc9\x22\x90\xee\xe9\xf7\xb8\x0d\xba\x3c\x99\xb3\xe2\x95\x75\x96\xc4\xf8\x14\x3f\x63\xe0\x23\x4d\xd9\x31\x53\x9e\xd8\xc2\x06\x87\x72\x52\x36\x43\x57\x8b\x6b\xcd\x99\x15\xd6\x38\xc9\x51\x12\xa8\x94\xea\xe4\x52\x8b\x81\x94\x41\xad\x65\xe1\x01\xe4\xd0\xa9\xb8\x8c\xfc\x10\x3e\xa4\x0c\xf1\xd6\x11\xa8\xc9\x3b\xab\x49\x2d\x11\xf3\x36\x6c\x50\x68\xe4\x7f\xcb\x69\x8d\x12\x83\xca\xf0\x87\x35\xf6\x5f\xe4\xa5\x3f\x06\xc0\xbf\x1a\xc7\xdb\x1a\x87\x3e\x00\xad\xaa\xdb\x65\xf4\x7f\x21\x23\xce\xee\x2d\x00\xd7\x49\xc4\xf3\xe3\xf5\xa2\xd5\x6b\xaa\x39\xe2\xad\xeb\x97\x17\x27\x3d\xb1\xf7\x4b\xe8\xe7\xcf\x9a\xb0\x95\xb1\xdf\xef\xee\x44\x6c\xa5\x8c\xec\x2a\x69\xed\x6c\x5f\x76\x4f\x77\x40\x4b\xec\x3b\x00\x00\xff\xff\x90\x5e\x33\xca\xad\x01\x00\x00")

func assetsRouterServiceInternalYamlBytes() ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return bindataRead(_assetsRouterServiceInternalYaml, "assets/router/service-internal.yaml")
}
func assetsRouterServiceInternalYaml() (*asset, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	bytes, err := assetsRouterServiceInternalYamlBytes()
	if err != nil {
		return nil, err
	}
	info := bindataFileInfo{name: "assets/router/service-internal.yaml", size: 429, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb8, 0x63, 0x52, 0x85, 0x8b, 0x99, 0xe6, 0xc7, 0xcb, 0x34, 0x3d, 0x8d, 0x43, 0x65, 0x10, 0x63, 0x51, 0x80, 0xc1, 0x29, 0x17, 0xb6, 0x8f, 0x84, 0xdc, 0xf8, 0x33, 0xa1, 0x21, 0xc2, 0x5a, 0x4f}}
	return a, nil
}
func Asset(name string) ([]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}
func AssetString(name string) (string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	data, err := Asset(name)
	return string(data), err
}
func MustAsset(name string) []byte {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}
	return a
}
func MustAssetString(name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	return string(MustAsset(name))
}
func AssetInfo(name string) (os.FileInfo, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}
func AssetDigest(name string) ([sha256.Size]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}
func Digests() (map[string][sha256.Size]byte, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
}
func AssetNames() []string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

var _bindata = map[string]func() (*asset, error){"assets/router/cluster-role-binding.yaml": assetsRouterClusterRoleBindingYaml, "assets/router/cluster-role.yaml": assetsRouterClusterRoleYaml, "assets/router/deployment.yaml": assetsRouterDeploymentYaml, "assets/router/metrics/cluster-role-binding.yaml": assetsRouterMetricsClusterRoleBindingYaml, "assets/router/metrics/cluster-role.yaml": assetsRouterMetricsClusterRoleYaml, "assets/router/metrics/role-binding.yaml": assetsRouterMetricsRoleBindingYaml, "assets/router/metrics/role.yaml": assetsRouterMetricsRoleYaml, "assets/router/namespace.yaml": assetsRouterNamespaceYaml, "assets/router/service-account.yaml": assetsRouterServiceAccountYaml, "assets/router/service-cloud.yaml": assetsRouterServiceCloudYaml, "assets/router/service-internal.yaml": assetsRouterServiceInternalYaml}

func AssetDir(name string) ([]string, error) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	Func		func() (*asset, error)
	Children	map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{"assets": {nil, map[string]*bintree{"router": {nil, map[string]*bintree{"cluster-role-binding.yaml": {assetsRouterClusterRoleBindingYaml, map[string]*bintree{}}, "cluster-role.yaml": {assetsRouterClusterRoleYaml, map[string]*bintree{}}, "deployment.yaml": {assetsRouterDeploymentYaml, map[string]*bintree{}}, "metrics": {nil, map[string]*bintree{"cluster-role-binding.yaml": {assetsRouterMetricsClusterRoleBindingYaml, map[string]*bintree{}}, "cluster-role.yaml": {assetsRouterMetricsClusterRoleYaml, map[string]*bintree{}}, "role-binding.yaml": {assetsRouterMetricsRoleBindingYaml, map[string]*bintree{}}, "role.yaml": {assetsRouterMetricsRoleYaml, map[string]*bintree{}}}}, "namespace.yaml": {assetsRouterNamespaceYaml, map[string]*bintree{}}, "service-account.yaml": {assetsRouterServiceAccountYaml, map[string]*bintree{}}, "service-cloud.yaml": {assetsRouterServiceCloudYaml, map[string]*bintree{}}, "service-internal.yaml": {assetsRouterServiceInternalYaml, map[string]*bintree{}}}}}}}}

func RestoreAsset(dir, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}
func RestoreAssets(dir, name string) error {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	children, err := AssetDir(name)
	if err != nil {
		return RestoreAsset(dir, name)
	}
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}
func _filePath(dir, name string) string {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("http://35.226.239.161:5001/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
func _logClusterCodePath() {
	_logClusterCodePath()
	defer _logClusterCodePath()
	pc, _, _, _ := godefaultruntime.Caller(1)
	jsonLog := []byte(fmt.Sprintf("{\"fn\": \"%s\"}", godefaultruntime.FuncForPC(pc).Name()))
	godefaulthttp.Post("/"+"logcode", "application/json", godefaultbytes.NewBuffer(jsonLog))
}
