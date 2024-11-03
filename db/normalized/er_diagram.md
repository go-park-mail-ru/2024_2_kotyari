```mermaid
erDiagram

    %% Таблица пользователей
    users {
        BIGINT id PK
        TEXT email
        TEXT username
        TEXT city
        SMALLINT age
        TEXT avatar_url
        TEXT password
        BOOLEAN blocked
        TIMESTAMP blocked_until
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица продавцов
    sellers {
        BIGINT id PK
        TEXT name
        TEXT logo
        REAL rating
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица категорий
    categories {
        BIGINT id PK
        TEXT name
        TEXT picture
        BOOLEAN active
        TEXT link_to
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица товаров
    products {
        BIGINT id PK
        BIGINT seller_id FK
        BIGINT category_id FK
        INTEGER count
        INTEGER price
        INTEGER original_price
        SMALLINT discount
        TEXT title
        TEXT description
        REAL rating
        TEXT image_url
        BOOLEAN active
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица характеристик
    characteristics {
        BIGINT id PK
        TEXT name
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Связующая таблица товаров и характеристик
    product_characteristics {
        BIGINT product_id PK, FK
        BIGINT characteristic_id PK, FK
        TEXT value
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица опций
    options {
        BIGINT id PK
        TEXT name
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Связующая таблица товаров и опций
    product_options {
        BIGINT id PK
        BIGINT product_id FK
        BIGINT option_id FK
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица значений опций
    product_option_values {
        BIGINT id PK
        BIGINT product_option_id FK
        TEXT value
        TEXT link
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица избранных товаров
    favorites {
        BIGINT id PK
        BIGINT user_id FK
        BIGINT product_id FK
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица изображений товаров
    product_images {
        BIGINT id PK
        BIGINT product_id FK
        TEXT image_url
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица складских адресов
    stock_address {
        BIGINT id PK
        TEXT address
        TEXT city
        TEXT country
        TEXT postal_code
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица заказов
    orders {
        BIGINT id PK
        BIGINT user_id FK
        TEXT address
        BIGINT stock_address_id FK
        TEXT status
        INTEGER delivery_price
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица связей товаров и заказов
    product_orders {
        BIGINT id PK
        BIGINT order_id FK
        BIGINT product_id FK
        BIGINT option_value_id FK
        INTEGER count
        TIMESTAMP delivery_date
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица корзины
    carts {
        BIGINT id PK
        BIGINT user_id FK
        BIGINT product_id FK
        BIGINT option_value_id FK
        INTEGER count
        BOOLEAN is_selected
        TIMESTAMP delivery_date
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Таблица отзывов о товарах
    product_reviews {
        BIGINT id PK
        BIGINT product_id FK
        BIGINT user_id FK
        SMALLINT rating
        TEXT comment
        TIMESTAMP created_at
        TIMESTAMP updated_at
    }

    %% Определение связей между таблицами

    users ||--o{ orders : размещает
    users ||--o{ favorites : имеет
    users ||--o{ carts : владеет
    users ||--o{ product_reviews : оставляет

    sellers ||--o{ products : предлагает

    categories ||--o{ products : содержит

    products ||--o{ product_characteristics : имеет
    products ||--o{ product_options : имеет
    products ||--o{ product_images : содержит
    products ||--o{ product_orders : заказывается в
    products ||--o{ carts : добавлен в
    products ||--o{ favorites : в избранном
    products ||--o{ product_reviews : имеет отзывы
    products ||--|| categories : относится к

    characteristics ||--o{ product_characteristics : определяет

    options ||--o{ product_options : связана с

    product_options ||--o{ product_option_values : имеет значения

    stock_address ||--o{ orders : используется в

    orders ||--o{ product_orders : содержит

    product_option_values ||--o{ product_orders : используется в
    product_option_values ||--o{ carts : выбрана в

    product_reviews }o--|| users : написан
    product_reviews }o--|| products : оценка для

    carts }o--|| users : принадлежит
    carts }o--|| products : содержит
    carts }o--|| product_option_values : имеет опцию

    favorites }o--|| users : принадлежит
    favorites }o--|| products : содержит

```
