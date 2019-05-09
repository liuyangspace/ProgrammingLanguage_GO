package encryption

import (
	"hash"
	"crypto/sha1"
)

/*
#Go 的标准库为加密提供了超过 30 个的包：
	hash 包：实现了 adler32、crc32、crc64 和 fnv 校验；
	crypto 包：实现了其它的 hash 算法，比如 md4、md5、sha1 等。以及完整地实现了 aes、blowfish、rc4、rsa、xtea 等加密算法。
 */

func New() hash.Hash { return sha1.New() }
