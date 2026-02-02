package proxy

import (
	"context"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Proxy struct {
	reverseProxy *httputil.ReverseProxy
}

func New(target string, timeout time.Duration)(*Proxy, error){
	backendUrl, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	rp := httputil.NewSingleHostReverseProxy(backendUrl)
	rp.Transport = &http.Transport{
		ResponseHeaderTimeout: timeout,
	}

	rp.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error){
		log.Printf("Proxy error : %v", err)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
	}

	return &Proxy{reverseProxy: rp}, nil
}

func (p *Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request){
	ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
	defer cancel()

	r = r.WithContext(ctx)
	p.reverseProxy.ServeHTTP(w, r)
}