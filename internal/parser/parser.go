// Package parser parses the input config file into a internal structure
package parser

import (
	"bufio"
	"github.com/tuxikus/env-forge/internal/env"
	"os"
	"strings"
)

var _ Parser = (*parser)(nil)

type Parser interface {
	Parse(string) []env.Env
}

type parser struct{}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) Parse(path string) []env.Env {
	envs := make([]env.Env, 0)
	lines := make([]string, 0)

	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	f.Close()

	for _, line := range lines {
		src := strings.Split(line, " ")[0]
		dst := strings.Split(line, " ")[1]
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		dst = strings.Replace(dst, "~", home, -1)
		conf := getConf(src)

		envs = append(envs, env.Env{
			Src:  src,
			Dst:  dst,
			Conf: conf,
		})
	}

	return envs
}

func getConf(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(content)
}
