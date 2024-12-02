# Create build directory if it doesn't exist
mkdir -p build

# Supported platforms
platforms=("windows/amd64" "linux/amd64" "darwin/amd64")

# Build for all platforms
for platform in "${platforms[@]}"; do
    # Split platform into OS and architecture
    IFS="/" read -r OS ARCH <<< "${platform}"
    
    # Create OS/arch specific directory
    mkdir -p "build/${OS}_${ARCH}"
    
    echo "Building for $OS/$ARCH..."
    
    # Set environment variables for cross-compilation
    export GOOS=$OS
    export GOARCH=$ARCH

    # Set output filename with .exe for Windows
    OUTPUT_FILE="ayana"
    if [ "$OS" = "windows" ]; then
        OUTPUT_FILE="${OUTPUT_FILE}.exe"
    fi

    if [ "$1" = "--release" ]; then
        go build -ldflags="-s -w" -o "build/${OS}_${ARCH}/${OUTPUT_FILE}" ./ayana.go
    else
        go build -gcflags="all=-N -l" -o "build/${OS}_${ARCH}/${OUTPUT_FILE}" ./ayana.go
    fi

    # Copy web folder and config.toml to build directory
    echo "Copying resources for $OS/$ARCH..."
    cp -r web "build/${OS}_${ARCH}/"
    cp config.toml "build/${OS}_${ARCH}/"
done

echo "Build complete!"