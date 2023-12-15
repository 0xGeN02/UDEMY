# Funcionalidades de las consultas SQL generadas por SQLC

Este documento describe las funcionalidades de las consultas SQL generadas por SQLC en la carpeta `simpleBank/db/sqlc`.

## [account.sql.go](account.sql.go)

### CreateRandomAccount

La función `CreateRandomAccount` crea una cuenta aleatoria utilizando los métodos `RandomOwner`, `RandomMoney` y `RandomCurrency` del paquete `util`. Luego, utiliza el método `CreateAccount` del objeto `queries` para insertar la cuenta en la base de datos.

```go
func CreateRandomAccount(queries *sqlc.Queries) (sqlc.Account, error)
```

### DeleteRandomAccount

La función `DeleteRandomAccount` elimina una cuenta aleatoria utilizando el método ``DeleteAccount`` del objeto ``queries``.

```go
func DeleteRandomAccount(queries *sqlc.Queries, account sqlc.Account) error
```

### GetRandomAccount

La función `GetRandomAccount` obtiene una cuenta aleatoria utilizando el método ``GetAccount`` del objeto ``queries``.

```go
func GetRandomAccount(queries *sqlc.Queries, account1 sqlc.Account) (sqlc.Account, error)
```

### ListRandomAccounts

La función `ListRandomAccounts` lista cuentas aleatorias utilizando el método ``ListAccounts`` del objeto ``queries``. Acepta parámetros `limit` y `offset` para la paginación.

```go
func ListRandomAccounts(queries *sqlc.Queries, limit, offset int32) ([]sqlc.Account, error)
```

## [db.go](db.go)

### DBTX

`DBTX` es una interfaz que define los métodos necesarios para ejecutar consultas SQL y transacciones en una base de datos. Los métodos incluyen `ExecContext`, `PrepareContext`, `QueryContext` y `QueryRowContext`.

```go
type DBTX interface {
    ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
    PrepareContext(context.Context, string) (*sql.Stmt, error)
    QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
    QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}
```

### New

La función `New` toma un objeto que implementa la interfaz `DBTX` y devuelve un nuevo objeto ``Queries`` que puede ser utilizado para ejecutar consultas SQL en la base de datos.

```go
func New(db DBTX) *Queries {
    return &Queries{db: db}
}
```

## [entries.sql.go](entries.sql.go)

Este documento describe las funcionalidades de las consultas SQL generadas por SQLC en el archivo `entries.sql.go`.

### CreateEntry

La función `CreateEntry` inserta una nueva entrada en la tabla `entries` con un `account_id` y `amount` especificados. Retorna el `id`, `account_id`, `amount` y `created_at` de la entrada creada.

```go
func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error)
```

### GetEntry

La función `GetEntry` selecciona una entrada de la tabla `entries` basada en un `id` especificado. Retorna el ``id``, ``account_id``, ``amount`` y ``created_at`` de la entrada seleccionada.

```go
func (q *Queries) GetEntry(ctx context.Context, id int64) (Entry, error)
```

### ListEntries

La función ``ListEntries`` selecciona una lista de entradas de la tabla ``entries`` basada en un ``account_id`` especificado. También acepta parámetros limit y offset para la paginación. Retorna una lista de entradas, cada una con ``id``, ``account_id``, ``amount`` y ``created_at``.

```go
    func (q *Queries) ListEntries(ctx context.Context, arg ListEntriesParams) ([]Entry, error)
```

## [models.go](models.go)

### Account

La estructura `Account` representa una ``cuenta`` en la base de datos.

```go
type Account struct {
    ID        int64     `json:"id"`
    Owner     string    `json:"owner"`
    Balance   int64     `json:"balance"`
    Currency  string    `json:"currency"`
    CreatedAt time.Time `json:"created_at"`
}
```

### Entry

La estructura `Entry` representa una ``entrada`` en la base de datos.

```go

type Entry struct {
    ID        int64 `json:"id"`
    AccountID int64 `json:"account_id"`
    Amount    int64     `json:"amount"` // can be positive or negative
    CreatedAt time.Time `json:"created_at"`
}
```

### Transfer

La estructura `Transfer` representa una ``transferencia`` en la base de datos.

```go

type Transfer struct {
    ID            int64 `json:"id"`
    FromAccountID int64 `json:"from_account_id"`
    ToAccountID   int64 `json:"to_account_id"`
    Amount int64 `json:"amount"` // must be positive
    CreatedAt time.Time `json:"created_at"` // the time when the transfer was created
}
```

## [store.go](store.go)

### Store

La estructura `Store` proporciona una forma de ejecutar consultas SQL y transacciones en una base de datos.

```go
type Store struct {
    *Queries
    db *sql.DB
}
```

Puedes crear una nueva Store con la función

```go
NewStore(db *sql.DB) *Store 
```

### execTx

La función `execTx` ejecuta una función dentro de una ``transacción`` de base de datos.

```go
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error
```

### TransferTxParams y TransferTxResult

TransferTxParams contiene los parámetros de entrada de la transacción de transferencia y TransferTxResult es el resultado de la transacción de transferencia.

```go

type TransferTxParams struct {
    FromAcccountID int64 `json:"from_account_id"`
    ToAccountID    int64 `json:"to_account_id"`
    Amount         int64 `json:"amount"`
}

type TransferTxResult struct {
    Transfer    Transfer `json:"transfer"`
    FromAccount Account  `json:"from_account"`
    ToAccount   Account  `json:"to_account"`
    FromEntry   Entry    `json:"from_entry"`
    ToEntry     Entry    `json:"to_entry"`
}
```

### TransferTx

La función `TransferTx` realiza una transferencia de dinero de una cuenta a otra. Crea un registro de ``transferencia`` y actualiza el ``saldo`` de la cuenta dentro de una única ``transacción`` de base de datos.

```go
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error)
```

## [transfers.sql.go](transfers.sql.go)

### CreateTransfer

La función `CreateTransfer` inserta una nueva transferencia en la base de datos. Acepta un objeto `CreateTransferParams` que contiene los detalles de la transferencia y devuelve un objeto `Transfer` con los detalles de la transferencia creada.

```go
func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error)
```

### GetTransfer

La función `GetTransfer` recupera los detalles de una ``transferencia`` específica de la base de datos. Acepta un ``id`` de transferencia y devuelve un objeto ``Transfer`` con los detalles de la ``transferencia``.

```go
    func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error)
```

### ListTransfers

La función `ListTransfers` recupera una lista de transferencias de la base de datos. Acepta un objeto ``ListTransfersParams`` que contiene los detalles para ``filtrar`` y ``paginar`` los resultados, y devuelve una ``lista`` de objetos Transfer.

```go
func (q *Queries) ListTransfers(ctx context.Context, arg ListTransfersParams) ([]Transfer, error)
```
