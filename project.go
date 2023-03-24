package kwclient

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CreateXPProject - Create new CreateXPProject
func (c *Client) CreateXPProject(ctx context.Context, xpProject XPProject) (projectId string, err error) {
	rb, err := json.Marshal(xpProject)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/createkwProject", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return "", err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return "", err
	}

	xpProjectId := projectId
	err = json.Unmarshal(body, &xpProjectId)
	if err != nil {
		return "", err
	}

	return xpProjectId, nil
}

// CreateSQSConnection - Create new SQS Connection
func (c *Client) CreateSQSConnection(ctx context.Context, sqsConnection SQSConnection) {
	rb, err := json.Marshal(sqsConnection)
	if err != nil {
		panic(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/createSQSConnection", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		panic(err)
	}

	// The code has been commented below as the createSQSConnection method does not return a response
	// body, err := c.doRequest(req)

	c.doRequest(req)

	/* The code commented below is to ensure that there is no error in processing the request as the createSQSConnection response path does not return a response.
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &sqsConnection)
	if err != nil {
		panic(err)
	}
	*/
}
