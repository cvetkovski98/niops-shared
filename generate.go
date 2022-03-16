package niops_shared

//go:generate protoc -I=. --go_out=. ./protocols/order.proto
//go:generate protoc -I=. --go_out=. ./protocols/production.proto
