# Api-Rest-Golang
This is a simple implementation of a Rest API using only the standard library (except the database drive) of the Golang language.

For simplicity the API has the following limitations:
- Each valid request is a connection to the database
- Competing requests were not handled

## The problem

The requirements specification can be found [here](problem/especificação)

- [ERD](db/ERD.png)
- [Physical data model](db/ini.sql)

## Routes
*/api/v1/gravadoras*
*/api/v1/albuns*
*/api/v1/faixas*
*/api/v1/compositores*
*/api/v1/interpretes*
*/api/v1/playlists*
*/api/v1/gravadora/id*
*/api/v1/gravadora/id/albuns*
*/api/v1/albuns/id*
*/api/v1/albuns/id/faixas*
*/api/v1/faixas/id*
*/api/v1/faixas/id/interpretes*
*/api/v1/faixas/id/compositores*
*/api/v1/compositores/id*
*/api/v1/compositores/id/*
*/api/v1/playlists/id*
*/api/v1/playlists/id/faixas*

[Here](example/main.py) is an example of consuming the api
