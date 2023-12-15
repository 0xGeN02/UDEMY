# Simple Bank

## Descripción General

Simple Bank es una aplicación de banca básica construida con el propósito de aprender y practicar el desarrollo de aplicaciones en Go. La aplicación permite a los usuarios crear cuentas, realizar transacciones y consultar sus balances.

## db/lib

En esta carpeta, encontramos las siguientes archivos:

- `entries_lib.go`: Esta libreria se encuentra en [entries_lib.go](lib/entries_lib.go). Contiene la funcionalidad para la query entries.
- `transfers_lib.go`: Esta libreria se encuentra en [transfers_lib.go](lib/transfers_lib.go). Contiene la funcionalidad para la query tranfers.
- `account_lib.go`: Esta libreria se encuentra en [transfers_lib.go](lib/account_lib.go). Contiene la funcionalidad para la query account.

## db/migration

Esta carpeta contiene los scripts de migración de la base de datos. Los scripts de migración son:

- `000001_init_schema.down.sql`: Este script se utiliza para revertir los cambios realizados por `000001_init_schema.up.sql`.
- `000001_init_schema.up.sql`: Este script se utiliza para inicializar el esquema de la base de datos.

## db/query

Esta carpeta contiene las consultas SQL que se utilizan en la aplicación. Las consultas son:

- `account.sql`: Este archivo contiene las consultas SQL para crear, obtener, listar, actualizar y eliminar cuentas.
- `entries.sql`: Este archivo contiene las consultas SQL para crear, obtener, listar, actualizar y eliminar entradas.
- `transfers.sql`: Este archivo contiene las consultas SQL para crear, obtener, listar, actualizar y eliminar transacciones.

## db/sqlc

Esta carpeta contiene el código Go generado por SQLC a partir de las consultas SQL. Los archivos generados son:

- `account.sql.go`: Este archivo contiene el código Go generado a partir de las consultas SQL en `account.sql`.
- `entries.sql.go`: Este archivo contiene el código Go generado a partir de las consultas SQL en `entries.sql`.
- `transfers.sql.go`: Este archivo contiene el código Go generado a partir de las consultas SQL en `transfers.sql`.

## db/test

Esta carpeta contiene los tests unitarios para las funciones en `db/lib` y `db/sqlc`. Los archivos de test son:

- `account_test.go`: Este archivo contiene los tests para las funciones en `account_lib.go` y `account.sql.go`.
- `entries_test.go`: Este archivo contiene los tests para las funciones en `entries_lib.go` y `entries.sql.go`.
- `main_test.go`: Este archivo contiene los tests para las funciones en `main.go`.

## db/util

Esta carpeta contiene funciones de utilidad como:

- `random.go`: Este archivo contiene funciones para generar datos aleatorios que se utilizan en los tests.
