package main

import (
	"flag"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
	Error      error
}

func main() {
	url := flag.String("url", "", "URL do serviço a ser testado")
	totalRequests := flag.Int("requests", 10, "Número total de requests")
	concurrency := flag.Int("concurrency", 5, "Número de chamadas simultâneas")
	flag.Parse()

	if *url == "" {
		fmt.Println("Erro: A URL é obrigatória (--url)")
		return
	}

	startTime := time.Now()
	results := make(chan Result, *totalRequests)

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, *concurrency)

	for i := 0; i < *totalRequests; i++ {
		wg.Add(1)
		semaphore <- struct{}{} // bloqueia se exceder concorrência
		go func() {
			defer wg.Done()
			defer func() { <-semaphore }()

			start := time.Now()
			resp, err := http.Get(*url)
			duration := time.Since(start)

			if err != nil {
				results <- Result{StatusCode: 0, Duration: duration, Error: err}
				return
			}
			defer resp.Body.Close()

			results <- Result{StatusCode: resp.StatusCode, Duration: duration}
		}()
	}

	wg.Wait()
	close(results)

	// Geração de relatório
	statusCounts := make(map[int]int)
	var successCount int
	var totalDuration time.Duration

	for res := range results {
		if res.StatusCode == 200 {
			successCount++
		}
		statusCounts[res.StatusCode]++
		totalDuration += res.Duration
	}

	endTime := time.Since(startTime)

	fmt.Println("\n=== Relatório de Teste de Carga ===")
	fmt.Printf("Tempo total do teste: %v\n", endTime)
	fmt.Printf("Quantidade totais de requests: %d\n", *totalRequests)
	fmt.Printf("Quantidade de requests com status 200: %d\n", successCount)
	fmt.Println("Quantidade de requests com outros status:")
	for code, count := range statusCounts {
		if code != 200 {
			fmt.Printf("  %d: %d\n", code, count)
		}
	}
}
