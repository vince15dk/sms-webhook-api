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

### For grafana alert webhook usage
* http://{sms-webhook-api}/v1/srep/grafana/sms

### For argocd notification webhook usage
* http://{sms-webhook-api}/v1/srep/argocd/sms

### Where to store config
```bash
vi /config/dep_users.json
    {
      "depGroup": [
        {
          "sender": "xx",
          "appKey": "xx",
          "secretKey": "xx",
          "groupName": "srep",
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

