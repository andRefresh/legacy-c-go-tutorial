package service

import (
	"fmt"
	"sync"
	"time"
)

type Watcher struct {
	mu sync.Mutex
	Publisher

	dataWatched map[any]any
	tickers     map[any]time.Ticker
}

type WatchData struct {
	Id   string
	Data any
}

func NewWatcher() (w *Watcher) {
	w = &Watcher{}
	w.dataWatched = make(map[any]any)
	w.tickers = make(map[any]time.Ticker)

	w.Publisher = NewPublisher()
	return
}

func (w *Watcher) StartWatch(id string, data any, ms time.Duration) {
	_, exists := w.dataWatched[id]
	if exists {
		w.StopWatch(id)
	}

	t := time.NewTicker(ms * time.Millisecond)
	go func() {
		for range t.C {
			w.watch(id, data)
		}
	}()

}

func (w *Watcher) StopWatch(data any) {
	_, exists := w.dataWatched[data]
	if exists {
		delete(w.dataWatched, data)
		ticker := w.tickers[data]
		ticker.Stop()
	}

}

func (w *Watcher) watch(id string, data any) {
	w.mu.Lock()

	oldDataStr := fmt.Sprintf("%v", w.dataWatched[id])
	newDataStr := fmt.Sprintf("%v", data)

	if oldDataStr != newDataStr {
		w.dataWatched[id] = newDataStr
		w.dataChanged(id, data)
	}

	w.mu.Unlock()
}

func (w *Watcher) dataChanged(id string, data any) {
	for _, subscriber := range w.subscribers {
		subscriber <- WatchData{Id: id, Data: data}
	}
}
