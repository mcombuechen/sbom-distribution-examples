package main

import "github.com/lib/pq"

func main() {
	_, err := pq.ParseURL("postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full")
	if err != nil {
		return
	}
}
