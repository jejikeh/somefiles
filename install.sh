#!/bin/bash

# Download the binary file
echo "Downloading binary file..."
if ! curl -o somefiles -L https://github.com/jejikeh/somefiles/releases/download/1.0/somefiles; then
    echo "Failed to download the binary file"
    exit 1
fi

# Make the binary file executable
chmod +x somefiles

# Run the binary file with arguments
echo "Running binary file with arguments..."
if ! ./somefiles -it=100 -chance=0.1; then
    echo "Failed to run the binary file"
    exit 1
fi

echo "Binary file execution completed successfully."
