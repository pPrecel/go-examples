package msf

import (
	"bytes"
	"fmt"
	"net/http"

	"gopkg.in/vmihailenco/msgpack.v2"
)

// RPC Api https://docs.rapid7.com/metasploit/standard-api-methods-reference

type Metasploit struct {
	host       string
	user, pass string
	token      string
}

func New(host, user, pass string) (*Metasploit, error) {
	msf := &Metasploit{
		host: host,
		user: user,
		pass: pass,
	}

	return msf, msf.Login()
}

func (m *Metasploit) Login() error {
	req := &loginReq{
		Method:   "auth.login",
		Username: m.user,
		Password: m.pass,
	}

	var res loginRes
	err := m.send(req, &res)
	if err != nil {
		return err
	}

	m.token = res.Token
	return nil
}

func (m *Metasploit) Logout() error {
	req := &logoutReq{
		Method: "auth.logout",
		Token:  m.token,
	}

	var res logoutRes
	err := m.send(req, &res)
	if err != nil {
		return err
	}

	m.token = ""
	return nil
}

func (m *Metasploit) SessionList() (VersionRes, error) {
	req := &versionReq{Method: "core.version", Token: m.token}
	var res VersionRes
	return res, m.send(req, &res)
}

func (m *Metasploit) send(req interface{}, res interface{}) error {
	buf := new(bytes.Buffer)
	msgpack.NewEncoder(buf).Encode(req)
	dest := fmt.Sprintf("http://%s/api", m.host)
	r, err := http.Post(dest, "binary/message-pack", buf)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := msgpack.NewDecoder(r.Body).Decode(&res); err != nil {
		return err
	}

	return nil
}
