package format

import (
	"encoding/json"
	"io"
)

/*

JSON 与 Go 类型对应如下：
	bool 对应 JSON 的 booleans
	float64 对应 JSON 的 numbers
	string 对应 JSON 的 strings
	nil 对应 JSON 的 null
不是所有的数据都可以编码为 JSON 类型：只有验证通过的数据结构才能被编码：
	JSON 对象只支持字符串类型的 key；要编码一个 Go map 类型，map 必须是 map[string]T（T是 json 包中支持的任何类型）
	Channel，复杂类型和函数类型不能被编码
	不支持循环数据结构；它将引起序列化进入一个无限循环
	指针可以被编码，实际上是对指针指向的值进行编码（或者指针是 nil）
 */

func Marshal(v interface{}) ([]byte, error) { return json.Marshal(v) } // json encoder
type Encoder struct {
	json.Encoder
}
func NewEncoder(w io.Writer) *json.Encoder { return json.NewEncoder(w) }
func (enc *Encoder) Encode(v interface{}) error { return enc.Encode(v) }

func Unmarshal(data []byte, v interface{}) error { return json.Unmarshal(data,v) } // json decoder
type Decoder struct {
	json.Decoder
}
func NewDecoder(r io.Reader) *json.Decoder { return json.NewDecoder(r) }
func (dec *Decoder) Decode(v interface{}) error { return dec.Decode(v)}