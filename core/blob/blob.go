package blob

import "git.sr.ht/~rockorager/go-jmap"

func init() {
	jmap.RegisterMethod("Blob/copy", newCopyResponse)
}
