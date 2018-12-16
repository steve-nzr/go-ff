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
		"192.168.99.100",
		[]channel{
			channel{
				"Channel 1",
				"192.168.99.100",
				500,
			},
		},
	},
}
