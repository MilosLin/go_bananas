package task

import (
	"flag"
	"fmt"
	"strings"

	"github.com/MilosLin/go_bananas/core/module/validator"
)

type ParseFlagStr struct {
	P1 string `validate:"omitempty"`
	P2 int    `validate:"omitempty"`
}

func NewParseFlagStr() *ParseFlagStr {
	return &ParseFlagStr{}
}

// run task ParseFlagStr
//
// task --name=parseflagstr --argu="-p1=abc -p2=150"
func (instance *ParseFlagStr) Run(argu *string) error {
	argus := flag.NewFlagSet("ask", flag.ExitOnError)
	argus.StringVar(&instance.P1, "p1", "default", "paramter 1")
	argus.IntVar(&instance.P2, "p2", 5, "paramter 2")
	argus.Parse(strings.Fields(*argu))

	if err := validator.Struct(instance); err != nil {
		return err
	}

	fmt.Printf("%+v\n", instance)
	return nil
}
