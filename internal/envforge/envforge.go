// Package envforge serves as the main API
package envforge

import (
	"github.com/tuxikus/env-forge/internal/confgen"
	"github.com/tuxikus/env-forge/internal/parser"
	"github.com/tuxikus/env-forge/internal/status"
)

var _ EnvForge = (*envForge)(nil)

type EnvForge interface {
	Forge()
}

type envForge struct {
	ef string
	p  parser.Parser
	cg confgen.ConfigGenerator
	sp status.StatusPrinter
}

func NewEnvForge(envFile string) *envForge {
	sp := status.NewStatusPrinter()

	return &envForge{
		ef: envFile,
		p:  parser.NewParser(),
		cg: confgen.NewConfigGenerator(sp),
		sp: sp,
	}
}

func (ef *envForge) Forge() {
	ef.cg.GenerateConfigs(ef.p.Parse(ef.ef))
}
