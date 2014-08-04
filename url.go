// Package gourl implements functions to parse different parts of a URL
package gourl

import (
	"bytes"
	"fmt"
)

// URL parts
type URL struct {
	URL string
	scheme string		//done
	user string		//done
	password string		//done
	hostname string		//done
	subdomain string	//done
	domain string		//done
	path string		//done
	query string		//done
	fragment string		//done

	userpass string
	tld string
}

// NewURL returns a new URL given a url string
func NewURL(url string) *URL {
	if url == "" {
		return nil
	}

	return &URL{URL: url}
}

// Scheme returns the scheme part of the url
func (u *URL) Scheme() string {
	if u != nil && u.scheme != "" {
		return u.scheme
	}

	burl := toByteArray(u.URL)
	if bytes.Contains(burl, toByteArray(`://`)) {
		u.scheme = string(bytes.Split(burl, toByteArray(`://`))[0])
	}

	return u.scheme
}

// User returns the user part of the url
func (u *URL) User() string {
	if u != nil && u.user != "" {
		return u.user
	}

	up := u.userPass()
	if up == "" {
		return u.user
	}

	buser := toByteArray(up)
	if bytes.Contains(buser, toByteArray(":")) {
		buser = bytes.Split(buser, toByteArray(":"))[0]
	}

	u.user = string(buser)
	return u.user
}

// Password returns the password part of the url
func (u *URL) Password() string {
	if u != nil && u.password != "" {
		return u.password
	}

	up := u.userPass()
	if up == "" {
		u.password = ""
	}

	bpass := toByteArray(up)
	if bytes.Contains(bpass, toByteArray(":")) {
		u.password = string(bytes.Split(bpass, toByteArray(":"))[1])
	}

	return u.password
}

// Domain returns the domain part of the url
func (u *URL) Domain() string {
	if u != nil && u.domain != "" {
		return u.domain
	}

	tld := u.TLD()
	if tld != "" {
		bhost := toByteArray(u.Hostname())
		ptld := "." + tld
		bdomain := bytes.TrimRight(bhost, ptld)
		c := bytes.Count(bdomain, toByteArray("."))

		if c > 0 {
			darray := bytes.Split(bdomain, toByteArray("."))
			bdomain = darray[len(darray)-1]
		}
		u.domain = fmt.Sprintf("%s.%s", bdomain, tld)
	}

	return u.domain
}

// Subdomain returns the subdomain part of the url
func (u *URL) Subdomain() string {
	if u != nil && u.subdomain != "" {
		return u.subdomain
	}

	if u.Hostname() != "" && u.Domain() != "" {
		subd := bytes.SplitN(toByteArray(u.Hostname()), toByteArray(u.Domain()), 2)[0]
		u.subdomain = string(bytes.TrimRight(subd, "."))
	}

	return u.subdomain
}

// Hostname returns the hostname part of the url
func (u *URL) Hostname() string {
	if u != nil && u.hostname != "" {
		return u.hostname
	}

	// Trim scheme
	burl := bytes.TrimLeft(toByteArray(u.URL), u.Scheme() + "://")

	// Trim path
	if bytes.Contains(burl, toByteArray("/")) {
		burl = bytes.Split(burl, toByteArray(`/`))[0]
	}

	// Trim user-password
	if bytes.Contains(burl, toByteArray("@")) {
                burl = bytes.Split(burl, toByteArray(`@`))[1]
        }

	u.hostname = string(burl)
	return u.hostname
}

// Path returns the path part of the url
func (u *URL) Path() string {
	if u != nil && u.path != "" {
		return u.path
	}

	burl := toByteArray(u.URL)
	if bytes.Contains(burl, toByteArray(`://`)) {
		burl = bytes.Split(burl, toByteArray(`://`))[1]
	}

	// Trim query
	if bytes.Contains(burl, toByteArray("?")) {
		burl = bytes.SplitN(burl, toByteArray("?"), 2)[0]
	}

	// Trim fragment
	if bytes.Contains(burl, toByteArray("#")) {
		burl = bytes.SplitN(burl, toByteArray("#"), 2)[0]
	}

	if bytes.Contains(burl, toByteArray(`/`)) {
		burl = bytes.SplitN(burl, toByteArray(`/`), 2)[1]
	} else {
		burl = toByteArray("") 
	}


	u.path = fmt.Sprintf("/%s", burl)
	return u.path
}

// Query returns the query part of the url
func (u *URL) Query() string {
	if u != nil && u.query != "" {
		return u.query
	}

	burl := toByteArray(u.URL)
	if bytes.Contains(burl, toByteArray("?")) {
		burl = bytes.SplitN(burl, toByteArray("?"), 2)[1]

		if u.Fragment() != "" {
			burl = bytes.TrimRight(burl, "#" + u.Fragment())
		}

		u.query = string(burl)
	}

	return u.query
}

// Fragment returns the fragment part of the url
func (u *URL) Fragment() string {
	if u != nil && u.fragment != "" {
		return u.fragment
	}

	burl := toByteArray(u.URL)
	if bytes.Contains(burl, toByteArray(`#`)) {
		u.fragment = string(bytes.SplitN(burl, toByteArray(`#`), 2)[1])
	}

	return u.fragment
}

// TLD returns the tld part of the url
func (u *URL) TLD() string {
	if u != nil && u.tld != "" {
		return u.tld
	}

	for _, tld := range bytes.Split(toByteArray(TLDs), toByteArray("\n")) {
		if bytes.Equal(tld, toByteArray("")) {
			continue
		}

		if bytes.HasSuffix(toByteArray(u.Hostname()), tld) {
			u.tld = string(tld)
			break
		}
	}
	return u.tld
}

// userPass returns the username-password part of the url
func (u *URL) userPass() string {
	if u != nil && u.userpass != "" {
		return u.userpass
	}

	burl := toByteArray(u.URL)
	bhost := toByteArray("@" + u.Hostname())
	if bytes.Contains(burl, bhost) {
		burl = bytes.TrimLeft(burl, u.Scheme() + "://")
		u.userpass = string(bytes.Split(burl, bhost)[0])
	}

	return u.userpass
}


func toByteArray(s string) []byte {
        b := []byte(s)

        return b
}

