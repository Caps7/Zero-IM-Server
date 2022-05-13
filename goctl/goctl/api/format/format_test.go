package format

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	notFormattedStr = `
type Request struct {
  Name string ` + "`" + `path:"name,options=you|me"` + "`" + `
}
type Response struct {
  Message string ` + "`" + `json:"message"` + "`" + `
  Students []Student ` + "`" + `json:"students"` + "`" + `
}
service A-api {
@server(
handler: GreetHandler
  )
  get /greet/from/:name(Request) returns (Response)
}
`

	formattedStr = `type Request {
	Name string ` + "`" + `path:"name,options=you|me"` + "`" + `
}
type Response {
	Message  string    ` + "`" + `json:"message"` + "`" + `
	Students []Student ` + "`" + `json:"students"` + "`" + `
}
service A-api {
	@server(
		handler: GreetHandler
	)
	get /greet/from/:name(Request) returns (Response)
}`
)

func TestFormat(t *testing.T) {
	r, err := apiFormat(notFormattedStr, true)
	assert.Nil(t, err)
	assert.Equal(t, formattedStr, r)
	_, err = apiFormat(notFormattedStr, false)
	assert.Errorf(t, err, " line 7:13 can not found declaration 'Student' in context")
}

func Test_apiFormatReader_issue1721(t *testing.T) {
	dir, err := os.MkdirTemp("", "goctl-api-format")
	require.NoError(t, err)
	defer os.RemoveAll(dir)
	subDir := path.Join(dir, "sub")
	err = os.MkdirAll(subDir, fs.ModePerm)
	require.NoError(t, err)

	importedFilename := path.Join(dir, "foo.api")
	err = ioutil.WriteFile(importedFilename, []byte{}, fs.ModePerm)
	require.NoError(t, err)

	filename := path.Join(subDir, "bar.api")
	err = ioutil.WriteFile(filename, []byte(fmt.Sprintf(`import "%s"`, importedFilename)), 0644)
	require.NoError(t, err)

	f, err := os.Open(filename)
	require.NoError(t, err)

	err = apiFormatReader(f, filename, false)
	assert.NoError(t, err)
}
