trigger
==================================================

A command to send a detected card ID to another commands or urls in json.

## building

```
$ make
```

## usage

### prepare

1. build this
2. copy `trigger` to raspberry pi

### example

For example, to send a card ID to `http://example.com/`:

```shell
./trigger get:http://example.com/
````

A following request will be sent to server:

```
GET /?readerData=%7B%22Time%22%3A1577360661%2C%22Type%22%3A%22TypeA%22%2C%22ID%22%3A%22aa%3Aaa%3Aaa%3Aaa%22%2C%22Random%22%3Atrue%7D HTTP/1.1
Host: example.com
User-Agent: Go-http-client/1.1
Accept-Encoding: gzip
```

Other exmaples:

```shell
# to send a card ID to example.com by POST
./trigger http://example.com/
```

```shell
# to send a card ID to example.com by GET
./trigger get:http://example.com/
```

```shell
# to send a card ID to an another program (./program)
./trigger ./program
```

```shell
# to send a card ID to example.com and to an another program also
./trigger get:http://example.com/ ./program
```

```shell
# detect felica card only
./trigger -type f get:http://example.com/
```

## limitations
* read only first 4 bytes when reading ID from nfc-a cards

