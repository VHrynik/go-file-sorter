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


## Roadmap
- [x] v1.0 Basic file sorting
- [ ] v1.1 More extensions
- [ ] v1.1 Unknown folder
- [ ] v1.1 Statistics
- [ ] v1.2 Dry Run
- [ ] v1.2 Filename conflict handling
- [ ] v1.5 JSON config
- [ ] v1.5 Logging


## Usage for example
```bash
go run main.go "D:\Downloads"
```


## Example
Before:
```text
test/
├── photo.png
├── song.mp3
├── archive.zip
└── notes.txt
```

After:
```text
test/
├── images/
│   └── photo.png
├── audio/
│   └── song.mp3
├── archives/
│   └── archive.zip
└── docs/
    └── notes.txt
```


## Version
Current version: v1.0
