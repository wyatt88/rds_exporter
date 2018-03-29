package client

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type Client struct {
	c *http.Client
}

func New(logger *logrus.Entry) *Client {
	return &Client{
		c: &http.Client{
			Transport: newTransport(logger),
			Timeout:   15 * time.Second,
		},
	}
}

func (c *Client) HTTP() *http.Client {
	return c.c
}

func (c *Client) Describe(chan<- *prometheus.Desc) {
	// TODO
	return
}

func (c *Client) Collect(chan<- prometheus.Metric) {
	// TODO
	return
}
