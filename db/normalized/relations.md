## Функциональные зависимости и нормальные формы

### **Таблица: users**

**Функциональные зависимости:**

- `{id} -> email, username, city, age, avatar_url, hashed_password, blocked, blocked_until, created_at, updated_at`
- `{email} -> id, username, city, age, avatar_url, hashed_password, blocked, blocked_until, created_at, updated_at`
- `{username} -> id, email, city, age, avatar_url, hashed_password, blocked, blocked_until, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Все атрибуты атомарны и содержат единичные значения.

- **2НФ:**
  Все неключевые атрибуты полностью зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей между неключевыми атрибутами.

- **НФБК:**
  Детерминанты `id`, `email`, `username` являются кандидатами в ключи (уникальные), поэтому НФБК соблюдается.

### **Таблица: sellers**

**Функциональные зависимости:**

- `{id} -> name, logo, rating, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Неключевые атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминант `id` является суперключом.

### **Таблица: categories**

**Функциональные зависимости:**

- `{id} -> name, picture, active, link_to, created_at, updated_at`
- `{name} -> id, picture, active, link_to, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Неключевые атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминанты `id` и `name` являются кандидатами в ключи (так как `name` уникален).

### **Таблица: products**

**Функциональные зависимости:**

- `{id} -> seller_id, category_id, count, price, original_price, discount, title, description, rating, image_url, active, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Неключевые атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминант `id` является суперключом.

### **Таблица: product_characteristics**

**Функциональные зависимости:**

- `{product_id, characteristic_id} -> value, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Составной первичный ключ `(product_id, characteristic_id)`. Атрибуты зависят от обоих ключей.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминант `(product_id, characteristic_id)` является суперключом.

### **Таблица: options**

**Функциональные зависимости:**

- `{id} -> name, created_at, updated_at`
- `{name} -> id, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Неключевые атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминанты `id` и `name` являются кандидатами в ключи.

### **Таблица: product_options**

**Функциональные зависимости:**

- `{id} -> product_id, option_id, created_at, updated_at`
- `{product_id, option_id} -> id, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминанты `id` и `(product_id, option_id)` являются кандидатами в ключи.

### **Таблица: product_option_values**

**Функциональные зависимости:**

- `{id} -> product_option_id, value, link, created_at, updated_at`
- `{product_option_id, value} -> id, link, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминанты `id` и `(product_option_id, value)` являются кандидатами в ключи.

### **Таблица: favorites**

**Функциональные зависимости:**

- `{id} -> user_id, product_id, created_at, updated_at`
- `{user_id, product_id} -> id, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминанты `id` и `(user_id, product_id)` являются кандидатами в ключи.

### **Таблица: product_images**

**Функциональные зависимости:**

- `{id} -> product_id, image_url, created_at, updated_at`
- `{product_id, image_url} -> id, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминанты `id` и `(product_id, image_url)` являются кандидатами в ключи.

### **Таблица: stock_address**

**Функциональные зависимости:**

- `{id} -> address, city, country, postal_code, created_at, updated_at`
- `{address, city, country, postal_code} -> id, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминанты `id` и `(address, city, country, postal_code)` являются кандидатами в ключи.

### **Таблица: orders**

**Функциональные зависимости:**

- `{id} -> user_id, address, stock_address_id, status, delivery_price, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминант `id` является суперключом.

### **Таблица: product_orders**

**Функциональные зависимости:**

- `{id} -> order_id, product_id, option_value_id, count, delivery_date, created_at, updated_at`
- `{order_id, product_id, option_value_id} -> count, delivery_date, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминанты `id` и `(order_id, product_id, option_value_id)` являются кандидатами в ключи.

### **Таблица: carts**

**Функциональные зависимости:**

- `{id} -> user_id, product_id, option_value_id, count, is_selected, delivery_date, created_at, updated_at`
- `{user_id, product_id, option_value_id} -> count, is_selected, delivery_date, created_at, updated_at`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминанты `id` и `(user_id, product_id, option_value_id)` являются кандидатами в ключи.

### **Таблица: product_reviews**

**Функциональные зависимости:**

- `{id} -> product_id, user_id, rating, comment, created_at, updated_at`
- `{product_id, user_id} -> rating, comment, created_at, updated_at, id`

**Объяснение нормализации:**

- **1НФ:**
  Атрибуты атомарны.

- **2НФ:**
  Атрибуты зависят от первичного ключа `id`.

- **3НФ:**
  Нет транзитивных зависимостей.

- **НФБК:**
  Детерминанты `id` и `(product_id, user_id)` являются кандидатами в ключи.