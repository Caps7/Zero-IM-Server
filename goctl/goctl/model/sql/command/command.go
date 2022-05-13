package command

import (
	"errors"
	"path/filepath"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/postgres"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/tools/goctl/config"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/command/migrationnotes"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/gen"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/model"
	"github.com/zeromicro/go-zero/tools/goctl/model/sql/util"
	file "github.com/zeromicro/go-zero/tools/goctl/util"
	"github.com/zeromicro/go-zero/tools/goctl/util/console"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

var (
	// VarStringSrc describes the source file of sql.
	VarStringSrc string
	// VarStringDir describes the output directory of sql.
	VarStringDir string
	// VarBoolCache describes whether the cache is enabled.
	VarBoolCache bool
	// VarBoolIdea describes whether is idea or not.
	VarBoolIdea bool
	// VarStringURL describes the dsn of the sql.
	VarStringURL string
	// VarStringSliceTable describes tables.
	VarStringSliceTable []string
	// VarStringTable describes a table of sql.
	VarStringTable string
	// VarStringStyle describes the style.
	VarStringStyle string
	// VarStringDatabase describes the database.
	VarStringDatabase string
	// VarStringSchema describes the schema of postgresql.
	VarStringSchema string
	// VarStringHome describes the goctl home.
	VarStringHome string
	// VarStringRemote describes the remote git repository.
	VarStringRemote string
	// VarStringBranch describes the git branch of the repository.
	VarStringBranch string
)

var errNotMatched = errors.New("sql not matched")

// MysqlDDL generates model code from ddl
func MysqlDDL(_ *cobra.Command, _ []string) error {
	migrationnotes.BeforeCommands(VarStringDir, VarStringStyle)
	src := VarStringSrc
	dir := VarStringDir
	cache := VarBoolCache
	idea := VarBoolIdea
	style := VarStringStyle
	database := VarStringDatabase
	home := VarStringHome
	remote := VarStringRemote
	branch := VarStringBranch
	if len(remote) > 0 {
		repo, _ := file.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}
	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}
	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}

	return fromDDL(src, dir, cfg, cache, idea, database)
}

// MySqlDataSource generates model code from datasource
func MySqlDataSource(_ *cobra.Command, _ []string) error {
	migrationnotes.BeforeCommands(VarStringDir, VarStringStyle)
	url := strings.TrimSpace(VarStringURL)
	dir := strings.TrimSpace(VarStringDir)
	cache := VarBoolCache
	idea := VarBoolIdea
	style := VarStringStyle
	home := VarStringHome
	remote := VarStringRemote
	branch := VarStringBranch
	if len(remote) > 0 {
		repo, _ := file.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}
	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}

	tableValue := VarStringSliceTable
	patterns := parseTableList(tableValue)
	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}

	return fromMysqlDataSource(url, dir, patterns, cfg, cache, idea)
}

type pattern map[string]struct{}

func (p pattern) Match(s string) bool {
	for v := range p {
		match, err := filepath.Match(v, s)
		if err != nil {
			console.Error("%+v", err)
			continue
		}
		if match {
			return true
		}
	}
	return false
}

func (p pattern) list() []string {
	var ret []string
	for v := range p {
		ret = append(ret, v)
	}
	return ret
}

func parseTableList(tableValue []string) pattern {
	tablePattern := make(pattern)
	for _, v := range tableValue {
		fields := strings.FieldsFunc(v, func(r rune) bool {
			return r == ','
		})
		for _, f := range fields {
			tablePattern[f] = struct{}{}
		}
	}
	return tablePattern
}

// PostgreSqlDataSource generates model code from datasource
func PostgreSqlDataSource(_ *cobra.Command, _ []string) error {
	migrationnotes.BeforeCommands(VarStringDir, VarStringStyle)
	url := strings.TrimSpace(VarStringURL)
	dir := strings.TrimSpace(VarStringDir)
	cache := VarBoolCache
	idea := VarBoolIdea
	style := VarStringStyle
	schema := VarStringSchema
	home := VarStringHome
	remote := VarStringRemote
	branch := VarStringBranch
	if len(remote) > 0 {
		repo, _ := file.CloneIntoGitHome(remote, branch)
		if len(repo) > 0 {
			home = repo
		}
	}
	if len(home) > 0 {
		pathx.RegisterGoctlHome(home)
	}

	if len(schema) == 0 {
		schema = "public"
	}

	pattern := strings.TrimSpace(VarStringTable)
	cfg, err := config.NewConfig(style)
	if err != nil {
		return err
	}

	return fromPostgreSqlDataSource(url, pattern, dir, schema, cfg, cache, idea)
}

func fromDDL(src, dir string, cfg *config.Config, cache, idea bool, database string) error {
	log := console.NewConsole(idea)
	src = strings.TrimSpace(src)
	if len(src) == 0 {
		return errors.New("expected path or path globbing patterns, but nothing found")
	}

	files, err := util.MatchFiles(src)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return errNotMatched
	}

	generator, err := gen.NewDefaultGenerator(dir, cfg, gen.WithConsoleOption(log))
	if err != nil {
		return err
	}

	for _, file := range files {
		err = generator.StartFromDDL(file, cache, database)
		if err != nil {
			return err
		}
	}

	return nil
}

func fromMysqlDataSource(url, dir string, tablePat pattern, cfg *config.Config, cache, idea bool) error {
	log := console.NewConsole(idea)
	if len(url) == 0 {
		log.Error("%v", "expected data source of mysql, but nothing found")
		return nil
	}

	if len(tablePat) == 0 {
		log.Error("%v", "expected table or table globbing patterns, but nothing found")
		return nil
	}

	dsn, err := mysql.ParseDSN(url)
	if err != nil {
		return err
	}

	logx.Disable()
	databaseSource := strings.TrimSuffix(url, "/"+dsn.DBName) + "/information_schema"
	db := sqlx.NewMysql(databaseSource)
	im := model.NewInformationSchemaModel(db)

	tables, err := im.GetAllTables(dsn.DBName)
	if err != nil {
		return err
	}

	matchTables := make(map[string]*model.Table)
	for _, item := range tables {
		if !tablePat.Match(item) {
			continue
		}

		columnData, err := im.FindColumns(dsn.DBName, item)
		if err != nil {
			return err
		}

		table, err := columnData.Convert()
		if err != nil {
			return err
		}

		matchTables[item] = table
	}

	if len(matchTables) == 0 {
		return errors.New("no tables matched")
	}

	generator, err := gen.NewDefaultGenerator(dir, cfg, gen.WithConsoleOption(log))
	if err != nil {
		return err
	}

	return generator.StartFromInformationSchema(matchTables, cache)
}

func fromPostgreSqlDataSource(url, pattern, dir, schema string, cfg *config.Config, cache, idea bool) error {
	log := console.NewConsole(idea)
	if len(url) == 0 {
		log.Error("%v", "expected data source of postgresql, but nothing found")
		return nil
	}

	if len(pattern) == 0 {
		log.Error("%v", "expected table or table globbing patterns, but nothing found")
		return nil
	}
	db := postgres.New(url)
	im := model.NewPostgreSqlModel(db)

	tables, err := im.GetAllTables(schema)
	if err != nil {
		return err
	}

	matchTables := make(map[string]*model.Table)
	for _, item := range tables {
		match, err := filepath.Match(pattern, item)
		if err != nil {
			return err
		}

		if !match {
			continue
		}

		columnData, err := im.FindColumns(schema, item)
		if err != nil {
			return err
		}

		table, err := columnData.Convert()
		if err != nil {
			return err
		}

		matchTables[item] = table
	}

	if len(matchTables) == 0 {
		return errors.New("no tables matched")
	}

	generator, err := gen.NewDefaultGenerator(dir, cfg, gen.WithConsoleOption(log), gen.WithPostgreSql())
	if err != nil {
		return err
	}

	return generator.StartFromInformationSchema(matchTables, cache)
}
