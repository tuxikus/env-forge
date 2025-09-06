// Package parser parses the input file into an env
// File structure:
// # files
// ./path/to/source/config.conf ~/path/to/destination/config.conf
// # directories
// /mnt/dir/must/exist
// # packages
// bash
// emacs
// go

package parser

import (
	"bufio"
	"os"
	"strings"

	"github.com/tuxikus/env-forge/internal/env"
)

const (
	LinksSectionHeader = "# links"
	dirsSectionHeader  = "# directories"
	pkgsSectionHeader  = "# packages"
)

const (
	linksSection = iota
	dirsSection
	pkgsSection
)

type Parser struct{}

func (p *Parser) Parse(path string) (*env.Env, error) {
	section := -1
	links := make([]env.Link, 0)
	dirs := make([]env.Dir, 0)
	pkgs := make([]env.Pkg, 0)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}
		switch line {
		case LinksSectionHeader:
			section = linksSection
			continue
		case dirsSectionHeader:
			section = dirsSection
			continue
		case pkgsSectionHeader:
			section = pkgsSection
			continue
		}

		switch section {
		case linksSection:
			src, dst, _ := parseLinks(line)
			links = append(links, env.Link{Src: src, Dst: dst})
		case dirsSection:
		case pkgsSection:
		}
	}

	return &env.Env{
		Links: links,
		Dirs:  dirs,
		Pkgs:  pkgs,
	}, nil
}

func replaceHome(input string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return strings.Replace(input, "~", home, 1)
}

func parseLinks(line string) (string, string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	src := strings.Split(line, " ")[0]
	src = pwd + "/" + src
	dst := replaceHome(strings.Split(line, " ")[1]) // no ~ in links
	return src, dst, nil
}

func getConf(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(content)
}
