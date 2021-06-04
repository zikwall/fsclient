### Client for Go file server

#### Install

`$ go get -u github.com/zikwall/fsclient@v0.0.1`

#### Install server

- Visit: https://github.com/zikwall/go-fileserver

### Usage

```go
func main() {
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

	if err != nil {
		log.Fatal(err)
	}

	f1, _ := os.Open("test.txt")
	f2, _ := os.Open("test2.txt")

	// .. check errors
	
	defer func() {
		_ = f1.Close()
		_ = f2.Close()
	}()

	if err := client.SendFile(context.Background(), f1, f2); err != nil {
		log.Fatal(err)
	}
}
```
