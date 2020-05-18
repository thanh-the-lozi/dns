package main

import (
	"log"
	"net"
	"strconv"

	"github.com/miekg/dns"
)

var domainToIPAddr map[string]string = map[string]string{
	"google.com.":   "172.217.17.110",
	"facebook.com.": "157.240.195.35",
	"linkedin.com.": "108.174.10.10",
	"kmin.edu.vn.":  "125.212.221.74",
}

type handler struct{}

func (this *handler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := dns.Msg{}
	msg.SetReply(r)
	switch r.Question[0].Qtype {
	case dns.TypeA:
		domain := msg.Question[0].Name
		address, ok := domainToIPAddr[domain]
		if ok {
			msg.Answer = append(msg.Answer, &dns.A{
				Hdr: dns.RR_Header{
					Name:   domain,
					Rrtype: dns.TypeA,
					Class:  dns.ClassINET,
					Ttl:    60,
				},
				A: net.ParseIP(address),
			})
		}
	}
	w.WriteMsg(&msg)
}

func main() {
	server := &dns.Server{
		Addr: ":" + strconv.Itoa(53),
		Net:  "udp",
	}

	server.Handler = &handler{}
	defer server.Shutdown()
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
