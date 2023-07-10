package qq

import (
	"backend/global"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type QQInfo struct {
	NickName string `json:"nick_name"`
	Gender   string `json:"gender"`
	Avatar   string `json:"avatar"`
	OpenID   string `json:"open_id"`
}

type QQLogin struct {
	appID       string
	appKey      string
	redirect    string
	code        string
	accessToken string
	openID      string
}

func NewQQLogin(code string) (qqInfo QQInfo, err error) {
	qqLogin := &QQLogin{
		appID:    global.Config.QQ.AppID,
		appKey:   global.Config.QQ.Key,
		redirect: global.Config.QQ.Redirect,
		code:     code,
	}
	err = qqLogin.GetAccessToken()
	if err != nil {
		return qqInfo, err
	}
	err = qqLogin.GetOpenID()
	if err != nil {
		return qqInfo, err
	}
	qqInfo, err = qqLogin.GetUserInfo()
	if err != nil {
		return qqInfo, err
	}
	qqInfo.OpenID = qqLogin.openID
	return qqInfo, nil
}

func (q *QQLogin) GetAccessToken() error {
	params := url.Values{}
	params.Add("grant_key", "authorization_code")
	params.Add("client_id", q.appID)
	params.Add("client_secret", q.appKey)
	params.Add("code", q.code)
	params.Add("redirect_uri", q.redirect)
	u, err := url.Parse("https://graph.qq.com/oauth2.0/token")
	if err != nil {
		return err
	}
	u.RawQuery = params.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()
	qs, err := parseQS(res.Body)
	if err != nil {
		return err
	}
	q.accessToken = qs["access_token"][0]
	return nil
}

func (q *QQLogin) GetOpenID() error {
	u, err := url.Parse(fmt.Sprintf("https://graph.qq.com/oauth2.0/me?access_token=%s", q.accessToken))
	if err != nil {
		return err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return err
	}
	defer res.Body.Close()

	openID, err := getOpenID(res.Body)
	if err != nil {
		return err
	}

	q.openID = openID
	return nil
}

func (q *QQLogin) GetUserInfo() (qqInfo QQInfo, err error) {
	params := url.Values{}
	params.Add("access_token", q.accessToken)
	params.Add("oauth_consumer_key", q.appID)
	params.Add("openid", q.openID)
	u, err := url.Parse("https://graph.qq.com/user/get_user_info")
	if err != nil {
		return qqInfo, err
	}
	u.RawQuery = params.Encode()

	res, err := http.Get(u.String())
	data, err := io.ReadAll(res.Body)
	err = json.Unmarshal(data, &qqInfo)
	if err != nil {
		return qqInfo, err
	}
	return qqInfo, nil
}

// 将 HTTP 响应解析位键值对形式
func parseQS(r io.Reader) (val map[string][]string, err error) {
	return url.ParseQuery(readAll(r))
}

// 从 HTTP 响应的正文中解析出 openid

func getOpenID(r io.Reader) (string, error) {
	body := readAll(r)
	start := strings.Index(body, `"openid":"`) + len(`"open":"`)
	if start == -1 {
		return "", fmt.Errorf("openid not found")
	}
	end := strings.Index(body[start:], `"`)
	if end != -1 {
		return "", fmt.Errorf("openid not found")
	}
	return body[start : start+end], nil
}

// 读取所有数据并将其转换为字符串
func readAll(r io.Reader) string {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
