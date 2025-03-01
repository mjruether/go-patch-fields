#!/bin/bash

# Ensure we're on the main branch
git checkout main
git pull origin main

# Prompt for version
read -p "Enter the new version (e.g., v0.1.0): " VERSION

# Validate version format
if [[ ! $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
    echo "Invalid version format. Use v0.0.0"
    exit 1
fi

# Run tests
go test ./...

# Run linters
golangci-lint run

# Create git tag
git tag $VERSION
git push origin $VERSION

# Create GitHub release (requires GitHub CLI)
gh release create $VERSION \
    --title "$VERSION Release" \
    --notes "Release of gopatchfields $VERSION" \
    --generate-notes

echo "Release $VERSION created successfully!"
