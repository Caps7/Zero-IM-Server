package env

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/zeromicro/go-zero/tools/goctl/internal/version"
	sortedmap "github.com/zeromicro/go-zero/tools/goctl/pkg/collection"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/protoc"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/protocgengo"
	"github.com/zeromicro/go-zero/tools/goctl/pkg/protocgengogrpc"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var goctlEnv *sortedmap.SortedMap

const (
	GoctlOS                = "GOCTL_OS"
	GoctlArch              = "GOCTL_ARCH"
	GoctlHome              = "GOCTL_HOME"
	GoctlDebug             = "GOCTL_DEBUG"
	GoctlCache             = "GOCTL_CACHE"
	GoctlVersion           = "GOCTL_VERSION"
	ProtocVersion          = "PROTOC_VERSION"
	ProtocGenGoVersion     = "PROTOC_GEN_GO_VERSION"
	ProtocGenGoGRPCVersion = "PROTO_GEN_GO_GRPC_VERSION"

	envFileDir = "env"
)

// init initializes the goctl environment variables, the environment variables of the function are set in order,
// please do not change the logic order of the code.
func init() {
	defaultGoctlHome, err := pathx.GetDefaultGoctlHome()
	if err != nil {
		log.Fatalln(err)
	}
	goctlEnv = sortedmap.New()
	goctlEnv.SetKV(GoctlOS, runtime.GOOS)
	goctlEnv.SetKV(GoctlArch, runtime.GOARCH)
	existsEnv := readEnv(defaultGoctlHome)
	if existsEnv != nil {
		goctlHome, ok := existsEnv.GetString(GoctlHome)
		if ok && len(goctlHome) > 0 {
			goctlEnv.SetKV(GoctlHome, goctlHome)
		}
		if debug := existsEnv.GetOr(GoctlDebug, "").(string); debug != "" {
			if strings.EqualFold(debug, "true") || strings.EqualFold(debug, "false") {
				goctlEnv.SetKV(GoctlDebug, debug)
			}
		}
		if value := existsEnv.GetStringOr(GoctlCache, ""); value != "" {
			goctlEnv.SetKV(GoctlCache, value)
		}
	}
	if !goctlEnv.HasKey(GoctlHome) {
		goctlEnv.SetKV(GoctlHome, defaultGoctlHome)
	}
	if !goctlEnv.HasKey(GoctlDebug) {
		goctlEnv.SetKV(GoctlDebug, "False")
	}

	if !goctlEnv.HasKey(GoctlCache) {
		cacheDir, _ := pathx.GetCacheDir()
		goctlEnv.SetKV(GoctlCache, cacheDir)
	}

	goctlEnv.SetKV(GoctlVersion, version.BuildVersion)
	protocVer, _ := protoc.Version()
	goctlEnv.SetKV(ProtocVersion, protocVer)

	protocGenGoVer, _ := protocgengo.Version()
	goctlEnv.SetKV(ProtocGenGoVersion, protocGenGoVer)

	protocGenGoGrpcVer, _ := protocgengogrpc.Version()
	goctlEnv.SetKV(ProtocGenGoGRPCVersion, protocGenGoGrpcVer)
}

func Print() string {
	return strings.Join(goctlEnv.Format(), "\n")
}

func Get(key string) string {
	return GetOr(key, "")
}

func GetOr(key, def string) string {
	return goctlEnv.GetStringOr(key, def)
}

func readEnv(goctlHome string) *sortedmap.SortedMap {
	envFile := filepath.Join(goctlHome, envFileDir)
	data, err := ioutil.ReadFile(envFile)
	if err != nil {
		return nil
	}
	dataStr := string(data)
	lines := strings.Split(dataStr, "\n")
	sm := sortedmap.New()
	for _, line := range lines {
		_, _, err = sm.SetExpression(line)
		if err != nil {
			continue
		}
	}
	return sm
}

func WriteEnv(kv []string) error {
	defaultGoctlHome, err := pathx.GetDefaultGoctlHome()
	if err != nil {
		log.Fatalln(err)
	}
	data := sortedmap.New()
	for _, e := range kv {
		_, _, err := data.SetExpression(e)
		if err != nil {
			return err
		}
	}
	data.RangeIf(func(key, value interface{}) bool {
		switch key.(string) {
		case GoctlHome, GoctlCache:
			path := value.(string)
			if !pathx.FileExists(path) {
				err = fmt.Errorf("[writeEnv]: path %q is not exists", path)
				return false
			}
		}
		if goctlEnv.HasKey(key) {
			goctlEnv.SetKV(key, value)
			return true
		} else {
			err = fmt.Errorf("[writeEnv]: invalid key: %v", key)
			return false
		}
	})
	if err != nil {
		return err
	}
	envFile := filepath.Join(defaultGoctlHome, envFileDir)
	return ioutil.WriteFile(envFile, []byte(strings.Join(goctlEnv.Format(), "\n")), 0o777)
}
