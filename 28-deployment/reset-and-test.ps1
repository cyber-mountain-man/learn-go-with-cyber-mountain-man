# reset-and-test.ps1

# Step 1: Stop and remove containers, networks, and volumes
Write-Host "⛔ Stopping and removing containers and volumes..." -ForegroundColor Yellow
docker compose down -v

# Step 2: Build and start containers
Write-Host "🔄 Rebuilding and starting containers..." -ForegroundColor Cyan
docker compose up --build -d

# Step 3: Wait for mssql to become healthy
Write-Host "⏳ Waiting for SQL Server to become healthy..." -ForegroundColor Yellow
$healthy = $false
for ($i = 0; $i -lt 20; $i++) {
    $status = docker inspect --format='{{.State.Health.Status}}' 28-deployment-mssql-1 2>$null
    if ($status -eq "healthy") {
        $healthy = $true
        break
    }
    Start-Sleep -Seconds 5
}
if (-not $healthy) {
    Write-Host "❌ SQL Server did not become healthy in time." -ForegroundColor Red
    exit 1
}

# Step 4: Run a test query using mssql-tools container
Write-Host "✅ Running SQL connection test..." -ForegroundColor Green
docker run --rm --network 28-deployment_go-net mcr.microsoft.com/mssql-tools `
    /opt/mssql-tools/bin/sqlcmd `
    -S mssql -U lesson28_user -P lesson28_pass -d lesson28_mssql `
    -Q "SELECT GETDATE() AS ConnectedTime;"
