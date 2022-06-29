package collector

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"

	"github.com/antikuz/xserver_exporter/pkg/logging"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	urnAuth = "/scalaboom/authUser?"
)

type xserver struct {
	logger       *logging.Logger
	url          string
	login        string
	passwd       string
	insecure bool
	client       *http.Client
}

func (xserver xserver) getSession(urn string) *http.Client {
	xserver.logger.Debugf("Start creating session url:%s\n", xserver.url)
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: nil})
	if err != nil {
		xserver.logger.Fatalln(err)
	}

	client := &http.Client{
		Jar: jar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: xserver.insecure,
			},
		},
	}

	url := xserver.url + urn
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		xserver.logger.Fatalln(err)
	}

	auth := base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%s:%s", xserver.login, xserver.passwd)))
	authorizationHeader := fmt.Sprintf("Basic %s", auth)

	request.Header.Add("content-type", "application/json")
	request.Header.Add("authorization", authorizationHeader)
	res, err := client.Do(request)
	if err != nil {
		xserver.logger.Fatalln(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		xserver.logger.Fatalf("Error Request: %s\n, status: %d\n", url, res.StatusCode)
	}
	xserver.logger.Debugf("Session created cookie:%s", res.Cookies())
	return client
}

func (xserver xserver) getJSON(urn string) ([]byte, error) {
	url := xserver.url + urn
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		xserver.logger.Fatalln(err)
	}

	auth := base64.StdEncoding.EncodeToString([]byte(
		fmt.Sprintf("%s:%s", xserver.login, xserver.passwd)))
	authorizationHeader := fmt.Sprintf("Basic %s", auth)

	request.Header.Add("content-type", "application/json")
	request.Header.Add("authorization", authorizationHeader)

	res, err := xserver.client.Do(request)
	if err != nil {
		xserver.logger.Fatalf("Failed request to: %s, due to err: %v\n", xserver.url, err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func NewXserverManager(logger *logging.Logger, url string, login string, passwd string, insecure bool, reg prometheus.Registerer) {
	x := &xserver{
		logger:       logger,
		url:          url,
		login:        login,
		passwd:       passwd,
		insecure: insecure,
	}
	x.client = x.getSession(urnAuth)

	xmc := &monitoringCollector{x}
	xuc := &usersCollector{x}
	xnc := &netstatCollector{x}
	xdc := &harddisksCollector{x}

	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xmc)
	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xuc)
	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xnc)
	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xdc)
}
