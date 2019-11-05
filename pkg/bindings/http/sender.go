package http

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	nethttp "net/http"
	"net/url"

	"github.com/cloudevents/sdk-go/pkg/binding"
)

type Sender struct {
	// Client is the HTTP client used to send events as HTTP requests
	Client *http.Client
	// Target is the URL to send event requests to.
	Target *url.URL
}

func (s *Sender) Send(ctx context.Context, m binding.Message) (err error) {
	defer func() { _ = m.Finish(err) }()
	if s.Client == nil || s.Target == nil {
		return fmt.Errorf("not initialized: %#v", s)
	}
	m, err = binding.Translate(m,
		BinaryEncoder{}.Encode,
		func(f string, b []byte) (binding.Message, error) { return NewStruct(f, b), nil },
	)
	if err != nil {
		return err
	}
	body := bytes.NewReader(m.(*Message).Body)
	req, err := http.NewRequest("POST", s.Target.String(), body)
	if err != nil {
		return err
	}
	req = req.WithContext(ctx)
	req.Header = m.(*Message).Header
	resp, err := s.Client.Do(req)
	if resp.StatusCode/100 != 2 {
		return fmt.Errorf("%d %s", resp.StatusCode, nethttp.StatusText(resp.StatusCode))
	}
	return nil
}

func NewSender(client *http.Client, target *url.URL) *Sender {
	return &Sender{Client: client, Target: target}
}
