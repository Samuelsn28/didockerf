# didockerf

didockerf is a simple dockerfiles and compose files manager made in Go, it provides a CLI interface to make operations like saving these files to improve the managing.

### Features
- Save dockerfiles and compose files;
- Manage the saved dockerfiles and compose files;
- Retrieve any saved file;
-  Import and export saves.

## Requirements
- Golang 1.26.4

## Installation

Clone the didockerf repository to your local machine.

```
$ git clone https://github.com/Samuelsn28/didockerf didockerf-project
```

Access the created folder by Git and build the project with Golang to create didockerf's binary (`didockerf`).

```
$ cd didockerf-project
$ go build
```

### Linux - Turning the binary executable

Created the `didockerf` binary, make the bellow commands to move it to its directory and turning it executable. **The commands will turn the didockerf visible only for your current user**.

```
$ mkdir ~/.local/share/didockerf
$ mv didockerf ~/.local/share/didockerf/
$ chmod 744 ~/.local/share/didockerf/didockerf
$ ln -s ~/.local/share/didockerf/didockerf ~/.local/bin/didockerf
```
## Basic usage

Save a dockerfile:

```
$ didockerf dfile save <Dockerfile path> <name>:<tag>
```

List all saved dockerfiles:

```
$ didockerf dfile ls
```

Remove a saved dockerfile:

```
$ didockerf dfile rm <name>:<tag>
```

Save a compose file:

```
$ didockerf composefile save <Compose file path> <name>:<tag>
```

List all saved dockerfiles:

```
$ didockerf composefile ls
```

Remove a saved dockerfile:

```
$ didockerf composefile rm <name>:<tag>
```

The `--help` option shows all possible commands:

```
$ didockerf --help
```



