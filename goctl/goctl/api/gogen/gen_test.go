package gogen

import (
	_ "embed"
	goformat "go/format"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/rpc/execx"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var (
	//go:embed testdata/test_api_template.api
	testApiTemplate string
	//go:embed testdata/test_multi_service_template.api
	testMultiServiceTemplate string
	//go:embed testdata/ap_ino_info.api
	apiNoInfo string
	//go:embed testdata/invalid_api_file.api
	invalidApiFile string
	//go:embed testdata/anonymous_annotation.api
	anonymousAnnotation string
	//go:embed testdata/api_has_middleware.api
	apiHasMiddleware string
	//go:embed testdata/api_jwt.api
	apiJwt string
	//go:embed testdata/api_jwt_with_middleware.api
	apiJwtWithMiddleware string
	//go:embed testdata/api_has_no_request.api
	apiHasNoRequest string
	//go:embed testdata/api_route_test.api
	apiRouteTest string
	//go:embed testdata/has_comment_api_test.api
	hasCommentApiTest string
	//go:embed testdata/has_inline_no_exist_test.api
	hasInlineNoExistTest string
	//go:embed testdata/import_api.api
	importApi string
	//go:embed testdata/has_import_api.api
	hasImportApi string
	//go:embed testdata/no_struct_tag_api.api
	noStructTagApi string
	//go:embed testdata/nest_type_api.api
	nestTypeApi string
)

func TestParser(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(testApiTemplate), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	api, err := parser.Parse(filename)
	assert.Nil(t, err)

	assert.Equal(t, len(api.Types), 2)
	assert.Equal(t, len(api.Service.Routes()), 2)

	assert.Equal(t, api.Service.Routes()[0].Path, "/greet/from/:name")
	assert.Equal(t, api.Service.Routes()[1].Path, "/greet/get")

	assert.Equal(t, api.Service.Routes()[1].RequestTypeName(), "Request")
	assert.Equal(t, api.Service.Routes()[1].ResponseType, nil)

	validate(t, filename)
}

func TestMultiService(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(testMultiServiceTemplate), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	api, err := parser.Parse(filename)
	assert.Nil(t, err)

	assert.Equal(t, len(api.Service.Routes()), 2)
	assert.Equal(t, len(api.Service.Groups), 2)

	validate(t, filename)
}

func TestApiNoInfo(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(apiNoInfo), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.Nil(t, err)

	validate(t, filename)
}

func TestInvalidApiFile(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(invalidApiFile), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.NotNil(t, err)
}

func TestAnonymousAnnotation(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(anonymousAnnotation), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	api, err := parser.Parse(filename)
	assert.Nil(t, err)

	assert.Equal(t, len(api.Service.Routes()), 1)
	assert.Equal(t, api.Service.Routes()[0].Handler, "GreetHandler")

	validate(t, filename)
}

func TestApiHasMiddleware(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(apiHasMiddleware), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.Nil(t, err)

	validate(t, filename)
}

func TestApiHasJwt(t *testing.T) {
	filename := "jwt.api"
	err := ioutil.WriteFile(filename, []byte(apiJwt), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.Nil(t, err)

	validate(t, filename)
}

func TestApiHasJwtAndMiddleware(t *testing.T) {
	filename := "jwt.api"
	err := ioutil.WriteFile(filename, []byte(apiJwtWithMiddleware), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.Nil(t, err)

	validate(t, filename)
}

func TestApiHasNoRequestBody(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(apiHasNoRequest), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.Nil(t, err)
}

func TestApiRoutes(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(apiRouteTest), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.Nil(t, err)

	validate(t, filename)
}

func TestHasCommentRoutes(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(hasCommentApiTest), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.Nil(t, err)

	validate(t, filename)
}

func TestInlineTypeNotExist(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(hasInlineNoExistTest), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.NotNil(t, err)
}

func TestHasImportApi(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(hasImportApi), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	importApiName := "importApi.api"
	err = ioutil.WriteFile(importApiName, []byte(importApi), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(importApiName)

	api, err := parser.Parse(filename)
	assert.Nil(t, err)

	var hasInline bool
	for _, ty := range api.Types {
		if ty.Name() == "ImportData" {
			hasInline = true
			break
		}
	}
	assert.True(t, hasInline)

	validate(t, filename)
}

func TestNoStructApi(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(noStructTagApi), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	spec, err := parser.Parse(filename)
	assert.Nil(t, err)
	assert.Equal(t, len(spec.Types), 5)

	validate(t, filename)
}

func TestNestTypeApi(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(nestTypeApi), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.NotNil(t, err)
}

func TestCamelStyle(t *testing.T) {
	filename := "greet.api"
	err := ioutil.WriteFile(filename, []byte(testApiTemplate), os.ModePerm)
	assert.Nil(t, err)
	defer os.Remove(filename)

	_, err = parser.Parse(filename)
	assert.Nil(t, err)

	validateWithCamel(t, filename, "GoZero")
}

func validate(t *testing.T, api string) {
	validateWithCamel(t, api, "gozero")
}

func validateWithCamel(t *testing.T, api, camel string) {
	dir := "workspace"
	defer func() {
		os.RemoveAll(dir)
	}()
	err := pathx.MkdirIfNotExist(dir)
	assert.Nil(t, err)
	err = initMod(dir)
	assert.Nil(t, err)
	err = DoGenProject(api, dir, camel)
	assert.Nil(t, err)
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".go") {
			code, err := ioutil.ReadFile(path)
			assert.Nil(t, err)
			assert.Nil(t, validateCode(string(code)))
		}
		return nil
	})
}

func initMod(mod string) error {
	_, err := execx.Run("go mod init "+mod, mod)
	return err
}

func validateCode(code string) error {
	_, err := goformat.Source([]byte(code))
	return err
}
