# Database Schema, DDL, DML
## Merancang Skema Database

1. Sistem dapat menyimpan data mengenai detail item product, yaitu : product, product_type, product_description, operator, payment_methods.
2. Sistem juga harus menyimpan data mengenai pelanggan yang akan membeli product tsb diantaranya : nama, alamat, tanggal lahir, status_user, gender, created_at, updated_at.
3. Sistem dapat mencatat transaksi pembelian dari pelanggan.
4. Sistem dapat mencatat detail transaksi pembelian dari pelanggan.

Berikut link draw.io : https://drive.google.com/file/d/1TPCH9P5Xu2KDDmtkbeQIl6m_yRAMTej_/view?usp=sharing

![skema_db](/11_DatabaseSchema_DDL_DML/screenshots/skema_db.png)

## Data Definition Language (DDL)

1. Create database alta_online_shop.
![create_db](/11_DatabaseSchema_DDL_DML/screenshots/create_db.JPG)

2. Dari schema Olshop yang telah kamu kerjakan di, Implementasikanlah menjadi table pada MySQL.
    a. Create table user.
![create_2a](/11_DatabaseSchema_DDL_DML/screenshots/create_2a.JPG)

    b. Create table product, product type, operators, product description, payment_method.
![create_2b](/11_DatabaseSchema_DDL_DML/screenshots/create_2b.JPG)

    c. Create table transaction, transaction detail.
![create_2c](/11_DatabaseSchema_DDL_DML/screenshots/create_2c.JPG)

3. Create tabel kurir dengan field id, name, created_at, updated_at.
![create_3](/11_DatabaseSchema_DDL_DML/screenshots/create_3.JPG)

4. Tambahkan ongkos_dasar column di tabel kurir.
![create_4](/11_DatabaseSchema_DDL_DML/screenshots/create_4.JPG)

5. Rename tabel kurir menjadi shipping.
![create_5](/11_DatabaseSchema_DDL_DML/screenshots/create_5.JPG)

6. Hapus / Drop tabel shipping karena ternyata tidak dibutuhkan.
![create_6](/11_DatabaseSchema_DDL_DML/screenshots/create_6.JPG)

7. Silahkan menambahkan entity baru dengan relation 1-to-1, 1-to-many, many-to-many. Seperti:
    a. 1-to-1: payment method description.
![create_7a](/11_DatabaseSchema_DDL_DML/screenshots/create_7a.JPG)

    b. 1-to-many: user dengan alamat.
![create_7b](/11_DatabaseSchema_DDL_DML/screenshots/create_7b.JPG)    
    
    c. many-to-many: user dengan payment method menjadi user_payment_method_detail.
![create_7c](/11_DatabaseSchema_DDL_DML/screenshots/create_7c.JPG)
