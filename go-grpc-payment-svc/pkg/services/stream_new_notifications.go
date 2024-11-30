package service

import (
	"fmt"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/models"
	"github.com/RonnieZad/nyumba-go-grpc-payment-svc/pkg/pb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

func (s *Server) StreamNewPaymentNotifications(stream pb.PaymentService_StreamNewPaymentNotificationsServer) error {
	// Read the user ID from the client request
	req, err := stream.Recv()
	if err != nil {
		// Handle error
		return err
	}

	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		// Handle error
		return err
	}

	// Create a channel for the client to receive new payment notifications
	notificationChannel := make(chan *pb.PaymentTransaction)

	// Simulate sending new payment notifications when they occur
	go s.listenForNewPaymentNotifications(userId, notificationChannel)

	for {
		select {
		case newPayment := <-notificationChannel:
			if err := stream.Send(newPayment); err != nil {
				// Handle error
				return err
			}
		case <-stream.Context().Done():
			// Client has closed the connection, cleanup and exit
			close(notificationChannel)
			return nil
		}
	}
}

func (s *Server) listenForNewPaymentNotifications(userId uuid.UUID, channel chan<- *pb.PaymentTransaction) {
	for {
		// Retrieve and generate new payment notifications
		newPayment, err := s.generateNewPaymentNotificationFromDB(userId)
		if err != nil {
			// Handle error
			fmt.Println("Error generating new payment notification:", err)
			time.Sleep(5 * time.Second) // Retry after a delay
			continue
		}

		// Simulate sending a new payment notification after 3 seconds
		time.Sleep(3 * time.Second)

		// Send the new payment notification to the client
		select {
		case channel <- newPayment:
		case <-time.After(5 * time.Second):
			// If sending is delayed, proceed to the next iteration
		}
	}
}

func (s *Server) generateNewPaymentNotificationFromDB(userId uuid.UUID) (*pb.PaymentTransaction, error) {
	// Query your database table to retrieve the latest new payment transaction for the user
	var paymentTransaction models.PaymentTransaction
	result := s.H.DB.Where("user_id = ?", userId).Order("created_at DESC").First(&paymentTransaction)
	if result.Error != nil {
		return nil, result.Error
	}

	// Convert the retrieved payment transaction to the gRPC message format
	newPayment := &pb.PaymentTransaction{
		Id:              paymentTransaction.Id.String(),
		UserId:          paymentTransaction.UserId.String(),
		Amount:          paymentTransaction.Amount,
		Currency:        paymentTransaction.Currency,
		TransactionType: paymentTransaction.TransactionType,
		Status:          paymentTransaction.Status,
		// Populate other fields accordingly
	}

	return newPayment, nil
}

func (s *Server) StartPaymentNotificationStreaming() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, s)

	fmt.Println("Payment Notification Service 🚀 on Port 50052")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
