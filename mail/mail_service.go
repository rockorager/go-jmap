package mail

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/client"
)

type MailService struct {
	Mailboxes *MailboxService
	Threads   *ThreadService
	Emails    *EmailService
	// Identities       *Identities
	// EmailSubmission  *EmailSubmission
	// VacationResponse *VacationResponse

	accts map[string]*mailAccount
	c     *client.Client
	ctx   context.Context
}

type mailAccount struct {
	jmap.Account
	ID             jmap.ID
	MailCapability *Capability
}

func NewService(username string, password string) (*MailService, error) {
	return NewServiceWithClient(http.DefaultClient, username, password)
}

func NewServiceWithClient(hc client.HTTPClient, username string, password string) (*MailService, error) {
	auth := authHeader(username, password)
	c, err := client.NewWithClient(hc, "https://jmap.fastmail.com:443/.well-known/jmap", auth)
	if err != nil {
		return nil, err
	}
	unmarshallers := jmap.RawUnmarshallers([]string{
		// Mailbox methods
		"Mailbox/get",
		"Mailbox/changes",
		"Mailbox/query",
		"Mailbox/queryChanges",
		"Mailbox/set",

		// Thread methods
		"Thread/get",
		"Thread/changes",

		// Email methods
		"Email/get",
		"Email/changes",
		"Email/query",
		"Email/queryChanges",
		"Email/set",
		"Email/copy",
		"Email/import",
		"Email/parse",
	})

	c.Enable(unmarshallers)
	m := &MailService{
		accts: make(map[string]*mailAccount),
		c:     c,
		ctx:   context.Background(),
	}
	for id, acct := range c.Session.Accounts {
		if raw, ok := acct.Capabilities[MailCapabilityName]; ok {
			mc := &Capability{}
			err = json.Unmarshal(raw, mc)
			if err != nil {
				return nil, err
			}
			mailAcct := &mailAccount{
				Account:        acct,
				ID:             id,
				MailCapability: mc,
			}
			m.accts[acct.Name] = mailAcct
		}
	}
	m.Mailboxes = newMailboxService(m.c, m.accts)
	m.Threads = newThreadService(m.c, m.accts)
	m.Emails = newEmailService(c, m.accts)
	return m, nil
}

// Context sets the default context to use for all Do methods
func (m *MailService) Context(ctx context.Context) {
	m.ctx = ctx
}

// Session returns the session object
func (m *MailService) Session() *jmap.Session {
	return m.c.Session
}

// Capability returns the mail Capability for the given account
func (m *MailService) Capability(acct string) (*Capability, bool) {
	a, ok := m.accts[acct]
	if !ok {
		return nil, false
	}
	return a.MailCapability, true
}

func (m *MailService) Download(acct string, blob jmap.ID) (io.ReadCloser, error) {
	ma, ok := m.accts[acct]
	if !ok {
		return nil, fmt.Errorf("could not find account %s", acct)
	}

	return m.c.Download(ma.ID, blob)
}

func authHeader(username string, password string) string {
	buf := bytes.NewBuffer(nil)
	encoder := base64.NewEncoder(base64.StdEncoding, buf)
	_, err := encoder.Write([]byte(username + ":" + password))
	if err != nil {
		return ""
	}
	encoder.Close()
	return "Basic " + buf.String()
}
