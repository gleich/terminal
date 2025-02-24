package main

import (
	"errors"
	"fmt"
	"io/fs"
	"net"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/joho/godotenv"
	"go.mattglei.ch/lcp-2/pkg/lcp"
	"go.mattglei.ch/terminal/internal/cmds"
	"go.mattglei.ch/terminal/internal/output"
	"go.mattglei.ch/timber"
)

func main() {
	setupLogger()

	if _, err := os.Stat(".env"); !errors.Is(err, fs.ErrNotExist) {
		err := godotenv.Load()
		if err != nil {
			timber.Fatal(err, "loading .env failed")
		}
	}

	homedir, err := os.UserHomeDir()
	if err != nil {
		timber.Fatal(err, "getting home directory failed")
	}

	client := lcp.Client{Token: os.Getenv("LCP_TOKEN")}

	srv, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort("0.0.0.0", "22")),
		wish.WithHostKeyPath(filepath.Join(homedir, ".ssh", "id_rsa")),
		wish.WithMiddleware(func(next ssh.Handler) ssh.Handler {
			return func(s ssh.Session) {
				ct := time.Now()
				timber.Info(
					fmt.Sprintf("login from user \"%s\" started connection", s.User()),
				)
				styles := output.LoadStyles(s)
				if os.Getenv("OUTPUT_WELCOME") == "true" {
					output.Welcome(s, styles)
				}
				cmds.Terminal(s, styles, &client)
				timber.Done(
					fmt.Sprintf("logout from user \"%s\"; spent %s", s.User(), time.Since(ct)),
				)
			}
		}, activeterm.Middleware()),
	)
	if err != nil {
		timber.Fatal(err, "creating server failed")
	}

	timber.Info("starting ssh server")
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		timber.Fatal(err, "starting server failed")
	}
}

func setupLogger() {
	ny, err := time.LoadLocation("America/New_York")
	if err != nil {
		timber.Fatal(err, "failed to load new york timezone")
	}
	timber.Timezone(ny)
	timber.TimeFormat("01/02 03:04:05 PM MST")
}
