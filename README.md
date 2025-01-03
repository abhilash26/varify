# Varify

Varify is a command-line tool that converts a JSON configuration file into CSS variables. It allows for flexible configuration, minification, and customization of output paths, making it ideal for developers working with design tokens or scalable design systems.

## Features

- Parse a JSON configuration to generate CSS variables.
- Minify the output CSS for production-ready usage.
- Flexible input and output file paths via command-line flags.
- Customizable via JSON and command-line options.

## Installation

### Prerequisites

- Go 2.23 or higher

### Build the binary

1. Clone the repository and build the binary

```sh
git clone https://github.com/abhilash26/varify.git
cd varify
go build -o varify main.go
```

2. Build the binary

```sh
cd varify
go build -o varify main.go
```

3. Optionally, move the binary to a directory in your `PATH`:

```sh
mv varify /usr/local/bin
```

## Usage

### Command Line Options

```sh
Usage:
  varify [options]

Options:
  -i <path>      Path to the input JSON file (default: ./varify.json)
  -o <path>      Path to the output CSS file (overrides JSON output setting)
  -m, --minify   Enable minification of the output CSS
  -h             Show this help message

Example:
  varify -i ./input.json -o ./output.css -m

```

### Import Valid JSON Structure

```json
{
  "output": "./varify.min.css",
  "selector": ":root",
  "modules": {
    "text": {
      "xs": "clamp(0.625rem, 0.5rem + 0.3125vw, 0.75rem)",
      "sm": "clamp(0.75rem, 0.625rem + 0.3125vw, 0.875rem)"
    },
    "color": {
      "primary": "#3498db",
      "secondary": "#2ecc71"
    }
  }
}
```

### Example Usage

1. Convert JSON to CSS without minification:

```sh
varify -i ./varify.json -o ./output.css
```

2. Minify the generated CSS:

```sh
varify -i ./varify.json -o ./output.min.css -m
```

3. Use default paths:

- Input: `./varify.json`
- Output: As specified in the JSON file.

```sh
varify
```

### Generated Output Example

Given the above JSON, the output would look like this:

```css
:root {
  --text-xs: clamp(0.625rem, 0.5rem + 0.3125vw, 0.75rem);
  --text-sm: clamp(0.75rem, 0.625rem + 0.3125vw, 0.875rem);
  --color-primary: #3498db;
  --color-secondary: #2ecc71;
}
```

If the --minify flag is used:

```css
:root {
  --text-xs: clamp(0.625rem, 0.5rem+0.3125vw, 0.75rem);
  --text-sm: clamp(0.75rem, 0.625rem+0.3125vw, 0.875rem);
  --color-primary: #3498db;
  --color-secondary: #2ecc71;
}
```

## Inspirations

- [Pollen](pollen.style)
- [Tailwindcss](https://tailwindcss.com/)
- [Open Props](https://open-props.style/)

## Contribution

Contributions are welcome! Feel free to fork the repository, create a feature branch, and submit a pull request.

## License

This project is licensed under the [MIT License](https://github.com/abhilash26/varify/blob/main/LICENSE).

## Support

For any questions or issues, please [open an issue](https://github.com/abhilash26/varify/issues) on GitHub.
