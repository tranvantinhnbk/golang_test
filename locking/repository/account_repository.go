package repository

import (
	"context"
	"errors"
	"fmt"
	"golang_test/locking/models"
	"log"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// AccountRepository handles database operations for Account model
type AccountRepository struct {
	db *gorm.DB
}

// NewAccountRepository creates a new AccountRepository instance
func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

// CreateAccount creates a new account if it doesn't exist
func (r *AccountRepository) CreateAccount(ctx context.Context, username string, initialBalance float64) (*models.Account, error) {
	// Start transaction
	tx := r.db.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return nil, fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Check if account exists
	var existingAccount models.Account
	err := tx.Where("username = ?", username).First(&existingAccount).Error
	if err == nil {
		tx.Rollback()
		return nil, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		return nil, fmt.Errorf("failed to check existing account: %w", err)
	}

	// Create new account
	account := &models.Account{
		Username: username,
		Balance:  initialBalance,
		Version:  1,
	}

	if err := tx.Create(account).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("failed to create account: %w", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return account, nil
}

// Deposit adds money to an account's balance
func (r *AccountRepository) DepositPessimistic(ctx context.Context, accountID uint, amount float64) error {
	// Start transaction
	tx := r.db.WithContext(ctx).Begin()
	if err := tx.Error; err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Get account with lock
	var account models.Account
	if err := tx.Clauses(clause.Locking{
		Strength: "UPDATE",
	}).First(&account, accountID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to get account: %w", err)
	}

	// Update balance
	account.Balance += amount
	account.Version++
	if err := tx.Save(&account).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update balance: %w", err)
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// DepositOptimistic updates the account balance using optimistic locking
func (r *AccountRepository) DepositOptimistic(ctx context.Context, accountId uint, amount float64) error {
	maxRetries := 3
	retryCount := 0

	for retryCount < maxRetries {
		// Start transaction
		tx := r.db.WithContext(ctx).Begin()
		if err := tx.Error; err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		// Get current account state within transaction
		var account models.Account
		if err := tx.First(&account, accountId).Error; err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to get account: %w", err)
		}

		// Update balance and version
		account.Balance += amount
		account.Version++

		// Try to save with version check within transaction
		// Use Updates instead of Save to ensure version check
		result := tx.Model(&models.Account{}).
			Where("id = ? AND version = ?", accountId, account.Version-1).
			Updates(map[string]interface{}{
				"balance": account.Balance,
				"version": account.Version,
			})

		if result.Error != nil {
			tx.Rollback()
			return fmt.Errorf("failed to save account: %w", result.Error)
		}

		// Check if update was successful
		if result.RowsAffected > 0 {
			// Commit transaction
			if err := tx.Commit().Error; err != nil {
				return fmt.Errorf("failed to commit transaction: %w", err)
			}
			return nil
		}

		// If no rows were affected, rollback and retry
		tx.Rollback()
		retryCount++
		if retryCount < maxRetries {
			log.Printf("Optimistic lock failed for account %d, attempt %d/%d, retrying...", accountId, retryCount, maxRetries)
			time.Sleep(time.Millisecond * 100) // Small delay before retry
		} else {
			return fmt.Errorf("optimistic lock failed after %d attempts for account %d", maxRetries, accountId)
		}
	}

	return fmt.Errorf("max retries exceeded for account %d", accountId)
}
