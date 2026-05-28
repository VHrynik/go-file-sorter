# Go File Sorter
Simple CLI utility written in Go for sorting files by extension.


## Features
- Sort files by extension
- Automatically create category folders
- Move files into:
  - images
  - audio
  - docs
  - archives


## Usage for example
```bash
go run main.go "D:\Downloads"
```


## Example
Before:
test/
├── photo.png
├── song.mp3
├── archive.zip
└── notes.txt

After:
test/
├── images/
│   └── photo.png
├── audio/
│   └── song.mp3
├── archives/
│   └── archive.zip
└── docs/
    └── notes.txt


## Version
Current version: v1.0
