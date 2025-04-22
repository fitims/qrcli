# qrcli

A simple and CLI tool to **encode** text into a QR code image and **decode** text from a QR code image.

---

## Features

- Encode any text into a QR code and save it as a PNG image.
- Decode text from an existing QR code PNG image.
- Minimal and easy-to-use interface.

---

## Installation

```bash
# Clone the repository
$ git clone https://github.com/fitims/qrcli.git

# Move into the project directory
$ cd qrcli

# Build the binary (example for Go)
$ go build -o qrcli
```

Or install via Go:

```bash
$ go install github.com/fitims/qrcli@latest
```

---

## Usage

### Encode
Encode a string into a QR code PNG image.

```bash
$ qrcli encode -t "Hello, World!" -o output.png
```

**Options:**
- `-t, --text`   : The text you want to encode.
- `-o, --output` : The path to save the generated QR code PNG image.

### Decode
Decode text from a QR code PNG image.

```bash
$ qrcli decode -i input.png
```

**Options:**
- `-i, --input`  : The path to the PNG image containing the QR code.

---

## Examples

### Encoding Example
```bash
$ qrcli encode -t "https://example.com" -o website_qr.png
```
Creates a QR code pointing to `https://example.com` saved as `website_qr.png`.

### Decoding Example
```bash
$ qrcli decode -i website_qr.png
```
Outputs:
```
https://example.com
```

---

## Requirements
- Go 1.20+
- Libraries for QR code generation and decoding (e.g., `github.com/skip2/go-qrcode`, `github.com/tuotoo/qrcode`)

---

## License

MIT License. See [LICENSE](./LICENSE) for more information.

---

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

---

## Author

Created by [Your Name](https://github.com/yourusername).

