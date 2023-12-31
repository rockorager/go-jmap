package jmap_test

import (
	"fmt"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/core/push"
	"git.sr.ht/~rockorager/go-jmap/mail"
	"git.sr.ht/~rockorager/go-jmap/mail/email"
	"git.sr.ht/~rockorager/go-jmap/mail/mailbox"
)

// Basic usage of the client, with chaining of methods
func Example() {
	// Create a new client. The SessionEndpoint must be specified for
	// initial connections.
	client := &jmap.Client{
		SessionEndpoint: "https://api.fastmail.com/jmap/session",
	}
	// Set the authentication mechanism. This also sets the HttpClient of
	// the jmap client
	client.WithAccessToken("my-access-token")

	// Authenticate the client. This gets a Session object. Session objects
	// are cacheable, and have their own state string clients can use to
	// decide when to refresh. The client can be initialized with a cached
	// Session object. If one isn't available, the first request will also
	// authenticate the client
	if err := client.Authenticate(); err != nil {
		// Handle the error
	}

	// Get the account ID of the primary mail account
	id := client.Session.PrimaryAccounts[mail.URI]

	// Create a new request
	req := &jmap.Request{}

	// Invoke a method. The CallID of this method will be returned to be
	// used when chaining calls
	req.Invoke(&mailbox.Get{
		Account: id,
	})

	// Invoke a changes call, let's save the callID and pass it to a Get
	// method
	callID := req.Invoke(&email.Changes{
		Account:    id,
		SinceState: "some-known-state",
	})

	// Invoke a result reference call
	req.Invoke(&email.Get{
		Account: id,
		ReferenceIDs: &jmap.ResultReference{
			ResultOf: callID,          // The CallID of the referenced method
			Name:     "Email/changes", // The name of the referenced method
			Path:     "/created",      // JSON pointer to the location of the reference
		},
	})

	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		// Handle the error
	}

	// Loop through the responses to invidividual invocations
	for _, inv := range resp.Responses {
		// Our result to individual calls is in the Args field of the
		// invocation
		switch r := inv.Args.(type) {
		case *mailbox.GetResponse:
			// A GetResponse contains a List of the objects
			// retrieved
			for _, mbox := range r.List {
				fmt.Printf("Mailbox name: %s", mbox.Name)
				fmt.Printf("Total email: %d", mbox.TotalEmails)
				fmt.Printf("Unread email: %d", mbox.UnreadEmails)
			}
		case *email.GetResponse:
			for _, eml := range r.List {
				fmt.Printf("Email subject: %s", eml.Subject)
			}
		}
		// There is a response in here to the Email/changes call, but we
		// don't care about the results since we passed them to the
		// Email/get call
	}
}

// Example usage of an eventsource push notification connection
func Example_eventsource() {
	client := &jmap.Client{
		SessionEndpoint: "https://api.fastmail.com/jmap/session",
	}
	myHandlerFunc := func(change *jmap.StateChange) {
		// handle the change
	}

	// If we don't set the Events field, all events will be subscribed to
	stream := &push.EventSource{
		Client:  client,
		Handler: myHandlerFunc,
	}
	if err := stream.Listen(); err != nil {
		// error occurs if the stream couldn't connect. Listen will
		// return when stream.Close is called
	}
}
