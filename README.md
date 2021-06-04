### Client for Go file server

#### Install

`$ go get -u github.com/zikwall/fsclient`

#### Install server

- Visit: https://github.com/zikwall/go-fileserver

#### Usage

```go
func main() {
	fsclient, err := WithCopyFsClient(FsClient{
		Uri:        "http://localhost:1337/",
		SecureType: TypeToken,
		TokenType:  requesters.TokenTypeQuery,
		Token:      "changemeplease123",
		User:       "qwx1337",
		Password:   "123456",
	})

	if err != nil {
		log.Fatal(err)
	}

	f1, _ := os.Open("test.txt")
	f2, _ := os.Open("test2.txt")

	defer func() {
		f1.Close()
		f2.Close()
	}()

	if err := fsclient.SendFile(context.Background(), f1, f2); err != nil {
		log.Fatal(err)
	}
}
```
