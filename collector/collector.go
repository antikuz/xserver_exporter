package collector

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	urnAuth = "/scalaboom/authUser?"
)

type xserver struct {
	url          string
	login        string
	passwd       string
	insecureSkip bool
	client       *http.Client
}

func (xserver xserver) getSession(urn string) *http.Client {
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: nil})
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: xserver.insecureSkip,
			},
		},
	}

	url := xserver.url + urn
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	auth := base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%s:%s", xserver.login, xserver.passwd)))
	authorizationHeader := fmt.Sprintf("Basic %s", auth)

	request.Header.Add("content-type", "application/json")
	request.Header.Add("authorization", authorizationHeader)
	res, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("Error Request: %s\n, status: %d\n", url, res.StatusCode)
	}

	return client
}

func (xserver xserver) getJSON(urn string) ([]byte, error) {
	url := xserver.url + urn
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	auth := base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%s:%s", xserver.login, xserver.passwd)))
	authorizationHeader := fmt.Sprintf("Basic %s", auth)

	request.Header.Add("content-type", "application/json")
	request.Header.Add("authorization", authorizationHeader)

	res, err := xserver.client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func NewXserverManager(url string, login string, passwd string, insecureSkip bool, reg prometheus.Registerer) {
	x := &xserver{
		url:          url,
		login:        login,
		passwd:       passwd,
		insecureSkip: insecureSkip,
	}
	x.client = x.getSession(urnAuth)
	
	xmc := &monitoringCollector{x}
	xuc := &usersCollector{x}
	xnc := &netstatCollector{x}
	
	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xmc)
	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xuc)
	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xnc)
}
