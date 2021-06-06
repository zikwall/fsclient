### Client for Go file server

#### Install

`$ go get -u github.com/zikwall/fsclient@v0.0.2`

#### Install server

- Visit: https://github.com/zikwall/go-fileserver

### Usage

#### Create client

```go
client, err := fsclient.WithCopyFsClient(fsclient.FsClient{
	// fileserver host
	Uri:        "http://localhost:1337/",
	SecureType: fsclient.TypeToken,
	// for token auth or JWT
	TokenType:  requesters.TokenTypeQuery,
	Token:      "changemeplease123",
	// for base auth
	User:       "qwx1337",
	Password:   "123456",
})
```

#### Call API for send files

```go
if err := client.SendFile(context.Background(), file, fileAnother); err != nil {
	log.Fatal(err)
}
```
