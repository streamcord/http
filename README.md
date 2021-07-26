# HTTP

This is the currently unnamed Discord HTTP API wrapper package that is used in several of Streamcord's microservices.

The package features automatic ratelimit handling to prevent 429s and uses structs to represent objects such as embeds in a friendly format.

While the package itself does not directly interact with Streamcord's Redis cache, it is designed to be fully compatible with it.