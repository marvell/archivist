# Archivist

## Version 1

`GET /servers`
`GET /servers?zone=ru`
`GET /servers?zone=ru&component=back`

`GET /servers/:ip`
`GET /servers/:ip?only=tags.zone`
`POST /servers`
`PUT /servers/:ip`
`DELETE /servers/:ip`

`GET /servers/my`
`GET /servers/my?only=addr`
`POST /servers`
`PUT /servers/my`
`DELETE /servers/my`

```
{
  "name": "",
  "addr": "",
  "tags": {
    "zone": "ru"
  }
}
```
