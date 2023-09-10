-- 1. Create database alta_online_shop.
CREATE DATABASE alta_online_shop;

-- 2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
-- 2. a. Create tabel user
CREATE TABLE users (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  alamat VARCHAR(255),
  dob DATE,
  status SMALLINT,
  gender CHAR(1),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 2. Create table product, product type, operators, product description, payment_method.
CREATE TABLE operators (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE product_types (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE products (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  product_type_id INT(11),
  operator_id INT(11),
  code VARCHAR(50),
  name VARCHAR(100),
  status SMALLINT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (product_type_id) REFERENCES product_types(id),
  FOREIGN KEY (operator_id) REFERENCES operators(id)
);

CREATE TABLE product_descriptions (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  description TEXT, 
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE payment_methods (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  name VARCHAR(255),
  status SMALLINT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 2. c. Create table transaction, transaction detail.
CREATE TABLE transactions (
  id INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  user_id INT(11),
  payment_method_id INT(11),
  status VARCHAR(10),
  total_qty INT(11),
  total_price NUMERIC(25,2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

CREATE TABLE transaction_details (
  transaction_id INT(11),
  product_id INT(11),
  status VARCHAR(10),
  qty INT(11),
  price NUMERIC(25,2),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  FOREIGN KEY (transaction_id) REFERENCES transactions(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);

-- 3. Create tabel kurir dengan field id, name, created_at, updated_at.
CREATE TABLE kurir (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255), 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- 4. Tambahkan ongkos_dasar column di tabel kurir.
ALTER TABLE kurir ADD ongkos_dasar INT;

-- 5. Rename tabel kurir menjadi shipping.
ALTER TABLE kurir RENAME TO shipping;

-- 6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
DROP TABLE shipping;

-- 7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
-- 7. a. 1-to-1: payment method description.
CREATE TABLE payment_method_description (
    id INT PRIMARY KEY AUTO_INCREMENT,
    payment_method_id INT,
    description TEXT,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);

-- 7. b. 1-to-many: user dengan alamat.
CREATE TABLE address (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    street_address TEXT,
    city VARCHAR(255),
    province VARCHAR(255),
    country VARCHAR(255),
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- 7. c. many-to-many: user dengan payment method menjadi user_payment_method_detail.
CREATE TABLE user_payment_method_detail (
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_id INT,
    payment_method_id INT,
    created_at DATETIME,
    updated_at DATETIME,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);
