## Generate certs

```bash
mkdir certs
openssl ecparam -genkey -name secp384r1 -out certs/tls.key
openssl req -new -x509 -sha256 -key certs/tls.key -out certs/tls.crt -days 3650
```

## Convert the cert to k8s format
```bash
cat certs/tls.key | base64 -w0
cat certs/tls.crt | base64 -w0
```

## Save into deployment template

```bash
data:
  tls.crt: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tC...
  tls.key: LS0tLS1CRUdJTiBFQyBQQVJBTUVURVJTLS0tL...
```