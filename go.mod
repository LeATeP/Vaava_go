module vaava/guard

go 1.18

replace vaava/psql => ./psql

replace vaava/utils => ./utils

require vaava/psql v0.0.0-00010101000000-000000000000

require github.com/lib/pq v1.10.4 // indirect
