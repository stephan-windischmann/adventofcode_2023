package day20

import (
	"strings"

	"github.com/jdavid5815/extmath"
)

type moduleType int

const (
	BROADCAST moduleType = iota
	FLIPLOP
	CONJUNCTION
	OUTPUT
	BUTTON
)

type input struct {
	Module string
	Input  int
}

type module struct {
	Type          moduleType
	Outputs       []string
	Inputs        []input
	FlipflopState bool
}

type pulse struct {
	Start       string
	Destination string
	Value       int
}

func (m *module) allInputsHigh() bool {
	for _, i := range m.Inputs {
		if i.Input == 0 {
			return false
		}
	}

	return true
}

func (m *module) ProcessInput(p pulse) []pulse {
	outputs := make([]pulse, 0)

	switch m.Type {
	case BROADCAST:
		for _, o := range m.Outputs {
			outputPulse := pulse{
				Start:       p.Destination,
				Destination: o,
				Value:       p.Value,
			}
			outputs = append(outputs, outputPulse)
		}
	case FLIPLOP:
		if p.Value == 0 {
			m.FlipflopState = !m.FlipflopState
			var outputVal int
			if m.FlipflopState {
				outputVal = 1
			} else {
				outputVal = 0
			}
			for _, o := range m.Outputs {
				outputPulse := pulse{
					Start:       p.Destination,
					Destination: o,
					Value:       outputVal,
				}
				outputs = append(outputs, outputPulse)
			}
		}
	case CONJUNCTION:
		for i := 0; i < len(m.Inputs); i++ {
			if m.Inputs[i].Module == p.Start {
				m.Inputs[i].Input = p.Value
			}
		}

		outputVal := 1
		if m.allInputsHigh() {
			outputVal = 0
		}
		for _, o := range m.Outputs {
			outputPulse := pulse{
				Start:       p.Destination,
				Destination: o,
				Value:       outputVal,
			}
			outputs = append(outputs, outputPulse)
		}
	}

	return outputs
}

func parseModules(inputTxt []string) map[string]module {
	modules := make(map[string]module)
	for _, l := range inputTxt {
		module := module{}

		name := strings.Split(l, " -> ")[0]
		connections := strings.Fields(strings.Replace(strings.Split(l, " -> ")[1], ",", "", -1))
		if name[0] == '%' {
			module.Type = FLIPLOP
			name = name[1:]
		} else if name[0] == '&' {
			module.Type = CONJUNCTION
			name = name[1:]
		} else {
			module.Type = BROADCAST
		}
		module.Outputs = connections
		module.Inputs = make([]input, 0)

		modules[name] = module
	}

	mapKeys := make([]string, len(modules))
	i := 0
	for k := range modules {
		mapKeys[i] = k
		i++
	}

	for _, k := range mapKeys {
		for _, out := range modules[k].Outputs {
			in := input{
				Module: k,
				Input:  0,
			}
			if o, ok := modules[out]; ok {
				o.Inputs = append(o.Inputs, in)
				modules[out] = o
			} else {
				o = module{
					Type:   OUTPUT,
					Inputs: []input{in},
				}
				modules[out] = o
			}
		}
	}

	button := module{
		Type:    BUTTON,
		Outputs: []string{"broadcaster"},
	}

	modules["start"] = button

	return modules
}

func pushButton(modules map[string]module) (int, int) {
	queue := make([]pulse, 0)

	lowPulses, highPulses := 0, 0

	p := pulse{
		Start:       "button",
		Destination: "broadcaster",
		Value:       0,
	}

	queue = append(queue, p)

	for len(queue) > 0 {
		pulses := len(queue)

		for i := 0; i < pulses; i++ {
			curPulse := queue[0]
			if curPulse.Value == 1 {
				highPulses++
			} else {
				lowPulses++
			}
			queue = queue[1:]
			destModule := modules[curPulse.Destination]
			nextPulses := destModule.ProcessInput(curPulse)
			modules[curPulse.Destination] = destModule

			queue = append(queue, nextPulses...)
		}

	}

	return lowPulses, highPulses
}

func SolvePart1(inputTxt []string) int {
	modules := parseModules(inputTxt)

	totalLowPulses, totalHighPulses := 0, 0
	for i := 0; i < 1000; i++ {
		lowPulses, highPulses := pushButton(modules)
		totalLowPulses += lowPulses
		totalHighPulses += highPulses
	}

	return totalLowPulses * totalHighPulses
}

func pushButtonUntilHighPulse(modules map[string]module, source string) bool {
	queue := make([]pulse, 0)

	p := pulse{
		Start:       "button",
		Destination: "broadcaster",
		Value:       0,
	}

	queue = append(queue, p)

	for len(queue) > 0 {
		pulses := len(queue)

		for i := 0; i < pulses; i++ {
			curPulse := queue[0]
			if curPulse.Value == 1 && curPulse.Start == source {
				return true
			}
			queue = queue[1:]
			destModule := modules[curPulse.Destination]
			nextPulses := destModule.ProcessInput(curPulse)
			modules[curPulse.Destination] = destModule

			queue = append(queue, nextPulses...)
		}

	}

	return false
}

func SolvePart2(inputTxt []string) int {
	modules := parseModules(inputTxt)

	sourceModules := modules["bb"].Inputs

	buttonPresses := make([]int, 0)
	for _, m := range sourceModules {
		newModules := parseModules(inputTxt)

		c := 0
		for {
			c++
			if pushButtonUntilHighPulse(newModules, m.Module) {
				break
			}
		}
		buttonPresses = append(buttonPresses, c)
	}

	lcm := uint(buttonPresses[0])
	for i := 1; i < len(buttonPresses); i++ {
		lcm = extmath.Lcm(lcm, uint(buttonPresses[i]))
	}

	return int(lcm)
}
