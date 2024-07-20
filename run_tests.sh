#!/bin/bash

# Set the output file for coverage data
COVERAGE_FILE="coverage.out"

# Run tests and collect coverage data
echo "Running tests and collecting coverage data..."
go test ./... -coverprofile=$COVERAGE_FILE

# Check if tests passed
if [ $? -ne 0 ]; then
  echo "Tests failed. Please fix the issues and try again."
  exit 1
fi

# Generate coverage report in HTML format
COVERAGE_HTML="coverage.html"
echo "Generating HTML coverage report..."
go tool cover -html=$COVERAGE_FILE -o $COVERAGE_HTML

# Open coverage report in the default browser (macOS)
if [[ "$OSTYPE" == "darwin"* ]]; then
  open $COVERAGE_HTML
# Open coverage report in the default browser (Linux)
elif [[ "$OSTYPE" == "linux-gnu"* ]]; then
  xdg-open $COVERAGE_HTML
# For other systems, print the location of the coverage report
else
  echo "Coverage report generated at $COVERAGE_HTML"
fi

echo "Done."
