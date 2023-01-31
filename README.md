# go-jmap

A JMAP client library. go-jmap is a library for interacting with JMAP servers.
It includes a client for making requests, and data structures for the Core and
Mail specifications.

Note: this library started as a fork of [github.com/foxcpp/go-jmap](https://github.com/foxcpp/go-jmap)
It has since undergone massive restructuring, and it only loosely based on the
original project.

## Usage

```go
package main

import (
	"fmt"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/mail"
	"git.sr.ht/~rockorager/go-jmap/mail/email"
	"git.sr.ht/~rockorager/go-jmap/mail/mailbox"
)

func main() {
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
		Account: id,
		SinceState: "some-known-state",
	})

	// Invoke a result reference call 
	req.Invoke(&email.Get{
		Account: id,
		ReferenceIDs: &jmap.ResultReference{
			ResultOf: callID, // The CallID of the referenced method
			Name: "Email/changes", // The name of the referenced method
			Path: "/created", // JSON pointer to the location of the reference
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
```

## Status

- [x] Client interface

  - [x] Access Token authentication (client.WithTokenAuth)
  - [x] Basic authentication (client.WithBasicAuth)
  - [x] Chain method calls (Request.Invoke(method...))
  - [x] BYO http.Client

- [ ] Core ([RFC 8620](https://tools.ietf.org/html/rfc8620))

  - [ ] Autodiscovery
  - [x] Session
  - [x] Account
  - [x] Core Capabilities
  - [x] Invocation
  - [x] Request
  - [x] Response
  - [x] Request-level errors
  - [x] Method-level errors
  - [x] Set-level errors

  - [x] Core/echo

  - [x] Blob/Downloading
  - [x] Blob/Uploading
  - [x] Blob/Copy method

  - [x] Push
    - [x] StateChange structure
    - [x] PushSubscription structure
    - [x] PushSubscription/get
    - [x] PushSubscription/set
    - [x] Event Source

- [ ] Mail ([RFC 8621](https://tools.ietf.org/html/rfc8621))

  - [x] Capability

  - [x] Mailbox

    - [x] Get
    - [x] Changes
    - [x] Query
    - [x] QueryChanges
    - [x] Set

  - [x] Threads

    - [x] Get
    - [x] Changes

  - [x] Emails

    - [x] Get
    - [x] Changes
    - [x] Query
    - [x] QueryChanges
    - [x] Set
    - [x] Copy
    - [x] Import
    - [x] Parse

  - [x] SearchSnippets

    - [x] Get

  - [ ] Identities

    - [ ] Get
    - [ ] Changes
    - [ ] Set

  - [ ] EmailSubmission

    - [ ] Get
    - [ ] Changes
    - [ ] Query
    - [ ] QueryChanges
    - [ ] Set

  - [ ] VacationResponse

    - [ ] Get

  - [ ] Client Macros
