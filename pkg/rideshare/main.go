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
  GetSchedule(...): Build a basic schedule with no rider balancing
*/
func GetSchedule(pickup_times []int) [][]int {
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