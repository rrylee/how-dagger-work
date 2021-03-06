package main

import (
	"context"
	"os"

	"deep_into_dagger"

	"cuelang.org/go/cue"
	"cuelang.org/go/tools/flow"
)

var r cue.Runtime // nolint

func main() {
	if len(os.Args) != 3 {
		panic("Usage: ./main.go *.cue isTask")
	}

	file := os.Args[1]
	flag := os.Args[2]

	b, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	inst, err := r.Parse(file, b) // nolint
	if err != nil {
		panic(err)
	}

	flow := flow.New(
		&flow.Config{},
		inst.Value(),
		func(flowVal cue.Value) (flow.Runner, error) {
			if !flowVal.LookupPath(cue.ParsePath(flag)).Exists() {
				return nil, nil
			}

			return flow.RunnerFunc(func(t *flow.Task) error {
				return nil
			}), nil
		},
	)

	deep_into_dagger.PrintTasks(flow.Tasks(), 0)

	err = flow.Run(context.TODO())
	if err != nil {
		panic(err)
	}
}
