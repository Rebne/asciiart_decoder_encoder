# Line Art Encoder/Decoder

This is a simple web application for encoding and decoding line art. It allows users to encode a line art drawing into a compressed format and decode it back to its original form.

## Features

- **Encoding:** Converts a line art drawing into a compressed format by replacing consecutive characters with a count and character combination, enclosed in square brackets.
- **Decoding:** Reverts the encoded line art back to its original form.
- **Web Interface:** Provides a user-friendly web interface for users to input their line art and choose between encoding and decoding.

## Installation

1. Clone the repository:  
   git clone https://gitea.kood.tech/rene-anterohogren/art.git
2. Navigate to project directory
3. Run the code:  
   go run main.go
4. Access the application in your web browser at http://localhost:80

## Usage

1. Input your line art into the provided text area.  
2. Choose whether you want to encode or decode the line art.  
3. Click the "Generate" button.
4. The result will be displayed on a new page.
5. Click "Generate another" to generate another one.

