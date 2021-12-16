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

// sendAPItoSMS post request to SMS API server
func (wh WebHook) sendAPItoSMS(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var g interface{}
	var ListRecipients []scheme.RecipientList
	var sms scheme.SMSRequest
	var smsurl string
	var sender string

	// generating a new header
	h := generateHeader(&http.Header{})

	//u, err := ParsingJsonData("/Users/nhn/Desktop/Linux/Go/sms-webhook-api/app/sms-api/external/dep_users.json")
	// Read config file
	u, err := ParsingJsonData("/config/dep_users.json")
	if err != nil{
		return err
	}

	// v1/:dep/:groups/sms
	for _, v := range u.DepGroup{
		if Param(r, "dep") == v.GroupName{
			// generates a full SMS API url
			smsurl = fmt.Sprintf("%s/%s/%s", url, v.AppKey,"sender/sms")
			// updates secret
			h.Set("X-Secret-Key", v.SecretKey)
			// add users from DepGroup whose GroupName is matched with :dep
			for _, n := range v.Users {
				ListRecipients = append(ListRecipients, scheme.RecipientList{
					RecipientNo: n.PhoneNo,
					CountryCode: "82",
				})
			}
			sender = v.Sender
		}
	}

	// parsing OOS struct to SMS Api request body struct
	// v1/:dep/:groups/sms
	switch Param(r, "groups") {
	case "grafana":
		s, err := Unmarshal(r, &scheme.GrafanaLog{})
		if err != nil {
			return err
		}
		g = s
		// Parsing grafanaLog's message into SMS API body struct originally [:81] for fit in size
		sms = scheme.SMSRequest{
			Body:          fmt.Sprintf("%s", g.(*scheme.GrafanaLog).Message),
			SendNo:        sender,
			RecipientList: ListRecipients,
		}

	case "argocd":
		s, err := Unmarshal(r, &scheme.Argocd{})
		if err != nil {
			return err
		}
		g = s
		// Parsing Argocd's message into SMS API body struct
		sms = scheme.SMSRequest{
			Body:          fmt.Sprintf("%s", g.(*scheme.Argocd).Text),
			SendNo:        sender,
			RecipientList: ListRecipients,
		}
	}

	// Call Post Handler
	PostHandlerFunc(smsurl, sms, h)

	return api.Respond(ctx, w, nil, http.StatusOK)
}

// generateHeader presets Content-Type to application/json
func generateHeader(h *http.Header) *http.Header {
	h.Set("Content-Type", "application/json")
	return h
}
