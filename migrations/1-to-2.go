package migrations

import (
	"fmt"
	"os"

	"github.com/filecoin-project/go-filecoin/config"
	"github.com/filecoin-project/go-filecoin/repo"
	"github.com/ipfs/go-ipfs-cmds"
	"github.com/pkg/errors"
	logging "github.com/ipfs/go-log"
)

const (
	Description      = "a test migrator that just updates the repo version"
	MigrationVersion = "0.2"
	PreviousVersion  = "0.1"
)

var (
	ErrMigrationFailed = errors.New("migrator failed")
)

//  Migration runner defines an interface which migrator code must satisfy.
//  Migrations are a pure function, given access to the input (read-only) and
//  output repos, or a read-write repo to be migrated in place.

type migrator struct {
	log     logging.EventLogger
	oldRepo os.File // read only
	newRepo os.File // read-write
	verbose bool
}

// 	GetNewRepoVersion() string

// fs access and we should know the location since we know the previous version.
func getOldFSRepo(req *cmds.Request) (repo.FSRepo, error) {
	return repo.FSRepo{}, nil
}

// fs access in the old repo
func getCurrentRepoVersion(req *cmds.Request) string {
	return ""
}

// also just fs access
func getOldConfig() *config.Config {}

// NewMigrator instantiates a new migrator
func NewMigrator(oldRepo, newRepo os.File, verbose bool) *migrator {
	logstr := fmt.Sprintf("Migration from %s to %s", PreviousVersion, MigrationVersion)
	return &migrator{
		log:          logging.Logger(logstr),
		verbose:      verbose,
	}
}

func (mig *migrator) Update() error {
	switch mig.req.Arguments[0] {
	case "Describe":
		mig.Describe()
		return nil
	case "DryRun":
		return mig.DryRun()
	case "Run":
		return mig.Migrate()
	}
}

// Describe emits a description of what this migrator will do.
// Verbose option is ignored; output is not logged.
func (mig *migrator) Describe() {
	mig.log.Info(Description)
	// use the emitter to output description
}

// Run runs the migrator steps on a copy of the repo
func (mig *migrator) Migrate() error {
	if err := mig.DryRun(); err != nil {
		return err
	}
	// return linkNewRepo()
}

// DryRun runs the migrator steps on a copy of the repo and stops there
func (mig *migrator) DryRun() error {
	// makeNewFSRepo()
	// copyData()
	// migrateStep1
	// migrateStep2
	// migrateStep3
}

// Validate returns error if migration tests failed, describing why
func (mig *migrator) Validate() error {
	return nil
}