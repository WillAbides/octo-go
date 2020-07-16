package octo

import (
	"bytes"
	"encoding/json"

	"github.com/willabides/octo-go/components"
)

// UnmarshalJSON wraps json in [] before unmarshalling it ... if necessary
func (r *ReposGetContentResponseBody) UnmarshalJSON(p []byte) error {
	var val []components.ContentFile
	p = bytes.TrimSpace(p)
	if !bytes.HasPrefix(p, []byte("[")) {
		p = append([]byte("["), p...)
	}
	if !bytes.HasSuffix(p, []byte("]")) {
		p = append(p, []byte("]")...)
	}

	err := json.Unmarshal(p, &val)
	if err != nil {
		return err
	}
	*r = val
	return nil
}
