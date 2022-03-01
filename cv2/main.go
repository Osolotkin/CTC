package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
  ST_GAS      = 0;
  ST_DIESEL   = 1;
	ST_LPG		  = 2;
	ST_ELECTRIC	= 3;
	ST_COUNT	  = 4;

	GAS_COUNT 		  = 4;
	DIESEL_COUNT 	  = 4;
	LPG_COUNT 		  = 4;
	ELECTRIC_COUNT 	= 8;

	STATION_COUNT = GAS_COUNT + DIESEL_COUNT + LPG_COUNT + ELECTRIC_COUNT;
	CACH_REGISTER_COUNT = 2;

	MAX_JOBS_PER_STATION = 10000;

	TIME_RESOLUTION = time.Microsecond;

	CARS_TO_PROCESS = 1000;
)

type void *int;

type ServicePoint struct {

	// in seconds
	minServiceTime float64;
	maxServiceTime float64;

	queue chan void;

	carServed int;

}

// to wait for join, when the time will come
var wgStations sync.WaitGroup;
var wgPaymentQueues sync.WaitGroup;

var stations [STATION_COUNT] *ServicePoint;
var paymentQueues [CACH_REGISTER_COUNT] *ServicePoint;

func makeStation(station *ServicePoint, stationType int) {

	switch stationType {

		case ST_GAS:

			station.minServiceTime = 1;
			station.maxServiceTime = 5;
			station.carServed = 0;

			break;

		case ST_DIESEL:

			station.minServiceTime = 1;
			station.maxServiceTime = 5;
			station.carServed = 0;

			break;

		case ST_LPG:

			station.minServiceTime = 1;
			station.maxServiceTime = 5;
			station.carServed = 0;

			break;

		case ST_ELECTRIC:

			station.minServiceTime = 3;
			station.maxServiceTime = 10;
			station.carServed = 0;

			break;

	}

	station.queue = make(chan void, MAX_JOBS_PER_STATION);

}

func sleep(min float64, max float64) {

	var delta float64 = min - max; 
	var duration float64 = min + rand.Float64() * delta;
	time.Sleep(time.Duration(duration) * TIME_RESOLUTION);

}

func processPayment(paymentQueue *ServicePoint, queue <- chan void) {

	defer wgPaymentQueues.Done();

	for range queue {
		sleep(paymentQueue.minServiceTime, paymentQueue.maxServiceTime);
		paymentQueue.carServed++;
	}

}

func processStation(station *ServicePoint, queue <- chan void) {

	defer wgStations.Done();

	for range queue {

		sleep(station.minServiceTime, station.maxServiceTime);
		
		var idx = rand.Intn(CACH_REGISTER_COUNT);
		paymentQueues[idx].queue <- nil;

		station.carServed++;

	}

}

func main() {
	
	// alloc
	for i := 0; i < STATION_COUNT; i++ {
		stations[i] = new(ServicePoint);
	}

	for i := 0; i < CACH_REGISTER_COUNT; i++ {
		paymentQueues[i] = new(ServicePoint);
	}

	// init all stations
	i := 0;
	offset := 0;

	for ; i < GAS_COUNT; i++ {
		makeStation(stations[i], ST_GAS);
	}
	offset = GAS_COUNT;

	for ; i < offset + DIESEL_COUNT; i++ {
		makeStation(stations[i], ST_DIESEL);
	}
	offset += DIESEL_COUNT;

	for ; i < offset + LPG_COUNT; i++ {
		makeStation(stations[i], ST_LPG);
	}
	offset += LPG_COUNT;

	for ; i < offset + ELECTRIC_COUNT; i++ {
		makeStation(stations[i], ST_ELECTRIC);
	}

	// init payment queues
	for i := 0; i < CACH_REGISTER_COUNT; i++ {
		paymentQueues[i].minServiceTime = 0.5;
		paymentQueues[i].maxServiceTime = 2;
		paymentQueues[i].carServed = 0;
		paymentQueues[i].queue = make(chan void, uint32((float64(STATION_COUNT) / CACH_REGISTER_COUNT) * MAX_JOBS_PER_STATION));
	}

	// ok, lets run all threads or whatever it is
	wgStations.Add(STATION_COUNT);
	wgPaymentQueues.Add(CACH_REGISTER_COUNT);
	
	for i := 0; i < STATION_COUNT; i++ {
		go processStation(stations[i], stations[i].queue);
	}

	for i := 0; i < CACH_REGISTER_COUNT; i++ {
		go processPayment(paymentQueues[i], paymentQueues[i].queue);
	}

	// car generation
	for i := 0; i < CARS_TO_PROCESS; i++ {

		// generate new car
		stationType := rand.Intn(ST_COUNT);

		// accordingly set index
		var idx = 0;
		switch stationType {

			case ST_GAS:

				idx = rand.Intn(GAS_COUNT);
				break;

			case ST_DIESEL:

				idx = GAS_COUNT + rand.Intn(DIESEL_COUNT);
				break;

			case ST_LPG:

				idx = GAS_COUNT + DIESEL_COUNT + rand.Intn(LPG_COUNT);
				break;

			case ST_ELECTRIC:

				idx = GAS_COUNT + DIESEL_COUNT + LPG_COUNT + rand.Intn(ELECTRIC_COUNT);
				break;

		}

		// add new job to the queue of the following station
		stations[idx].queue <- nil;

		// sleep some time, pepeChill
		time.Sleep(TIME_RESOLUTION / 100);

	}

	// close jobs
	for i := 0; i < STATION_COUNT; i++ {
		close(stations[i].queue);
	}

	// wait till all jobs finish
	wgStations.Wait();

	for i := 0; i < CACH_REGISTER_COUNT; i++ {
		close(paymentQueues[i].queue);
	}

	wgPaymentQueues.Wait();

	fmt.Printf("Summary");
	fmt.Printf("-------\n\n");

	sum := 0;
	for i := 0; i < STATION_COUNT; i++ {
		sum += stations[i].carServed;
		fmt.Printf("Station[%d]:\t%d car served\n", i + 1, stations[i].carServed);
	}

	fmt.Printf("\nControl sum:\t%d\n\n", sum);

	sum = 0;
	for i := 0; i < CACH_REGISTER_COUNT; i++ {
		sum += paymentQueues[i].carServed;
		fmt.Printf("Cash register[%d]:\t%d car served\n", i + 1, paymentQueues[i].carServed);
	}

	fmt.Printf("\nControl sum:\t%d\n", sum);

}
