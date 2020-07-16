# Implementation of Fermat Prime Test in ~~Go~~ TinyGo for WebAssembly

This is part of a set of projects to compare how different programming languages build to WebAssembly. Other projects in this set include:

* [as-primecheck](https://github.com/ianmcodes/as-primecheck)
* [rust-primecheck](https://github.com/ianmcodes/rust-primecheck)

## Why?

WebAssembly is an exciting new technology in web development. It's a byte code that many languages can compile to, and that can run in the browser. This allows for doing processing on the client side that would run too slowly in pure JS.

This also gives me an opportunity to learn and try out languages that may not have found the time to otherwise.

## Why the Fermat Prime Test?

I think of the Fermat Prime Test as a nice middle point. It's more complex than a simple "Hello World" tutorial, but not so complex that it cannot be explained in a few minutes to someone with some programming experience and understanding of math(s).

## Anything else I should know?

Yes. If you are an expert in this language or in WebAssembly you can probably find a number of optimizations that can be made that would not necessarily be obvious to someone that is just starting. Feel free to point these out in an issue or pull request, but please don't be offended if I don't do anything about them.

Part of the comparison that I am trying to to make is how easy/hard it is for someone with out a deep understanding of the specific language to build something that works, and how well that thing works. This isn't intended to be mission critical production code. It's intended to be a tool to help make (as close as possible) an "apples to apples" comparison across very different languages and tool chains.

# Building

## Prerequisites

* tinygo >= 0.13

## Command

`tinygo build -o main.wasm -target wasm ./main.go`
