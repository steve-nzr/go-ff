package main

type channel struct {
	name      string
	ip        string
	maxPlayer uint32
}

type server struct {
	name     string
	ip       string
	channels []channel
}

var servers = []server{
	server{
		"Server 1",
		"127.0.0.1",
		[]channel{
			channel{
				"Channel 1",
				"127.0.0.1",
				500,
			},
		},
	}}
