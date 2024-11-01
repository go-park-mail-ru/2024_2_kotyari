erDiagram
    %% Таблица пользователей
    users {
        bigint id PK
        text email
        text username
        text city
        smallint age
        text avatar_url
        text password
        boolean blocked
        timestamp blocked_until
        timestamp created_at
        timestamp updated_at
    }

    %% Таблица продавцов
    sellers {
        bigint id PK
        text name
        text logo
        real rating
        timestamp created_at
        timestamp updated_at
    }

    %% Таблица товаров
    products {
        bigint id PK
        bigint seller_id FK
        integer count
        integer price
        integer original_price
        smallint discount
        text title
        text description
        real rating
        text image_url
        boolean active
        timestamp created_at
        timestamp updated_at
    }

    %% Таблица характеристик
    characteristics {
        bigint id PK
        text name
    }

    %% Связующая таблица товаров и характеристик
    product_characteristics {
        bigint product_id PK, FK
        bigint characteristic_id PK, FK
        text value
    }

    %% Таблица опций
    options {
        bigint id PK
        text name
    }

    %% Связующая таблица товаров и опций
    product_options {
        bigint id PK
        bigint product_id FK
        bigint option_id FK
    }

    %% Таблица значений опций
    product_option_values {
        bigint id PK
        bigint product_option_id FK
        text value
        text link
    }

    %% Таблица избранных товаров
    favorites {
        bigint id PK
        bigint user_id FK
        bigint product_id FK
        timestamp created_at
    }

    %% Таблица изображений товаров
    product_images {
        bigint id PK
        bigint product_id FK
        text image_url
        timestamp created_at
    }

    %% Таблица складских адресов
    stock_address {
        bigint id PK
        text address
        text city
        text country
        text postal_code
    }

    %% Таблица заказов
    orders {
        bigint id PK
        bigint user_id FK
        text address
        bigint stock_address_id FK
        integer total_price
        order_status status
        timestamp created_at
        timestamp updated_at
    }

    %% Таблица связей товаров и заказов
    product_orders {
        bigint id PK
        bigint order_id FK
        bigint product_id FK
        bigint option_value_id FK
        integer count
        timestamp delivery_date
    }

    %% Таблица корзины
    carts {
        bigint id PK
        bigint user_id FK
        bigint product_id FK
        bigint option_value_id FK
        integer count
        boolean is_selected
        timestamp created_at
        timestamp updated_at
        timestamp delivery_date
    }

    %% Таблица категорий
    categories {
        bigint id PK
        text name
        text picture
        boolean active
        text link_to
    }

    %% Связующая таблица товаров и категорий
    product_categories {
        bigint id PK
        bigint product_id FK
        bigint category_id FK
        boolean active
    }

    %% Таблица отзывов о товарах
    product_reviews {
        bigint id PK
        bigint product_id FK
        bigint user_id FK
        smallint rating
        text comment
        timestamp created_at
        timestamp updated_at
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
    products ||--o{ product_orders : "входит в"
    products ||--o{ carts : "добавлен в"
    products ||--o{ favorites : "в избранном"
    products ||--o{ product_reviews : "имеет отзывы"
    products }o--o{ categories : "относится к"

    characteristics ||--o{ product_characteristics : "определяет"

    options ||--o{ product_options : "связана с"

    product_options ||--o{ product_option_values : "имеет значения"

    stock_address ||--o{ orders : "используется в"

    orders ||--o{ product_orders : "содержит"

    product_option_values ||--o{ product_orders : "используется в"
    product_option_values ||--o{ carts : "выбрана в"

    categories ||--o{ product_categories : "содержит"

    product_categories }o--|| products : "связан с"

    product_reviews }o--|| users : "написан"
    product_reviews }o--|| products : "оценка для"

    carts }o--|| users : "принадлежит"
    carts }o--|| products : "содержит"
    carts }o--|| product_option_values : "имеет опцию"

    favorites }o--|| users : "принадлежит"
    favorites }o--|| products : "содержит"
