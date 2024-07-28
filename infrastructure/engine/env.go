package engine

import (
	"context"
	"github.com/simon28082/mengine/infrastructure/errors"
	cliv3 "github.com/urfave/cli/v3"
)

var envs = make(map[string]*Env)

type Env struct {
	Name  string
	Setup func(ctx context.Context, cli *cliv3.Command) error
}

// RegisterEnv a env profile
func RegisterEnv(p *Env) error {
	if _, ok := envs[p.Name]; ok {
		return errors.Errorf(`the env [%s] already exists`, p.Name)
	}

	envs[p.Name] = p
	return nil
}

func LoadEnv(name string) (*Env, error) {
	if v, ok := envs[name]; ok {
		return v, nil
	}

	return nil, errors.Errorf(`the env [%s] not found`, name)
}

func LoadEnvFromCli(ctx context.Context, cli *cliv3.Command) error {
	runningEnv := cli.String(`env`)
	if len(runningEnv) == 0 {
		return nil
	}

	env, err := LoadEnv(runningEnv)
	if err == nil {
		return env.Setup(ctx, cli)
	}

	return nil
}
