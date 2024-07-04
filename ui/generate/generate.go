package generate

import (
	"bufio"
	"bytes"
	"fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"github.com/valentinRog/sba-todo/utils"
)

var uiPath = filepath.Join(utils.GetCurrentDirectory(), "..")

type Cmd struct {
	cmd    *exec.Cmd
	stdout bytes.Buffer
	stderr bytes.Buffer
}

func NewCmd(cmd *exec.Cmd) *Cmd {
	c := &Cmd{cmd: cmd}
	c.cmd.Stdout = &c.stdout
	c.cmd.Stderr = &c.stderr
	return c
}

func (c *Cmd) Start() error {
	return c.cmd.Start()
}

func (c *Cmd) Wait() string {
	err := c.cmd.Wait()
	if err != nil {
		log.Fatal(c.stderr.String())
	}
	return c.stdout.String()
}

func Generate() {
	idCount := 1

	cmds := []*Cmd{}

	postcss := filepath.Join(uiPath, "node_modules", ".bin", "postcss")

	cmds = append(cmds, NewCmd(exec.Command(postcss, filepath.Join(uiPath, "style", "global", "style.css"))))

	err := filepath.Walk(filepath.Join(uiPath, "templates"), func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() && info.Name() == "style.css" {
			idString := fmt.Sprintf("id-%d", idCount)
			idCount += 1

			f, _ := os.Create(filepath.Join(filepath.Dir(path), "generated_id.go"))
			defer f.Close()
			_, _ = f.WriteString(fmt.Sprintf("package %s\n\n", filepath.Base(filepath.Dir(path))))
			_, _ = f.WriteString(fmt.Sprintf("const id = \"%s\"\n", idString))

			cmd := exec.Command(postcss, path)
			cmd.Env = append(os.Environ(), fmt.Sprintf("PREFIX=%s", idString))
			cmds = append(cmds, NewCmd(cmd))
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	for _, cmd := range cmds {
		cmd.Start()
	}

	distPath := filepath.Join(uiPath, "dist")
	_ = os.RemoveAll(distPath)
	_ = os.Mkdir(distPath, os.ModePerm)
	file, err := os.Create(filepath.Join(distPath, "style.css"))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for _, cmd := range cmds {
		fmt.Println(cmd.cmd.Args)
		out := cmd.Wait()
		fmt.Println(out)
		writer.WriteString(out)
	}

	err = writer.Flush()
	if err != nil {
		log.Fatal(err)
	}
}
