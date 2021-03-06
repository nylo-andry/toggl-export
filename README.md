# Togglsheet

[![Go Report Card](https://goreportcard.com/badge/github.com/nylo-andry/togglsheet)](https://goreportcard.com/report/github.com/nylo-andry/togglsheet)

> A simple command-line tool that exports your toggl timesheet to CSV

## Installation

`go install github.com/nylo-andry/togglsheet/cmd/togglsheet`

## Getting started

1. Create a `.togglsheet` file by copying/renaming the `togglsheet.example`.

    - Move that file to your home directory.
    - Fill in three values (`api_token`, `workspace_id`, `user_name`).
    - You can find your API token in your [Profile settings](https://toggl.com/app/profile)
    - If you do not know your `workspace_id`, you can find it by running the `workspace` command (e.g. `togglsheet workspace`). You will still need to provide your API Token in your configuration file before you can find it.

2. Run `togglsheet export` provide the timeframe you want to export (e.g. `export --start 2018-12-17 --end 2018-12-21`)
