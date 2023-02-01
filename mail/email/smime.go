package email

import "git.sr.ht/~rockorager/go-jmap"

const SMIMEVerify jmap.URI = "urn:ietf:params:jmap:smimeverify"

type smimeVerify struct{}

func (s *smimeVerify) URI() jmap.URI { return SMIMEVerify }

func (s *smimeVerify) New() jmap.Capability { return &smimeVerify{} }
