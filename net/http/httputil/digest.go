package xhttputil

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// NewDigestTransport returns an http.RoundTripper that uses the given http.RoundTripper for
// the first trip. If the response of the first trip results in an HTTP Digest Authentication
// challenge, it uses the given username and password to attempt the challenge on a second
// trip.
func NewDigestTransport(transport http.RoundTripper, username, password string) http.RoundTripper {
	if transport == nil {
		transport = http.DefaultTransport
	}

	return &digestTransport{transport, username, password}
}

type digestTransport struct {
	transport http.RoundTripper
	username, password  string
}

type challenge struct {
	realm     string
	domain    string
	nonce     string
	opaque    string
	stale     string
	algorithm string
	qop       string
}

func parseChallenge(input string) (*challenge, error) {
	s := strings.TrimSpace(input)
	if !strings.HasPrefix(s, "Digest ") {
		return nil, fmt.Errorf("bad challenge")
	}
	s = strings.TrimSpace(s[7:])

	var (
		sl = strings.Split(s, ", ")
		c = &challenge{
			algorithm: "MD5",
		}
	)

	const qs = `"`

	for i := range sl {
		if r := strings.SplitN(sl[i], "=", 2); len(r) != 2 {
			return nil, fmt.Errorf("bad challenge")
		} else {
			switch r[0] {
			case "realm":
				c.realm = strings.Trim(r[1], qs)
			case "domain":
				c.domain = strings.Trim(r[1], qs)
			case "nonce":
				c.nonce = strings.Trim(r[1], qs)
			case "opaque":
				c.opaque = strings.Trim(r[1], qs)
			case "stale":
				c.stale = strings.Trim(r[1], qs)
			case "algorithm":
				c.algorithm = strings.Trim(r[1], qs)
			case "qop":
				c.qop = strings.Trim(r[1], qs)
			default:
				return nil, fmt.Errorf("bad challenge")
			}
		}
	}

	return c, nil
}

type credentials struct {
	username   string
	realm      string
	nonce      string
	digestURI  string
	algorithm  string
	cnonce     string
	opaque     string
	messageQop string
	nonceCount int
	method     string
	password   string
}

func sumMD5(data string) string {
	hf := md5.New()
	_, _ = io.WriteString(hf, data)
	return fmt.Sprintf("%x", hf.Sum(nil))
}

func kd(secret, data string) string {
	return sumMD5(fmt.Sprintf("%s:%s", secret, data))
}

func (c *credentials) ha1() string {
	return sumMD5(fmt.Sprintf("%s:%s:%s", c.username, c.realm, c.password))
}

func (c *credentials) ha2() string {
	return sumMD5(fmt.Sprintf("%s:%s", c.method, c.digestURI))
}

func (c *credentials) resp(cnonce string) (string, error) {
	c.nonceCount++

	switch c.messageQop {
	case "auth":
		if cnonce != "" {
			c.cnonce = cnonce
		} else {
			b := make([]byte, 8)
			if n, err := io.ReadFull(rand.Reader, b); err != nil {
				return "", err
			} else if n != 8 {
				return "", fmt.Errorf("generated wrong amount of bytes for cnonce: %d", n)
			}

			c.cnonce = fmt.Sprintf("%x", b)[:16]
		}
		return kd(c.ha1(), fmt.Sprintf("%s:%08x:%s:%s:%s",
			c.nonce, c.nonceCount, c.cnonce, c.messageQop, c.ha2())), nil
	case "":
		return kd(c.ha1(), fmt.Sprintf("%s:%s", c.nonce, c.ha2())), nil
	}

	return "", fmt.Errorf("algorithm not implemented: %s", c.messageQop)
}

func (c *credentials) authorize() (string, error) {
	// Only implemented for MD5 and NOT MD5-sess.
	if c.algorithm != "MD5" {
		return "", fmt.Errorf("algorithm not implemented: %s", c.algorithm)
	}

	// Not implemented for "qop=auth-int".
	if c.messageQop != "auth" && c.messageQop != "" {
		return "", fmt.Errorf("algorithm not implemented: %s", c.messageQop)
	}

	resp, err := c.resp("")
	if err != nil {
		return "", fmt.Errorf("algorithm not implemented")
	}

	sl := []string{
		fmt.Sprintf(`username="%s"`, c.username),
		fmt.Sprintf(`realm="%s"`, c.realm),
		fmt.Sprintf(`nonce="%s"`, c.nonce),
		fmt.Sprintf(`uri="%s"`, c.digestURI),
		fmt.Sprintf(`response="%s"`, resp),
	}

	if c.algorithm != "" {
		sl = append(sl, fmt.Sprintf(`algorithm="%s"`, c.algorithm))
	}
	if c.opaque != "" {
		sl = append(sl, fmt.Sprintf(`opaque="%s"`, c.opaque))
	}
	if c.messageQop != "" {
		sl = append(sl,
			fmt.Sprintf("qop=%s", c.messageQop),
			fmt.Sprintf("nc=%08x", c.nonceCount),
			fmt.Sprintf(`cnonce="%s"`, c.cnonce),
		)
	}

	return fmt.Sprintf("Digest %s", strings.Join(sl, ", ")), nil
}

func (t *digestTransport) newCredentials(req *http.Request, c *challenge) *credentials {
	return &credentials{
		username:   t.username,
		realm:      c.realm,
		nonce:      c.nonce,
		digestURI:  req.URL.RequestURI(),
		algorithm:  c.algorithm,
		opaque:     c.opaque,
		messageQop: c.qop,
		nonceCount: 0,
		method:     req.Method,
		password:   t.password,
	}
}

// RoundTrip implements http.RoundTripper.
func (t *digestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if t == nil {
		t = &digestTransport{}
	}

	if t.transport == nil {
		t.transport = http.DefaultTransport
	}

	req2 := req.Clone(req.Context())

	res, err := t.transport.RoundTrip(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode != http.StatusUnauthorized {
		return res, nil
	}

	var (
		wwwAuthenticate = strings.TrimSpace(res.Header.Get("WWW-Authenticate"))
		digestAuthenticationPrefix = "Digest "
	)
	if !strings.HasPrefix(wwwAuthenticate, digestAuthenticationPrefix) {
		return nil, fmt.Errorf("bad challenge")
	}

	wwwAuthenticate = strings.TrimSpace(wwwAuthenticate[len(digestAuthenticationPrefix):])

	chal, err := parseChallenge(wwwAuthenticate)
	if err != nil {
		return res, err
	}

	auth, err := t.newCredentials(req2, chal).authorize()
	if err != nil {
		return res, err
	}

	_ = res.Body.Close()
	req2.Header.Set("Authorization", auth)

	return t.transport.RoundTrip(req2)
}
