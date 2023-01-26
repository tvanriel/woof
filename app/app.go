package app

import (
	"fmt"
	"net"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type Application struct {
	file         string
	port         int
	max          int
	maxMutex     sync.Mutex
	current      int
	currentMutex sync.Mutex
}

func NewApplication() *Application {
	return new(Application)
}

func (a *Application) SetHostFile(file string) {
	a.file = file
}

func (a *Application) SetMax(max int) {
	a.max = max
}

func (a *Application) SetPort(port int) {
	a.port = port
}
func (a *Application) IncreasePort() {
	a.port += 1
}

func (a *Application) GetAddr() string {
	var sb strings.Builder
	sb.WriteString("0.0.0.0:")
	sb.WriteString(strconv.Itoa(a.port))
	return sb.String()
}

func (a *Application) ServeHttp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		a.maxMutex.Lock()
		a.max -= 1
		max := a.max
		a.maxMutex.Unlock()
		a.currentMutex.Lock()
		a.current += 1
		current := a.current
		a.currentMutex.Unlock()

		fmt.Println(logStyle.Render(fmt.Sprintf("GET [%s] (max=%d,current=%d) User-Agent: %s", r.RemoteAddr, max, current, r.UserAgent())))

		w.Header().Set("Content-Disposition", "attachment; filename="+strconv.Quote(path.Base(a.file)))
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, a.file)
		a.currentMutex.Lock()
		a.current -= 1
		a.currentMutex.Unlock()
	})

	fmt.Println(servingStyle.Render(fmt.Sprintf("Serving file: %s for a max of %d connections.", a.file, a.max)))
	printAddressesWithPort(a.port)
	fmt.Println()
	var srv http.Server
	srv.Addr = a.GetAddr()
	srv.Handler = mux
	go srv.ListenAndServe()
	for {
		time.Sleep(1 * time.Second)
		a.maxMutex.Lock()
		max := a.max
		a.maxMutex.Unlock()
		a.currentMutex.Lock()
		current := a.current
		a.currentMutex.Unlock()
		if max == 0 && current == 0 {
			srv.Close()
			return
		}
	}
}

func printAddressesWithPort(port int) {
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(errStyle.Render(fmt.Sprintf("Err: Unable to find interfaces of this machine: %v\n", err)))
	} else {
		for i := range ifaces {
			if ifaces[i].Name == "lo" {
				continue
			}
			var tmpl strings.Builder
			tmpl.WriteString(fmt.Sprintf("%v:\n", ifaces[i].Name))
			addrs, _ := ifaces[i].Addrs()
			for j := range addrs {
				if strings.Contains(addrs[j].String(), ":") {
					continue
				}
				ip, _, _ := strings.Cut(addrs[j].String(), "/")

				tmpl.WriteString(fmt.Sprintf("- %v:%d\n", ip, port))

			}
			fmt.Println(interfaceStyle.Render(strings.TrimSpace(tmpl.String())))
		}
	}
}

var interfaceStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("63")).
	Foreground(lipgloss.Color("#FAFAFA")).
	Padding(2).Margin(1)

var servingStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("#3079CF")).
	Foreground(lipgloss.Color("#FAFAFA")).
	Padding(2).Margin(1)

var errStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("#f55050")).
	Foreground(lipgloss.Color("#FAFAFA")).
	Padding(2).
	Margin(1)

var logStyle = lipgloss.NewStyle().
	Bold(true).
	Background(lipgloss.Color("#64C964")).
	Foreground(lipgloss.Color("#FAFAFA")).
	Padding(0).Margin(0)
