package logs

import (
	"log"
	"log/syslog"
	"os"
)

// Init log level to print in syslog
func Init(tag string, debug bool) {
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
