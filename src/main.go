package main

import (
  "os"
  "log"
  "flag"
  "sort"
  "strings"
  "strconv"
  "github.com/gabbottron/rideshare-app/pkg/rideshare"
)

// clean pickup times list
var pickup_times []int

/*
  must_provide_times(): If times are bad, print help and exit
*/
func must_provide_times() {
  log.Println("You must provide at least one pickup time!")
  log.Println("e.g. rideshare_app --pickup_times 2,7,5")
  os.Exit(1)
}

/*
  init(): Process and validate flag data
*/
func init() {
  // flag variables
  var ride_time int
  var recovery_time int
  var balance_load bool
  var pickup_times_raw string

  // optional flags
  flag.IntVar(&ride_time, "ride_time", 15, "the ride time in minutes")
  flag.IntVar(&recovery_time, "recovery_time", 5, "the driver recovery time")
  flag.BoolVar(&balance_load, "balance_load", false, "balance load between drivers")

  // load list of pickup times
  flag.StringVar(&pickup_times_raw, "pickup_times", "", "list of pickup times, e.g. 5,15,35")

  // parse the command line flags
  flag.Parse()

  // Check the raw pickup times for data
  if len(pickup_times_raw) < 1 {
    must_provide_times()
  }
  // split the ride time values out
  s := strings.Split(pickup_times_raw, ",")
  
  // process the raw pickup times
  for i := 0; i < len(s); i++ {
    if len(s[i]) > 0 { // 
    // attempt to convert time to int
      val, err := strconv.Atoi(s[i])
      if err != nil {
        log.Printf("Error processing pickup time (%s) at index (%d), error: %s", s[i], i, err.Error())
        continue
      } else {
        // value must be positive
        if val >= 0 {
          // this pickup time is ok, add it
          pickup_times = append(pickup_times, val)
        }
      }
    }
  }

  // make sure we got at least one valid pickup time
  if len(pickup_times) < 1 {
    must_provide_times()
  }

  // sort pickup times (lowest to highest)
  // This simplifies driver distribution
  sort.Ints(pickup_times)

  // values look good, initialize the rideshare
  rideshare.Init(ride_time, recovery_time, balance_load)

  log.Printf("Processing these pickup times: %v", pickup_times)
}

func main() {
  // get a schedule for the drivers & riders
  driver_schedule := rideshare.GetSchedule(pickup_times)

  log.Printf("The final schedule: %v", driver_schedule)
}