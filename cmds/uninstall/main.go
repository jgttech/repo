package purge

import (
	"context"
	"os"
	"slices"

	"github.com/jgttech/repo/src/cfg"
	"github.com/jgttech/repo/src/fs"
	"github.com/jgttech/repo/src/self"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	return &cli.Command{
		Name:  "uninstall",
		Usage: "Unlinks the CLI config from the system but does NOT remove the CLI, use 'purge' for that.",
		Action: func(ctx context.Context, c *cli.Command) (err error) {
			deletedBackups := []string{}
			update := []*cfg.StateBackup{}

			for _, backup := range self.State.Backups {
				_, err = os.Stat(backup.To)

				if err == nil {
					err = os.Remove(backup.From)
					err = fs.CopyFile(backup.To, backup.From)
					err = os.Remove(backup.To)
					deletedBackups = append(deletedBackups, backup.To)
				} else if os.IsNotExist(err) {
					deletedBackups = append(deletedBackups, backup.To)
				}
			}

			for _, backup := range self.State.Backups {
				if !slices.Contains(deletedBackups, backup.To) {
					update = append(update, backup)
				}
			}

			self.State.Backups = update
			self.State.Save()

			return nil
		},
	}
}
