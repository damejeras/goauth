# goauth
goauth is a JWT issuing microservice

You will need RS256 keypair. Here is how to generate key-pair:
```bash
cd keys
ssh-keygen -t rsa -b 4096 -m PEM -f jwtRS256.key
# Don't add passphrase
openssl rsa -in jwtRS256.key -pubout -outform PEM -out jwtRS256.key.pub
cat jwtRS256.key
cat jwtRS256.key.pub
```
