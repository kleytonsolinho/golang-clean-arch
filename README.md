# GOLang Clean Architecture APP

Este programa em Go é um servidor HTTP/GraphQL e gRPC rodando em Multithreads com Clean Archictecture

## Funcionalidades

- **Cadastro de Orders:**
  - O servidor permite cadastrar Ordens utilizando HTTP, GraphQL e gRPC.

- **Listagem de Orders:**
  - O servidor permite listar todas as Ordens utilizando HTTP, GraphQL e gRPC.


## Como Utilizar localmente:

1. **Requisitos:** 
   - Certifique-se de ter o Go instalado em sua máquina.
   - Certifique-se de ter Evans e gRPC instalado em sua máquina.
   - Certifique-se de ter GraphQL (gqlgen) instalado em sua máquina.
   - Certifique-se de ter o Docker instalado em sua máquina.
&nbsp;
2. **Clonar o Repositório:**
&nbsp;

```bash
git clone https://github.com/kleytonsolinho/golang-clean-arch.git
```
&nbsp;
3. **Acesse a pasta do app:**
&nbsp;

```bash
cd golang-clean-arch
```
&nbsp;
4. **Rode o docker para subir o serviço RabbitMQ e MySQL:**
&nbsp;

```bash 
docker-compose up -d
```
5. **Acesse a pasta cmd/orderssystem e rode o main.go e wire_gen.go:**
&nbsp;

```bash 
cd cmd/ordersystem
```

```bash 
go run main.go wire_gen.go
```

&nbsp;
5. **Rode o Evans para acessar o gRPC:**
&nbsp;

```bash
evans -r repl
```

## Como testar localmente:

### Portas
HTTP server on port :8000 <br />
gRPC server on port 50051 <br />
GraphQL server on port 8080

### HTTP
 - Acesse a pasta api/ e dispare os dois arquivos
** Necessário instalar o plugin REST Client no VSCode. **

### GraphQL
 - Abra a página do GraphQL na porta 8080 
 <a href="http:localhost:8080" target="_blank">http:localhost:8080</a>

### gRPC
 - Rode o evans:

```bash
evans -r repl
```
```bash
evans package pb
```
```bash
evans service OrderService
```
```bash
evans call ListOrders
```
