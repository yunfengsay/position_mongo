package tools

import (
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var (
	ACCESS_KEY = "5H53Vzt51TXbJjmOPJUfDdgMuD3WBnTzYuvbdSN0"
	SECRET_KEY = "dHCVS1xG0ue_pRNQFwMMJokyz_A34_eAU8jU1HPC"
)

func GetQiniuToken() (upToken string) {
	bucket := "location"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(ACCESS_KEY, SECRET_KEY)
	upToken = putPolicy.UploadToken(mac)
	return
}
