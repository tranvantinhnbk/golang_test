package main

import (
	"context"
	"golang_test/locking/db"
	"golang_test/locking/repository"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// Initialize database connection
	config := db.DBConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "your_password",
		DBName:   "testdb",
	}

	dbConn, err := db.GetDBInstance(config)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	// Initialize schema
	err = dbConn.InitSchema()
	if err != nil {
		log.Fatalf("Failed to initialize schema: %v", err)
	}

	// Create repository
	accountRepo := repository.NewAccountRepository(dbConn.DB)

	// Create a new account
	ctx := context.Background()
	account, err := accountRepo.CreateAccount(ctx, "test_user", 0)
	if err != nil {
		log.Fatalf("Failed to create account: %v", err)
	}
	log.Printf("Created account with ID: %d, Initial balance: %.2f\n", account.ID, account.Balance)

	// Simulate 100 concurrent deposits
	var wg sync.WaitGroup
	depositAmount := 1.0 // Each deposit adds 1.0
	numTransactions := 100
	var failedCount int32 // Atomic counter for failed transactions

	// Start timing
	startTime := time.Now()

	// Launch concurrent deposits
	for i := 0; i < numTransactions; i++ {
		wg.Add(1)
		go func(transactionID int) {
			defer wg.Done()
			err := accountRepo.DepositOptimistic(ctx, account.ID, depositAmount)
			if err != nil {
				atomic.AddInt32(&failedCount, 1)
				log.Printf("Transaction %d failed: %v\n", transactionID, err)
			} else {
				log.Printf("Transaction %d completed successfully\n", transactionID)
			}
		}(i)
	}

	// Wait for all transactions to complete
	wg.Wait()

	// Calculate final balance and time taken
	duration := time.Since(startTime)
	log.Printf("\nTransaction Summary:")
	log.Printf("Total transactions: %d", numTransactions)
	log.Printf("Failed transactions: %d", atomic.LoadInt32(&failedCount))
	log.Printf("Successful transactions: %d", numTransactions-int(failedCount))
	log.Printf("Time taken: %v\n", duration)
	log.Printf("Expected final balance: %.2f\n", float64(numTransactions)*depositAmount)
}
