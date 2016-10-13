package io

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"math"
)

// ReadXML is a helper function to reads remote xml responses data
func ReadXML(body io.ReadCloser, result interface{}) error {
	data, err := ioutil.ReadAll(io.LimitReader(body, math.MaxInt64))

	if err != nil {
		return err
	}

	if e := xml.Unmarshal(data, &result); e != nil {
		return e
	}

	return nil
}
