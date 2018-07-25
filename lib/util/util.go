package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"log/syslog"
	"net/http"
	"os"

	"github.com/MihaiLupoiu/house-bot/models"
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

// GetConfigurationFile parshe json configuration file.
func GetConfigurationFile(configFile string) models.Config {
	configuration := models.Config{}
	file, err := os.Open(configFile)
	if err != nil {
		log.Println("error:", err)
	} else {
		decoder := json.NewDecoder(file)
		err := decoder.Decode(&configuration)
		if err != nil {
			log.Println("error:", err)
		}
	}
	return configuration
}
