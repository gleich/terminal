package main

import (
	"errors"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/activeterm"
	"github.com/charmbracelet/wish/logging"
	"github.com/gleich/lumber/v2"
	"github.com/gleich/terminal/internal/cmds"
	"github.com/gleich/terminal/internal/output"
	"github.com/joho/godotenv"
)

func main() {
	setupLogger()

	err := godotenv.Load()
	if err != nil {
		lumber.Fatal(err, "loading .env failed")
	}

	go startHTTP()
	startSSH()
}

func setupLogger() {
	logger := lumber.NewCustomLogger()
	nytime, err := time.LoadLocation("America/New_York")
	if err != nil {
		lumber.Fatal(err, "failed to load new york timezone")
	}
	logger.Timezone = nytime
	lumber.SetLogger(logger)
}

func startHTTP() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://github.com/gleich/terminal", http.StatusTemporaryRedirect)
	})

	lumber.Info("starting http server")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		lumber.Fatal(err, "starting http server failed")
	}
}

func startSSH() {
	homedir, err := os.UserHomeDir()
	if err != nil {
		lumber.Fatal(err, "getting home directory failed")
	}

	srv, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort("0.0.0.0", "22")),
		wish.WithHostKeyPath(filepath.Join(homedir, ".ssh", "id_rsa")),
		wish.WithMiddleware(func(next ssh.Handler) ssh.Handler {
			return func(s ssh.Session) {
				colors := output.LoadColors(s)
				output.Welcome(s, colors)
				cmds.Terminal(s, colors)
			}
		}, logging.Middleware(), activeterm.Middleware()),
	)
	if err != nil {
		lumber.Fatal(err, "creating server failed")
	}

	lumber.Info("starting ssh server")
	if err = srv.ListenAndServe(); err != nil && !errors.Is(err, ssh.ErrServerClosed) {
		lumber.Fatal(err, "starting server failed")
	}
}
