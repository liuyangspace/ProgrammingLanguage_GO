package format

import (
	"io"
	"encoding/gob"
)

/*
Gob 是 Go 自己的以二进制形式序列化和反序列化程序数据的格式。
Gob 通常用于远程方法调用（RPCs,例如两个用 Go 写的服务之间的通信）参数和结果的传输，以及应用程序和机器之间的数据传输。


 */
type GobEncoder struct {
	gob.Encoder
}
func GobNewEncoder(w io.Writer) *gob.Encoder { return gob.NewEncoder(w) }
func (enc *Encoder) GobEncode(e interface{}) error { return enc.Encode(e) }