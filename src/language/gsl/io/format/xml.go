package format

import (
	"encoding/xml"
	"io"
)

/*

 */

type Token interface{
	xml.Token
}
func XmlNewDecoder(r io.Reader) *xml.Decoder { return xml.NewDecoder(r)}
func (d *Decoder) Token() (Token, error) { return d.Token() }
