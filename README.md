gourl
=====
[![wercker status](https://app.wercker.com/status/bb67622f0e5529877beed36a65bcc400/s "wercker status")](https://app.wercker.com/project/bykey/bb67622f0e5529877beed36a65bcc400)

URL Parser for Go

#### How to install:
```bash
go get github.com/asalvador/gourl
```
#### How to use:
gourl is a URL Parser that's very easy to use!

Just construct a new URL:
```go
url, err := gourl.Parse("http://example.com")
```

Parse the following parts:
```go
  //Scheme
  url.Scheme

  //User
  url.User

  //Password
  url.Password

  //Domain
  url.Domain

  //Subdomain
  url.Subdomain

  //Hostname
  url.Hostname

  //Port
  url.Port

  //Path
  url.Path
  
  //Query
  url.Query
  
  //Fragment
  url.Fragment
```

Get normalized version of a URL:
```go
url.String()
```
