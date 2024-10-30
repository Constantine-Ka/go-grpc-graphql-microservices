## Структура проекта

Проект состоит из следующих основных компонентов:

- Служба учетных записей
- Служба каталогов
- Служба заказов
- GraphQL API Gateway

Каждая служба имеет свою собственную базу данных:
- Службы учетных записей и заказов используют PostgreSQL
- Служба каталогов использует Elasticsearch

## Начало работы

1. Клонируем репозиторий:
    ```
   git clone <repository-url>
   cd <project-directory>
   ```

2. Запустите службы с помощью Docker Compose:
    ```
   docker-compose up -d --build
   ```

3. Зайдите в GraphQL playground по ссылке `http://localhost:8000/playground`

## Использование GraphQL API

GraphQL API предоставляет единый интерфейс для взаимодействия со всеми микросервисами.
### Запрос учетных записей

```graphql
query {
  accounts {
    id
    name
  }
}
```

### Создать учетную запись

```graphql
mutation {
  createAccount(account: {name: "New Account"}) {
    id
    name
  }
}
```

### Запрос продуктов
```graphql
query {
  products {
    id
    name
    price
  }
}
```

### Создать продукт
```graphql
mutation {
  createProduct(product: {name: "New Product", description: "A new product", price: 19.99}) {
    id
    name
    price
  }
}
```

### Создать заказ
```graphql
mutation {
  createOrder(order: {accountId: "account_id", products: [{id: "product_id", quantity: 2}]}) {
    id
    totalPrice
    products {
      name
      quantity
    }
  }
}
```

### Запрос учетной записи с заказами
```graphql
query {
  accounts(id: "account_id") {
    name
    orders {
      id
      createdAt
      totalPrice
      products {
        name
        quantity
        price
      }
    }
  }
}
```

## Расширенные запросы
### Разбивка на страницы и фильтрация
```graphql
query {
  products(pagination: {skip: 0, take: 5}, query: "search_term") {
    id
    name
    description
    price
  }
}
```
### Подсчитать общую сумму, потраченную по счету
```graphql
query {
  accounts(id: "account_id") {
    name
    orders {
      totalPrice
    }
  }
}
```