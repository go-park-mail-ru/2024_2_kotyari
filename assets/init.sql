-- Таблица пользователей, содержащая данные о зарегистрированных пользователях (почта, логин, пароль, аватар и др.)
CREATE TABLE IF NOT EXISTS "users" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"email" text UNIQUE NOT NULL,
	"username" text NOT NULL,
	"city" text NOT NULL DEFAULT 'Москва',
	"age" smallint CHECK (age >= 0 AND age <= 120),
	"avatar_url" text,
	"password" text NOT NULL,
	"blocked" boolean NOT NULL DEFAULT false,
	"blocked_until" timestamp with time zone DEFAULT NULL,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id")
);
CREATE TYPE seller_type AS ENUM ('individual', 'company');

-- Таблица продавцов, с указанием типа продавца (физическое лицо или компания), проверенности и рейтинга
CREATE TABLE IF NOT EXISTS "sellers" (
     "id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
     "name" text NOT NULL,  -- Название продавца (для компаний) или ФИО (для физических лиц)
     "logo" text not null,
     "rating" real CHECK (rating >= 0 AND rating <= 5) DEFAULT 0 NOT NULL,  -- Рейтинг продавца от 0 до 5
     "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
     "updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY ("id")
);

-- Таблица продуктов, хранящая информацию о товаре (цена, описание, скидка, изображения, характеристики и др.)
CREATE TABLE IF NOT EXISTS "products" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
    "seller_id" bigint NOT NULL,
	"count" integer NOT NULL DEFAULT '1' CHECK (count >= 0),
	"price" integer NOT NULL CHECK (price > 0),  -- Новая цена
	"original_price" integer CHECK (original_price > 0),  -- Оригинальная цена
	"discount" smallint CHECK (discount >= 0 AND discount < 100),  -- Скидка
    "title" text NOT NULL,
	"description" text NOT NULL ,
	"rating" real DEFAULT 0 NOT NULL CHECK (rating >= 0 AND rating <= 5),
	"image_url" text NOT NULL,
	"active" boolean NOT NULL DEFAULT true,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"characteristics" jsonb,  -- Характеристики продукта в формате JSON
--     например:
--     {"size": "123", "color": "red"}
	PRIMARY KEY ("id"),
    FOREIGN KEY ("seller_id") REFERENCES "sellers"("id") ON DELETE CASCADE
);

-- Таблица избранных продуктов для пользователей
CREATE TABLE IF NOT EXISTS "favorites" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"user_id" bigint NOT NULL,
	"product_id" bigint NOT NULL,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("user_id") REFERENCES "users"("id"),
	FOREIGN KEY ("product_id") REFERENCES "products"("id"),
    UNIQUE ("user_id", "product_id")
);

-- Таблица для хранения изображений продукта
CREATE TABLE IF NOT EXISTS "product_images" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"product_id" bigint NOT NULL,
	"image_url" text NOT NULL,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);

-- Таблица складских адресов, которая содержит информацию о складе (адрес, город, страна, почтовый индекс)
CREATE TABLE IF NOT EXISTS "stock_address" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"address" text NOT NULL,
	"city" text NOT NULL,
	"country" text NOT NULL,
	"postal_code" text NOT NULL,
	PRIMARY KEY ("id")
);


CREATE TYPE order_status AS ENUM ('awaiting_payment', 'paid', 'delivered', 'cancelled');
-- Таблица заказов, связанная с пользователями и складскими адресами
CREATE TABLE IF NOT EXISTS "orders" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"user_id" bigint NOT NULL,
	"address" text NOT NULL DEFAULT NULL,  -- Адрес доставки (если не используется стоковый адрес)
	"stock_address_id" bigint,  -- Ссылка на таблицу стоковых адресов
	"total_price" integer NOT NULL CHECK ("total_price" > 0),
	"status" order_status DEFAULT 'awaiting_payment',
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("user_id") REFERENCES "users"("id"),
	FOREIGN KEY ("stock_address_id") REFERENCES "stock_address"("id")  -- Внешний ключ для стоковых адресов
);

-- Таблица для опций продукта, таких как цвет или размер, с возможностью указать доступные значения
CREATE TABLE IF NOT EXISTS "product_options" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"product_id" bigint NOT NULL,  -- Ссылка на товар
	"values" jsonb NOT NULL,  -- Возможные значения опции в формате JSON
--     [
--          {
--              "type": "size",
--              "title": "Размер",
--              "options": [
--                  {
--                      "link": "",
--                      "value": "XL"
--                  },
--                  {
--                      "link": "",
--                      "value": "XS"
--                  },
--              ]
--         }
--    ]
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);

-- Таблица связывает продукты с заказами, хранит количество и дату доставки продукта в заказе
CREATE TABLE IF NOT EXISTS "product_orders" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"order_id" bigint NOT NULL,
	"product_id" bigint NOT NULL,
    "option_id" bigint,  -- Ссылка на опцию
	"count" integer NOT NULL DEFAULT '1' CHECK (count > 0),
	"delivery_date" timestamp with time zone NOT NULL,  -- Дата доставки продукта
	PRIMARY KEY ("id"),
	FOREIGN KEY ("order_id") REFERENCES "orders"("id"),
	FOREIGN KEY ("product_id") REFERENCES "products"("id"),
    FOREIGN KEY ("option_id") REFERENCES "product_options"("id")
);

CREATE UNIQUE INDEX unique_order_index ON product_orders(order_id, product_id, COALESCE(option_id, -1));

-- Таблица корзины пользователя, где хранится информация о продуктах в корзине
CREATE TABLE IF NOT EXISTS "carts" (
   "id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
   "user_id" bigint NOT NULL,
   "product_id" bigint NOT NULL,
   "count" integer NOT NULL,
   "is_selected" boolean DEFAULT false,
   "created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
   "updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
   "delivery_date" timestamp with time zone,  -- Дата доставки для товара в корзине
   PRIMARY KEY ("id"),
   FOREIGN KEY ("user_id") REFERENCES "users"("id"),
   FOREIGN KEY ("product_id") REFERENCES "products"("id")
);

CREATE UNIQUE INDEX unique_cart_index ON carts(user_id, product_id);

-- Таблица тегов для продуктов
CREATE TABLE IF NOT EXISTS "categories" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"name" text NOT NULL,
    "picture" text NOT NULL,
	"active" boolean NOT NULL DEFAULT true,
	PRIMARY KEY ("id")
);

-- Таблица, связывающая продукты и теги
CREATE TABLE IF NOT EXISTS "product_categories" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"product_id" bigint NOT NULL,
	"category_id" bigint NOT NULL,
	"active" boolean NOT NULL DEFAULT true,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("product_id") REFERENCES "products"("id"),
	FOREIGN KEY ("category_id") REFERENCES "categories"("id")
);
-- Индекс для ускорения поиска товаров по продавцу
CREATE INDEX idx_product_seller ON products(seller_id);

-- Индекс для ускорения поиска товаров по опциям
CREATE INDEX idx_product_option_product ON product_options(product_id);

-- Индексы для обеспечения быстрого поиска
CREATE INDEX idx_product_created_at ON products(created_at);
CREATE INDEX idx_product_price ON products(price);
CREATE INDEX idx_product_characteristics ON products USING gin (characteristics);

-- Не уверен насколько это нужно
ALTER TABLE orders
ADD CONSTRAINT status_check CHECK (status IN ('awaiting_payment', 'paid', 'delivered', 'cancelled'));

ALTER TABLE categories
ADD COLUMN link_to text;