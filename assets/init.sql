-- Таблица пользователей, содержащая данные о зарегистрированных пользователях (почта, логин, пароль, аватар и др.)
CREATE TABLE IF NOT EXISTS "users" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"email" text NOT NULL,
	"username" text NOT NULL,
	"age" smallint CHECK (age >= 0 AND age <= 120),
	"avatar_url" text,
	"password" text NOT NULL,
	"blocked" boolean NOT NULL DEFAULT false,
	"block_until" timestamp with time zone DEFAULT NULL,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id")
);

-- Таблица продуктов, хранящая информацию о товаре (цена, описание, скидка, изображения, характеристики и др.)
CREATE TABLE IF NOT EXISTS "products" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"count" integer NOT NULL DEFAULT '1' CHECK (count >= 0),
	"price" integer NOT NULL CHECK (price >= 0),  -- Новая цена
	"old_price" integer CHECK (old_price >= 0),  -- Старая цена
	"currency" text NOT NULL DEFAULT 'rub',
	"discount" smallint CHECK (discount >= 0 AND discount <= 100),  -- Скидка
	"description" text,
	"short_description" text NOT NULL,
	"rating" real DEFAULT 0 NOT NULL CHECK (rating >= 0 AND rating <= 5),
	"image_url" text NOT NULL,
	"active" boolean NOT NULL DEFAULT true,
	"days_to_delivery" smallint NOT NULL,  -- Дата доставки продукта
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"characteristics" jsonb,  -- Характеристики продукта в формате JSON
	PRIMARY KEY ("id")
);

-- Таблица избранных продуктов для пользователей
CREATE TABLE IF NOT EXISTS "favorites" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"user_id" bigint NOT NULL,
	"product_id" bigint NOT NULL,
	"is_favorite" boolean NOT NULL DEFAULT true,  -- Флаг, указывающий, что продукт в избранном
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("user_id") REFERENCES "users"("id"),
	FOREIGN KEY ("product_id") REFERENCES "products"("id")
);

-- Таблица для хранения изображений продукта
CREATE TABLE IF NOT EXISTS "product_images" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"product_id" bigint NOT NULL,
	"image_url" text NOT NULL,
	"active" boolean NOT NULL DEFAULT true,
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

-- Таблица заказов, связанная с пользователями и складскими адресами
CREATE TABLE IF NOT EXISTS "orders" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"user_id" bigint NOT NULL,
	"address" text NOT NULL DEFAULT 'NONE',  -- Адрес доставки (если не используется стоковый адрес)
	"stock_address_id" bigint,  -- Ссылка на таблицу стоковых адресов
	"total_price" integer NOT NULL CHECK ("total_price" >= 0),
	"status" text NOT NULL DEFAULT 'paid',
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
	"option_name" text NOT NULL,  -- Название опции (например, цвет, размер)
	"option_values" jsonb NOT NULL,  -- Возможные значения опции в формате JSON (например, ["красный", "синий", "зеленый"])
	"count" integer NOT NULL DEFAULT '1' CHECK (count >= 0),
	"price" integer NOT NULL CHECK (price >= 0),  -- Новая цена
	"old_price" integer CHECK (old_price >= 0),  -- Старая цена
	"currency" text NOT NULL DEFAULT 'rub',
	"discount" smallint CHECK (discount >= 0 AND discount <= 100),  -- Скидка
	"image_url" text,
	"active" boolean NOT NULL DEFAULT true,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("product_id") REFERENCES "products"("id") ON DELETE CASCADE
);

-- Таблица для хранения изображений продукта
CREATE TABLE IF NOT EXISTS "product_option_images" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"option_id" bigint NOT NULL,
	"image_url" text NOT NULL,
	"active" boolean NOT NULL DEFAULT true,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("option_id") REFERENCES "product_options"("id") ON DELETE CASCADE
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
    "option_id" bigint,  -- Ссылка на опцию
	"count" integer NOT NULL CHECK (count > 0),
	"is_selected" boolean DEFAULT false,
	"active" boolean NOT NULL DEFAULT true,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"delivery_date" timestamp with time zone,  -- Дата доставки для товара в корзине
	PRIMARY KEY ("id"),
	FOREIGN KEY ("user_id") REFERENCES "users"("id"),
	FOREIGN KEY ("product_id") REFERENCES "products"("id"),
    FOREIGN KEY ("option_id") REFERENCES "product_options"("id")
);

CREATE UNIQUE INDEX unique_cart_index ON carts(user_id, product_id, COALESCE(option_id, -1));

-- Таблица тегов для продуктов
CREATE TABLE IF NOT EXISTS "categories" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"name" text NOT NULL,
	"resultant_limitation" smallint NOT NULL DEFAULT '0' CHECK (resultant_limitation >= 0 AND resultant_limitation <= 120),
	"active" boolean NOT NULL DEFAULT true,
	"color" text DEFAULT 'rgb(160, 60, 236)',
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

-- Таблица продавцов, с указанием типа продавца (физическое лицо или компания), проверенности и рейтинга
CREATE TABLE IF NOT EXISTS "sellers" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"seller_name" text NOT NULL,  -- Название продавца (для компаний) или ФИО (для физических лиц)
	"seller_type" text NOT NULL CHECK (seller_type IN ('individual', 'company')),  -- Тип продавца (физ лицо или компания)
	"verified" boolean NOT NULL DEFAULT false,  -- Флаг, указывающий, что продавец проверен
	"rating" real CHECK (rating >= 0 AND rating <= 5) DEFAULT 0 NOT NULL,  -- Рейтинг продавца от 0 до 5
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id")
);

-- Добавление продавца к продукту (один продукт - один продавец)
ALTER TABLE "products" ADD COLUMN "seller_id" bigint NOT NULL;
ALTER TABLE "products" ADD FOREIGN KEY ("seller_id") REFERENCES "sellers"("id") ON DELETE CASCADE;

-- Индекс для ускорения поиска товаров по продавцу
CREATE INDEX idx_product_seller ON products(seller_id);

-- Индекс для ускорения поиска товаров по опциям
CREATE INDEX idx_product_option_product ON product_options(product_id);

-- Индексы для обеспечения быстрого поиска
CREATE INDEX idx_product_created_at ON products(created_at);
CREATE INDEX idx_product_price ON products(price);
CREATE INDEX idx_product_characteristics ON products USING gin (characteristics);


-- Функция для автоматического обновления поля updated_at при изменении записи
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = CURRENT_TIMESTAMP;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_user
BEFORE UPDATE ON "users"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_product
BEFORE UPDATE ON "products"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_cart
BEFORE UPDATE ON "carts"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_seller
BEFORE UPDATE ON "sellers"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_seller
BEFORE UPDATE ON "orders"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
