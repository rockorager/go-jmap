# go-jmap

Note: this is an actively maintained fork of [github.com/foxcpp/go-jmap](https://github.com/foxcpp/go-jmap)

JMAP Core client Go library.

## Status
Reference: https://jmap.io/spec-core.html List below is nowhere complete.

* [x]  Fundamental types

	* [x]  Int
	* [x]  UnsignedInt
	* [x]  Date
	* [x]  UTCDate
	* [x]  Id

* [x]  Autodiscovery

* [ ]  Structures for base JMAP Core objects

	* [x]  Session
	* [x]  Account
	* [x]  Core Capabilities
	* [x]  Request-level errors
	* [x]  Invocation

	* [x]  Method-level errors
	* [x]  Decode Invocation "subtypes" using type -> decoder mapping.
	* [x]  Request
	* [x]  Response
	* [ ]  Set-level errors

* [x]  Core request objects

	* [x]  Object/get
	* [x]  Object/changes
	* [x]  Object/set
	* [x]  Object/copy
	* [x]  Object/query
	* [x]  Object/queryChanges

	* [x]  Core/echo

* [ ]  Binary data I/O

	* [x]  Downloading
	* [x]  Uploading
	* [ ]  Blob/Copy method

* [ ]  Push

	* [ ]  StateChange structure
	* [ ]  PushSubscription structure

	* [ ]  PushSubscription/get
	* [ ]  PushSubscription/set
	* [ ]  Event Source

* [x]  Client interface

	* [x]  Get Session object
	* [x]  Interface to send request objects directly
	* [x]  "Call batch builder" interface
	* [x]  Binary I/O interface

* [ ]  Server interface

	* [ ]  Base backend interface
	* [ ]  Binary I/O backend interface
	* [ ]  Method call dispatching



## Related standards

- [RFC 8620]
  JSON Meta Application Protocol
- [RFC 8621]
  JMAP for Mail
- [RFC 2782], [RFC 6186], [RFC 6764]
  DNS-based service auto-discovery.
- [RFC 5785]
  .well-known URIs
- [RFC 7807]
  Problem details for HTTP APIs
- [RFC 6901]
  JavaScript Object Notation (JSON) Pointer

## License 

The code is under MIT license.

Documentation strings for most of the protocol objects are taken from (or based
on) contents of draft-ietf-jmap-core-17 and is subject to the IETF Trust
Provisions. See https://trustee.ietf.org/trust-legal-provisions.html for
details. See included draft-ietf-jmap-core-17.txt for related copyright
notices.


[RFC 8620]: https://tools.ietf.org/html/rfc8620
[RFC 8621]: https://tools.ietf.org/html/rfc8621
[RFC 2782]: https://tools.ietf.org/html/rfc2782
[RFC 6186]: https://tools.ietf.org/html/rfc6186
[RFC 6764]: https://tools.ietf.org/html/rfc6764
[RFC 5785]: https://tools.ietf.org/html/rfc5785
[RFC 7807]: https://tools.ietf.org/html/rfc7807
[RFC 6901]: https://tools.ietf.org/html/rfc6901
