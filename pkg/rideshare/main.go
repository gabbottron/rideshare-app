package rideshare

// the settings
var RideTime int
var RecoveryTime int
var BalanceLoad bool

/*
  Init(): Set the settings for scheduling
*/
func Init(ride_time int, recovery_time int, balance_load bool) {
  RideTime = ride_time
  RecoveryTime = recovery_time
  BalanceLoad = balance_load
}

/*
  GetSimpleSchedule(...): Build a basic schedule with no rider balancing
*/
func GetSimpleSchedule(pickup_times []int) [][]int {
  // start with a single driver
  driver_schedule := make([][]int, 1)

  // no care about even ordering
  var placed = false
  for i := 0; i < len(pickup_times); i++ {
    placed = false
    for d := 0; d < len(driver_schedule); d++ {
      // Driver has no riders yet
      if len(driver_schedule[d]) < 1 {
        // Rider will be the first for this driver
        driver_schedule[d] = make([]int, 1)
        driver_schedule[d][0] = pickup_times[i]
        placed = true
        break // rider is placed, move on to next
      }
      // Current driver has enough time to get this rider after last dropoff
      if (driver_schedule[d][len(driver_schedule[d]) - 1] + (RideTime + RecoveryTime)) <= pickup_times[i] {
        driver_schedule[d] = append(driver_schedule[d], pickup_times[i])
        placed = true
        break // rider is placed, move on to next
      }
    }
    if !placed {
      // we need a new driver, no current drivers could accomodate the rider
      driver_schedule = append(driver_schedule, []int{pickup_times[i]})
    }
  }

  return driver_schedule
}


/*
  GetBalancedSchedule(...): Build a driver load balanced schedule
*/
func GetBalancedSchedule(pickup_times []int) [][]int {
  // start with a single driver
  driver_schedule := make([][]int, 1)

  placed := false
  found_c := false
  // tracks the index of the driver that should take
  // the current rider
  best_driver_index := 0
  // tracks the current lowest rider load of a driver yet found
  lowest_load := len(pickup_times)
  for i := 0; i < len(pickup_times); i++ {
    placed = false
    found_c = false
    best_driver_index = 0
    lowest_load = len(pickup_times)
    for d := 0; d < len(driver_schedule); d++ {
      // Driver has no riders yet
      if len(driver_schedule[d]) < 1 {
        // Rider will be the first for this driver
        driver_schedule[d] = make([]int, 1)
        driver_schedule[d][0] = pickup_times[i]
        placed = true
        // we still break in this case because the driver
        // had no riders, so no other driver could have fewer
        break // rider is placed, move on to next
      }
      // Current driver has enough time to get this rider after last dropoff
      if (driver_schedule[d][len(driver_schedule[d]) - 1] + (RideTime + RecoveryTime)) <= pickup_times[i] {
        // Is this driver the current best candidate to take this rider?
        if len(driver_schedule[d]) < lowest_load {
          found_c = true
          best_driver_index = d
          lowest_load = len(driver_schedule[d])
        }
      }
    }
    if !placed {
      if found_c {
        // add the pickup time to the least loaded driver
        driver_schedule[best_driver_index] = append(driver_schedule[best_driver_index], pickup_times[i])
      } else {
        // we need a new driver, no current drivers could accomodate the rider
        driver_schedule = append(driver_schedule, []int{pickup_times[i]})
      }
    }
  }

  return driver_schedule
}

/*
  GetSchedule(...): Build a driver schedule
*/
func GetSchedule(pickup_times []int) [][]int {
  if BalanceLoad {
    return GetBalancedSchedule(pickup_times)
  } else {
    return GetSimpleSchedule(pickup_times)
  }
}
