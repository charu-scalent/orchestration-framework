package helper

import (
	"net"
	"net/http"
	"time"
)

var netTransport = &http.Transport{
	Dial: (&net.Dialer{
		Timeout:   10 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial,
	TLSHandshakeTimeout: 10 * time.Second,
	MaxIdleConns:        100,
	MaxIdleConnsPerHost: 100,
}
var netClient = &http.Client{
	Timeout:   60 * time.Second,
	Transport: netTransport,
}

func Request(request *http.Request) (*http.Response, error) {
	resp, err := netClient.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func RequestWithTimeout(request *http.Request, timeoutInSecond int) (*http.Response, error) {
	netClient.Timeout = time.Duration(timeoutInSecond) * time.Second
	resp, err := netClient.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
