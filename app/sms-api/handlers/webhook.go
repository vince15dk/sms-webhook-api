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
)

type WebHook struct {
	build string
}

// Param takes url parameters into map
func Param(r *http.Request, key string) string {
	m := httptreemux.ContextParams(r.Context())
	return m[key]
}

// Unmarshal unmarshals json into empty interface
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

// ParsingJsonData opens config, and unmarshals json structure into DepUsers
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
	var sender string
	h := generateHeader(&http.Header{})
	//u, err := ParsingJsonData("/Users/nhn/Desktop/Linux/Go/sms-webhook-api/app/sms-api/external/dep_users.json")
	u, err := ParsingJsonData("/config/dep_users.json")
	if err != nil{
		return err
	}

	for _, v := range u.DepGroup{
		if Param(r, "dep") == v.GroupName{
			smsurl = fmt.Sprintf("%s/%s/%s", url, v.AppKey,"sender/sms")
			h.Set("X-Secret-Key", v.SecretKey)
			for _, n := range v.Users {
				ListRecipients = append(ListRecipients, scheme.RecipientList{
					RecipientNo: n.PhoneNo,
					CountryCode: "82",
				})
			}
			sender = v.Sender
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
