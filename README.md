# Orchestra FM Scores Service

Recordkeeping.

![GitHub](https://img.shields.io/github/license/orchestrafm/scores?style=flat-square) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/orchestrafm/scores?style=flat-square) ![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/orchestrafm/scores?style=flat-square) ![GitHub language count](https://img.shields.io/github/languages/count/orchestrafm/scores?style=flat-square) ![GitHub top language](https://img.shields.io/github/languages/top/orchestrafm/scores?style=flat-square) [![Go Report Card](https://goreportcard.com/badge/github.com/orchestrafm/scores?style=flat-square)](https://goreportcard.com/report/github.com/orchestrafm/scores) [![tickgit](https://badgen.net/https/api.tickgit.com/badgen/github.com/orchestrafm/scores?style=flat-square)](https://badgen.net/https/api.tickgit.com/badgen/github.com/orchestrafm/scores)

## Requirements
- Go 1.12.17+
- [Task 2+](https://taskfile.dev/) 
- [packr 1.30.1+](https://github.com/gobuffalo/packr/)
- MySQL 8+
- [Profiles Service](https://github.com/orchestrafm/profiles)
- [Tracks Service](https://github.com/orchestrafm/tracks)
- [Scores Service](https://github.com/orchestrafm/scores)

## Configuration
The following variables MUST be defined by the environment in order for the application to run properly.
```
MYSQL_DB
MYSQL_HOST
MYSQL_USER
MYSQL_PASS
```

## Development Setup
1. Run `task buiild`, this will automatically pack and embed migrations into the final binary.
2. Ensure the following environment variables listed in [Configuration](#configuration).
3. Execute the application, the migrations will run at startup.