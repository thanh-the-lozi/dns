package main

import (
	"log"
	"net"
	"strconv"

	"github.com/miekg/dns"
)

var domainToIPAddr map[string]string = map[string]string{
	"kmin.edu.vn.":      "1.2.3.4",
	"jameshfisher.com.": "104.198.14.52",
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
	srv := &dns.Server{
		Addr: ":" + strconv.Itoa(53),
		Net:  "udp",
	}

	srv.Handler = &handler{}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("Failed to set udp listener %s\n", err.Error())
	}
}
