package client

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type transport struct {
	t *http.Transport
	l *logrus.Entry
}

func newTransport(logger *logrus.Entry) *transport {
	return &transport{
		t: &http.Transport{
			MaxIdleConnsPerHost: 5,
			IdleConnTimeout:     2 * time.Minute,
		},
		l: logger,
	}
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	// TODO use net/http/httptrace to get a number of open connections
	resp, err := t.RoundTrip(req)
	t.l.Debugf("%s %s -> %d %v", req.Method, req.URL.String(), resp.StatusCode, err)
	return resp, err
}

// check interface
var _ http.RoundTripper = (*transport)(nil)
