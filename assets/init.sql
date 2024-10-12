-- Таблица пользователей, содержащая данные о зарегистрированных пользователях (почта, логин, пароль, аватар и др.)
CREATE TABLE IF NOT EXISTS "user" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"mail" text NOT NULL,
	"username" text NOT NULL,
	"age" smallint CHECK (age >= 0 AND age <= 120),
	"avatar_url" text,
	"password" text NOT NULL,
	"active" boolean NOT NULL DEFAULT true,
	"blocked" boolean NOT NULL DEFAULT false,
	"block_until" timestamp with time zone DEFAULT NULL,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id")
);

-- Таблица продуктов, хранящая информацию о товаре (цена, описание, скидка, изображения, характеристики и др.)
CREATE TABLE IF NOT EXISTS "product" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"count" integer NOT NULL DEFAULT '1' CHECK (count >= 0),
	"price" integer NOT NULL CHECK (price >= 0),  -- Новая цена
	"old_price" integer CHECK (old_price >= 0),  -- Старая цена
	"currency" text NOT NULL DEFAULT 'rub',
	"discount" smallint CHECK (discount >= 0 AND discount <= 100),  -- Скидка
	"description" text,
	"short_description" text NOT NULL,
	"image_url" text NOT NULL,
	"active" boolean NOT NULL DEFAULT true,
	"delivery_date" timestamp with time zone,  -- Дата доставки продукта
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"characteristics" jsonb,  -- Характеристики продукта в формате JSON
	PRIMARY KEY ("id")
);

-- Таблица избранных продуктов для пользователей
CREATE TABLE IF NOT EXISTS "favorite" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"user_id" bigint NOT NULL,
	"product_id" bigint NOT NULL,
	"is_favorite" boolean NOT NULL DEFAULT true,  -- Флаг, указывающий, что продукт в избранном
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("user_id") REFERENCES "user"("id"),
	FOREIGN KEY ("product_id") REFERENCES "product"("id")
);

-- Таблица для хранения изображений продукта
CREATE TABLE IF NOT EXISTS "product_image" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"product_id" bigint NOT NULL,
	"image_url" text NOT NULL,
	"active" boolean NOT NULL DEFAULT true,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("product_id") REFERENCES "product"("id") ON DELETE CASCADE
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
CREATE TABLE IF NOT EXISTS "order" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"user_id" bigint NOT NULL,
	"address" text NOT NULL DEFAULT 'NONE',  -- Адрес доставки (если не используется стоковый адрес)
	"stock_address_id" bigint,  -- Ссылка на таблицу стоковых адресов
	"price" integer NOT NULL,
	"status" text NOT NULL DEFAULT 'paid',
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("user_id") REFERENCES "user"("id"),
	FOREIGN KEY ("stock_address_id") REFERENCES "stock_address"("id")  -- Внешний ключ для стоковых адресов
);

-- Таблица для опций продукта, таких как цвет или размер, с возможностью указать доступные значения
CREATE TABLE IF NOT EXISTS "product_option" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"product_id" bigint NOT NULL,  -- Ссылка на товар
	"option_name" text NOT NULL,  -- Название опции (например, цвет, размер)
	"option_values" jsonb NOT NULL,  -- Возможные значения опции в формате JSON (например, ["красный", "синий", "зеленый"])
	"active" boolean NOT NULL DEFAULT true,
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("product_id") REFERENCES "product"("id") ON DELETE CASCADE
);

-- Таблица связывает продукты с заказами, хранит количество и дату доставки продукта в заказе
CREATE TABLE IF NOT EXISTS "product_and_order" (
	"order_id" bigint NOT NULL,
	"product_id" bigint NOT NULL,
    "option_id" bigint NOT NULL,  -- Ссылка на опцию
	"option_value" text NOT NULL,  -- Конкретное выбранное значение опции (например, красный или XL)
	"count" integer NOT NULL DEFAULT '1' CHECK (count > 0),
	"delivery_date" timestamp with time zone NOT NULL,  -- Дата доставки продукта
	PRIMARY KEY ("order_id", "product_id", "option_id"),
	FOREIGN KEY ("order_id") REFERENCES "order"("id"),
	FOREIGN KEY ("product_id") REFERENCES "product"("id"),
    FOREIGN KEY ("option_id") REFERENCES "product_option"("id")
);

-- Таблица корзины пользователя, где хранится информация о продуктах в корзине
CREATE TABLE IF NOT EXISTS "cart" (
	"user_id" bigint NOT NULL,
	"product_id" bigint NOT NULL,
    "option_id" bigint NOT NULL,  -- Ссылка на опцию
	"count" integer NOT NULL CHECK (count > 0),
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"delivery_date" timestamp with time zone,  -- Дата доставки для товара в корзине
	PRIMARY KEY ("user_id", "product_id", "option_id"),  -- Комбинированный первичный ключ, чтобы каждая пара была уникальной
	FOREIGN KEY ("user_id") REFERENCES "user"("id"),
	FOREIGN KEY ("product_id") REFERENCES "product"("id"),
    FOREIGN KEY ("option_id") REFERENCES "product_option"("id")
);

-- Таблица тегов для продуктов
CREATE TABLE IF NOT EXISTS "tag" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"tag" text NOT NULL,
	"resultant_limitation" smallint NOT NULL DEFAULT '0' CHECK (resultant_limitation >= 0 AND resultant_limitation <= 120),
	"active" boolean NOT NULL DEFAULT true,
	"color" text DEFAULT 'rgb(160, 60, 236)',
	PRIMARY KEY ("id")
);

-- Таблица, связывающая продукты и теги
CREATE TABLE IF NOT EXISTS "product_tag" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"product_id" bigint NOT NULL,
	"tag_id" bigint NOT NULL,
	"active" boolean NOT NULL DEFAULT true,
	PRIMARY KEY ("id"),
	FOREIGN KEY ("product_id") REFERENCES "product"("id"),
	FOREIGN KEY ("tag_id") REFERENCES "tag"("id")
);

-- Таблица продавцов, с указанием типа продавца (физическое лицо или компания), проверенности и рейтинга
CREATE TABLE IF NOT EXISTS "seller" (
	"id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY,
	"seller_name" text NOT NULL,  -- Название продавца (для компаний) или ФИО (для физических лиц)
	"seller_type" text NOT NULL CHECK (seller_type IN ('individual', 'company')),  -- Тип продавца (физ лицо или компания)
	"verified" boolean NOT NULL DEFAULT false,  -- Флаг, указывающий, что продавец проверен
	"rating" smallint CHECK (rating >= 0 AND rating <= 5),  -- Рейтинг продавца от 0 до 5
	"created_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	"updated_at" timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY ("id")
);

-- Добавление продавца к продукту (один продукт - один продавец)
ALTER TABLE "product" ADD COLUMN "seller_id" bigint NOT NULL;
ALTER TABLE "product" ADD FOREIGN KEY ("seller_id") REFERENCES "seller"("id") ON DELETE CASCADE;

-- Индекс для ускорения поиска товаров по продавцу
CREATE INDEX idx_product_seller ON product(seller_id);

-- Индекс для ускорения поиска товаров по опциям
CREATE INDEX idx_product_option_product ON product_option(product_id);

-- Индексы для обеспечения быстрого поиска
CREATE INDEX idx_product_created_at ON product(created_at);
CREATE INDEX idx_product_price ON product(price);
CREATE INDEX idx_product_characteristics ON product USING gin (characteristics);

ALTER TABLE "user" 
ADD CONSTRAINT chk_mail_format 
CHECK (mail ~* '^[A-Za-zА-Яа-яЁё0-9._%+-]+@[A-Za-zА-Яа-яЁё0-9.-]+\.[A-Za-zА-Яа-яЁё]{2,}$');


-- Функция для автоматического обновления поля updated_at при изменении записи
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
	NEW.updated_at = CURRENT_TIMESTAMP;
	RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updated_at_user
BEFORE UPDATE ON "user"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_product
BEFORE UPDATE ON "product"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_cart
BEFORE UPDATE ON "cart"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_updated_at_seller
BEFORE UPDATE ON "seller"
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
