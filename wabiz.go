package nwabiz

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type InitOpt struct {
	BaseUrl     string
	Username    string
	Password    string
	Timeout     int64
	InsecureSSL bool
}

func NewWhatsAppBiz(opt InitOpt) *WhatsAppBiz {
	// Generate base 64 credential
	rc := opt.Username + ":" + opt.Password
	cred := base64.StdEncoding.EncodeToString([]byte(rc))

	// Create REST Client
	timeout := time.Duration(opt.Timeout) * time.Millisecond
	cl := http.Client{
		Timeout: timeout,
	}

	// If skip ssl check true
	if opt.InsecureSSL {
		cl.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	return &WhatsAppBiz{
		BaseUrl:    opt.BaseUrl,
		Credential: cred,
		Client:     cl,
	}
}

type WhatsAppBiz struct {
	BaseUrl         string
	Credential      string
	AccessToken     string
	AccessExpiredAt int64
	Client          http.Client
}

func (c *WhatsAppBiz) IsAccessValid(t time.Time) bool {
	// If access token is empty, return false
	if c.AccessToken == "" {
		return false
	}

	// If access expired at is less than now, return false
	if c.AccessExpiredAt < t.Unix() {
		return false
	}

	return true
}

func (c *WhatsAppBiz) Login() error {
	// Create Request
	url := c.BaseUrl + "/users/login"

	payload := strings.NewReader("")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return NewUnhandledError(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+c.Credential)

	// Send Request
	res, err := c.Client.Do(req)
	if err != nil {
		return NewUnhandledError(err)
	}

	defer res.Body.Close()

	// Parse body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return NewUnhandledError(err)

	}
	if res.StatusCode == http.StatusUnauthorized {
		return NewError(InvalidCredentialsError)
	}

	if res.StatusCode != http.StatusOK {
		return NewUnhandledError(fmt.Errorf("unexpected http status: %d", res.StatusCode))
	}

	var respBody LoginResp
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return NewUnhandledError(err)
	}

	// Get access token
	if len(respBody.Users) == 0 {
		return NewError(EmptyDataError)
	}
	access := respBody.Users[0]

	// Parse expire
	expiredAt, err := time.Parse(TimeRespLayout, access.ExpiresAfter)
	if err != nil {
		return NewUnhandledError(err)
	}

	// Set Access Token
	c.AccessToken = access.Token
	c.AccessExpiredAt = expiredAt.Unix()

	return nil
}

func (c *WhatsAppBiz) CheckContact(msisdn string) (*ContactsResp, error) {
	// Check if access valid
	if !c.IsAccessValid(time.Now()) {
		err := c.Login()
		if err != nil {
			return nil, err
		}
	}

	// Create Request
	url := c.BaseUrl + "/contacts"

	reqBody := CheckContactsReq{
		Blocking:   "wait",
		Contacts:   []string{msisdn},
		ForceCheck: false,
	}
	buffer, err := json.Marshal(reqBody)
	if err != nil {
		return nil, NewUnhandledError(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(buffer))
	if err != nil {
		return nil, NewUnhandledError(err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.AccessToken)

	// Send Request
	res, err := c.Client.Do(req)
	if err != nil {
		return nil, NewUnhandledError(err)
	}

	defer res.Body.Close()

	// Parse body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, NewUnhandledError(err)

	}
	if res.StatusCode == http.StatusUnauthorized {
		return nil, NewError(InvalidCredentialsError)
	}

	if res.StatusCode != http.StatusOK {
		return nil, NewUnhandledError(fmt.Errorf("unexpected http status: %d", res.StatusCode))
	}

	var respBody CheckContactsResp
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		return nil, NewUnhandledError(err)
	}

	// Get access token
	if len(respBody.Contacts) == 0 {
		return nil, NewError(EmptyDataError)
	}

	contact := respBody.Contacts[0]

	return &contact, nil
}
