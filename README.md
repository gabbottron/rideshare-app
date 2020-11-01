# rideshare-app
a rideshare app for an interview challenge question

## usage
go run ./src/main.go -help

## specify ride times and use defaults
go run ./src/main.go -pickup_times 5,15,55,205,208,213

## specify ride times with custom ride time and recovery time
go run ./src/main.go -pickup_times 5,15,55,205,208,213 -ride_time 1 -recovery_time 55
