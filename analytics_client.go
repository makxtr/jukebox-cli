package main

import (
	"context"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "jukebox/cli/proto"
)

type AnalyticsClient struct {
	client pb.AnalyticsServiceClient
	conn   *grpc.ClientConn
}

func NewAnalyticsClient(addr string) (*AnalyticsClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := pb.NewAnalyticsServiceClient(conn)
	return &AnalyticsClient{
		client: client,
		conn:   conn,
	}, nil
}

func (c *AnalyticsClient) Close() {
	c.conn.Close()
}

func (c *AnalyticsClient) LogPlayback(trackID int, amountPaid float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.client.LogPlayback(ctx, &pb.LogPlaybackRequest{
		TrackId:    int32(trackID),
		AmountPaid: amountPaid,
	})
	return err
}
