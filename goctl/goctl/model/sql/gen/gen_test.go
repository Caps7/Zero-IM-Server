package gen

import (
	"database/sql"
	_ "embed"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stringx"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/builderx"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/parser"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

//go:embed testdata/user.sql
var source string

func TestCacheModel(t *testing.T) {
	logx.Disable()
	_ = Clean()

	sqlFile := filepath.Join(pathx.MustTempDir(), "tmp.sql")
	err := ioutil.WriteFile(sqlFile, []byte(source), 0o777)
	assert.Nil(t, err)

	dir := filepath.Join(pathx.MustTempDir(), "./testmodel")
	cacheDir := filepath.Join(dir, "cache")
	noCacheDir := filepath.Join(dir, "nocache")
	g, err := NewDefaultGenerator(cacheDir, &config.Config{
		NamingFormat: "GoZero",
	})
	assert.Nil(t, err)

	err = g.StartFromDDL(sqlFile, true, "go_zero")
	assert.Nil(t, err)
	assert.True(t, func() bool {
		_, err := os.Stat(filepath.Join(cacheDir, "TestUserModel.go"))
		return err == nil
	}())
	g, err = NewDefaultGenerator(noCacheDir, &config.Config{
		NamingFormat: "gozero",
	})
	assert.Nil(t, err)

	err = g.StartFromDDL(sqlFile, false, "go_zero")
	assert.Nil(t, err)
	assert.True(t, func() bool {
		_, err := os.Stat(filepath.Join(noCacheDir, "testusermodel.go"))
		return err == nil
	}())
}

func TestNamingModel(t *testing.T) {
	logx.Disable()
	_ = Clean()

	sqlFile := filepath.Join(pathx.MustTempDir(), "tmp.sql")
	err := ioutil.WriteFile(sqlFile, []byte(source), 0o777)
	assert.Nil(t, err)

	dir, _ := filepath.Abs("./testmodel")
	camelDir := filepath.Join(dir, "camel")
	snakeDir := filepath.Join(dir, "snake")
	defer func() {
		_ = os.RemoveAll(dir)
	}()
	g, err := NewDefaultGenerator(camelDir, &config.Config{
		NamingFormat: "GoZero",
	})
	assert.Nil(t, err)

	err = g.StartFromDDL(sqlFile, true, "go_zero")
	assert.Nil(t, err)
	assert.True(t, func() bool {
		_, err := os.Stat(filepath.Join(camelDir, "TestUserModel.go"))
		return err == nil
	}())
	g, err = NewDefaultGenerator(snakeDir, &config.Config{
		NamingFormat: "go_zero",
	})
	assert.Nil(t, err)

	err = g.StartFromDDL(sqlFile, true, "go_zero")
	assert.Nil(t, err)
	assert.True(t, func() bool {
		_, err := os.Stat(filepath.Join(snakeDir, "test_user_model.go"))
		return err == nil
	}())
}

func TestWrapWithRawString(t *testing.T) {
	assert.Equal(t, "``", wrapWithRawString("", false))
	assert.Equal(t, "``", wrapWithRawString("``", false))
	assert.Equal(t, "`a`", wrapWithRawString("a", false))
	assert.Equal(t, "a", wrapWithRawString("a", true))
	assert.Equal(t, "`   `", wrapWithRawString("   ", false))
}

func TestFields(t *testing.T) {
	type Student struct {
		ID         int64           `db:"id"`
		Name       string          `db:"name"`
		Age        sql.NullInt64   `db:"age"`
		Score      sql.NullFloat64 `db:"score"`
		CreateTime time.Time       `db:"create_time"`
		UpdateTime sql.NullTime    `db:"update_time"`
	}
	var (
		studentFieldNames          = builderx.RawFieldNames(&Student{})
		studentRows                = strings.Join(studentFieldNames, ",")
		studentRowsExpectAutoSet   = strings.Join(stringx.Remove(studentFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
		studentRowsWithPlaceHolder = strings.Join(stringx.Remove(studentFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"
	)

	assert.Equal(t, []string{"`id`", "`name`", "`age`", "`score`", "`create_time`", "`update_time`"}, studentFieldNames)
	assert.Equal(t, "`id`,`name`,`age`,`score`,`create_time`,`update_time`", studentRows)
	assert.Equal(t, "`name`,`age`,`score`", studentRowsExpectAutoSet)
	assert.Equal(t, "`name`=?,`age`=?,`score`=?", studentRowsWithPlaceHolder)
}

func Test_genPublicModel(t *testing.T) {
	var err error
	dir := pathx.MustTempDir()
	modelDir := path.Join(dir, "model")
	err = os.MkdirAll(modelDir, 0777)
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	modelFilename := filepath.Join(modelDir, "foo.sql")
	err = ioutil.WriteFile(modelFilename, []byte(source), 0777)
	require.NoError(t, err)

	g, err := NewDefaultGenerator(modelDir, &config.Config{
		NamingFormat: config.DefaultFormat,
	})
	require.NoError(t, err)

	tables, err := parser.Parse(modelFilename, "")
	require.Equal(t, 1, len(tables))

	code, err := g.genModelCustom(*tables[0], false)
	assert.NoError(t, err)
	assert.True(t, strings.Contains(code, "package model"))
	assert.True(t, strings.Contains(code, "TestUserModel interface {\n\t\ttestUserModel\n\t}\n"))
	assert.True(t, strings.Contains(code, "customTestUserModel struct {\n\t\t*defaultTestUserModel\n\t}\n"))
	assert.True(t, strings.Contains(code, "func NewTestUserModel(conn sqlx.SqlConn) TestUserModel {"))
}
