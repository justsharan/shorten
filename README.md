# shorten

Shorten is a tiny, speedy URL shortener written in pure Go, using just the standard library and no outside modules whatsoever. It stores the list of sites in a file called `routes.txt` and serves them.

### Installation

You can get the appropriate executable for your system by visiting the [Releases](https://github.com/justsharan/shorten/releases) page.

### API

- `GET`
  - `GET /`: Simple Hello World check to see if the program is online.
  - `GET /:key`: Go to the URL with that key.
- `POST`
  - `POST /:key`: Create a new shortened URL with that key.
- `DELETE /:key`: Delete the URL with that key.

`POST` requests are treated as upserts. If you `POST` to a pre-existing URL, that key will simply be updated. This is simpler than having separate `POST` and `PUT` endpoints for adding and updating keys.

### Usage

When the program starts up, it'll give you an auth token to use for `POST` and `DELETE` endpoints. Be sure to save this in a safe place. A new one will be generated each time the program restarts. Here's an example using my own instance:

```sh
# Creating a new shortened URL
$ curl -X POST cybg.cf/nyt \
       -H "Authorization: AUTH_TOKEN_HERE" \
       -d "https://nytimes.com/section/todayspaper"

# Deleting a pre-existing URL
$ curl -X DELETE cybg.cf/nyt \
       -H "Authorization: AUTH_TOKEN_HERE"
```
