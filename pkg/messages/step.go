package messages

import (
	"fmt"

	"github.com/Matt-Gleich/logoru"
	"github.com/gliderlabs/ssh"
)

// Outline for a step
type Step struct {
	Name    string
	Session ssh.Session
}

// Start a step
func (s Step) Start() {
	logoru.Info("Starting", s.Name)
	fmt.Fprint(s.Session, s.Name, " ...")
}

// Finish a step
func (s Step) Done() {
	logoru.Success("Done with", s.Name)
	fmt.Fprintln(s.Session, " done")
}

// Output a step that finishes instantly
func FakeStep(s ssh.Session, name string) {
	step := Step{Name: name, Session: s}
	step.Start()
	step.Done()
}
