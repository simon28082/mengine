package migration

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	migrate "github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/simon28082/mengine/infrastructure/engine"
	"github.com/simon28082/mengine/infrastructure/errors"
	"github.com/simon28082/mengine/infrastructure/logger"
	os2 "github.com/simon28082/mengine/infrastructure/support/os"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"time"
)

type process struct{}

func NewProcess() *process {
	return &process{}
}

func (m *process) Name() string {
	return `migration`
}

func (m *process) Global() bool {
	return false
}

func (m *process) Dependencies() []string {
	return nil
}

func (m *process) Prepare(engine engine.Engine) error {
	return nil
}

func (m *process) Shutdown(engine engine.Engine) error {
	return nil
}

func (m *process) Cobra() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "migration",
		Example: `migration --source=/var/source/migrations.sql --host={} --username={} --password={} --database={} [up|down|stepN]`,
	}
	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		var sourceDirectory = cmd.Flag(`source`).Value.String()
		if !os2.IsDir(sourceDirectory) {
			if err := os2.MkdirDefault(sourceDirectory); err != nil {
				return err
			}
		}

		return nil
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		var (
			username   = cmd.Flag(`user`).Value.String()
			password   = cmd.Flag(`password`).Value.String()
			host       = cmd.Flag(`host`).Value.String()
			database   = cmd.Flag(`database`).Value.String()
			sourceFile = cmd.Flag(`source`).Value.String()
			//step       = cmd.Flag(`step`).Value.String()
		)

		db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?multiStatements=true", username, password, host, database))
		if err != nil {
			logger.Error(`open database failed`, `error`, err)
			return err
		}
		driver, err := mysql.WithInstance(db, &mysql.Config{})
		if err != nil {
			logger.Error(`select mysql instance failed`, `error`, err)
			return err
		}

		m, err := migrate.NewWithDatabaseInstance(
			fmt.Sprintf("file://%s", sourceFile),
			"mysql",
			driver,
		)
		if err != nil {
			logger.Error(`open migrate database failed`, `error`, err)
			return err
		}

		var action string
		if len(args) == 0 {
			return errors.WithErrorf(`migration params must required`, errors.ErrEmptyValue)
		}
		action = args[0]
		switch action {
		case `up`, `UP`:
			m.Up()
		case `down`, `DOWN`:
			m.Down()
		default:
			stepInt, err := strconv.Atoi(args[0])
			if err != nil {
				logger.Error(`covert step strint to int failed`, `error`, err)
				return err
			}
			if err := m.Steps(stepInt); err != nil {
				logger.Error(`execute migration step failed`, `error`, err)
				return err
			}
		}

		return nil
	}
	cmd.Flags().String(`host`, `localhost:3306`, `mysql host like 'localhost:3306'`)
	cmd.Flags().String(`database`, ``, `mysql database`)
	cmd.Flags().String(`user`, `root`, `mysql username`)
	cmd.Flags().String(`password`, ``, `mysql password`)
	cmd.PersistentFlags().String(`source`, ``, `migrate source file`)
	//cmd.Flags().String(`step`, `1`, `migrate step must be integer`)
	//cmd.Flags().Bool(`up`, false, `migrate up`)
	//cmd.Flags().Bool(`down`, false, `migrate rollback`)

	cmd.AddCommand(
		&cobra.Command{
			Use: `generate`,
			RunE: func(cmd *cobra.Command, args []string) error {
				var sourceDirectory = cmd.Flag(`source`).Value.String()
				if !os2.IsDir(sourceDirectory) {
					if err := os2.MkdirDefault(sourceDirectory); err != nil {
						return err
					}
				}
				var file string
				if len(args) == 0 {
					return errors.WithErrorf(`filename must required`, errors.ErrEmptyValue)
				}
				file = args[0]
				_, err := os.Create(fmt.Sprintf("%s/%s_%s", sourceDirectory, strconv.FormatInt(time.Now().Unix(), 10), file))
				if err != nil {
					return err
				}
				return nil
			},
		},
	)
	return cmd
}
