CREATE TABLE IF NOT EXISTS users (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    email TEXT UNIQUE NOT NULL,
    username TEXT NOT NULL,
    city TEXT NOT NULL DEFAULT 'Москва',
    age SMALLINT,
    avatar_url TEXT,
    password TEXT NOT NULL,
    blocked BOOLEAN NOT NULL DEFAULT FALSE,
    blocked_until TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Таблица: sellers
-- Содержит информацию о продавцах, включая рейтинг и логотип.

CREATE TABLE IF NOT EXISTS sellers (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL,
    logo_url TEXT NOT NULL,
    rating REAL DEFAULT 0 NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Таблица: products
-- Хранит информацию о товарах.

CREATE TABLE IF NOT EXISTS products (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    seller_id BIGINT NOT NULL REFERENCES sellers(id) ON DELETE CASCADE,
    category_id BIGINT NOT NULL REFERENCES categories(id),
    count INTEGER NOT NULL DEFAULT 1,
    price INTEGER NOT NULL,
    original_price INTEGER,
    discount SMALLINT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    rating REAL DEFAULT 0 NOT NULL,
    image_url TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Таблица: characteristics и product_characteristics
-- Хранит характеристики товаров.

CREATE TABLE IF NOT EXISTS characteristics (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_characteristics (
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    characteristic_id BIGINT NOT NULL REFERENCES characteristics(id),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    value TEXT NOT NULL,
    PRIMARY KEY (product_id, characteristic_id)
);

-- Таблица: options, product_options и product_option_values
-- Хранит опции товаров, такие как размер или цвет.

CREATE TABLE IF NOT EXISTS options (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS product_options (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    option_id BIGINT NOT NULL REFERENCES options(id),
    UNIQUE (product_id, option_id)
);

CREATE TABLE IF NOT EXISTS product_option_values (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_option_id BIGINT NOT NULL REFERENCES product_options(id) ON DELETE CASCADE,
    value TEXT NOT NULL,
    link TEXT,
    UNIQUE (product_option_id, value)
);

-- Таблица: favorites
-- Хранит избранные товары пользователей.

CREATE TABLE IF NOT EXISTS favorites (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    product_id BIGINT NOT NULL REFERENCES products(id),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (user_id, product_id)
);

-- Таблица: product_images
-- Хранит дополнительные изображения товаров.

CREATE TABLE IF NOT EXISTS product_images (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    image_url TEXT NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (product_id, image_url)
);

-- Таблица: stock_address
-- Хранит адреса складов для доставки.

CREATE TABLE IF NOT EXISTS stock_address (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    address TEXT NOT NULL,
    city TEXT NOT NULL,
    country TEXT NOT NULL,
    postal_code TEXT NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (address, city, country, postal_code)
);

-- Тип данных: order_status
-- Представляет статус заказа.

CREATE TYPE order_status AS ENUM ('awaiting_payment', 'paid', 'delivered', 'cancelled');

-- Таблица: orders
-- Хранит информацию о заказах.

CREATE TABLE IF NOT EXISTS orders (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    address TEXT,
    stock_address_id BIGINT REFERENCES stock_address(id),
    delivery_price INTEGER NOT NULL,
    status order_status DEFAULT 'awaiting_payment',
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Таблица: product_orders
-- Связывает товары с заказами.

CREATE TABLE IF NOT EXISTS product_orders (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    order_id BIGINT NOT NULL REFERENCES orders(id),
    product_id BIGINT NOT NULL REFERENCES products(id),
    option_value_id BIGINT REFERENCES product_option_values(id),
    count INTEGER NOT NULL DEFAULT 1,
    delivery_date TIMESTAMP WITH TIME ZONE NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- Таблица: carts
-- Хранит товары в корзине пользователя.

CREATE TABLE IF NOT EXISTS carts (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id BIGINT NOT NULL REFERENCES users(id),
    product_id BIGINT NOT NULL REFERENCES products(id),
    option_value_id BIGINT REFERENCES product_option_values(id),
    count INTEGER NOT NULL,
    is_selected BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    delivery_date TIMESTAMP WITH TIME ZONE
);

-- Таблица: categories
-- Хранит категории товаров.

CREATE TABLE IF NOT EXISTS categories (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    picture TEXT NOT NULL,
    active BOOLEAN NOT NULL DEFAULT TRUE,
    link_to TEXT
);

-- Таблица: product_reviews
-- Хранит отзывы пользователей о товарах.

CREATE TABLE IF NOT EXISTS product_reviews (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id),
    rating SMALLINT NOT NULL,
    comment TEXT,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (product_id, user_id)
);

alter table public.users
    rename column password to hashed_password;