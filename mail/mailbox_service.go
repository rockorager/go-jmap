package mail

import (
	"context"
	"encoding/json"
	"fmt"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/client"
)

// MailboxService represents the set of methods available for the Mailbox object
type MailboxService struct {
	accts map[string]*mailAccount
	c     *client.Client
}

func newMailboxService(c *client.Client, accts map[string]*mailAccount) *MailboxService {
	return &MailboxService{
		accts: accts,
		c:     c,
	}
}

// Get lists the mailboxes in a users account. Leave the IDs property empty to
// list all mailboxes
func (m *MailboxService) Get(account string) *MailboxGetCall {
	return &MailboxGetCall{
		acct:  account,
		accts: m.accts,
		args:  &MailboxGetRequest{},
		c:     m.c,
		ctx:   context.Background(),
	}
}

// Changes lists mailbox IDs which have changed since the state supplied by the
// client
func (m *MailboxService) Changes(account string, state string) *MailboxChangesCall {
	return &MailboxChangesCall{
		acct:  account,
		accts: m.accts,
		args: &MailboxChangesRequest{
			SinceState: state,
		},
		c:   m.c,
		ctx: context.Background(),
	}
}

// Query lists mailbox IDs matching a filter criteria. The list can optionally
// be sorted
func (m *MailboxService) Query(account string) *MailboxQueryCall {
	return &MailboxQueryCall{
		acct:  account,
		accts: m.accts,
		args:  &MailboxQueryRequest{},
		c:     m.c,
		ctx:   context.Background(),
	}
}

// QueryChanges lists mailbox IDs which have changed since the state supplied
// by the client, assuming the mailbox matches a filter. Can be optionally
// sorted
func (m *MailboxService) QueryChanges(account string, state string) *MailboxQueryChangesCall {
	return &MailboxQueryChangesCall{
		acct:  account,
		accts: m.accts,
		args: &MailboxQueryChangesRequest{
			SinceQueryState: state,
		},
		c:   m.c,
		ctx: context.Background(),
	}
}

// Set creates, updates, or deletes mailboxes in the account
func (m *MailboxService) Set(account string) *MailboxSetCall {
	return &MailboxSetCall{
		acct:  account,
		accts: m.accts,
		args:  &MailboxSetRequest{},
		c:     m.c,
		ctx:   context.Background(),
	}
}

// NewFilter returns an empty MailboxFilter object
func (m *MailboxService) NewFilter() *MailboxFilter {
	return &MailboxFilter{}
}

type MailboxGetCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *MailboxGetRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Mailbox/get method
func (c *MailboxGetCall) Do() (*MailboxGetResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Mailbox/get",
		CallID: c.id,
		Args:   c.args,
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

	raw := invResp.Args.(json.RawMessage)
	getResp := &MailboxGetResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *MailboxGetCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *MailboxGetCall) CallID(id string) {
	c.id = id
}

// Properties allows partial responses to be retrieved. Properties specifies
// the Mailbox Properties to be returned. If Properties is nil, all properties
// will be returned. Mailbox ID is always returned. Allowed Properties are:
// - name
// - parentId
// - role
// - sortOrder
// - totalEmails
// - unreadEmails
// - totalThreads
// - unreadThreads
// - myRights
// - isSubscribed
func (c *MailboxGetCall) Properties(props ...string) {
	c.args.Properties = append(c.args.Properties, props...)
}

// IDs sets the list of IDs to return from the call
func (c *MailboxGetCall) IDs(ids ...jmap.ID) {
	c.args.IDs = append(c.args.IDs, ids...)
}

type MailboxChangesCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *MailboxChangesRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Mailbox/changes method
func (c *MailboxChangesCall) Do() (*MailboxChangesResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Mailbox/changes",
		Args:   c.args,
		CallID: c.id,
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
	chResp := &MailboxChangesResponse{}
	err = json.Unmarshal(raw, chResp)
	if err != nil {
		return nil, err
	}
	return chResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *MailboxChangesCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *MailboxChangesCall) CallID(id string) {
	c.id = id
}

// MaxChanges sets the maximum number of changes the server will return in one
// response
func (c *MailboxChangesCall) MaxChanges(max jmap.UnsignedInt) {
	c.args.MaxChanges = max
}

type MailboxQueryCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *MailboxQueryRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Mailbox/query method
func (c *MailboxQueryCall) Do() (*MailboxQueryResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Mailbox/query",
		Args:   c.args,
		CallID: c.id,
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
	getResp := &MailboxQueryResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// CallID sets the call-id within the invocation
func (c *MailboxQueryCall) CallID(id string) {
	c.id = id
}

// Context sets the context to use for the calls Do method.
func (c *MailboxQueryCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// Filter applies a filter to the returned results.
func (c *MailboxQueryCall) Filter(filter *MailboxFilter) {
	if filter.Operator == "" {
		c.args.Filter = filter.Conditions[0]
		return
	}
	c.args.Filter = filter
}

// If true, when sorting the query results and comparing Mailboxes A and B:
//
//     If A is an ancestor of B, it always comes first regardless of the sort
//     comparators. Similarly, if A is descendant of B, then B always comes
//     first.
//
//     Otherwise, if A and B do not share a parentId, find the nearest
//     ancestors of each that do have the same parentId and compare the sort
//     properties on those Mailboxes instead.
//
// The result of this is that the Mailboxes are sorted as a tree according to
// the parentId properties, with each set of children with a common parent
// sorted according to the standard sort comparators.
func (c *MailboxQueryCall) SortAsTree(b bool) {
	c.args.SortAsTree = b
}

// If true, a Mailbox is only included in the query if all its ancestors are
// also included in the query according to the filter.
func (c *MailboxQueryCall) FilterAsTree(b bool) {
	c.args.FilterAsTree = b
}

// Servers MUST support sorting by the following properties:
// - sortOrder
// - name
// Multiple calls to sort are supported for secondary sort options
func (c *MailboxQueryCall) Sort(s *MailboxSortComparator) {
	c.args.Sort = append(c.args.Sort, s)
}

// The zero-based index of the first id in the full list of results to return.
//
// If a negative value is given, it is an offset from the end of the list.
// Specifically, the negative value MUST be added to the total number of
// results given the filter, and if still negative, it’s clamped to 0. This is
// now the zero-based index of the first id to return.
//
// If the index is greater than or equal to the total number of objects in the
// results list, then the ids array in the response will be empty, but this is
// not an error.
func (c *MailboxQueryCall) Position(p jmap.Int) {
	c.args.Position = p
}

// A Foo id. If supplied, the position argument is ignored. The index of this
// id in the results will be used in combination with the anchorOffset argument
// to determine the index of the first result to return (see below for more
// details).
//
// If an anchor argument is given, the anchor is looked for in the results
// after filtering and sorting. If found, the anchorOffset is then added to its
// index. If the resulting index is now negative, it is clamped to 0. This
// index is now used exactly as though it were supplied as the position
// argument. If the anchor is not found, the call is rejected with an
// anchorNotFound error.
//
// If an anchor is specified, any position argument supplied by the client MUST
// be ignored. If no anchor is supplied, any anchorOffset argument MUST be
// ignored.
//
// A client can use anchor instead of position to find the index of an id
// within a large set of results.
func (c *MailboxQueryCall) Anchor(a jmap.ID) {
	c.args.Anchor = a
}

// The index of the first result to return relative to the index of the anchor,
// if an anchor is given. This MAY be negative. For example, -1 means the Foo
// immediately preceding the anchor is the first result in the list returned
// (see below for more details).
func (c *MailboxQueryCall) AnchorOffset(ao jmap.Int) {
	c.args.AnchorOffset = ao
}

// The maximum number of results to return. If null, no limit presumed. The
// server MAY choose to enforce a maximum limit argument. In this case, if a
// greater value is given (or if it is null), the limit is clamped to the
// maximum; the new limit is returned with the response so the client is aware.
// If a negative value is given, the call MUST be rejected with an
// invalidArguments error.
func (c *MailboxQueryCall) Limit(l jmap.UnsignedInt) {
	c.args.Limit = l
}

// Does the client wish to know the total number of results in the query? This
// may be slow and expensive for servers to calculate, particularly with
// complex filters, so clients should take care to only request the total when
// needed.
func (c *MailboxQueryCall) CanCalculateTotal(b bool) {
	c.args.CalculateTotal = b
}

type MailboxQueryChangesCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *MailboxQueryChangesRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Mailbox/queryChanges method
func (c *MailboxQueryChangesCall) Do() (*MailboxQueryChangesResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Mailbox/queryChanges",
		Args:   c.args,
		CallID: c.id,
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
	getResp := &MailboxQueryChangesResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// CallID sets the call-id within the invocation
func (c *MailboxQueryChangesCall) CallID(id string) {
	c.id = id
}

// Context sets the context to use for the calls Do method.
func (c *MailboxQueryChangesCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// MaxChanges sets the maximum number of changes the server will return in one
// response
func (c *MailboxQueryChangesCall) MaxChanges(max jmap.UnsignedInt) {
	c.args.MaxChanges = max
}

// The last (highest-index) id the client currently has cached from the query
// results. When there are a large number of results, in a common case, the
// client may have only downloaded and cached a small subset from the beginning
// of the results. If the sort and filter are both only on immutable
// properties, this allows the server to omit changes after this point in the
// results, which can significantly increase efficiency. If they are not
// immutable, this argument is ignored.
func (c *MailboxQueryChangesCall) UpToID(id jmap.ID) {
	c.args.UpToID = id
}

// Servers MUST support sorting by the following properties:
// - sortOrder
// - name
// Multiple calls to sort are supported for secondary sort options
func (c *MailboxQueryChangesCall) Sort(s *MailboxSortComparator) {
	c.args.Sort = append(c.args.Sort, s)
}

// Filter applies a filter to the returned results.
func (c *MailboxQueryChangesCall) Filter(filter *MailboxFilter) {
	if filter.Operator == "" {
		c.args.Filter = filter.Conditions[0]
		return
	}
	c.args.Filter = filter
}

type MailboxSetCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *MailboxSetRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Mailbox/set method
func (c *MailboxSetCall) Do() (*MailboxSetResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Mailbox/set",
		Args:   c.args,
		CallID: c.id,
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
	getResp := &MailboxSetResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *MailboxSetCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *MailboxSetCall) CallID(id string) {
	c.id = id
}

// Create creates a mailbox. A temporary ID must be given. The response will
// map the temporary ID to the actual ID of the created mailbox
//
// Multiple mailboxes may be created in a single call
func (c *MailboxSetCall) Create(tempId jmap.ID, mbox *Mailbox) {
	if c.args.Create == nil {
		c.args.Create = make(map[jmap.ID]*Mailbox)
	}
	c.args.Create[tempId] = mbox
}

// Update updates a given mailbox. Multiple mailboxes may be updated in a
// single call
func (c *MailboxSetCall) Update(id jmap.ID, patch jmap.Patch) {
	if c.args.Update == nil {
		c.args.Update = make(map[jmap.ID]map[string]interface{})
	}
	p := map[string]interface{}{
		patch.Path: patch.Value,
	}
	c.args.Update[id] = p
}

// Destroy deletes an entire mailbox. See OnDestroyRemoveEmails. Multiple
// mailboxes may be destroyed in a single call. Mailboxes cannot be destroyed
// recursively: if a mailbox has children, it will not be destroyed and an
// error will be returned on execution
func (c *MailboxSetCall) Destroy(id jmap.ID) {
	c.args.Destroy = append(c.args.Destroy, id)
}

// If false, any attempt to destroy a Mailbox that still has Emails in it will
// be rejected with a mailboxHasEmail SetError. If true, any Emails that were
// in the Mailbox will be removed from it, and if in no other Mailboxes, they
// will be destroyed when the Mailbox is destroyed.
func (c *MailboxSetCall) OnDestroyRemoveEmails(b bool) {
	c.args.OnDestroyRemoveEmails = b
}
