package gourl

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	URLTest, e := Parse("http://example.com")

	if URLTest == nil {
		t.Fatalf("got nil from Parse()")
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}
	fmt.Println("Parse test passed...")
}

func TestURLSchemePositive(t *testing.T) {
	url := "http://example.com"
	URLTest, e := Parse(url)

	exp := "http"
	got := URLTest.Scheme

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("scheme test 1 passed...")
}

func TestURLSchemeNegative(t *testing.T) {
	url := "example.com"
	URLTest, e := Parse(url)

	exp := ""
	got := URLTest.Scheme

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("scheme test 2 passed...")
}

func TestURLDomainPositive(t *testing.T) {
	url := "http://www.example.com/"
	URLTest, e := Parse(url)

	exp := "example.com"
	got := URLTest.Domain

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("domain test 1 passed...")
}

func TestURLDomainNegative(t *testing.T) {
	url := ".com"
	URLTest, e := Parse(url)

	exp := ".com"
	got := URLTest.Domain

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("domain test 2 passed...")
}

func TestURLSubdomainPositive1(t *testing.T) {
	url := "http://www.example.com/"
	URLTest, e := Parse(url)

	exp := "www"
	got := URLTest.Subdomain

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("subdomain test 1 passed...")
}

func TestURLSubdomainPositive2(t *testing.T) {
	url := "http://register.hub.example.com/"
	URLTest, e := Parse(url)

	exp := "register.hub"
	got := URLTest.Subdomain

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("subdomain test 2 passed...")
}

func TestURLSubdomainNegative(t *testing.T) {
	url := "google.com"
	URLTest, e := Parse(url)

	exp := ""
	got := URLTest.Subdomain

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("subdomain test 3 passed...")
}

func TestURLHostnamePositive1(t *testing.T) {
	url := "http://user:pass@www.example.com/"
	URLTest, e := Parse(url)

	exp := "www.example.com"
	got := URLTest.Hostname

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("hostname test 1 passed...")
}

func TestURLHostnamePositive2(t *testing.T) {
	url := "http://www.example.com/"
	URLTest, e := Parse(url)

	exp := "www.example.com"
	got := URLTest.Hostname

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("hostname test 2 passed...")
}

func TestURLHostnameNegative(t *testing.T) {
	url := ".com"
	URLTest, e := Parse(url)

	exp := ".com"
	got := URLTest.Hostname

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("hostname test 3 passed...")
}

func TestURLUserPositive1(t *testing.T) {
	url := "http://user:pass@www.example.com/"
	URLTest, e := Parse(url)

	exp := "user"
	got := URLTest.User

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("user test 1 passed...")
}

func TestURLUserPositive2(t *testing.T) {
	url := "http://user@www.example.com/"
	URLTest, e := Parse(url)

	exp := "user"
	got := URLTest.User

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("user test 2 passed...")
}

func TestURLUserNegative(t *testing.T) {
	url := ".com"
	URLTest, e := Parse(url)

	exp := ""
	got := URLTest.User

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("user test 3 passed...")
}

func TestURLPasswordPositive(t *testing.T) {
	url := "http://user:pass@www.example.com/"
	URLTest, e := Parse(url)

	exp := "pass"
	got := URLTest.Password

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("password test 1 passed...")
}

func TestURLPasswordNegative(t *testing.T) {
	url := "http://user@www.example.com/"
	URLTest, e := Parse(url)

	exp := ""
	got := URLTest.Password

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("password test 2 passed...")
}

func TestURLPathPositive1(t *testing.T) {
	url := "http://user:pass@www.example.com/"
	URLTest, e := Parse(url)

	exp := "/"
	got := URLTest.Path

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("path test 1 passed...")
}

func TestURLPathPositive2(t *testing.T) {
	url := "http://www.example.com/path/goes/here.html"
	URLTest, e := Parse(url)

	exp := "/path/goes/here.html"
	got := URLTest.Path

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("path test 2 passed...")
}

func TestURLPathPositive3(t *testing.T) {
	url := "http://www.example.com/path/goes/here.html?query1=1&param2=2"
	URLTest, e := Parse(url)

	exp := "/path/goes/here.html"
	got := URLTest.Path

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("path test 3 passed...")
}

func TestURLPathPositive4(t *testing.T) {
	url := "http://www.example.com/path/goes/here.html#fragment"
	URLTest, e := Parse(url)

	exp := "/path/goes/here.html"
	got := URLTest.Path

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("path test 4 passed...")
}

func TestURLPathPositive5(t *testing.T) {
	url := "http://www.example.com/path/goes/here.html?query1=1&param2=2#fragment"
	URLTest, e := Parse(url)

	exp := "/path/goes/here.html"
	got := URLTest.Path

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("path test 5 passed...")
}

func TestURLPathNegative(t *testing.T) {
	url := "example.com"
	URLTest, e := Parse(url)

	exp := "/"
	got := URLTest.Path

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("path test 6 passed...")
}

func TestURLQueryPositive1(t *testing.T) {
	url := "http://www.example.com/path/goes/here.html?query1=1&param2=2#fragment"
	URLTest, e := Parse(url)

	exp := "query1=1&param2=2"
	got := URLTest.Query

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("query test 1 passed...")
}

func TestURLQueryPositive2(t *testing.T) {
	url := "http://www.example.com/path/goes/here.html?query"
	URLTest, e := Parse(url)

	exp := "query"
	got := URLTest.Query

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("query test 2 passed...")
}

func TestURLQueryNegative(t *testing.T) {
	url := "http://www.example.com/path/goes/here.html"
	URLTest, e := Parse(url)

	exp := ""
	got := URLTest.Query

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("query test 3 passed...")
}

func TestURLWithPort(t *testing.T) {
	url := "http://www.example.com:80"
	URLTest, e := Parse(url)

	exp := "80"
	got := URLTest.Port

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("port test 1 passed...")
}

func TestURLString(t *testing.T) {
	url := "www.example.com:80"
	URLTest, e := Parse(url)

	exp := "http://www.example.com:80/"
	got := URLTest.String()

	if got != exp {
		t.Fatalf("got %s, expecting %s\n", got, exp)
	}
	if e != nil {
		t.Fatalf("error encountered:", e)
	}

	fmt.Println("url normalization test 1 passed...")
}
