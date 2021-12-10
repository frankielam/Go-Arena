package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	p := fmt.Println
	// s := "postgres://user:password@host.com:1234/path?key=value#f"
	s := "https://github.com/frankielam/Go-Arena.git?key=yy#zz"
	p(s)
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	p("scheme", u.Scheme)

	p(u.User)
	p(u.User.Username())
	p(u.User.Password())
	pass, _ := u.User.Password()
	p(pass)

	p("host", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	p(host)
	p(port)

	p("path", u.Path)
	p(u.Fragment)
	p(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	p(m)
	p(m["key"][0])

}