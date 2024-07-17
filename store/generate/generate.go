package generate

import (
	"fmt"
	"io/fs"
	"os/exec"
	"path/filepath"

	"github.com/valentinRog/sba-todo/utils"
)

func Generate() {
	storePath := filepath.Join(filepath.Join(utils.GetCurrentDirectory(), ".."))
	_ = filepath.Walk(storePath, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == "sqlc.yaml" {
			cmd := exec.Command("sqlc", "generate")
			cmd.Dir = filepath.Join(path, "..")
			res, _ := cmd.CombinedOutput()
			fmt.Println(string(res))
		}
		return nil
	})
}
