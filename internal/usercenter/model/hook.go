package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/mindmatterlab/go-pro/internal/pkg/zid"
	"github.com/mindmatterlab/go-pro/pkg/authn"
)

// User status constants.
const (
	StatusUserDisabled = iota // Status used for disabling a user.
	StatusUserNormal          // Status used for enabling a user.
)

// Secret status constants.
const (
	StatusSecretDisabled = iota // Status used for disabling a secret.
	StatusSecretNormal          // Status used for enabling a secret.
)

// BeforeCreate runs before creating a SecretM database record and initializes various fields.
func (s *SecretM) BeforeCreate(tx *gorm.DB) (err error) {
	// Supports custom SecretKey, but users need to ensure the uniqueness of the SecretKey themselves.
	// gopro-cacheserver will use this feature to set secret.
	if s.SecretID == "" {
		// Generate a new UUID for SecretKey.
		s.SecretID = uuid.New().String()
	}

	// Generate a new UUID for SecretID.
	s.SecretKey = uuid.New().String()

	// Set the default status for the secret as normal.
	s.Status = StatusSecretNormal

	return nil
}

// AfterCreate runs after creating a UserM database record and updates the UserID field.
func (u *UserM) AfterCreate(tx *gorm.DB) (err error) {
	u.UserID = zid.User.New(uint64(u.ID)) // Generate and set a new user ID.

	return tx.Save(u).Error // Save the updated user record to the database.
}

// BeforeCreate runs before creating a UserM database record and initializes various fields.
func (u *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	u.Password, err = authn.Encrypt(u.Password) // Encrypt the user password.
	if err != nil {
		return err // Return error if there's a problem with encryption.
	}

	u.Status = StatusUserNormal // Set the default status for the user as normal.

	return nil
}
