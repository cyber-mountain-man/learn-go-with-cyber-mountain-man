# List of folders to create
$folders = @(
  "01-introduction-and-setup",
  "02-your-first-go-file",
  "03-variables-strings-numbers",
  "04-arrays-slices",
  "05-structs",
  "06-maps",
  "07-loops",
  "08-conditionals",
  "09-methods",
  "10-pointers",
  "11-packages",
  "12-standard-library",
  "13-writing-files",
  "14-reading-files",
  "15-go-routines",
  "16-channels",
  "17-web-servers",
  "18-json",
  "19-project-setup",
  "20-handlers",
  "21-routing",
  "22-controllers",
  "23-middleware",
  "24-templates",
  "25-forms",
  "26-databases",
  "27-sessions",
  "28-deployment"
)

foreach ($folder in $folders) {
    New-Item -Path $folder -ItemType Directory -Force | Out-Null

    # Create main.go
    $goCode = @"
package main

import "fmt"

func main() {
    fmt.Println("Placeholder for $folder")
}
"@
    Set-Content -Path "$folder\main.go" -Value $goCode

    # Create README.md
    $readme = "# $folder`n`nNotes and explanations for this lesson."
    Set-Content -Path "$folder\README.md" -Value $readme
}

# Create challenges folder
New-Item -Path "challenges" -ItemType Directory -Force | Out-Null
Set-Content -Path "challenges\README.md" -Value "# Challenges`n`nPractice problems go here."

Write-Host "✅ Folder structure and boilerplate files created successfully!"
