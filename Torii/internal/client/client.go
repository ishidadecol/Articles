package client

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	conn   net.Conn
	reader *bufio.Reader
}

func New(address string) (*Client, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}

	return &Client{
		conn:   conn,
		reader: bufio.NewReader(conn),
	}, nil
}

func (c *Client) Send(message string) (string, error) {
	// send message (ensure newline for server)
	_, err := fmt.Fprintf(c.conn, "%s\n", message)
	if err != nil {
		return "", err
	}

	// read response
	resp, err := c.reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return resp, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}

