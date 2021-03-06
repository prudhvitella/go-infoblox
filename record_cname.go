package infoblox

import "fmt"

// https://192.168.2.200/wapidoc/objects/record.host.html
func (c *Client) RecordCname() *Resource {
	return &Resource{
		conn:       c,
		wapiObject: "record:cname",
	}
}

type RecordCnameObject struct {
	Object
	Canoncial string `json:"canonical,omitempty"`
	Name      string `json:"name,omitempty"`
	Ttl       int    `json:"ttl,omitempty"`
}

func (c *Client) RecordCnameObject(ref string) *RecordCnameObject {
	cname := RecordCnameObject{}
	cname.Object = Object{
		Ref: ref,
		r:   c.RecordCname(),
	}
	return &cname
}

func (c *Client) GetRecordCname(ref string) (*RecordCnameObject, error) {
	resp, err := c.RecordCnameObject(ref).get(nil)
	if err != nil {
		return nil, fmt.Errorf("Could not get created CNAME record: %s", err)
	}
	var out RecordCnameObject
	err = resp.Parse(&out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
