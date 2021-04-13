package service

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/arijitnayak92/go-lang-projects/BIGRPC/proto"
	"github.com/twinj/uuid"
)

type NotificationService interface {
	PushNotification() (string, error)
}

type NotificationCient struct {
	client pb.NotificationClient
}

func NewNotificationClient(client pb.NotificationClient) *NotificationCient {
	return &NotificationCient{client: client}
}

func (n *NotificationCient) PushNotification() (string, error) {

	stream, err := n.client.PushNotification(context.Background())
	if err != nil {
		log.Fatalf("openn stream error %v", err)
	}

	ctx := stream.Context()
	done := make(chan bool)

	// first goroutine sends random increasing numbers to stream
	// and closes int after 10 iterations
	go func() {
		for i := 1; i <= 10; i++ {
			// generate random nummber and send it to stream

			req := pb.NotificationRequest{
				UserID:   uuid.NewV4().String(),
				TargetID: uuid.NewV4().String(),
				Message:  "new notification",
			}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	// second goroutine receives data from stream
	// and saves result in max variable
	//
	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			log.Println("new received", resp.NotificationID)
		}
	}()

	// third goroutine closes done channel
	// if context is done
	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()

	<-done
	log.Println("done !")
	return "", nil
}
