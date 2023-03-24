package app

import (
	"fmt"
	"net/http"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
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
		w.Header().Set("Cache-Control", "no-cache")
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
		if max <= 0 && current == 0 {
			srv.Close()
			return
		}
	}
}
