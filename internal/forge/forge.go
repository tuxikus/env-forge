// Package forge serves as the main API
package forge

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"

	"github.com/tuxikus/env-forge/internal/env"
	"github.com/tuxikus/env-forge/internal/parser"
)

type Forge struct {
	ef string
	p  parser.Parser
}

func NewForge(envFile string) *Forge {
	return &Forge{
		ef: envFile,
		p:  parser.Parser{},
	}
}

func (ef *Forge) Forge() {
	env, err := ef.p.Parse(ef.ef)
	if err != nil {
		panic(err)
	}

	forgeLinks(env.Links)
	checkDirs(env.Dirs)
	checkPkgs(env.Pkgs)
}

func forgeLinks(links []env.Link) {
	fmt.Println("Forging links...")

	for _, link := range links {
		err := os.Remove(link.Dst)
		if err != nil {
			panic(err)
		}

		err = os.Symlink(link.Src, link.Dst)
		if err != nil {
			switch err.(type) {
			case *os.LinkError:
				fmt.Println(link.Dst, "already exists")
			default:
				panic(err)
			}
		}
	}
}

func checkDirs(dirs []env.Dir) {
	fmt.Println("Checking directories...")

	for _, dir := range dirs {
		_, err := os.Stat(dir.Path)
		if err == nil {
			fmt.Println(dir.Path, "exists")
			continue
		}
		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println(dir.Path, "not exists")
			continue
		}

		panic(err)
	}

}

func removeTailingNewline(input string) string {
	if input[len(input)-1] == '\n' {
		return input[:len(input)-1]
	}
	return input
}

func checkPkgs(pkgs []env.Pkg) {
	fmt.Println("Checking packages...")
	var whichCmd *exec.Cmd
	for _, pkg := range pkgs {
		whichCmd = exec.Command("which", pkg.Name)
		whichOut, err := whichCmd.Output()
		if err != nil {
			panic(err)
		}

		whichOutStr := string(whichOut)
		whichOutStr = removeTailingNewline(whichOutStr)
		fmt.Println(whichOutStr)
	}
}
