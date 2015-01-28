package caldav

import (
	"github.com/taviti/caldav-go/icalendar"
	"github.com/taviti/caldav-go/utils"
	"github.com/taviti/caldav-go/webdav"
	"io/ioutil"
)

// a WebDAV response object
type Response webdav.Response

// downcasts the response to the WebDAV interface
func (r *Response) WebDAV() *webdav.Response {
	return (*webdav.Response)(r)
}

// decodes a CalDAV iCalendar response into the provided interface
func (r *Response) Decode(into interface{}) error {
	if body := r.Body; body == nil {
		return nil
	} else if encoded, err := ioutil.ReadAll(body); err != nil {
		return utils.NewError(r.Decode, "unable to read response body", r, err)
	} else if err := icalendar.Unmarshal(string(encoded), into); err != nil {
		return utils.NewError(r.Decode, "unable to decode response body", r, err)
	} else {
		return nil
	}
}

// creates a new WebDAV response object
func NewResponse(response *webdav.Response) *Response {
	return (*Response)(response)
}