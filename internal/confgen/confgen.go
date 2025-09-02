// Package confgen generates the final config files
package confgen

import (
	"github.com/tuxikus/env-forge/internal/env"
	"github.com/tuxikus/env-forge/internal/status"
	"os"
	"path/filepath"
)

var _ ConfigGenerator = (*configGenerator)(nil)

type ConfigGenerator interface {
	GenerateConfigs([]env.Env)
}

type configGenerator struct {
	sp status.StatusPrinter
}

func NewConfigGenerator(sp status.StatusPrinter) *configGenerator {
	return &configGenerator{
		sp: sp,
	}
}

func (cg *configGenerator) GenerateConfigs(envs []env.Env) {
	for _, env := range envs {
		dst := env.Dst
		dstDir := filepath.Dir(dst)

		if _, err := os.Stat(dstDir); os.IsNotExist(err) {
			err := os.MkdirAll(dstDir, 0755)
			if err != nil {
				panic(err)
			}
		}

		f, err := os.Create(env.Dst)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		cg.sp.Print(status.StatusWritingFile, env.Dst)

		_, err = f.Write([]byte(env.Conf))
		if err != nil {
			panic(err)
		}

		cg.sp.Print(status.StatusDoneWritingFile, env.Dst)
	}
}
