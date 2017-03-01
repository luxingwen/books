package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func EncodePasswd(str string) string {
	t := md5.New()
	io.WriteString(t, str)
	io.WriteString(t, `这里的每一句话都是加密的盐,我也不知道写什么好,随便写点吧!!`)
	return hex.EncodeToString(t.Sum(nil))
}
