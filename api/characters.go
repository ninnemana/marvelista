package api

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/url"
	"time"
)

const (
	PUBLIC_KEY  = "344dd9e52bab584a2b99ce0c60a3e99b"
	PRIVATE_KEY = "030c359d71af304dbcc8b9d80c45eb18b4bed9eb"
)

func List() string {
	url, _ := requestMux("/v1/public/characters", nil)
	return url
}

func requestMux(endpoint string, qs url.Values) (string, error) {
	mux := new(url.URL)
	mux.Host = "gateway.marvel.com"
	mux.Scheme = "https"
	mux.Path = endpoint

	ts := time.Now().String()

	h := md5.New()
	io.WriteString(h, ts)
	io.WriteString(h, PRIVATE_KEY)
	io.WriteString(h, PUBLIC_KEY)
	if qs == nil {
		qs = make(map[string][]string, 0)
	}
	qs.Add("ts", ts)
	qs.Add("apikey", PUBLIC_KEY)
	qs.Add("hash", hex.EncodeToString(h.Sum(nil)))

	mux.RawQuery = qs.Encode()

	return mux.String(), nil

}
