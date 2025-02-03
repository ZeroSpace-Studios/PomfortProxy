# PomfortProxy

A proxy service that enables compatibility between Pomfort and disguise systems, providing seamless integration for production workflows.

## Overview
PomfortProxy is a specialized proxy service that resolves compatibility issues between Pomfort software and disguise systems. It ensures smooth data flow and communication between these production tools.

## Features
- Seamless Pomfort-disguise integration
- Real-time data translation
- Protocol compatibility handling
- Production workflow optimization
- Low-latency operation

## Technical Details
Built using:
- Go for high-performance networking
- Custom protocol handling
- Real-time data processing
- Error recovery mechanisms

## Installation
```bash
go install github.com/ZeroSpace-Studios/PomfortProxy@latest
```

## Usage
```bash
# Start the proxy service
pomfort-proxy [options]
```

## Configuration
The proxy can be configured for:
- Custom port settings
- Protocol specifications
- Connection parameters
- Logging preferences

## Project Structure
```
├── main.go     # Core proxy implementation
└── go.mod     # Project dependencies
```

## Status
- Currently active and working in production
- Maintains compatibility with latest versions
- Regular updates for new features

Download, unzip and run from a terminal as follows from the machine running livegrade.
```
./PomfortProxy <ip-of-director>
```

Now just add a disguise device to a slot in livegrade studio and point it to 127.0.0.1.
