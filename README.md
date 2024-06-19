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

Observação: Ao iniciar a aplicação, o processo de migração é executado automaticamente, não sendo necessário realizá-lo manualmente.

## Como testar localmente:

### Portas
HTTP server on port :8000 <br />
gRPC server on port :50051 <br />
GraphQL server on port :8080

### HTTP
 - Acesse a pasta api/ e dispare os dois arquivos
** Necessário instalar o plugin REST Client no VSCode. **

### GraphQL
 - Abra a página do GraphQL na porta 8080 e execute a mutation ou query abaixo:
 <a href="http://localhost:8080/" target="_blank">http://localhost:8080/</a>

 ```graphql
  mutation createOrder {
    createOrder (input:{ id: "teste", Price: 110, Tax: 2 }) {
      id
      Price
      Tax
    }
  }


  query queryOrders {
      orders {
          id
          Price
          Tax
          FinalPrice
      }
  }
 ```

### gRPC
 - Rode o evans:

```bash
evans -r repl
```
```bash
package pb
```
```bash
service OrderService
```

Para listar as orders utilize
```bash
call ListOrders
```

Para criar orders utilize
```bash
call CreateOrder
```