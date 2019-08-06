package cryptoguess // import "moul.io/cryptoguess/cryptoguess"
import (
	"fmt"
	"reflect"
	"strings"
)

//
// Question
//

type Question struct {
	Experiments []Experiment
}

func (q *Question) Short() string {
	parts := []string{}
	for _, experiment := range q.Experiments {
		for _, result := range experiment.Results() {
			if part := result.Short(); part != "" {
				parts = append(parts, part)
			}
		}
	}
	switch len(parts) {
	case 0:
		return "unknown format"
	case 1:
		return parts[0]
	default:
		return fmt.Sprintf("potential candidates: %s", strings.Join(parts, ", "))
	}
}

func New(input []byte) Question {
	question := Question{
		Experiments: make([]Experiment, len(AvailableExperiments)),
	}
	for idx, fn := range AvailableExperiments {
		exp := fn(input)
		question.Experiments[idx] = exp
		question.Experiments[idx].Run()
	}
	return question
}

//
// Experiment
//

type Experiment interface {
	Name() string
	Input() []byte
	Results() []Result
	Short() string
	String() string
	Run()
}

type ExperimentFunc func(input []byte) Experiment

var AvailableExperiments []ExperimentFunc

func (exp baseExperiment) Short() string {
	parts := []string{}
	for _, result := range exp.Results() {
		if result.Err() == nil {
			parts = append(parts, result.Name())
		}
	}
	return strings.Join(parts, ", ")
}

func (exp baseExperiment) Name() string {
	return exp.name
}

func (exp baseExperiment) Input() []byte {
	return exp.input
}

func (exp baseExperiment) Results() []Result {
	return exp.results
}

func (exp *baseExperiment) Run() {
	exp.results = exp.run(exp)
}

func (exp baseExperiment) String() string {
	output := ""
	for _, result := range exp.Results() {
		output += fmt.Sprintf("- %s\n", result)
	}
	return output
}

type baseExperiment struct {
	input   []byte
	name    string
	run     func(Experiment) []Result // this function takes a new Experiment as argument to allow usage of more complex golang inheritance & composition patterns
	results []Result
}

//
// Result
//

type Result interface {
	Name() string
	Data() interface{}
	Err() error
	Confidence() float64
	String() string
	Short() string
	Rest() []byte
}

type baseResult struct {
	exp        Experiment
	name       string
	data       interface{}
	err        error
	rest       []byte
	confidence float64
}

func (res baseResult) Name() string {
	if res.name == "" {
		return res.exp.Name()
	}
	return fmt.Sprintf("%s: %s", res.exp.Name(), res.name)
}

func (res baseResult) Short() string {
	if res.Err() == nil {
		return res.Name()
	}
	return ""
}

func (res baseResult) Data() interface{} {
	return res.data
}

func (res baseResult) Err() error {
	return res.err
}

func (res baseResult) Confidence() float64 {
	return res.confidence
}

func (res baseResult) String() string {
	if res.err != nil {
		return fmt.Sprintf("%s: err: %v", res.Name(), res.err)
	}
	if res.data == nil {
		return fmt.Sprintf("%s: assert: result and err are empty", res.Name())
	}
	return fmt.Sprintf(
		"%s: %s: %v",
		res.Name(),
		reflect.TypeOf(res.data).String(),
		res.data,
	)
}

func (res *baseResult) Rest() []byte {
	return res.rest
}
