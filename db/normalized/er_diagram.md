```mermaid
erDiagram

    %% Таблица пользователей
    users {
        bigint id PK
        email
        username
        city
        age
        avatar_url
        password
        blocked
        blocked_until
        created_at
        updated_at
    }

    %% Таблица продавцов
    sellers {
        bigint id PK
        name
        logo
        rating
        created_at
        updated_at
    }

    %% Таблица категорий
    categories {
        bigint id PK
        name
        picture
        active
        link_to
        created_at
        updated_at
    }

    %% Таблица товаров
    products {
        bigint id PK
        bigint seller_id FK
        bigint category_id FK
        count
        price
        original_price
        discount
        title
        description
        rating
        image_url
        active
        created_at
        updated_at
    }

    %% Таблица характеристик
    characteristics {
        bigint id PK
        name
        created_at
        updated_at
    }

    %% Связующая таблица товаров и характеристик
    product_characteristics {
        bigint product_id PK, FK
        bigint characteristic_id PK, FK
        value
        created_at
        updated_at
    }

    %% Таблица опций
    options {
        bigint id PK
        name
        created_at
        updated_at
    }

    %% Связующая таблица товаров и опций
    product_options {
       bigint  id PK
        product_id FK
        option_id FK
        created_at
        updated_at
    }

    %% Таблица значений опций
    product_option_values {
        bigint id PK
        product_option_id FK
        value
        link
        created_at
        updated_at
    }

    %% Таблица избранных товаров
    favorites {
        bigint id PK
        user_id FK
        product_id FK
        created_at
        updated_at
    }

    %% Таблица изображений товаров
    product_images {
        bigint id PK
        product_id FK
        image_url
        created_at
        updated_at
    }

    %% Таблица складских адресов
    stock_address {
        bigint id PK
        address
        city
        country
        postal_code
        created_at
        updated_at
    }

    %% Тип данных: order_status
    %% Создание типа ENUM в ER-диаграмме не отображается, поэтому статус указан как атрибут

    %% Таблица заказов
    orders {
        bigint id PK
        user_id FK
        address
        stock_address_id FK
        status
        delivery_price
        created_at
        updated_at
    }

    %% Таблица связей товаров и заказов
    product_orders {
        bigint id PK
        order_id FK
        product_id FK
        option_value_id FK
        count
        delivery_date
        created_at
        updated_at
    }

    %% Таблица корзины
    carts {
        bigint id PK
        user_id FK
        product_id FK
        option_value_id FK
        count
        is_selected
        delivery_date
        created_at
        updated_at
    }

    %% Таблица отзывов о товарах
    product_reviews {
        bigint id PK
        product_id FK
        user_id FK
        rating
        comment
        created_at
        updated_at
    }

    %% Определение связей между таблицами

    users ||--o{ orders : "размещает"
    users ||--o{ favorites : "имеет"
    users ||--o{ carts : "владеет"
    users ||--o{ product_reviews : "оставляет"

    sellers ||--o{ products : "предлагает"

    products ||--o{ product_characteristics : "имеет"
    products ||--o{ product_options : "имеет"
    products ||--o{ product_images : "содержит"
    products ||--o{ product_orders : "заказывается в"
    products ||--o{ carts : "добавлен в"
    products ||--o{ favorites : "в избранном"
    products ||--o{ product_reviews : "имеет отзывы"
    products ||--|| categories : "относится к"

    characteristics ||--o{ product_characteristics : "определяет"

    options ||--o{ product_options : "связана с"

    product_options ||--o{ product_option_values : "имеет значения"

    stock_address ||--o{ orders : "используется в"

    orders ||--o{ product_orders : "содержит"

    product_option_values ||--o{ product_orders : "используется в"
    product_option_values ||--o{ carts : "выбрана в"

    categories ||--o{ products : "содержит"

    product_reviews }o--|| users : "написан"
    product_reviews }o--|| products : "оценка для"

    carts }o--|| users : "принадлежит"
    carts }o--|| products : "содержит"
    carts }o--|| product_option_values : "имеет опцию"

    favorites }o--|| users : "принадлежит"
    favorites }o--|| products : "содержит"
```
