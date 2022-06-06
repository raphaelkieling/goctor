package main

type Runner struct {
	Configuration Configuration
}

type OutputStep struct {
	Name         string
	Description  string
	Error        bool
	Possiblities []string
}

type OutputExam struct {
	Name        string
	Description string
	AnyError    bool
}

type OutputRunner struct {
	Exams []OutputExam
}

func (r *Runner) NewRunner(configuration Configuration) *Runner {
	return &Runner{
		Configuration: configuration,
	}
}

func (r *Runner) StartProcess() OutputRunner {
	output := OutputRunner{
		Exams: []OutputExam{},
	}

	for _, exam := range r.Configuration.Exams {
		outputExam := OutputExam{
			Name:        exam.Name,
			Description: exam.Description,
		}
		output.Exams = append(output.Exams, outputExam)

		for _, step := range exam.Steps {
			outputStep := OutputStep{
				Name:        step.Name,
				Description: step.Description,
				Error:       false,
			}
			exitCode, goctorcodes, err := step.execute()

			if err != nil {
				outputExam.AnyError = true
				for _, possibility := range step.Possibilities {
					if possibility.matchError(exitCode, goctorcodes) {
						outputStep.Error = true
						outputStep.Possiblities = append(outputStep.Possiblities, possibility.Message)
					}
				}
			}
		}
	}

	return output
}
