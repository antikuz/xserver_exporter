package main

import (
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type UsersWidget struct {
	Active  int    `json:"active,omitempty"`
	Blocked int    `json:"blocked,omitempty"`
	Enabled int    `json:"enabled,omitempty"`
	Users   int    `json:"users,omitempty"`
	Vpn     int    `json:"vpn,omitempty"`
}

type StatWidget struct {
	Stat  map[string]int `json:"stat,omitempty"`
}

type IfacesWidget struct {
	Today struct {
		Ping           string `json:"ping,omitempty"`
		VpnConnections int    `json:"vpn_connections,omitempty"`
		ChannelUsage   struct {
			Rp float32 `json:"rp,omitempty"`
			Tp float32 `json:"tp,omitempty"`
			Rx float32 `json:"rx,omitempty"`
			Tx float32 `json:"tx,omitempty"`
		} `json:"channelUsage,omitempty"`
	} `json:"today,omitempty"`
}

type MonitoringWidget struct {
	Cpu   int     `json:"cpu,omitempty"`
	Ram   int     `json:"ram,omitempty"`
	Swap  int     `json:"swap,omitempty"`
	MaxLa int     `json:"maxLa,omitempty"`
	La    float32 `json:"la,omitempty"`
}

type Envelope struct {
	Oid string `json:"oid"`
}

var (
    usersWidget UsersWidget
    statWidget StatWidget
    ifacesWidget IfacesWidget
    monitoringWidget MonitoringWidget
)

func parseJson(arrayJsonRaw []*json.RawMessage) {
    var env Envelope
    for _, element := range arrayJsonRaw {
        if err := json.Unmarshal(*element, &env); err != nil {
            log.Fatal(err)
        }
        switch env.Oid {
        case "UsersWidget":
            if err := json.Unmarshal(*element, &usersWidget); err != nil {
                log.Fatal(err)
            }
        case "StatWidget":
            if err := json.Unmarshal(*element, &statWidget); err != nil {
                log.Fatal(err)
            }
        case "IfacesWidget":
            if err := json.Unmarshal(*element, &ifacesWidget); err != nil {
                log.Fatal(err)
            }
        case "MonitoringWidget":
            if err := json.Unmarshal(*element, &monitoringWidget); err != nil {
                log.Fatal(err)
            }
        }
    }
}

func getJson() []byte {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	url := "http://localhost:8000"

	request, err := http.NewRequest("GET", url, nil)

    if err != nil {
        log.Fatal(err)
    }

	login := "login"
	passwd := "passwd"
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", login, passwd)))
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
        log.Fatal(err)
    }
    return body
}

func main() {
	reg := prometheus.NewPedanticRegistry()
	NewClusterManager("db", reg)

	reg.MustRegister(
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
		collectors.NewGoCollector(),
	)

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
