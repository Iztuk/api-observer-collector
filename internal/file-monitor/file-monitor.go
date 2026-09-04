// Package filemonitor focuses on monitoring files for any writes
package filemonitor

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func MonitorFiles() error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Has(fsnotify.Write) {
					log.Printf("File modified inside directory: %s\n", event.Name)
				}

				if event.Has(fsnotify.Create) {
					log.Printf("New file created inside directory: %s", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("Watcher error:", err)
			}
		}
	}()

	err = watcher.Add("./test/logs")
	if err != nil {
		return err
	}

	select {}
}
