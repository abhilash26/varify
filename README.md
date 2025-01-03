# Varify

Varify is a command-line tool that converts a JSON configuration file into CSS variables. It allows for flexible configuration, minification, and customization of output paths, making it ideal for developers working with design tokens or scalable design systems.

## Features

* Parse a JSON configuration to generate CSS variables.
* Minify the output CSS for production-ready usage.
* Flexible input and output file paths via command-line flags.
* Customizable via JSON and command-line options.

## Installation

### Prerequisites
* Go 2.23 or higher

### Build the binary
1. Clone the repository
```
git clone https://github.com/abhilash26/varify.git
cd varify
go build -o varify main.go
```

