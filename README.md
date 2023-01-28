# go-jmap

A JMAP client library. go-jmap is a library for interacting with JMAP servers.
It includes a client for making requests, and data structures for the Core and
Mail specifications.

Note: this library started as a fork of [github.com/foxcpp/go-jmap](https://github.com/foxcpp/go-jmap)
It has since undergone massive restructuring, and it only loosely based on the
original project.

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
  - [ ] Request-level errors
  - [ ] Method-level errors
  - [ ] Set-level errors

  - [x] Core/echo

  - [x] Blob/Downloading
  - [x] Blob/Uploading
  - [ ] Blob/Copy method

  - [ ] Push
    - [ ] StateChange structure
    - [ ] PushSubscription structure
    - [ ] PushSubscription/get
    - [ ] PushSubscription/set
    - [ ] Event Source

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
