package handlers

import (
	"context"
	"github.com/vince15dk/sms-webhook-api/foundation/api"
	"io/ioutil"
	"log"
	"net/http"
)

type WebHook struct {
	build string
}

func (wh WebHook)sendAPItoSMS(ctx context.Context, w http.ResponseWriter, r *http.Request)error{

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil{
		return err
	}
	defer r.Body.Close()
	log.Printf("\n%v", string(bytes))
	return api.Respond(ctx, w, nil, http.StatusOK)
}
