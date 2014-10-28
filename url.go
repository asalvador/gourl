package gourl

import (
	"errors"
	nurl "net/url"
	"strconv"
	"strings"
)

//GoURL contains the parts of URL after parsing
type GoURL struct {
	URL       string
	Scheme    string
	User      string
	Password  string
	Hostname  string
	Subdomain string
	Domain    string
	TLD       string
	Port      string
	Path      string
	Query     string
	Fragment  string
}

//Parse accepts (or treats the input as) an absolute URL string
func Parse(url string) (*GoURL, error) {
	if url == "" {
		return nil, errors.New("gourl: empty url")
	}

	hasScheme := true
	if strings.Contains(url, "://") == false {
		url = "http://" + url
		hasScheme = false
	}

	u, e := nurl.Parse(url)

	if e != nil {
		return nil, errors.New("gourl: url parse error")
	}

	gourl := &GoURL{URL: url, Scheme: u.Scheme, Hostname: u.Host, Path: u.Path, Query: u.RawQuery, Fragment: u.Fragment}

	if hasScheme == false {
		gourl.Scheme = ""
	}

	if gourl.Path == "" {
		gourl.Path = "/"
	}

	gourl.Port, e = port(gourl.Hostname)
	if e != nil {
		return nil, e
	}

	gourl.Subdomain, e = subdomain(gourl.Hostname)
	if e != nil {
		return nil, e
	}

	gourl.Domain, e = domain(gourl.Hostname)
	if e != nil {
		return nil, e
	}

	gourl.TLD, e = tld(gourl.Hostname)
	if e != nil {
		return nil, e
	}

	//gourl.User, gourl.Password = userpass(rawup)
	if u.User != nil {
		gourl.User = u.User.Username()
		gourl.Password, _ = u.User.Password()
	}

	return gourl, nil
}

//Returns normalized version of url based on GoURL parts
func (url *GoURL) String() string {
	var u string
	u = url.Scheme
	if u == "" {
		u = "http"
	}

	u = u + "://" + url.Hostname + url.Path
	if url.Query != "" {
		u = u + "?" + url.Query
	}
	if url.Fragment != "" {
		u = u + "#" + url.Fragment
	}

	return u
}

// subdomain returns the subdomain part of the url
func subdomain(host string) (string, error) {
	sub := ""
	if host != "" {
		t, e := tld(host)
		if e != nil {
			return "", errors.New("gourl: error parsing subdomain")
		}
		psub := "." + t
		p, e := port(host)
		if e != nil {
			return "", errors.New("gourl: error parsing subdomain")
		}
		if p != "" {
			psub = psub + ":" + p
		}
		if t != "" {
			subdom := strings.TrimSuffix(host, psub)
			c := strings.Count(subdom, ".")

			if c >= 1 {
				array := strings.Split(subdom, ".")
				sub = array[0]
				if len(array) > 1 {
					sub = strings.Join(array[0:len(array)-1], ".")
				}
			}
		}
	}

	return sub, nil
}

// domainreturns the domain part of the url
func domain(host string) (string, error) {
	dom := ""
	if host != "" {
		t, e := tld(host)
		if e != nil {
			return "", errors.New("gourl: error parsing domain")
		}
		pdom := "." + t
		p, e := port(host)
		if e != nil {
			return "", errors.New("gourl: error parsing domain")
		}
		if p != "" {
			pdom = pdom + ":" + p
		}
		if t != "" {
			d := strings.TrimSuffix(host, pdom)
			c := strings.Count(d, ".")

			if c >= 0 {
				array := strings.Split(d, ".")
				dom = array[len(array)-1] + pdom
			}
		}
	}

	return dom, nil
}

// TLD returns the tld part of the url
func tld(host string) (string, error) {
	if host != "" {
		if strings.Contains(host, ":") {
			p, e := port(host)
			if e != nil {
				return "", errors.New("gourl: error parsing tld")
			} else if p != "" {
				pport := ":" + p
				host = strings.TrimSuffix(host, pport)
			}
		}
		for _, t := range strings.Split(TLDs, "\n") {
			if t == "" {
				continue
			}

			if strings.HasSuffix(host, t) {
				return t, nil
			}
		}
	}
	return "", nil
}

func port(host string) (string, error) {
	var e error
	p := ""
	if host != "" {
		if strings.Contains(host, ":") {
			array := strings.SplitN(host, ":", 2)
			if len(array) == 2 && array[1] != "" {
				_, err := strconv.Atoi(array[1])
				if err != nil {
					e = errors.New("gourl: invalid url")
				} else {
					p = array[1]
				}
			} else {
				e = errors.New("gourl: error parsing port")
			}
		}
	}

	return p, e
}

func userpass(raw string) (user, password string) {
	if raw != "" {
		up := strings.Split(raw, ":")
		if len(up) == 1 {
			user = up[0]
		} else if len(up) == 2 {
			user = up[0]
			password = up[1]
		}
	}
	return user, password
}
