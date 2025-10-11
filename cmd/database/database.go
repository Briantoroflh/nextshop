package database

import (
	"fmt"
	"log"

	"nextshop/cmd/config"
	"nextshop/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	cfg := config.GetDBConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Port, cfg.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to database: %v", err)
	}

	// Run auto migration
	err = AutoMigrate(db)
	if err != nil {
		log.Fatalf("❌ Failed to run auto migration: %v", err)
	}

	log.Println("✅ Database connection successful and migrations completed!")

	DB = db
}

// AutoMigrate runs GORM auto migration for all entities
func AutoMigrate(db *gorm.DB) error {
	log.Println("🔄 Running auto migration...")

	// Create custom types first (enums)
	err := db.Exec(`
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'user_role') THEN
				CREATE TYPE user_role AS ENUM ('buyer', 'seller', 'admin');
			END IF;
		END $$;
	`).Error
	if err != nil {
		log.Printf("⚠️  Warning: Could not create user_role enum (might already exist): %v", err)
	}

	err = db.Exec(`
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status') THEN
				CREATE TYPE order_status AS ENUM ('pending', 'paid', 'shipped', 'delivered', 'cancelled');
			END IF;
		END $$;
	`).Error
	if err != nil {
		log.Printf("⚠️  Warning: Could not create order_status enum (might already exist): %v", err)
	}

	err = db.Exec(`
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_status') THEN
				CREATE TYPE payment_status AS ENUM ('pending', 'success', 'failed', 'refunded');
			END IF;
		END $$;
	`).Error
	if err != nil {
		log.Printf("⚠️  Warning: Could not create payment_status enum (might already exist): %v", err)
	}

	// Auto migrate all entities in correct order (respecting foreign key dependencies)
	entities := []interface{}{
		&entities.Users{},
		&entities.Categories{},
		&entities.Products{},
		&entities.ProductImages{},
		&entities.Addresses{},
		&entities.PaymentMethods{},
		&entities.ShippingMethods{},
		&entities.Orders{},
		&entities.OrderItems{},
		&entities.Reviews{},
	}

	for _, entity := range entities {
		if err := db.AutoMigrate(entity); err != nil {
			return fmt.Errorf("failed to migrate %T: %w", entity, err)
		}
		log.Printf("✅ Migrated: %T", entity)
	}

	log.Println("🎉 Auto migration completed successfully!")
	return nil
}

// DropAllTables drops all tables (useful for development/testing)
func DropAllTables(db *gorm.DB) error {
	log.Println("🗑️  Dropping all tables...")

	// Drop tables in reverse order to respect foreign key constraints
	entities := []interface{}{
		&entities.Reviews{},
		&entities.OrderItems{},
		&entities.Orders{},
		&entities.ShippingMethods{},
		&entities.PaymentMethods{},
		&entities.Addresses{},
		&entities.ProductImages{},
		&entities.Products{},
		&entities.Categories{},
		&entities.Users{},
	}

	for _, entity := range entities {
		if err := db.Migrator().DropTable(entity); err != nil {
			log.Printf("⚠️  Warning: Could not drop table for %T: %v", entity, err)
		} else {
			log.Printf("🗑️  Dropped table: %T", entity)
		}
	}

	// Drop custom types
	db.Exec("DROP TYPE IF EXISTS payment_status CASCADE;")
	db.Exec("DROP TYPE IF EXISTS order_status CASCADE;")
	db.Exec("DROP TYPE IF EXISTS user_role CASCADE;")

	log.Println("🎉 All tables dropped successfully!")
	return nil
}

// ResetDatabase drops all tables and runs auto migration again
func ResetDatabase(db *gorm.DB) error {
	log.Println("🔄 Resetting database...")

	if err := DropAllTables(db); err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	if err := AutoMigrate(db); err != nil {
		return fmt.Errorf("failed to auto migrate after reset: %w", err)
	}

	log.Println("🎉 Database reset completed successfully!")
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
