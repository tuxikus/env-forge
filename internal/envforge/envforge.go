// Package envforge serves as the main API
package envforge

import (
	"fmt"
	"os"
	"github.com/tuxikus/env-forge/internal/parser"
)

type EnvForge struct {
	ef string
	p  parser.Parser
}

func NewEnvForge(envFile string) *EnvForge {
	return &EnvForge{
		ef: envFile,
		p:  parser.Parser{},
	}
}

func (ef *EnvForge) Forge() {
	env, err := ef.p.Parse(ef.ef)
	if err != nil {
		panic(err)
	}
	
	for _, link := range env.Links {
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
