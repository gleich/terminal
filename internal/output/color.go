package output

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/creack/pty"
	"github.com/muesli/termenv"
)

type Styles struct {
	Renderer *lipgloss.Renderer
	Blue     lipgloss.Style
	Green    lipgloss.Style
	Grey     lipgloss.Style
	Red      lipgloss.Style
}

func LoadStyles(s ssh.Session) Styles {
	clientOutput := outputFromSession(s)
	r := lipgloss.NewRenderer(s)
	r.SetOutput(clientOutput)
	return Styles{
		Renderer: r,
		Blue:     r.NewStyle().Foreground(lipgloss.Color("#2B95FF")),
		Green:    r.NewStyle().Foreground(lipgloss.Color("#30CE75")),
		Grey:     r.NewStyle().Foreground(lipgloss.Color("#424242")),
		Red:      r.NewStyle().Foreground(lipgloss.Color("#F30928")),
	}
}

// Bridge Wish and Termenv so we can query for a user's terminal capabilities.
type sshOutput struct {
	ssh.Session
	tty *os.File
}

func (s *sshOutput) Write(p []byte) (int, error) {
	return s.Session.Write(p)
}

func (s *sshOutput) Read(p []byte) (int, error) {
	return s.Session.Read(p)
}

func (s *sshOutput) Fd() uintptr {
	return s.tty.Fd()
}

type sshEnviron struct {
	environ []string
}

func (s *sshEnviron) Getenv(key string) string {
	for _, v := range s.environ {
		if strings.HasPrefix(v, key+"=") {
			return v[len(key)+1:]
		}
	}
	return ""
}

func (s *sshEnviron) Environ() []string {
	return s.environ
}

// Create a termenv.Output from the session.
func outputFromSession(sess ssh.Session) *termenv.Output {
	sshPty, _, _ := sess.Pty()
	_, tty, err := pty.Open()
	if err != nil {
		log.Fatal(err)
	}
	o := &sshOutput{
		Session: sess,
		tty:     tty,
	}
	environ := sess.Environ()
	environ = append(environ, fmt.Sprintf("TERM=%s", sshPty.Term))
	e := &sshEnviron{environ: environ}
	// We need to use unsafe mode here because the ssh session is not running
	// locally and we already know that the session is a TTY.
	return termenv.NewOutput(o, termenv.WithUnsafe(), termenv.WithEnvironment(e))
}
