package mail

import (
	"context"
	"encoding/json"
	"fmt"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/client"
)

// ThreadService represents the set of methods available for the Thread object
type ThreadService struct {
	accts map[string]*mailAccount
	c     *client.Client
}

func newThreadService(c *client.Client, accts map[string]*mailAccount) *ThreadService {
	return &ThreadService{
		accts: accts,
		c:     c,
	}
}

// Get lists the threads in a users account.
func (s *ThreadService) Get(account string) *ThreadGetCall {
	return &ThreadGetCall{
		acct:  account,
		accts: s.accts,
		args:  &ThreadGetRequest{},
		c:     s.c,
		ctx:   context.Background(),
	}
}

// Changes lists thread IDs which have changed since the state supplied by the client
func (s *ThreadService) Changes(account string, state string) *ThreadChangesCall {
	return &ThreadChangesCall{
		acct:  account,
		accts: s.accts,
		args: &ThreadChangesRequest{
			SinceState: state,
		},
		c:   s.c,
		ctx: context.Background(),
	}
}

type ThreadGetCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *ThreadGetRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Thread/get method
func (c *ThreadGetCall) Do() (*ThreadGetResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name: "Thread/get",
		Args: c.args,
	}
	if c.id != "" {
		inv.CallID = c.id
	}
	req := &jmap.Request{
		Using: []string{jmap.CoreCapabilityName, MailCapabilityName},
		Calls: []jmap.Invocation{inv},
	}
	resp, err := c.c.RawSendWithContext(c.ctx, req)
	if err != nil {
		return nil, err
	}
	invResp := resp.Responses[0]
	if invResp.Name == "error" {
		errArgs, ok := invResp.Args.(jmap.MethodErrorArgs)
		if !ok {
			return nil, fmt.Errorf("error converting method error")
		}
		return nil, errArgs
	}
	raw, ok := invResp.Args.(json.RawMessage)
	if !ok {
		return nil, err
	}
	getResp := &ThreadGetResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *ThreadGetCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *ThreadGetCall) CallID(id string) {
	c.id = id
}

// Properties allows partial responses to be retrieved. Properties specifies
// the Thread Properties to be returned. If Properties is nil, all properties
// will be returned. Thread ID is always returned. Allowed Properties are:
// - emailIds
//
// This is mostly useless for the Thread object
func (c *ThreadGetCall) Properties(props ...string) {
	c.args.Properties = append(c.args.Properties, props...)
}

// IDs sets the list of IDs to return from the call
func (c *ThreadGetCall) IDs(ids ...jmap.ID) {
	c.args.IDs = append(c.args.IDs, ids...)
}

type ThreadChangesCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *ThreadChangesRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Thread/changes method
func (c *ThreadChangesCall) Do() (*ThreadChangesResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name: "Thread/changes",
		Args: c.args,
	}
	if c.id != "" {
		inv.CallID = c.id
	}
	req := &jmap.Request{
		Using: []string{jmap.CoreCapabilityName, MailCapabilityName},
		Calls: []jmap.Invocation{inv},
	}
	resp, err := c.c.RawSendWithContext(c.ctx, req)
	if err != nil {
		return nil, err
	}
	invResp := resp.Responses[0]
	if invResp.Name == "error" {
		errArgs, ok := invResp.Args.(jmap.MethodErrorArgs)
		if !ok {
			return nil, fmt.Errorf("error converting method error")
		}
		return nil, errArgs
	}
	raw, ok := invResp.Args.(json.RawMessage)
	if !ok {
		return nil, err
	}
	chResp := &ThreadChangesResponse{}
	err = json.Unmarshal(raw, chResp)
	if err != nil {
		return nil, err
	}
	return chResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *ThreadChangesCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *ThreadChangesCall) CallID(id string) {
	c.id = id
}

// MaxChanges sets the maximum number of changes the server will return in one
// response
func (c *ThreadChangesCall) MaxChanges(max jmap.UnsignedInt) {
	c.args.MaxChanges = max
}
