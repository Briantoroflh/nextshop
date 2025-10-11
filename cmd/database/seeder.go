package database

import (
	"log"
	"nextshop/entities"
	"time"

	"gorm.io/gorm"
)

// SeedDatabase inserts sample data for development/testing
func SeedDatabase(db *gorm.DB) error {
	log.Println("🌱 Seeding database with sample data...")

	// Seed Users (use entity's HashPassword function)
	users := []entities.Users{
		{
			FullName: stringPtr("Admin User"),
			Email:    "admin@example.com",
			Phone:    stringPtr("+1234567890"),
			Role:     entities.UserRoleAdmin,
			IsActive: true,
		},
		{
			FullName: stringPtr("John Seller"),
			Email:    "seller@example.com",
			Phone:    stringPtr("+1234567891"),
			Role:     entities.UserRoleSeller,
			IsActive: true,
		},
		{
			FullName: stringPtr("Jane Buyer"),
			Email:    "buyer@example.com",
			Phone:    stringPtr("+1234567892"),
			Role:     entities.UserRoleBuyer,
			IsActive: true,
		},
	}

	// Default seed password (change as needed)
	defaultPassword := "password123"

	for i := range users {
		// Hash password using entity method
		if err := users[i].HashPassword(defaultPassword); err != nil {
			return err
		}

		// Use FirstOrCreate with email as unique key
		if err := db.FirstOrCreate(&users[i], entities.Users{Email: users[i].Email}).Error; err != nil {
			return err
		}
	}

	// Seed Categories
	categories := []entities.Categories{
		{
			Name:        "Electronics",
			Slug:        "electronics",
			Description: stringPtr("Electronic devices and gadgets"),
		},
		{
			Name:        "Fashion",
			Slug:        "fashion",
			Description: stringPtr("Clothing and accessories"),
		},
		{
			Name:        "Books",
			Slug:        "books",
			Description: stringPtr("Books and educational materials"),
		},
	}

	for _, category := range categories {
		if err := db.FirstOrCreate(&category, entities.Categories{Slug: category.Slug}).Error; err != nil {
			return err
		}
	}

	// Seed Payment Methods
	paymentMethods := []entities.PaymentMethods{
		{
			Name:        "Bank Transfer",
			Description: stringPtr("Transfer via online banking"),
			IsActive:    true,
		},
		{
			Name:        "Credit Card",
			Description: stringPtr("Payment via credit/debit card"),
			IsActive:    true,
		},
		{
			Name:        "E-Wallet",
			Description: stringPtr("Digital wallet payment"),
			IsActive:    true,
		},
	}

	for _, method := range paymentMethods {
		if err := db.FirstOrCreate(&method, entities.PaymentMethods{Name: method.Name}).Error; err != nil {
			return err
		}
	}

	// Seed Shipping Methods
	shippingMethods := []entities.ShippingMethods{
		{
			Name:          "Standard Shipping",
			Description:   stringPtr("3-5 business days delivery"),
			Cost:          10.00,
			EstimatedDays: intPtr(5),
			IsActive:      true,
		},
		{
			Name:          "Express Shipping",
			Description:   stringPtr("1-2 business days delivery"),
			Cost:          25.00,
			EstimatedDays: intPtr(2),
			IsActive:      true,
		},
		{
			Name:          "Same Day Delivery",
			Description:   stringPtr("Same day delivery within city"),
			Cost:          50.00,
			EstimatedDays: intPtr(1),
			IsActive:      true,
		},
	}

	for _, method := range shippingMethods {
		if err := db.FirstOrCreate(&method, entities.ShippingMethods{Name: method.Name}).Error; err != nil {
			return err
		}
	}

	log.Println("🎉 Database seeding completed successfully!")
	return nil
}

// Helper functions for pointer types
func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

func timePtr(t time.Time) *time.Time {
	return &t
}
