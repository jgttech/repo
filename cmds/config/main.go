package config

import (
	"context"
	"fmt"
	"strconv"

	"github.com/jgttech/repo/src/self"
	"github.com/urfave/cli/v3"
)

func Command() *cli.Command {
	dir := self.Context.Pwd

	return &cli.Command{
		Name: "config",
		Action: func(ctx context.Context, c *cli.Command) error {
			key := c.Args().First()

			for _, value := range self.State.Repositories {
				if value.Name == dir.Name {
					data := map[string]string{
						"created_at": strconv.Itoa(int(value.CreatedAt)),
						"updated_at": strconv.Itoa(int(value.UpdatedAt)),
						"dir":        value.Dir,
						"name":       value.Name,
						"alias":      value.Alias,
						"branch":     value.Branch,
						"primary":    value.Primary,
						"ssh":        value.Ssh,
					}

					if key != "" {
						value, ok := data[key]

						if !ok {
							fmt.Printf("No key found called '%s'.", value)
						}

						fmt.Println(value)
						return nil
					}

					if data["dir"] == "" {
						data["dir"] = "-"
					}

					if data["name"] == "" {
						data["name"] = "-"
					}

					if data["alias"] == "" {
						data["alias"] = "-"
					}

					if data["branch"] == "" {
						data["branch"] = "-"
					}

					if data["primary"] == "" {
						data["primary"] = "-"
					}

					if data["ssh"] == "" {
						data["ssh"] = "-"
					}

					fmt.Println()
					fmt.Printf("[REPO CONFIG]\n")
					fmt.Println("created_at.:", data["created_at"])
					fmt.Println("updated_at.:", data["updated_at"])
					fmt.Println("dir........:", data["dir"])
					fmt.Println("name.......:", data["name"])
					fmt.Println("primary....:", data["primary"])
					fmt.Println("alias......:", data["alias"])
					fmt.Println("branch.....:", data["branch"])
					fmt.Println("ssh........:", data["ssh"])
					return nil
				}
			}

			fmt.Printf("Repo not found.\n\n")
			fmt.Println("To add your repo to the CLI, run:")
			fmt.Println("repo add")

			return nil
		},
	}
}
