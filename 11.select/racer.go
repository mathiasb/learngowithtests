package selecter

import (
	"fmt"
	"net/http"
	"time"
)

const (
	tenSecondTimeout = time.Second * 10
)

func Racer(a, b string) (selectedUrl string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (selectedUrl string, err error) {
	select {
	case <-ping(a):
		selectedUrl = a
	case <-ping(b):
		selectedUrl = b
	case <-time.After(timeout):
		err = fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
	return
}

func ping(a string) chan struct{} {
	ch := make(chan struct{})
	go func(url string) {
		http.Get(url)
		close(ch)
	}(a)
	return ch
}
