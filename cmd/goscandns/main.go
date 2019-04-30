package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

type scannable []string

func main() {
	cidrs := genPool()
	log.Printf("About to scan (%v) CIDRs, or (%v) IPs...", len(cidrs), len(cidrs)*254)

	// Don't clog the system,
	// watch for max open files.
	maxScan := ulimit() / 2
	sem := semaphore.NewWeighted(int64(maxScan))
	foundTargets := newRegistry()
	now := time.Now()

	// Fire.
	for _, targetCIDR := range cidrs {
		ts := time.Now()
		targets, err := hostsFromCIDR(targetCIDR)
		if err != nil {
			log.Fatalf("could not retrieve IP pool: %v", err)
		}

		log.Println("---------------------------------------------------------")
		log.Printf("Scanning (%v) targets ulimit (%v)...", len(targets), maxScan)
		log.Println("---------------------------------------------------------")
		processTargets(sem, targets, foundTargets)

		log.Println("---------------------------------------------------------")
		log.Printf("Finished scanning (%v) IP addresses in (%v).", len(targets), time.Since(ts))
		if foundTargets.len() > 0 {
			log.Println("---- [RESULTS] ----")
			for k := range foundTargets.retrieve() {
				fmt.Println(k)
			}
			log.Println("---- [END] ----")
		}
		log.Println("---------------------------------------------------------")
	}

	log.Println("---------------------------------------------------------")
	log.Printf("DONE: (%v) IP addresses in (%v) and found (%v) possible candidates.", len(cidrs)*254, time.Since(now), foundTargets.len())
	log.Println("---------------------------------------------------------")

	if foundTargets.len() > 0 {
		log.Println("---- [RESULTS] ----")
		for k := range foundTargets.retrieve() {
			fmt.Println(k)
		}
		log.Println("---- [END] ----")
	}
}

func processTargets(sem *semaphore.Weighted, targets []string, foundTargets *registry) {
	var wg sync.WaitGroup
	wg.Add(len(targets))
	ctx := context.TODO()

	for _, target := range targets {
		go func(ipRaw string) {
			if err := sem.Acquire(ctx, 1); err != nil {
				log.Fatalf("could not aquire semaphore: %v", err)
			}
			defer sem.Release(1)
			defer wg.Done()

			ip := net.ParseIP(ipRaw)
			if ip == nil || ip.To4() == nil {
				log.Fatalf("invalid IPv4: %v", ipRaw)
			}
			ips, err := scanIP(context.Background(), ip)
			if err != nil {
				log.Printf("could not scan (%v): %v", ip, err)
			}

			// Good target.
			if len(ips) > 0 {
				foundTargets.add(ipRaw)
			}
		}(target)
	}

	wg.Wait()
}

func mustGetEnv(value string) string {
	v := os.Getenv(value)
	if v == "" {
		log.Fatalf("could not retrieve needed value (%v) from the environment", value)
	}

	return v
}
