# `SMS-Webhook-api`

### Execute it in local env
```bash
make run
```
<br/>

### Webhook Post Request URL

```bash
Post :8080/{dep}/{groups}/sms
```
<br/>

### For grafana alert webhook usage
* http://{sms-webhook-api}/v1/srep/grafana/sms

<br/>

### For argocd notification webhook usage
* http://{sms-webhook-api}/v1/srep/argocd/sms

<br/>

### Where to store config
```bash
vi /config/dep_users.json
    {
      "depGroup": [
        {
          "sender": "xx",
          "appKey": "xx",
          "secretKey": "xx",
          "groupName": "{dep}",
          "users": [
            {
              "name": "rabbit",
              "phoneNo": "123",
              "email": "rabbit@animal.com"
            },
            {
              "name": "mouse",
              "phoneNo": "123",
              "email": "mouse@animal.com"
            }
          ]
        }
      ]
    }
```

