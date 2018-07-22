package util

import (
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"
	"os"
)

// InitLog level to print in syslog
func InitLog(tag string, debug bool) {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	if debug {
		log.SetOutput(os.Stdout)
	} else {
		var logWriter, logErr = syslog.New(syslog.LOG_ERR, tag)
		if logErr == nil {
			log.SetOutput(logWriter)
		}
	}
}

// Get body from URL
func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	kk, err := ioutil.ReadAll(resp.Body)
	return kk, err
}
