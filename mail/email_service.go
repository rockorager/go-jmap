package mail

import (
	"context"
	"encoding/json"
	"fmt"

	"git.sr.ht/~rockorager/go-jmap"
	"git.sr.ht/~rockorager/go-jmap/client"
)

// EmailService represents the set of methods available for the Email object
type EmailService struct {
	accts map[string]*mailAccount
	c     *client.Client
}

func newEmailService(c *client.Client, accts map[string]*mailAccount) *EmailService {
	return &EmailService{
		accts: accts,
		c:     c,
	}
}

// Get lists the emails in a users account. Leave the IDs property empty to
// list all emails
func (s *EmailService) Get(account string) *EmailGetCall {
	return &EmailGetCall{
		acct:  account,
		accts: s.accts,
		args:  &EmailGetRequest{},
		c:     s.c,
		ctx:   context.Background(),
	}
}

// Changes lists email IDs which have changed since the state supplied by the
// client
func (s *EmailService) Changes(account string, state string) *EmailChangesCall {
	return &EmailChangesCall{
		acct:  account,
		accts: s.accts,
		args: &EmailChangesRequest{
			SinceState: state,
		},
		c:   s.c,
		ctx: context.Background(),
	}
}

// Query lists email IDs matching a filter criteria. The list can optionally
// be sorted
func (s *EmailService) Query(account string) *EmailQueryCall {
	return &EmailQueryCall{
		acct:  account,
		accts: s.accts,
		args:  &EmailQueryRequest{},
		c:     s.c,
		ctx:   context.Background(),
	}
}

// QueryChanges lists email IDs which have changed since the state supplied
// by the client, assuming the email matches a filter. Can be optionally
// sorted
func (s *EmailService) QueryChanges(account string, state string) *EmailQueryChangesCall {
	return &EmailQueryChangesCall{
		acct:  account,
		accts: s.accts,
		args: &EmailQueryChangesRequest{
			SinceQueryState: state,
		},
		c:   s.c,
		ctx: context.Background(),
	}
}

// The "Email/set" method encompasses:
//    -  Creating a draft
//    -  Changing the keywords of an Email (e.g., unread/flagged status)
//    -  Adding/removing an Email to/from Mailboxes (moving a message)
//    -  Deleting Emails
func (s *EmailService) Set(account string) *EmailSetCall {
	return &EmailSetCall{
		acct:  account,
		accts: s.accts,
		args:  &EmailSetRequest{},
		c:     s.c,
		ctx:   context.Background(),
	}
}

// This is a standard "/copy" method as described in [RFC8620], Section 5.4,
// except only the "mailboxIds", "keywords", and "receivedAt" properties may be
// set during the copy.  This method cannot modify the message represented by
// the Email.
func (s *EmailService) Copy(accountFrom string, accountTo string) *EmailCopyCall {
	return &EmailCopyCall{
		acctTo:   accountTo,
		acctFrom: accountFrom,
		accts:    s.accts,
		args: &EmailCopyRequest{
			Create: make(map[jmap.ID]*Email),
		},
		c:   s.c,
		ctx: context.Background(),
	}
}

// The "Email/import" method adds messages [RFC5322] to the set of Emails in an
// account.  The server MUST support messages with Email Address
// Internationalization (EAI) headers [RFC6532].  The messages must first be
// uploaded as blobs using the standard upload mechanism.
func (s *EmailService) Import(account string) *EmailImportCall {
	return &EmailImportCall{
		acct:  account,
		accts: s.accts,
		args: &EmailImportRequest{
			Emails: make(map[jmap.ID]*EmailImport),
		},
		c:   s.c,
		ctx: context.Background(),
	}
}

// This method allows you to parse blobs as messages [RFC5322] to get
// Email objects.  The server MUST support messages with EAI headers
// [RFC6532].  This can be used to parse and display attached messages
// without having to import them as top-level Email objects in the mail
// store in their own right.
func (s *EmailService) Parse(account string) *EmailParseCall {
	return &EmailParseCall{
		acct:  account,
		accts: s.accts,
		args:  &EmailParseRequest{},
		c:     s.c,
		ctx:   context.Background(),
	}
}

// NewFilter returns an empty Filter object
func (s *EmailService) NewFilter() *EmailFilter {
	return &EmailFilter{}
}

// TODO add additional methods:\
// - BodyProperties
// - FetchTextBodyValues
// - FetchHTMLBodyValues
// - FetchAllBodyValues
// - MaxBodyValueBytes
type EmailGetCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *EmailGetRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Email/get method
func (c *EmailGetCall) Do() (*EmailGetResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Email/get",
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
	getResp := &EmailGetResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *EmailGetCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *EmailGetCall) CallID(id string) {
	c.id = id
}

// Properties allows partial responses to be retrieved. Properties
// specifies the Email Properties to be returned. If Properties is nil,
// all properties will be returned. Email ID is always returned. The
// following properties are expected to be fast to return, however any
// email property is valid
// - id
// - blobId
// - threadId
// - mailboxIds
// - keywords
// - size
// - receivedAt
// - messageId
// - inReplyTo
// - sender
// - from
// - to
// - cc
// - bcc
// - replyTo
// - subject
// - sentAt
// - hasAttachment
// - preview
func (c *EmailGetCall) Properties(props ...string) {
	c.args.Properties = append(c.args.Properties, props...)
}

// IDs sets the list of IDs to return from the call
func (c *EmailGetCall) IDs(ids ...jmap.ID) {
	c.args.IDs = append(c.args.IDs, ids...)
}

// A list of properties to fetch for each EmailBodyPart returned.  If
// omitted, this defaults to:
//
//    [ "partId", "blobId", "size", "name", "type", "charset",
//      "disposition", "cid", "language", "location" ]
func (c *EmailGetCall) BodyProperties(props ...string) {
	c.args.BodyProperties = append(c.args.BodyProperties, props...)
}

// If true, the "bodyValues" property includes any "text/*" part in
// the "textBody" property.
func (c *EmailGetCall) FetchTextBodyValues(b bool) {
	c.args.FetchTextBodyValues = b
}

// If true, the "bodyValues" property includes any "text/*" part in
// the "htmlBody" property.
func (c *EmailGetCall) FetchHTMLBodyValues(b bool) {
	c.args.FetchHTMLBodyValues = b
}

// If true, the "bodyValues" property includes any "text/*" part in
// the "bodyStructure" property.
func (c *EmailGetCall) FetchAllBodyValues(b bool) {
	c.args.FetchAllBodyValues = b
}

// If greater than zero, the "value" property of any EmailBodyValue
// object returned in "bodyValues" MUST be truncated if necessary so
// it does not exceed this number of octets in size.  If 0 (the
// default), no truncation occurs.
func (c *EmailGetCall) MaxBodyValueBytes(b jmap.UnsignedInt) {
	c.args.MaxBodyValueBytes = b
}

type EmailChangesCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *EmailChangesRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Email/changes method
func (c *EmailChangesCall) Do() (*EmailChangesResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Email/changes",
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
	chResp := &EmailChangesResponse{}
	err = json.Unmarshal(raw, chResp)
	if err != nil {
		return nil, err
	}
	return chResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *EmailChangesCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *EmailChangesCall) CallID(id string) {
	c.id = id
}

// MaxChanges sets the maximum number of changes the server will return in one
// response
func (c *EmailChangesCall) MaxChanges(max jmap.UnsignedInt) {
	c.args.MaxChanges = max
}

type EmailQueryCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *EmailQueryRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Email/query method
func (c *EmailQueryCall) Do() (*EmailQueryResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Email/query",
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
	getResp := &EmailQueryResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// CallID sets the call-id within the invocation
func (c *EmailQueryCall) CallID(id string) {
	c.id = id
}

// Context sets the context to use for the calls Do method.
func (c *EmailQueryCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// Filter applies a filter to the returned results.
func (c *EmailQueryCall) Filter(filter *EmailFilter) {
	if filter.Operator == "" {
		c.args.Filter = filter.Conditions[0]
		return
	}
	c.args.Filter = filter
}

// Servers MUST support sorting by the following properties:
// - sortOrder
// - name
// Multiple calls to sort are supported for secondary sort options
func (c *EmailQueryCall) Sort(s *EmailSortComparator) {
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
func (c *EmailQueryCall) Position(p jmap.Int) {
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
func (c *EmailQueryCall) Anchor(a jmap.ID) {
	c.args.Anchor = a
}

// The index of the first result to return relative to the index of the anchor,
// if an anchor is given. This MAY be negative. For example, -1 means the Foo
// immediately preceding the anchor is the first result in the list returned
// (see below for more details).
func (c *EmailQueryCall) AnchorOffset(ao jmap.Int) {
	c.args.AnchorOffset = ao
}

// The maximum number of results to return. If null, no limit presumed. The
// server MAY choose to enforce a maximum limit argument. In this case, if a
// greater value is given (or if it is null), the limit is clamped to the
// maximum; the new limit is returned with the response so the client is aware.
// If a negative value is given, the call MUST be rejected with an
// invalidArguments error.
func (c *EmailQueryCall) Limit(l jmap.UnsignedInt) {
	c.args.Limit = l
}

// Does the client wish to know the total number of results in the query? This
// may be slow and expensive for servers to calculate, particularly with
// complex filters, so clients should take care to only request the total when
// needed.
func (c *EmailQueryCall) CanCalculateTotal(b bool) {
	c.args.CalculateTotal = b
}

// If true, Emails in the same Thread as a previous Email in the list
// (given the filter and sort order) will be removed from the list. This
// means only one Email at most will be included in the list for any
// given Thread.
func (c *EmailQueryCall) CollapseThreads(b bool) {
	c.args.CollapseThreads = b
}

type EmailQueryChangesCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *EmailQueryChangesRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Email/queryChanges method
func (c *EmailQueryChangesCall) Do() (*EmailQueryChangesResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Email/queryChanges",
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
	getResp := &EmailQueryChangesResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// CallID sets the call-id within the invocation
func (c *EmailQueryChangesCall) CallID(id string) {
	c.id = id
}

// Context sets the context to use for the calls Do method.
func (c *EmailQueryChangesCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// MaxChanges sets the maximum number of changes the server will return in one
// response
func (c *EmailQueryChangesCall) MaxChanges(max jmap.UnsignedInt) {
	c.args.MaxChanges = max
}

// The last (highest-index) id the client currently has cached from the query
// results. When there are a large number of results, in a common case, the
// client may have only downloaded and cached a small subset from the beginning
// of the results. If the sort and filter are both only on immutable
// properties, this allows the server to omit changes after this point in the
// results, which can significantly increase efficiency. If they are not
// immutable, this argument is ignored.
func (c *EmailQueryChangesCall) UpToID(id jmap.ID) {
	c.args.UpToID = id
}

// Servers MUST support sorting by the following properties:
// - sortOrder
// - name
// Multiple calls to sort are supported for secondary sort options
func (c *EmailQueryChangesCall) Sort(s *EmailSortComparator) {
	c.args.Sort = append(c.args.Sort, s)
}

// Filter applies a filter to the returned results.
func (c *EmailQueryChangesCall) Filter(filter *EmailFilter) {
	if filter.Operator == "" {
		c.args.Filter = filter.Conditions[0]
		return
	}
	c.args.Filter = filter
}

// If true, Emails in the same Thread as a previous Email in the list
// (given the filter and sort order) will be removed from the list. This
// means only one Email at most will be included in the list for any
// given Thread.
func (c *EmailQueryChangesCall) CollapseThreads(b bool) {
	c.args.CollapseThreads = b
}

type EmailSetCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *EmailSetRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Email/set method
func (c *EmailSetCall) Do() (*EmailSetResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Email/set",
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
	getResp := &EmailSetResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *EmailSetCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *EmailSetCall) CallID(id string) {
	c.id = id
}

// Create creates an email. A temporary ID must be given. The response will
// map the temporary ID to the actual ID of the created email
//
// Multiple emails may be created in a single call
func (c *EmailSetCall) Create(tempId jmap.ID, em *Email) {
	if c.args.Create == nil {
		c.args.Create = make(map[jmap.ID]*Email)
	}
	c.args.Create[tempId] = em
}

// Update updates a given email. Multiple emails may be updated in a
// single call
func (c *EmailSetCall) Update(id jmap.ID, patch jmap.Patch) {
	if c.args.Update == nil {
		c.args.Update = make(map[jmap.ID]map[string]interface{})
	}
	p := map[string]interface{}{
		patch.Path: patch.Value,
	}
	c.args.Update[id] = p
}

// Destroying an Email removes it from all Mailboxes to which it
// belonged.  To just delete an Email to trash, simply change the
// "mailboxIds" property, so it is now in the Mailbox with a "role"
// property equal to "trash", and remove all other Mailbox ids.
func (c *EmailSetCall) Destroy(ids []jmap.ID) {
	c.args.Destroy = append(c.args.Destroy, ids...)
}

type EmailCopyCall struct {
	acctFrom string
	acctTo   string
	accts    map[string]*mailAccount
	args     *EmailCopyRequest
	c        *client.Client
	ctx      context.Context
	id       string
}

// Do executes the Email/copy method
func (c *EmailCopyCall) Do() (*EmailCopyResponse, error) {
	acctTo, ok := c.accts[c.acctTo]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acctTo)
	}
	c.args.AccountID = acctTo.ID
	acctFrom, ok := c.accts[c.acctFrom]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acctFrom)
	}
	c.args.FromAccountID = acctFrom.ID
	inv := jmap.Invocation{
		Name:   "Email/copy",
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
	getResp := &EmailCopyResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *EmailCopyCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *EmailCopyCall) CallID(id string) {
	c.id = id
}

// This is a state string as returned by the Foo/get method. If
// supplied, the string must match the current state of the account
// referenced by the fromAccountId when reading the data to be copied;
// otherwise, the method will be aborted and a stateMismatch error
// returned. If null, the data will be read from the current state.
func (c *EmailCopyCall) IfFrominState(state string) {
	c.args.IfFromInState = state
}

// This is a state string as returned by the Foo/get method. If
// supplied, the string must match the current state of the account
// referenced by the accountId; otherwise, the method will be aborted
// and a stateMismatch error returned. If null, any changes will be
// applied to the current state.
func (c *EmailCopyCall) IfinState(state string) {
	c.args.IfInState = state
}

// A map of the creation id to a Foo object. The Foo object MUST
// contain an id property, which is the id (in the fromAccount) of the
// record to be copied. When creating the copy, any other properties
// included are used instead of the current value for that property on
// the original.
func (c *EmailCopyCall) Create(id jmap.ID, em *Email) {
	c.args.Create[id] = em
}

// If true, an attempt will be made to destroy the original records
// that were successfully copied: after emitting the Foo/copy response,
// but before processing the next method, the server MUST make a single
// call to Foo/set to destroy the original of each successfully copied
// record; the output of this is added to the responses as normal, to
// be returned to the client.
func (c *EmailCopyCall) OnSuccessDestroyOriginal(b bool) {
	c.args.OnSuccessDestroyOriginal = b
}

// This argument is passed on as the ifInState argument to the implicit
// Foo/set call, if made at the end of this request to destroy the
// originals that were successfully copied.
func (c *EmailCopyCall) DestroyFromIfInState(state string) {
	c.args.DestroyFromIfInState = state
}

type EmailImportCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *EmailImportRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Email/import method
func (c *EmailImportCall) Do() (*EmailImportResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Email/import",
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
	getResp := &EmailImportResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *EmailImportCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *EmailImportCall) CallID(id string) {
	c.id = id
}

// This is a state string as returned by the Foo/get method
// (representing the state of all objects of this type in the account).
// If supplied, the string must match the current state; otherwise, the
// method will be aborted and a stateMismatch error returned. If null,
// any changes will be applied to the current state.
func (c *EmailImportCall) IfInState(state string) {
	c.args.IfInState = state
}

// The email to import. blobId MUST be set and at least ONE mailbox ID must be supplied
func (c *EmailImportCall) Import(blobId jmap.ID, mailboxes []jmap.ID, keywords []string, rcvd jmap.Date, tmpId jmap.ID) error {
	if blobId == "" || mailboxes == nil {
		return fmt.Errorf("jmap-mail: invalid properties")
	}
	m := make(map[jmap.ID]bool, len(mailboxes))
	for _, id := range mailboxes {
		m[id] = true
	}
	k := make(map[string]bool, len(keywords))
	for _, keyword := range keywords {
		k[keyword] = true
	}
	if tmpId == "" {
		tmpId, _ = jmap.RandomID()
	}
	c.args.Emails[tmpId] = &EmailImport{
		BlobID:     blobId,
		MailboxIDs: m,
		Keywords:   k,
		ReceivedAt: rcvd,
	}
	return nil
}

type EmailParseCall struct {
	acct  string
	accts map[string]*mailAccount
	args  *EmailParseRequest
	c     *client.Client
	ctx   context.Context
	id    string
}

// Do executes the Email/import method
func (c *EmailParseCall) Do() (*EmailParseResponse, error) {
	acct, ok := c.accts[c.acct]
	if !ok {
		return nil, fmt.Errorf("mail: account %s not found", c.acct)
	}
	c.args.AccountID = acct.ID
	inv := jmap.Invocation{
		Name:   "Email/parse",
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
	getResp := &EmailParseResponse{}
	err = json.Unmarshal(raw, getResp)
	if err != nil {
		return nil, err
	}
	return getResp, nil
}

// Context sets the context to use for the calls Do method.
func (c *EmailParseCall) Context(ctx context.Context) {
	c.ctx = ctx
}

// CallID sets the call-id within the invocation
func (c *EmailParseCall) CallID(id string) {
	c.id = id
}

// Properties allows partial responses to be retrieved. Properties
// specifies the Email Properties to be returned. If Properties is nil,
// all properties will be returned. Email ID is always returned. The
// following properties are expected to be fast to return, however any
// email property is valid except id, mailboxIds, keywords, and receivedAt
// - blobId
// - threadId
// - size
// - messageId
// - inReplyTo
// - sender
// - from
// - to
// - cc
// - bcc
// - replyTo
// - subject
// - sentAt
// - hasAttachment
// - preview
func (c *EmailParseCall) Properties(props ...string) {
	c.args.Properties = append(c.args.Properties, props...)
}

// The ids of the blobs to parse.
func (c *EmailParseCall) BlobIDs(ids ...jmap.ID) {
	c.args.BlobIDs = append(c.args.BlobIDs, ids...)
}

// A list of properties to fetch for each EmailBodyPart returned.  If
// omitted, this defaults to:
//
//    [ "partId", "blobId", "size", "name", "type", "charset",
//      "disposition", "cid", "language", "location" ]
func (c *EmailParseCall) BodyProperties(props ...string) {
	c.args.BodyProperties = append(c.args.BodyProperties, props...)
}

// If true, the "bodyValues" property includes any "text/*" part in
// the "textBody" property.
func (c *EmailParseCall) FetchTextBodyValues(b bool) {
	c.args.FetchTextBodyValues = b
}

// If true, the "bodyValues" property includes any "text/*" part in
// the "htmlBody" property.
func (c *EmailParseCall) FetchHTMLBodyValues(b bool) {
	c.args.FetchHTMLBodyValues = b
}

// If true, the "bodyValues" property includes any "text/*" part in
// the "bodyStructure" property.
func (c *EmailParseCall) FetchAllBodyValues(b bool) {
	c.args.FetchAllBodyValues = b
}

// If greater than zero, the "value" property of any EmailBodyValue
// object returned in "bodyValues" MUST be truncated if necessary so
// it does not exceed this number of octets in size.  If 0 (the
// default), no truncation occurs.
func (c *EmailParseCall) MaxBodyValueBytes(b jmap.UnsignedInt) {
	c.args.MaxBodyValueBytes = b
}
