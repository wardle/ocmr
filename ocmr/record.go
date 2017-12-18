package ocmr

import (
	"encoding/json"
	"fmt"
	"github.com/wardle/ocmr/snomed"
	"strings"
)

// Record is made up of a scenario containing findings and one single best answer.
type Record struct {
	Age      int
	Sex      Sex
	Findings []*ClinicalFinding // a list of clinical findings derived from the scenario (stem)
	Answer   *snomed.Concept    // the single best answer
	Parents  []*snomed.Concept  // cache parents of correct answer for ease-of-use in subsequent processing
}

func (q Record) String() string {
	findings := make([]string, 0)
	for _, finding := range q.Findings {
		findings = append(findings, finding.String())
	}
	return fmt.Sprintf("[%s] --> %s", strings.Join(findings, ", "), q.Answer.FullySpecifiedName)
}

// Sex of patient in the question
type Sex int

// Possible values for Sex
const (
	Male   Sex = 1
	Female Sex = 2
)

// MarshalJSON marshalls sex into JSON
func (sex Sex) MarshalJSON() ([]byte, error) {
	var s string
	switch sex {
	default:
		s = "unknown"
	case Male:
		s = "male"
	case Female:
		s = "female"
	}
	return json.Marshal(s)
}

// Duration reflects the temporal course of a clinical finding
type Duration int

// Valid types of Duration
const (
	Unknown  Duration = iota // symptom onset is unknown / not specified
	Acute                    // the symptom came on acutely
	Subacute                 // the symptom came on subacutely
	Chronic                  // the symptom has been chronic - e.g. co-morbidities
	Episodic                 // the symptom has been intermittent or episodic
)

func (d Duration) String() string {
	switch {
	case d == Acute:
		return "Acute"
	case d == Subacute:
		return "Subacute"
	case d == Chronic:
		return "Chronic"
	case d == Episodic:
		return "Episodic"
	default:
		return "Unknown"
	}
}

// MarshalJSON marshalls duration into JSON
func (d Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// ClinicalFinding combines a clinical finding SNOMED-CT concept and a duration
// e.g. acute chest pain
type ClinicalFinding struct {
	Concept  *snomed.Concept
	Parents  []*snomed.Concept
	Duration Duration
}

func (cf ClinicalFinding) String() string {
	return cf.Duration.String() + " " + cf.Concept.FullySpecifiedName
}
