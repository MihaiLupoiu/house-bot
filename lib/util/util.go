package util

import (
	"io/ioutil"
	"net/http"
)

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	kk, err := ioutil.ReadAll(resp.Body)
	return kk, err
}
