package constant

import (
	"github.com/wenzhenxi/gorsa"
	"log"
)

const PrivateKey = `-----BEGIN PRIVATE KEY-----
MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQCwm5r88cGvnP7+
/rK/nsytVkVjs6EKblg6qjCwpDOu7/7t0Rk7I8DDCS2yAc46987yzz3i+waNhhS9
gcU63wqTXhtRzzjpHBXMzrKwVM+YDPSmEK+h5gd9oYccnzrG5oqnpmHsUyA6MwPD
UwydnnmusFvGTqxbCocJtcCX0MvMEIIxnTIN8bi2TrzavhQOp88nC12WIf7PSHQY
8s5kCDWMFElbRUxb8xjcERtdKit2uxHbJnZwmK3/uCWAtzON4KbhYhLT7d+TQ+Zj
ERjB3ipNZF5kNZlNnXzq1lkM+YWtfkJsPV4xDwpIOXe+H9wilhJ9jhSyRMT2kx/V
W5KwAs3/AgMBAAECggEAXVeAzRvBsN2bcI6RTvHRX/Pu7Mkt3mN/wlHdNh8zKcyy
wmL3oOg+o1agjiqiI0yCVpbuMJ9YUUiCgb6qzXjBEC525gBOJfawFvVQQB4zMjv2
uiarjTgU0XOkA56s5B9JPV6U9uDgcaIw5PVp/4yY55lSiDdgpSWlRZ6IUjUK4akf
QR8cTIgHT7ZcpV7wsm4hOq8OyDNtZVLU4V+rTf9VNOTdI9tezmdEB59MPwmi3dUq
PYlfuRG9w0wtjm6xYHNGNbxrkKy7n2iCI2bMrXnl7E5CC5lclA+Dq8tUZUJMXOLj
X6HpsDmpPCbnhqiZa1Uc4WcVMAhmGkxTbk7aF5ahEQKBgQDWWNvoB52gD8DX6WW0
8eHATadMjxikfUvlzmXevdz3rfdyTewjMfTcmtHsEtxMDcVuaGR80ORPm5gytram
RpcRrBKsyyJVMGkux9BGcdHTcLvjvyVdHaWVNssydVhHTToRmpMTbHTGmyBtIyD3
lv7s9Hlg4IJPjYaWGexOjdKaTQKBgQDS7VP/1tkZ2Y1+GGpnhxcuXzwdvrlfk7A8
2AjYEzyOf7Aua4RkQw/aBoozmpzZyD2USsbX+lz5gADIZC+HW4DvPA+ErtkE4Qcj
ipKi5EmCaBcZ08q/BbUf3PdMQEf3apMXmRZKVpqUKZ2iKbWt/MhpYOe31Ok+gcKp
jUy0QyzXewKBgQDPzyKUmB2Xuu4p5hw3C3Lk7+RM80be4qMcz7gv9T8v2bfXV/ii
PqK45lizWa2/BnLtEcSqxhMzoxQxxrbaCgIYng0U/XPxx5eSRSKFyX+2UMu4h+7/
keig+5iQ+Pd3WUX9HzGMWU/15RvQ84fYwsmwjWzIIaH04q8jA0r49cLC2QKBgHN9
lzzCOuU4hYzBeaFuCndNlfEDeSSoDRuDfnQxNDsiFymTnU34mjla01kmEzP7vBvy
RaYEY9wJBmSy1oNayuIzFJwl0JYnUG1ZdqjSLp3EuMFfgefODDqCFrYdjEoghKGr
JNVXThRKJRGhnOTDQEjS0o2JepzLF/cWcyVQksD5AoGAIS1FBDMqmi5G1Q4p1M9s
Ug0pr6oecsR2GVHCfQujCkoq2oCodtI8GiDsYhnFWoMtApXdM8cGatlUZ7gyPPZf
mxvKd7oOrPTJMgBP/GTZcLqmuICASNCHbQVe12Ranfz7snuGS9mOQx1tnM2TJ5B5
euJ5W/Vnj5SVFmCEdsJ2f4g=
-----END PRIVATE KEY-----`

const PublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAsJua/PHBr5z+/v6yv57M
rVZFY7OhCm5YOqowsKQzru/+7dEZOyPAwwktsgHOOvfO8s894vsGjYYUvYHFOt8K
k14bUc846RwVzM6ysFTPmAz0phCvoeYHfaGHHJ86xuaKp6Zh7FMgOjMDw1MMnZ55
rrBbxk6sWwqHCbXAl9DLzBCCMZ0yDfG4tk682r4UDqfPJwtdliH+z0h0GPLOZAg1
jBRJW0VMW/MY3BEbXSordrsR2yZ2cJit/7glgLczjeCm4WIS0+3fk0PmYxEYwd4q
TWReZDWZTZ186tZZDPmFrX5CbD1eMQ8KSDl3vh/cIpYSfY4UskTE9pMf1VuSsALN
/wIDAQAB
-----END PUBLIC KEY-----`

func init() {
	if err := gorsa.RSA.SetPublicKey(PublicKey); err != nil {
		log.Fatalln(`set public key :`, err)
	}
	if err := gorsa.RSA.SetPrivateKey(PrivateKey); err != nil {
		log.Fatalln(`set private key :`, err)
	}
}