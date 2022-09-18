---
sidebar_position: 1
---

# S3D Club CLI Tool
The `s3d` CLI tool is inspired by the many other great CLI commands we use on a
daily basis.

Our CLI tool is often a layer on top of other tools where we provide functions
that we feel can be represented in different ways or simplified based on the
opinionated view the S3D Club has for software development.

Our CLI uses [spf13/cobra](https://github.com/spf13/cobra#overview). All
commands have help screens explaining what they do. We will evententually also
have document pages on the [s3d.club](https://s3d.club/) page with a discussion
of how the commands support our opinionated workflow.

## Output Format
All output from our commands is JSON format _(with the exception of help and
errors)_.
