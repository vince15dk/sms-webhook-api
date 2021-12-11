package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/vince15dk/sms-webhook-api/business/data/scheme"
	"github.com/vince15dk/sms-webhook-api/foundation/api"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	url    = "https://api-sms.cloud.toast.com/sms/v3.0/appKeys"
	sender    = "01045745984"
)

type WebHook struct {
	build string
}

func Param(r *http.Request, key string) string {
	m := httptreemux.ContextParams(r.Context())
	return m[key]
}

func Unmarshal(r *http.Request, g interface{}) (interface{}, error) {
	//json.NewDecoder(r.Body).Decode(g)
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	err = json.Unmarshal(bytes, g)
	return g, nil
}

func ParsingJsonData(name string) (*scheme.DepUsers, error) {
	jsonFile, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	var u scheme.DepUsers
	bytes, err := ioutil.ReadAll(jsonFile)
	err = json.Unmarshal(bytes, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (wh WebHook) sendAPItoSMS(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var g interface{}
	var ListRecipients []scheme.RecipientList
	var sms scheme.SMSRequest
	var smsurl string

	h := generateHeader(&http.Header{})
	//u, err := ParsingJsonData("/Users/nhn/Desktop/Linux/Go/sms-webhook-api/app/sms-api/external/dep_users.json")
	u, err := ParsingJsonData("/service/dep_users.json")
	if err != nil{
		return err
	}
	if Param(r, "dep") == u.DepGroup{
		smsurl = fmt.Sprintf("%s/%s/%s", url, u.AppKey,"sender/sms")
		h.Set("X-Secret-Key", u.SecretKey)
		for _, v := range u.Users {
			ListRecipients = append(ListRecipients, scheme.RecipientList{
				RecipientNo: v.PhoneNo,
				CountryCode: "82",
			})
		}
	}

	switch Param(r, "groups") {
	case "grafana":
		s, err := Unmarshal(r, &scheme.GrafanaLog{})
		if err != nil {
			return err
		}
		g = s
		sms = scheme.SMSRequest{
			Body:          fmt.Sprintf("%s", g.(*scheme.GrafanaLog).Message[:81]),
			SendNo:        sender,
			RecipientList: ListRecipients,
		}

	case "argocd":
		s, err := Unmarshal(r, &scheme.Argocd{})
		if err != nil {
			return err
		}
		g = s
		sms = scheme.SMSRequest{
			Body:          fmt.Sprintf("%s", g.(*scheme.Argocd).Text),
			SendNo:        sender,
			RecipientList: ListRecipients,
		}
	}

	PostHandlerFunc(smsurl, sms, h)

	return api.Respond(ctx, w, nil, http.StatusOK)
}

func generateHeader(h *http.Header) *http.Header {
	h.Set("Content-Type", "application/json")
	return h
}
