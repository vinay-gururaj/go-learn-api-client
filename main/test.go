package main

import (
	"context"
	"kwclient"
)

func main() {
	c, err := kwclient.NewClient("http://localhost:8001")

	if err != nil {
		panic(err)
	}

	println(c.SessionId)
	println(c.CsrfToken)

	projectId, error := c.CreateXPProject(context.Background(), kwclient.XPProject{Name: "Vinay Library Test", Description: "Test project created from visual studio"})

	if error != nil {
		panic(error)
	}

	println(projectId)

	c.CreateSQSConnection(context.Background(), kwclient.SQSConnection{Name: "SQS Connection", Key: "AXIA123456789YWB", Secret: "XvghjkoiQHJKLYYYYYT", ReadQueue: "read_queue", WriteQueue: "write_queue", ProjectId: projectId})

}
