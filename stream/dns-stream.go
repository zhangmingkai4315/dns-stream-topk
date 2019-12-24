package stream

import (
	"errors"
	"math/rand"
	"time"
)

// DNSFlow Manager
type DNSFlow struct {
	domains     map[string]int
	total       int
	duration    time.Duration
	flatDomains []string
	stream      chan string
}

// NewDNSFlow create a new dns flow manager
func NewDNSFlow(domains map[string]int, total int, duration time.Duration) (*DNSFlow, error) {
	if len(domains) == 0 {
		return nil, errors.New("configuration is empty")
	}
	dnsFlow := DNSFlow{
		domains:  domains,
		total:    total,
		duration: duration,
		stream:   make(chan string, 100),
	}
	dnsFlow.preCompile()
	return &dnsFlow, nil
}

func (df *DNSFlow) preCompile() {
	res := []string{}
	for k, v := range df.domains {
		for i := 0; i < v; i++ {
			res = append(res, k)
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(res), func(i, j int) { res[i], res[j] = res[j], res[i] })
	df.flatDomains = res
}

// Start the domain stream
func (df *DNSFlow) Start() chan string {
	total := 0
	stoppedAt := time.Now().Add(df.duration).Unix()
	go func() {
		for {
			if df.duration > 0 && time.Now().Unix() > stoppedAt {
				close(df.stream)
				return
			}
			for _, domain := range df.flatDomains {
				df.stream <- domain
				total++
				if df.total > 0 && total > df.total {
					close(df.stream)
					return
				}
			}
		}
	}()
	return df.stream
}
