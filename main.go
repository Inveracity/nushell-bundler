package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"

	"github.com/amenzhinsky/go-memexec"
)

//go:embed nu
var mybinary []byte

//go:embed all:src
var srcEmbed embed.FS

func main() {
	// Create a temporary directory to write the embedded nushell code to
	source := "src"
	tmpdir, err := os.MkdirTemp("", "gonu*")
	if err != nil {
		panic(err)
	}

	// Write the nushell code to /tmp/gonu*/src
	err = render(tmpdir, source)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Execute the code
	out, err := run(tmpdir, source)
	if err != nil {
		fmt.Println(err.Error())
	}

	// Output results
	fmt.Printf("%+v", string(out))

	// clean up rendered code
	os.RemoveAll(tmpdir)
}

// Run the nushell code
func run(tmpdir string, source string) (string, error) {
	exe, err := memexec.New(mybinary)
	if err != nil {
		return "", err
	}

	defer exe.Close()

	argsWithoutProg := os.Args[1:]
	args := []string{"-n", path.Join(tmpdir, source, "main.nu")}
	args = append(args, argsWithoutProg...)
	cmd := exe.Command(args...)

	out, err := cmd.CombinedOutput()
	return string(out), err
}

// Write the embedded nushell scripts to disk
func render(tmpdir string, src string) error {
	if err := fs.WalkDir(srcEmbed, src, func(file string, d fs.DirEntry, err error) error {

		if d.IsDir() {
			newdir := path.Join(tmpdir, file)
			err := os.MkdirAll(newdir, 0777)
			if err != nil {
				return err
			}
		} else {
			content, err := fs.ReadFile(srcEmbed, file)
			if err != nil {
				return err
			}
			newfile := path.Join(tmpdir, file)
			if err := os.WriteFile(newfile, content, 0666); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}
