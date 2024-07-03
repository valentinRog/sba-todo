package generate

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/valentinRog/sba-todo/utils"
)

var uiPath = filepath.Join(utils.GetCurrentDirectory(), "..")

func Generate() {
	idCount := 1
	distPath := filepath.Join(uiPath, "dist")
	_ = os.RemoveAll(distPath)
	_ = os.Mkdir(distPath, os.ModePerm)
	file, err := os.Create(filepath.Join(distPath, "style.css"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	cmd := exec.Command("npx", "postcss", filepath.Join(uiPath, "style", "global", "style.css"))
	cmd.Dir = uiPath
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		log.Fatal(err.Error())
	}
	writer.Write(output)

	err = filepath.Walk(filepath.Join(uiPath, "templates"), func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == "style.css" {
			idString := fmt.Sprintf("id-%d", idCount)
			idCount += 1

			f, _ := os.Create(filepath.Join(filepath.Dir(path), "generated_id.go"))
			defer f.Close()
			_, _ = f.WriteString(fmt.Sprintf("package %s\n\n", filepath.Base(filepath.Dir(path))))
			_, _ = f.WriteString(fmt.Sprintf("const id = \"%s\"\n", idString))

			cmd := exec.Command("npx", "postcss", path)
			cmd.Env = append(os.Environ(), fmt.Sprintf("PREFIX=%s", idString))
			fmt.Println(idString)
			cmd.Dir = uiPath
			output, err := cmd.CombinedOutput()
			fmt.Println(string(output))
			if err != nil {
				log.Fatal(err.Error())
			}
			writer.Write(output)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
