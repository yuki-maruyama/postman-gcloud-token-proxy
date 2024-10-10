# postman-gcloud-token-proxy

## HOW TO USE

build
```bash
docker build -t postman-gcloud-token-proxy .
```

run
```bash
docker run -d -p 8080:8080 postman-gcloud-token-proxy
```

postman pre script
```javascript
var tokenDate = new Date(2010,1,1);
var tokenTimestamp = pm.collectionVariables.get("GCLOUD_TOKEN_TIMESTAMP");
if(tokenTimestamp){
  tokenDate = Date.parse(tokenTimestamp);
}

if((new Date() - tokenDate) >= 36000) 
{
   pm.sendRequest({
      url:  'http://localhost:8080', 
      method: 'GET',
  }, function (err, res) {
        pm.collectionVariables.set("GCLOUD_TOKEN", res.text());
        pm.collectionVariables.set("GCLOUD_TOKEN_TIMESTAMP", new Date());
  });
}
```