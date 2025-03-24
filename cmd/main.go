package main

import (
	"log"
	"strconv"
	"sync"

	"github.com/rcsolis/apiconcurrentsample/internal"
)

func processRequest(city string, ch chan struct{}, res chan string) {
	// Release the token back to the channel.
	defer func() {
		ch <- struct{}{}
	}()
	//Acquire a token from the limiter channel.
	_, ok := <-ch
	if !ok {
		return
	}
	//Call then API
	weather, err := internal.GetWeather(city)
	//Print the results
	if err != nil {
		res <- "Error fetching weather for " + city + ": " + err.Error()
	} else {
		res <- "Weather in " + weather.City.Name + ", " + weather.Country + ": " +
			strconv.FormatFloat(weather.Temperature.Value, 'f', 2, 64) + " " + weather.Temperature.Unit
	}
}

func main() {
	cities := []string{
		"Seattle",
		"Chicago",
		"Houston",
		"Phoenix",
		"Philadelphia",
		"Toronto",
		"Vancouver",
		"Montreal",
		"London",
		"Berlin",
		"Paris",
		"Rome",
		"Madrid",
		"Tokyo",
		"Beijing",
		"Seoul",
		"Bangkok",
		"Sydney",
		"Melbourne",
	}

	log.Println("Total cities:", len(cities))
	// Maximum number of concurrent requests
	maxConcurrentRequests := 5
	// Create a buffered channel to control the number of concurrent requests
	wg := sync.WaitGroup{}
	limiter := make(chan struct{}, maxConcurrentRequests)
	// Create the result channel
	result := make(chan string)

	for _, city := range cities {
		wg.Add(1)
		go processRequest(city, limiter, result)
		limiter <- struct{}{}
	}

	go func(cities []string, res chan string, w *sync.WaitGroup) {
		for range cities {
			weatherResult := <-res
			log.Println(weatherResult)
			w.Done()
		}
	}(cities, result, &wg)

	wg.Wait()
	close(limiter)
	close(result)
	log.Println("All done!")
}
