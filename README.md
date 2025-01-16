# CodePort

CodePort is a CLI tool for sharing and downloading code snippets. It allows users to upload code files to a server and retrieve them using a unique ID.

## Features

- Upload code files to a server with syntax highlighting.
- Download code files using a unique ID.
- Password protection for uploads.
- Customizable titles and descriptions for code snippets.

## Installation

To install CodePort, clone the repository and build the CLI tool:

```bash
git clone https://github.com/xspoilt-dev/codeport.git
cd codeport
chmod +x install.sh
./install.sh
```

## Usage

### Uploading Code

To upload a code file to the server, use the following command:

```bash
./codeport -f <file_path> -l <language> -t <title> -d <description> -p <password>
```

- `-f`: Path to the file to upload (required)
- `-l`: Programming language for syntax highlighting (default: python)
- `-t`: Title for the paste (optional)
- `-d`: Description of the paste (optional)
- `-p`: Password to protect the paste (optional)

Example:

```bash
./codeport -f example.py -l python -t "Example Code" -d "This is an example code snippet." -p "mypassword"
```

### Downloading Code

To download a code file from the server, use the following command:

```bash
./codeport -g <paste_id> -o <output_file>
```

- `-g`: Paste ID to download the file (required)
- `-o`: Path to save the downloaded file (default: output.txt)

Example:

```bash
./codeport -g abc123 -o downloaded_code.py
```

## Contributors

- Minhajul Islam
- Farhan Ali

## License

This project is licensed under the MIT License.

## Contact

For more information, visit [https://pasteit.ftoolz.xyz](https://pasteit.ftoolz.xyz).

---

You can create the README.md file using this content in your repository. Would you like me to proceed with creating the file for you?
