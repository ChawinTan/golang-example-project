module main

go 1.15

replace server/services => ../services

require (
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gorilla/mux v1.8.0
	server/services v0.0.0-00010101000000-000000000000
)
