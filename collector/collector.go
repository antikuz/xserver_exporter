package collector

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
)

type xserverCollector struct {
	xserver
}

type xserver struct {
	url    string
	login  string
	passwd string
	insecureSkip bool
}

type Envelope struct {
	Oid string `json:"oid"`
}

func (xserver xserver) getJson() ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: xserver.insecureSkip,
			},
		},
	}

	url := xserver.url + "/scalaboom/widgets?sitesFeedCount=0&searchFeedCount=0&systemJournalCount=0"
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

func (xserver xserver) parseJson(arrayJsonRaw []*json.RawMessage) {
    var env Envelope
    for _, element := range arrayJsonRaw {
        if err := json.Unmarshal(*element, &env); err != nil {
            log.Fatal(err)
        }
        switch env.Oid {
        case "UsersWidget":
            if err := json.Unmarshal(*element, &UsersWidget); err != nil {
                log.Fatal(err)
            }
        case "StatWidget":
            if err := json.Unmarshal(*element, &Netstat); err != nil {
                log.Fatal(err)
            }
        case "IfacesWidget":
            if err := json.Unmarshal(*element, &Netstat); err != nil {
                log.Fatal(err)
            }
        case "MonitoringWidget":
            if err := json.Unmarshal(*element, &MonitoringWidget); err != nil {
                log.Fatal(err)
            }
        }
    }
}

func (xserver xserverCollector) Describe(ch chan<- *prometheus.Desc) {
	prometheus.DescribeByCollect(xserver, ch)
}

func (xserver xserverCollector) Collect(ch chan<- prometheus.Metric) {
    request, err := xserver.getJson()
	if err != nil {
		log.Fatal(err)
	}

	var arrayJsonRaw []*json.RawMessage
    json.Unmarshal(request, &arrayJsonRaw)
	xserver.parseJson(arrayJsonRaw)

	UsersWidget.Collect(ch)
	Netstat.Collect(ch)
	MonitoringWidget.Collect(ch)
}

func NewXserverManager(url string, login string, passwd string, insecureSkip bool, reg prometheus.Registerer) *xserverCollector {
	x := xserver{
			url: url,
			login: login,
			passwd: passwd,
			insecureSkip: insecureSkip,
		}
	xc := &xserverCollector{x}
	prometheus.WrapRegistererWith(prometheus.Labels{"url": url}, reg).MustRegister(xc)
	return xc
}