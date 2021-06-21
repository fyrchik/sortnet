package internal

//go:generate go run ../main.go -n 4 -out networks/batcher4.go -sort batcher -name Batcher4
//go:generate go run ../main.go -n 4 -out networks/bitonic4.go -sort bitonic -name Bitonic4
//go:generate go run ../main.go -n 8 -out networks/batcher8.go -sort batcher -name Batcher8
//go:generate go run ../main.go -n 8 -out networks/bitonic8.go -sort bitonic -name Bitonic8
//go:generate go run ../main.go -n 16 -out networks/batcher16.go -sort batcher -name Batcher16
//go:generate go run ../main.go -n 16 -out networks/bitonic16.go -sort bitonic -name Bitonic16
