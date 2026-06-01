# Go File Sorter
Simple CLI utility written in Go for sorting files by extension.


## Features
- Sort files by extension
- Automatically create category folders
- Handle unknown file types
- Case-intensitive extension matching
- Display sorting statistics

- Move files into:
  - images
  - audio
  - docs
  - archives
  - video
  - unknown


## Roadmap
- [x] v1.0 Basic file sorting

- [x] v1.1 More extensions
- [x] v1.1 Unknown folder
- [x] v1.1 Statistics

- [ ] v1.2 Dry Run
- [ ] v1.2 Filename conflict handling

- [ ] v1.3 Confirmation before execution
- [ ] v1.3 JSON configuration
- [ ] v1.3 Logging with timestamps

- [ ] v1.4 Copy + delete mode
- [ ] v1.4 Undo last sorting
- [ ] v1.4 Recursive folder sorting

- [ ] v1.5 Watch mode


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
Current version: v1.1