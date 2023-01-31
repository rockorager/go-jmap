package push

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"git.sr.ht/~rockorager/go-jmap"
)

// A subscription to an event stream
type EventSource struct {
	// The JMAP client to use for the stream
	Client *jmap.Client

	// The function to pass state change events to
	Handler func(*jmap.StateChange)

	// The events to subscribe to. If left unset, will default to AllEvents
	Events []jmap.EventType

	// Interval the server should ping the client at, in seconds. The server
	// may choose to ignore this value. Set to 0 to disable pinging (which
	// the server may also ignore)
	Ping uint

	// Whether to close the connection after a state event
	CloseAfterState bool

	// The response of the request.
	resp *http.Response
}

// Connect to the server
func (e *EventSource) connect() error {
	// Create the URL for the subscription
	url, err := url.Parse(e.Client.Session.EventSourceURL)
	if err != nil {
		return err
	}
	q := url.Query()

	if len(e.Events) == 0 {
		e.Events = []jmap.EventType{jmap.AllEvents}
	}
	// types field
	types := []string{}
	for _, e := range e.Events {
		types = append(types, string(e))
	}
	typeStr := strings.Join(types, ",")
	q.Set("types", typeStr)

	// ping field
	q.Set("ping", fmt.Sprintf("%d", e.Ping))

	// close after field
	closeAfter := "no"
	if e.CloseAfterState {
		closeAfter = "state"
	}
	q.Set("closeafter", closeAfter)

	// set the query string
	url.RawQuery = q.Encode()

	// make the request
	e.resp, err = e.Client.HttpClient.Get(url.String())
	if err != nil {
		return err
	}
	if e.resp.StatusCode != 200 {
		return fmt.Errorf("invalid request, response code: %d", e.resp.StatusCode)
	}
	return nil
}

// Starts listening for events from the source. Listen will block when called
// and return when the source has been disconnected or closed via a call to
// Close()
func (e *EventSource) Listen() error {
	err := e.connect()
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(e.resp.Body)
	var event string

	for scanner.Scan() {
		// https://html.spec.whatwg.org/multipage/server-sent-events.html#event-stream-interpretation
		line := scanner.Text()
		switch {
		case strings.HasPrefix(line, ":"):
			// ignore the line, it's a comment
		case strings.Contains(line, ":"):
			fields := strings.SplitN(line, ":", 2)
			if len(fields) < 2 {
				continue
			}
			k := fields[0]
			v := strings.TrimSpace(fields[1])
			switch k {
			case "event":
				event = v
			case "data":
				switch event {
				case "state":
					state := &jmap.StateChange{}
					err := json.Unmarshal([]byte(v), state)
					if err != nil {
						return err
					}
					e.Handler(state)
				}
			}
		}
	}
	return nil
}

// Closes the stream
func (e *EventSource) Close() {
	if e.resp != nil && e.resp.Body != nil {
		e.resp.Body.Close()
	}
}
