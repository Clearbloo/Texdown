# Texdown

A simple markdown to latex compiler in Go.

## Why?

For a long time, I wrote my notes in LaTeX, this was great and looked beautiful when compiled. However it is a bit of a faff to write, and on my old machine, it was quite slow to compile.

Recently I have started taking my daily notes in Obsidian, and I would love a unified way to take my notes so I don't have to switch between the two. It's also nice to have the on the spot markdown preview from Obsidian.

This compiler aims to help a workflow of doing everything in plain markdown. The compiler then handles the conversion of any more formal notes into a tex and then into a pdf. I also like using Neovim to edit my notes, but I don't love the VimTex plugin, so writing in markdown is a nice alternative.

The idea is that you can segment your book however you like. I like using separate pages as chapters. You would the pass the markdown documents to the compiler which converts each page into a chapter in a single .tex file.

## Current status
Only have a bare minimum example that converts simple markdown

# Known issues
When passing the src as `./foo.md` rather than `foo.md` the inferred output is split on the first period, which outputs on file `.tex`

## Todo
- [ ] Remove the need for a flag, and just pass a positional argument
- [ ] Add multiline math compilation
- [ ] Add better errors
- [ ] Support for backlinks (from obsidian)
- [ ] Support for multiple files into one tex document
- [ ] Add a LaTeX template
