package main

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

var projectRoot = filepath.Join(utils.GetCurrentDirectory(), "..", "..")

func main() {
	idCount := 1
	distPath := filepath.Join(projectRoot, "ui", "dist")
	_ = os.RemoveAll(distPath)
	_ = os.Mkdir(distPath, os.ModePerm)
	file, err := os.OpenFile(filepath.Join(distPath, "style.css"), os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	err = filepath.Walk(filepath.Join(projectRoot, "ui", "templates"), func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == "style.css" {
			idString := fmt.Sprintf("id-%d", idCount)
			idCount += 1

			f, _ := os.OpenFile(filepath.Join(filepath.Dir(path), "id.go"), os.O_CREATE|os.O_TRUNC, 0644)
			defer f.Close()
			_, _ = f.WriteString(fmt.Sprintf("package %s\n\n", filepath.Base(filepath.Dir(path))))
			_, _ = f.WriteString(fmt.Sprintf("const id = \"%s\"\n", idString))

			cmd := exec.Command("npx", "postcss", path)
			cmd.Env = append(os.Environ(), fmt.Sprintf("PREFIX=%s", idString))
			cmd.Dir = filepath.Join(projectRoot, "ui")
			output, err := cmd.Output()
			fmt.Println(path)
			fmt.Println(string(output))
			if err != nil {
				fmt.Println("yo la team")
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
