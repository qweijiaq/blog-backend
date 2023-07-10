package qiniu

import (
	"backend/config"
	"backend/global"
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"time"
)

func getToken(q config.QiNiu) string {
	accessKey := q.AccessKey
	secretKey := q.SecretKey
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func getCfg(q config.QiNiu) storage.Config {
	cfg := storage.Config{}
	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false
	return cfg
}

func UploadImage(data []byte, imageName string, prefix string) (filePath string, err error) {
	//if !global.Config.QiNiu.Enable {
	//	return "", errors.New("请启用七牛云存储")
	//}
	q := global.Config.QiNiu
	if q.AccessKey == "" || q.SecretKey == "" {
		errors.New("请配置AccessKey和SecretKey")
	}
	if float64(len(data))/1024/1024 > q.Size {
		errors.New("文件超过设定大小")
	}
	upToken := getToken(q)
	cfg := getCfg(q)

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))

	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s__%s", prefix, now, imageName)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", q.CDN, ret.Key), nil
}
