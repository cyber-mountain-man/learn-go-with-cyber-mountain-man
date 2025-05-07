
# Create a new user
Write-Host "🆕 Creating a new user..."
curl -X POST http://localhost:8080/users `
  -H "Content-Type: application/json" `
  -d '{"name": "Alice", "email": "alice@example.com"}'
Write-Host "`n"

# Get all users
Write-Host "📋 Getting all users..."
curl http://localhost:8080/users
Write-Host "`n"

# Get user by ID
Write-Host "🔍 Getting user with ID 1..."
curl http://localhost:8080/users/1
Write-Host "`n"

# Update user by ID
Write-Host "✏️ Updating user with ID 1..."
curl -X PUT http://localhost:8080/users/1 `
  -H "Content-Type: application/json" `
  -d '{"name": "Alice Updated", "email": "alice.updated@example.com"}'
Write-Host "`n"

# Delete user by ID
Write-Host "❌ Deleting user with ID 1..."
curl -X DELETE http://localhost:8080/users/1
Write-Host "`n"

# Get all users again
Write-Host "📋 Getting all users (after delete)..."
curl http://localhost:8080/users
Write-Host "`n"
