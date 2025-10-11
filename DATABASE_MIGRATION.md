# Database Migration Guide

This guide explains how to use GORM auto migration in the NextShop e-commerce application.

## Overview

The application uses GORM auto migration to automatically create and update database schema based on the entity definitions. This eliminates the need for manual SQL migration files for basic schema changes.

## Features

- **Auto Migration**: Automatically creates tables and updates schema
- **Enum Support**: Handles PostgreSQL custom types (enums)
- **Foreign Key Constraints**: Maintains referential integrity
- **Soft Deletes**: Supports soft delete functionality
- **Seeding**: Provides sample data for development/testing
- **CLI Tools**: Command-line interface for database operations

## Entity Files

The following entities are automatically migrated:

1. `Users` - User accounts with roles (admin, seller, buyer)
2. `Categories` - Product categories with hierarchical structure
3. `Products` - Product listings with seller and category relations
4. `ProductImages` - Product image management
5. `Addresses` - User delivery addresses
6. `PaymentMethods` - Available payment methods
7. `ShippingMethods` - Shipping options with costs
8. `Orders` - Order management with status tracking
9. `OrderItems` - Individual items within orders
10. `Reviews` - Product reviews and ratings

## Usage

### Automatic Migration (Recommended)

The application automatically runs migration when starting:

```bash
go run main.go
```

This will:

- Connect to the database
- Create necessary PostgreSQL enums
- Auto-migrate all entities
- Start the web server

### Manual Migration Commands

Use the CLI tool for advanced database operations:

```bash
# Run migration only
go run migrate.go -migrate

# Reset database (drop all tables and recreate)
go run migrate.go -reset

# Seed database with sample data
go run migrate.go -seed

# Drop all tables
go run migrate.go -drop

# Reset and seed in one command
go run migrate.go -reset -seed
```

### Environment Variables

Make sure your `.env` file contains the database configuration:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=nextshop
DB_SSLMODE=disable
```

## Database Schema

### Enums Created

- `user_role`: 'buyer', 'seller', 'admin'
- `order_status`: 'pending', 'paid', 'shipped', 'delivered', 'cancelled'
- `payment_status`: 'pending', 'success', 'failed', 'refunded'

### Key Relationships

- Users can be buyers, sellers, or admins
- Categories support parent-child hierarchy
- Products belong to sellers and categories
- Orders link buyers, addresses, payment, and shipping methods
- Order items connect orders to products
- Reviews link users to products

## Migration Order

Entities are migrated in the correct order to respect foreign key dependencies:

1. Users (no dependencies)
2. Categories (self-referencing)
3. Products (depends on Users, Categories)
4. ProductImages (depends on Products)
5. Addresses (depends on Users)
6. PaymentMethods (no dependencies)
7. ShippingMethods (no dependencies)
8. Orders (depends on Users, Addresses, PaymentMethods, ShippingMethods)
9. OrderItems (depends on Orders, Products)
10. Reviews (depends on Products, Users)

## Development Workflow

### Initial Setup

```bash
# 1. Set up environment variables
cp .env.example .env
# Edit .env with your database credentials

# 2. Create database
createdb nextshop

# 3. Run migration and seed
go run migrate.go -reset -seed

# 4. Start application
go run main.go
```

### Schema Changes

When you modify entity definitions:

```bash
# Auto migration will detect changes
go run main.go

# Or run migration explicitly
go run migrate.go -migrate
```

### Testing with Clean Data

```bash
# Reset database and add sample data
go run migrate.go -reset -seed
```

## Best Practices

1. **Always backup** production databases before running migrations
2. **Test migrations** in development environment first
3. **Use transactions** for complex data migrations
4. **Monitor performance** on large tables during migration
5. **Version control** your entity changes

## Troubleshooting

### Common Issues

**Error: relation does not exist**

- Run migration: `go run migrate.go -migrate`

**Error: type does not exist**

- Reset database: `go run migrate.go -reset`

**Error: foreign key constraint**

- Check entity relationships and migration order

**Error: database connection failed**

- Verify database credentials in `.env`
- Ensure PostgreSQL is running

### Logging

The migration system provides detailed logging:

- ✅ Success operations
- ⚠️ Warnings for non-critical issues
- ❌ Errors that stop migration

## Production Considerations

For production deployments:

1. **Backup first**: Always backup before migration
2. **Staged deployment**: Test on staging environment
3. **Downtime planning**: Some migrations may require downtime
4. **Rollback plan**: Have a rollback strategy ready
5. **Monitoring**: Monitor application after migration

## Advanced Usage

### Custom Migration Logic

If you need custom migration logic beyond auto-migration:

```go
// Add to database/database.go
func CustomMigration(db *gorm.DB) error {
    // Custom migration logic here
    return nil
}
```

### Data Migration

For complex data transformations:

```go
// Add to database/seeder.go
func MigrateExistingData(db *gorm.DB) error {
    // Data migration logic here
    return nil
}
```
