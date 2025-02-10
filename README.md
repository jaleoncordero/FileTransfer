# File Transfer

Just a quick utility to transfer files recursively from a source directory to a target directory.

---
### How to Use

Compile:

`make build`

Then, run the executable:

`./bin/FileTransfer [src] [dest] [mode]`

---
### Args

`[src]` - source directory. It can be absolute or relative.

`[dest]` - destination directory. It can be absoluate or relative.

`[mode]` (optional) - determines what file extensions to copy.

The following are currently supported modes:

- all
- image
- video
- audio
- pdf
- text
- presentation
- spreadsheet
- archive

If no mode is provided, the utility will default to mode `all`.

---
### Future Plans (no particular order)

1. Add custom mode where user can provide an explicit list of file extensions to copy.

2. Support separating files by type (in `all` or `custom` mode only).

3. Support persisting source directory structure.

3. Support compression of output.