responses
=========

`responses` is a common library for creating `net/http` responses. Useful for web services using the builtin `/net/http` library or with Gorilla Mux.

The errors returned by this library uses the JsonError format.

## Note

Exporting this publicly so that legacy code can be easily pulled by people who don't have access to the repository.

This will be replaced later on using the [terrors](https://gitlab.com/talino-labs/core/terrors) library.