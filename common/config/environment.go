// https://gist.github.com/hatappo/1b06eb02168e835ada591c4551143f0b
package config

import (
	"os"
	"strings"
)

type Environment int

var Env Environment

type EnvironmentAttr struct {
	V     Environment
	Name  string
	Label string
}

const (
	Local Environment = iota
	Production
)

const cnt = 2

var environments = [cnt]EnvironmentAttr{
	{Local, "local", "ローカル環境"},
	{Production, "production", ""},
}

func (e Environment) Name() string {
	return e.get().Name
}

func (e Environment) Label() string {
	return e.get().Label
}

func (e Environment) String() string {
	return e.Label()
}

func (e Environment) IsLocal() bool {
	return e == Local
}

func EnvironmentValueOf(name string) Environment {
	for _, environment := range environments {
		if strings.EqualFold(name, environment.Name) {
			return environment.V
		}
	}
	return Local
}

func (e Environment) get() EnvironmentAttr {
	if e < 0 || e >= cnt {
		panic("Illegal enum value.")
	}
	return environments[e]
}

func init() {
	Env = EnvironmentValueOf(os.Getenv("ENV"))
}
