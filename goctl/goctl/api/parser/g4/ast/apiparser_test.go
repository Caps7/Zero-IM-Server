package ast

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

func Test_ImportCycle(t *testing.T) {
	const (
		mainFilename = "main.api"
		subAFilename = "a.api"
		subBFilename = "b.api"
		mainSrc      = `import "./a.api"`
		subASrc      = `import "./b.api"`
		subBSrc      = `import "./a.api"`
	)
	var err error
	dir := pathx.MustTempDir()
	defer os.RemoveAll(dir)

	mainPath := filepath.Join(dir, mainFilename)
	err = ioutil.WriteFile(mainPath, []byte(mainSrc), 0777)
	require.NoError(t, err)
	subAPath := filepath.Join(dir, subAFilename)
	err = ioutil.WriteFile(subAPath, []byte(subASrc), 0777)
	require.NoError(t, err)
	subBPath := filepath.Join(dir, subBFilename)
	err = ioutil.WriteFile(subBPath, []byte(subBSrc), 0777)
	require.NoError(t, err)

	_, err = NewParser().Parse(mainPath)
	assert.ErrorIs(t, err, ErrImportCycleNotAllowed)
}

func Test_MultiImportedShouldAllowed(t *testing.T) {
	const (
		mainFilename = "main.api"
		subAFilename = "a.api"
		subBFilename = "b.api"
		mainSrc      = "import \"./b.api\"\n" +
			"import \"./a.api\"\n" +
			"type Main { b B `json:\"b\"`}"
		subASrc = "import \"./b.api\"\n" +
			"type A { b B `json:\"b\"`}\n"
		subBSrc = `type B{}`
	)
	var err error
	dir := pathx.MustTempDir()
	defer os.RemoveAll(dir)

	mainPath := filepath.Join(dir, mainFilename)
	err = ioutil.WriteFile(mainPath, []byte(mainSrc), 0777)
	require.NoError(t, err)
	subAPath := filepath.Join(dir, subAFilename)
	err = ioutil.WriteFile(subAPath, []byte(subASrc), 0777)
	require.NoError(t, err)
	subBPath := filepath.Join(dir, subBFilename)
	err = ioutil.WriteFile(subBPath, []byte(subBSrc), 0777)
	require.NoError(t, err)

	_, err = NewParser().Parse(mainPath)
	assert.NoError(t, err)
}

func Test_RedundantDeclarationShouldNotBeAllowed(t *testing.T) {
	const (
		mainFilename = "main.api"
		subAFilename = "a.api"
		subBFilename = "b.api"
		mainSrc      = "import \"./a.api\"\n" +
			"import \"./b.api\"\n"
		subASrc = `import "./b.api"
							 type A{}`
		subBSrc = `type A{}`
	)
	var err error
	dir := pathx.MustTempDir()
	defer os.RemoveAll(dir)

	mainPath := filepath.Join(dir, mainFilename)
	err = ioutil.WriteFile(mainPath, []byte(mainSrc), 0777)
	require.NoError(t, err)
	subAPath := filepath.Join(dir, subAFilename)
	err = ioutil.WriteFile(subAPath, []byte(subASrc), 0777)
	require.NoError(t, err)
	subBPath := filepath.Join(dir, subBFilename)
	err = ioutil.WriteFile(subBPath, []byte(subBSrc), 0777)
	require.NoError(t, err)

	_, err = NewParser().Parse(mainPath)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "duplicate type declaration")
}
