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
)

const (
	smsurl = "https://api-sms.cloud.toast.com/sms/v3.0/appKeys/g00TSSkVtzDp4qwX/sender/sms"
	key    = "UhUXiDfu"
	sender = "01045745984"
)

var (
	recipients = []string{"01045745984"}
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

func (wh WebHook) sendAPItoSMS(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	var g interface{}
	h := generateHeader(&http.Header{})
	var ListRecipients []scheme.RecipientList
	var sms scheme.SMSRequest

	for _, v := range recipients {
		ListRecipients = append(ListRecipients, scheme.RecipientList{
			RecipientNo: v,
			CountryCode: "82",
		})
	}

	switch Param(r, "groups") {
	case "grafana":
		s, err := Unmarshal(r, &scheme.GrafanaLog{})
		if err != nil {
			return err
		}
		g = s
		gsms := scheme.SMSRequest{
			Body:          fmt.Sprintf("%s", g.(*scheme.GrafanaLog).Message[:81]),
			SendNo:        sender,
			RecipientList: ListRecipients,
		}
		sms = gsms
	case "argocd":
	}

	PostHandlerFunc(smsurl, sms, h)

	return api.Respond(ctx, w, nil, http.StatusOK)
}

func generateHeader(h *http.Header) *http.Header {
	h.Set("Content-Type", "application/json")
	h.Set("X-Secret-Key", key)
	return h
}
