package collector

import (
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type xserver struct {
	url    string
	login  string
	passwd string
	insecureSkip bool
}

func (xserver xserver) getJson(urn string) ([]byte, error) {
	client := &http.Client{
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
	
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
        return nil, err
    }

    return body, nil
}

func NewXserverManager(url string, login string, passwd string, insecureSkip bool, reg prometheus.Registerer) {
	x := xserver{
			url: url,
			login: login,
			passwd: passwd,
			insecureSkip: insecureSkip,
		}
	xmc := &monitoringCollector{x}
	xuc := &usersCollector{x}
	xnc := &netstatCollector{x}
	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xmc)
	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xuc)
	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xnc)
}