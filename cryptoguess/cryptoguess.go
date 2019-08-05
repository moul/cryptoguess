package cryptoguess // import "moul.io/cryptoguess/cryptoguess"
import (
	"fmt"
	"reflect"
	"strings"
)

type ExperimentFunc func(input []byte) Experiment

var AvailableExperiments []ExperimentFunc

type Question struct {
	Experiments []Experiment
}

func (q *Question) Short() string {
	valids := []string{}
	for _, experiment := range q.Experiments {
		if experiment.Err() == nil {
			valids = append(valids, experiment.Name())
		}
	}
	switch len(valids) {
	case 0:
		return "unknown format"
	case 1:
		return valids[0]
	default:
		return fmt.Sprintf("potential candidates: %s", strings.Join(valids, ", "))
	}
}

type Experiment interface {
	Name() string
	Run() (interface{}, error)
	Result() interface{}
	Err() error
	Confidence() float64
	String() string
}

func New(input []byte) Question {
	question := Question{
		Experiments: make([]Experiment, len(AvailableExperiments)),
	}
	for idx, fn := range AvailableExperiments {
		question.Experiments[idx] = fn(input)
		_, _ = question.Experiments[idx].Run()
	}
	return question
}

//
// base
//

type base struct {
	input  []byte
	result interface{}
	err    error
	name   string
	run    func([]byte) (interface{}, error)
}

func (exp base) Name() string {
	return exp.name
}

func (exp base) Result() interface{} {
	return exp.result
}

func (exp base) Err() error {
	return exp.err
}

func (exp base) Confidence() float64 {
	return 0.5
}

func (exp base) String() string {
	if exp.err != nil {
		return fmt.Sprintf("%s: err: %v", exp.name, exp.err)
	}
	if exp.result == nil {
		return fmt.Sprintf("%s: assert: result and err are empty", exp.name)
	}
	return fmt.Sprintf(
		"%s: %s: %v",
		exp.name,
		reflect.TypeOf(exp.result).String(),
		exp.result,
	)
}

func (exp *base) Run() (interface{}, error) {
	exp.result, exp.err = exp.run(exp.input)
	return exp.result, exp.err
}
