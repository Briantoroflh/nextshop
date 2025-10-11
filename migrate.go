package main

import (
	"flag"
	"fmt"
	"log"
	"nextshop/cmd/config"
	"nextshop/cmd/database"
	"os"
)

func main() {
	// Define command line flags
	var (
		migrate = flag.Bool("migrate", false, "Run auto migration")
		reset   = flag.Bool("reset", false, "Reset database (drop all tables and recreate)")
		seed    = flag.Bool("seed", false, "Seed database with sample data")
		drop    = flag.Bool("drop", false, "Drop all tables")
	)

	flag.Parse()

	// Load environment variables
	config.LoadEnv()

	// Initialize database connection
	database.InitDB()
	db := database.GetDB()

	// Execute based on flags
	switch {
	case *reset:
		fmt.Println("🔄 Resetting database...")
		if err := database.ResetDatabase(db); err != nil {
			log.Fatalf("❌ Failed to reset database: %v", err)
		}
		fmt.Println("✅ Database reset completed!")

	case *drop:
		fmt.Println("🗑️  Dropping all tables...")
		if err := database.DropAllTables(db); err != nil {
			log.Fatalf("❌ Failed to drop tables: %v", err)
		}
		fmt.Println("✅ All tables dropped!")

	case *migrate:
		fmt.Println("🔄 Running auto migration...")
		if err := database.AutoMigrate(db); err != nil {
			log.Fatalf("❌ Failed to run migration: %v", err)
		}
		fmt.Println("✅ Migration completed!")

	case *seed:
		fmt.Println("🌱 Seeding database...")
		if err := database.SeedDatabase(db); err != nil {
			log.Fatalf("❌ Failed to seed database: %v", err)
		}
		fmt.Println("✅ Database seeded!")

	default:
		fmt.Println("Database Migration CLI")
		fmt.Println("Usage:")
		fmt.Println("  go run migrate.go -migrate    # Run auto migration")
		fmt.Println("  go run migrate.go -reset      # Reset database (drop & recreate)")
		fmt.Println("  go run migrate.go -seed       # Seed with sample data")
		fmt.Println("  go run migrate.go -drop       # Drop all tables")
		fmt.Println("")
		fmt.Println("Examples:")
		fmt.Println("  go run migrate.go -reset -seed  # Reset and seed database")
		os.Exit(1)
	}

	// If seed flag is also provided with other commands
	if *seed && (*reset || *migrate) {
		fmt.Println("🌱 Seeding database...")
		if err := database.SeedDatabase(db); err != nil {
			log.Fatalf("❌ Failed to seed database: %v", err)
		}
		fmt.Println("✅ Database seeded!")
	}
}
