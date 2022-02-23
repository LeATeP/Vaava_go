module github.com/leatep/mining

go 1.17

replace vaava/psql => ../../psql

replace vaava/utils => ../../utils

replace vaava/drop => ../../package/drop

require (
	vaava/drop v0.0.0-00010101000000-000000000000
	vaava/psql v0.0.0-00010101000000-000000000000
)

require (
	github.com/lib/pq v1.10.4 // indirect
	vaava/utils v0.0.0-00010101000000-000000000000 // indirect
)
