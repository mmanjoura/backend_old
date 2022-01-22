--
-- File generated with SQLiteStudio v3.3.3 on Thu Jan 20 18:23:12 2022
--
-- Text encoding used: System
--
PRAGMA foreign_keys = off;
BEGIN TRANSACTION;

-- Table: address
CREATE TABLE address (
    id             INTEGER        PRIMARY KEY AUTOINCREMENT,
    Shop_id        [INT]          REFERENCES shop (id),
    zip            [VARCHAR] (50),
    city           [VARCHAR] (50),
    state          [VARCHAR] (50),
    country        [VARCHAR] (50),
    street_address [VARCHAR] (50),
    type           [VARCHAR] (50),
    title          [NCHAR] (50),
    [default]      [INT]
);

INSERT INTO address (id, Shop_id, zip, city, state, country, street_address, type, title, "default") VALUES (1, 1, '08753', 'East Dover', 'New Jersey', 'USA', '588  Finwood Road', NULL, NULL, NULL);

-- Table: attribute
CREATE TABLE attribute (
    id         INTEGER        PRIMARY KEY AUTOINCREMENT,
    Shop_id    [INT]          REFERENCES shop (id),
    slug       [VARCHAR] (50),
    name       [VARCHAR] (50),
    created_at TEXT,
    updated_at TEXT
);

INSERT INTO attribute (id, Shop_id, slug, name, created_at, updated_at) VALUES (1, 1, 'color', 'Color', '2021-05-09T16:10:31.000000Z', '2021-05-09T16:10:31.000000Z');

-- Table: auth
CREATE TABLE auth (
    Id            INTEGER PRIMARY KEY AUTOINCREMENT,
    user_Id       INTEGER NOT NULL
                          REFERENCES user (Id) ON DELETE CASCADE,
    source        TEXT    NOT NULL,
    source_Id     TEXT    NOT NULL,
    access_token  TEXT    NOT NULL,
    refresh_token TEXT    NOT NULL,
    expiry        TEXT,
    created_at    TEXT    NOT NULL,
    updated_at    TEXT    NOT NULL,
    UNIQUE (
        user_Id,
        source
    ),-- one source per user
    UNIQUE (
        source,
        source_Id
    )-- one auth per source user
);

INSERT INTO auth (Id, user_Id, source, source_Id, access_token, refresh_token, expiry, created_at, updated_at) VALUES (1, 1, 'github', '5604914', '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', '', NULL, '2022-01-07T12:37:57Z', '2022-01-10T13:38:33Z');

-- Table: avatar
CREATE TABLE avatar (
    id        INTEGER         PRIMARY KEY AUTOINCREMENT,
    user_id   [INT]           REFERENCES user (id),
    original  [VARCHAR] (124),
    thumbnail [VARCHAR] (146) 
);

INSERT INTO avatar (id, user_id, original, thumbnail) VALUES (1, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/881/aatik-tasneem-7omHUGhhmZ0-unsplash%402x.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/881/conversions/aatik-tasneem-7omHUGhhmZ0-unsplash%402x-thumbnail.jpg');

-- Table: banner
CREATE TABLE banner (
    id               INTEGER         PRIMARY KEY AUTOINCREMENT,
    category_type_Id [INT],
    title            [VARCHAR] (100),
    description      [VARCHAR] (500),
    created_at       TEXT,
    updated_at       TEXT
);

INSERT INTO banner (id, category_type_Id, title, description, created_at, updated_at) VALUES (1, 1, 'Groceries Delivered in 90 Minute', 'Get your healthy foods & snacks delivered at your doorsteps all day everyday', '2021-07-17T13:21:55.000000Z', '2021-07-17T13:21:55.000000Z');
INSERT INTO banner (id, category_type_Id, title, description, created_at, updated_at) VALUES (2, 2, 'Get Your Bakery Items Delivered', 'Get your favorite bakery items baked and delivered to your doorsteps at any time', '2021-07-17T13:22:34.000000Z', '2021-07-17T13:22:34.000000Z');
INSERT INTO banner (id, category_type_Id, title, description, created_at, updated_at) VALUES (3, 3, 'Branded & imported makeups', 'Easiest and cheapest way to get your branded & imported makeups', '2021-07-17T13:23:05.000000Z', '2021-07-17T13:23:05.000000Z');
INSERT INTO banner (id, category_type_Id, title, description, created_at, updated_at) VALUES (4, 4, 'Exclusive Branded bags', 'Get your exclusive & branded bags delivered to you in no time', '2021-07-17T13:23:32.000000Z', '2021-07-17T13:23:32.000000Z');
INSERT INTO banner (id, category_type_Id, title, description, created_at, updated_at) VALUES (5, 5, 'Shop your designer dresses', 'Ready to wear dresses tailored for you online. Hurry up while stock lasts.', '2021-07-17T13:24:28.000000Z', '2021-07-17T13:24:28.000000Z');
INSERT INTO banner (id, category_type_Id, title, description, created_at, updated_at) VALUES (6, 6, 'Exclusive furniture on cheap price', 'Make your house a home with our wide collection of beautiful furniture', '2021-08-18T18:45:54.000000Z', '2021-08-18T18:45:54.000000Z');
INSERT INTO banner (id, category_type_Id, title, description, created_at, updated_at) VALUES (7, 6, 'Furniter 2', NULL, '2021-08-18T18:45:54.000000Z', '2021-08-18T18:45:54.000000Z');
INSERT INTO banner (id, category_type_Id, title, description, created_at, updated_at) VALUES (8, 7, 'You Deserve To Eat Fresh', 'We source the best healthy foods for you.', '2021-10-03T10:28:29.000000Z', '2021-10-03T10:28:29.000000Z');

-- Table: banner_image
CREATE TABLE banner_image (
    id               INTEGER         PRIMARY KEY AUTOINCREMENT,
    category_type_id [INT],
    original         [VARCHAR] (500),
    thumbnail        [VARCHAR] (500) 
);

INSERT INTO banner_image (id, category_type_id, original, thumbnail) VALUES (1, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/grocery.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/conversions/grocery-thumbnail.jpg');
INSERT INTO banner_image (id, category_type_id, original, thumbnail) VALUES (2, 2, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/bakery.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/conversions/bakery-thumbnail.jpg');
INSERT INTO banner_image (id, category_type_id, original, thumbnail) VALUES (3, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/makeup.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/conversions/makeup-thumbnail.jpg');
INSERT INTO banner_image (id, category_type_id, original, thumbnail) VALUES (4, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/907/bags.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/907/conversions/bags-thumbnail.jpg');
INSERT INTO banner_image (id, category_type_id, original, thumbnail) VALUES (5, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/908/cloths.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/908/conversions/cloths-thumbnail.jpg');
INSERT INTO banner_image (id, category_type_id, original, thumbnail) VALUES (6, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/922/furniture-banner-1.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/922/conversions/furniture-banner-1-thumbnail.jpg');
INSERT INTO banner_image (id, category_type_id, original, thumbnail) VALUES (7, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/923/furniture-banner-2.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/923/conversions/furniture-banner-2-thumbnail.jpg');
INSERT INTO banner_image (id, category_type_id, original, thumbnail) VALUES (8, 7, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1344/shutterstock_389040853-%281%29.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1344/conversions/shutterstock_389040853-%281%29-thumbnail.jpg');

-- Table: billing_address
CREATE TABLE billing_address (
    id                INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_order_Id [INT],
    zip               [VARCHAR] (50),
    city              [VARCHAR] (50),
    state             [VARCHAR] (50),
    country           [VARCHAR] (50),
    street_address    [VARCHAR] (50) 
);

INSERT INTO billing_address (id, customer_order_Id, zip, city, state, country, street_address) VALUES (1, 1, '40391', 'Winchester', 'KY', 'United States', '2148  Straford Park');

-- Table: category
CREATE TABLE category (
    id               INTEGER        PRIMARY KEY AUTOINCREMENT,
    category_type_id [INT],
    parent_Id        INTEGER,
    shop_id          INT,
    name             [VARCHAR] (50),
    slug             [VARCHAR] (50),
    icon             [VARCHAR] (50),
    details          VARCHAR (50),
    deleted_at       TEXT,
    created_at       TEXT,
    updated_at       TEXT
);

INSERT INTO category (id, category_type_id, parent_Id, shop_id, name, slug, icon, details, deleted_at, created_at, updated_at) VALUES (1, 1, 0, 1, 'Fruits & Vegetables', 'fruits-vegetables', 'FruitsVegetable', NULL, '2021-03-08T07:21:31.000000Z', '2021-03-08T07:21:31.000000Z', '2021-03-08T07:21:31.000000Z');
INSERT INTO category (id, category_type_id, parent_Id, shop_id, name, slug, icon, details, deleted_at, created_at, updated_at) VALUES (2, 1, 1, 1, 'Fruits', 'fruits-vegetables', NULL, NULL, '2021-03-08T07:21:31.000000Z', '2021-03-08T07:21:31.000000Z', '2021-03-08T07:21:31.000000Z');
INSERT INTO category (id, category_type_id, parent_Id, shop_id, name, slug, icon, details, deleted_at, created_at, updated_at) VALUES (3, 1, 1, 1, 'Vegetables', 'vegetables', NULL, NULL, '2021-03-08T07:21:31.000000Z', '2021-03-08T07:21:31.000000Z', '2021-03-08T07:21:31.000000Z');

-- Table: category_children
CREATE TABLE category_children (
    id          INTEGER        PRIMARY KEY AUTOINCREMENT,
    category_Id [INT],
    parent_id   [INT],
    name        [VARCHAR] (50),
    slug        [VARCHAR] (50),
    icon        [VARCHAR] (50),
    created_at  TEXT,
    updated_at  TEXT,
    deleted_at  TEXT
);

INSERT INTO category_children (id, category_Id, parent_id, name, slug, icon, created_at, updated_at, deleted_at) VALUES (1, 1, 1, 'Fruits', 'fruits', 'icon', '2021-03-08T07:22:04.000000Z', '2021-03-08T07:22:04.000000Z', '2021-03-08T07:22:04.000000Z');

-- Table: category_image
CREATE TABLE category_image (
    id          INTEGER        PRIMARY KEY AUTOINCREMENT,
    category_Id [INT],
    original    [VARCHAR] (50),
    thumbnail   [VARCHAR] (50) 
);


-- Table: category_type
CREATE TABLE category_type (
    id          INTEGER        PRIMARY KEY AUTOINCREMENT,
    category_Id [INT],
    shop_id     INT,
    name        [VARCHAR] (50),
    slug        [VARCHAR] (50),
    icon        [VARCHAR] (50),
    created_at  TEXT,
    updated_at  TEXT
);

INSERT INTO category_type (id, category_Id, shop_id, name, slug, icon, created_at, updated_at) VALUES (1, 1, 6, 'Grocery', 'grocery', 'FruitsVegetable', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO category_type (id, category_Id, shop_id, name, slug, icon, created_at, updated_at) VALUES (2, 1, 5, 'Bakery', 'bakery', 'Bakery', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO category_type (id, category_Id, shop_id, name, slug, icon, created_at, updated_at) VALUES (3, 1, 4, 'Makeup', 'makeup', 'FacialCare', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO category_type (id, category_Id, shop_id, name, slug, icon, created_at, updated_at) VALUES (4, 1, 3, 'Bags', 'bags', 'Handbag', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO category_type (id, category_Id, shop_id, name, slug, icon, created_at, updated_at) VALUES (5, 1, 2, 'Clothing', 'clothing', 'DressIcon', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO category_type (id, category_Id, shop_id, name, slug, icon, created_at, updated_at) VALUES (6, 1, 1, 'Furniture', 'furniture', 'FurnitureIcon', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');

-- Table: contactDetail
CREATE TABLE contactDetail (
    id         INTEGER        PRIMARY KEY AUTOINCREMENT,
    setting_id [INT],
    shop_id    INT,
    contact    [VARCHAR] (50),
    website    [VARCHAR] (50) 
);

INSERT INTO contactDetail (id, setting_id, shop_id, contact, website) VALUES (1, 1, 1, '+129290122122', 'https://redq.io');

-- Table: coupon
CREATE TABLE coupon (
    id          INTEGER        PRIMARY KEY AUTOINCREMENT,
    code        [VARCHAR] (50),
    description VARCHAR (100),
    type        [VARCHAR] (50),
    amount      [INT],
    active_from TEXT,
    expire_at   TEXT,
    created_at  TEXT,
    updated_at  TEXT,
    deleted_at  TEXT,
    is_valid    [INT]
);

INSERT INTO coupon (id, code, description, type, amount, active_from, expire_at, created_at, updated_at, deleted_at, is_valid) VALUES (1, 'EID2', 'long sescription', 'fixed', 2, '2021-03-10T21:48:15.000Z', '2024-07-27T21:48:15.000Z', '2021-03-10T21:49:46.000000Z', '2021-08-19T04:00:12.000000Z', '2021-08-19T04:00:12.000000Z', 1);
INSERT INTO coupon (id, code, description, type, amount, active_from, expire_at, created_at, updated_at, deleted_at, is_valid) VALUES (2, 'EID2', 'long sescription', 'fixed', 2, '2021-03-10T21:48:15.000Z', '2024-07-27T21:48:15.000Z', '2021-03-10T21:49:46.000000Z', '2021-08-19T04:00:12.000000Z', '2021-08-19T04:00:12.000000Z', 1);

-- Table: cover_image
CREATE TABLE cover_image (
    id        INTEGER         PRIMARY KEY AUTOINCREMENT,
    shop_id   [INT]           REFERENCES shop (id),
    original  [VARCHAR] (500),
    thumbnail [VARCHAR] (500) 
);

INSERT INTO cover_image (id, shop_id, original, thumbnail) VALUES (1, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/Untitled-6.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/conversions/Untitled-6-thumbnail.jpg');
INSERT INTO cover_image (id, shop_id, original, thumbnail) VALUES (2, 2, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/Untitled-6.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/conversions/Untitled-6-thumbnail.jpg');
INSERT INTO cover_image (id, shop_id, original, thumbnail) VALUES (3, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/Untitled-6.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/conversions/Untitled-6-thumbnail.jpg');
INSERT INTO cover_image (id, shop_id, original, thumbnail) VALUES (4, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/Untitled-6.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/conversions/Untitled-6-thumbnail.jpg');
INSERT INTO cover_image (id, shop_id, original, thumbnail) VALUES (5, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/Untitled-6.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/conversions/Untitled-6-thumbnail.jpg');
INSERT INTO cover_image (id, shop_id, original, thumbnail) VALUES (6, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/Untitled-6.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/conversions/Untitled-6-thumbnail.jpg');

-- Table: customer
CREATE TABLE customer (
    id                INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_order_Id [INT],
    shop_id           [INT],
    name              [VARCHAR] (50),
    email             [VARCHAR] (50),
    email_verified_at TEXT,
    is_active         [INT],
    created_at        TEXT,
    updated_at        TEXT
);

INSERT INTO customer (id, customer_order_Id, shop_id, name, email, email_verified_at, is_active, created_at, updated_at) VALUES (1, 1, 1, 'Customer', 'customer@demo.com', '2021-08-18T10:30:29.000000Z', 1, '2021-08-18T10:30:29.000000Z', '2021-08-18T13:17:53.000000Z');

-- Table: customer_order
CREATE TABLE customer_order (
    id                 INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_id        [INT],
    coupon_id          [INT],
    parent_id          [INT],
    shop_id            [INT],
    payment_id         [INT],
    tracking_number    [VARCHAR] (50),
    customer_contact   [VARCHAR] (50),
    amount             [DOUBLE],
    sales_tax          [DOUBLE],
    paid_total         [DOUBLE],
    total              [DOUBLE],
    discount           [INT],
    payment_gateway    [VARCHAR] (50),
    logistics_provider [VARCHAR] (50),
    delivery_fee       [INT],
    delivery_time      TEXT,
    deleted_at         TEXT,
    created_at         TEXT,
    updated_at         TEXT
);

INSERT INTO customer_order (id, customer_id, coupon_id, parent_id, shop_id, payment_id, tracking_number, customer_contact, amount, sales_tax, paid_total, total, discount, payment_gateway, logistics_provider, delivery_fee, delivery_time, deleted_at, created_at, updated_at) VALUES (1, 1, 1, 1, 1, 1, 'KN72GQqD4jJ0', '19365141641631', 21.6, 0.432, 72.032, 72.032, 0, 'cod', 'NULL', 50, '11.00 AM - 2.00 PM', '2021-08-26T11:24:26.000000Z', '2021-08-26T11:24:26.000000Z', '2021-08-26T11:24:26.000000Z');

-- Table: deliveryTime
CREATE TABLE deliveryTime (
    id          INTEGER        PRIMARY KEY AUTOINCREMENT,
    setting_id  [INT],
    shop_id     INT,
    title       [VARCHAR] (50),
    description [VARCHAR] (50) 
);

INSERT INTO deliveryTime (id, setting_id, shop_id, title, description) VALUES (1, 1, 1, 'Express Delivery', '90 min express delivery');
INSERT INTO deliveryTime (id, setting_id, shop_id, title, description) VALUES (2, 1, 1, 'Morning', '8.00 AM - 11.00 AM');
INSERT INTO deliveryTime (id, setting_id, shop_id, title, description) VALUES (3, 1, 1, 'Noon', '11.00 AM - 2.00 PM');
INSERT INTO deliveryTime (id, setting_id, shop_id, title, description) VALUES (4, 1, 1, 'Afternoon', '2.00 PM - 5.00 PM');
INSERT INTO deliveryTime (id, setting_id, shop_id, title, description) VALUES (5, 1, 1, 'Evening', '5.00 PM - 8.00 PM');

-- Table: location
CREATE TABLE location (
    id               INTEGER        PRIMARY KEY AUTOINCREMENT,
    shop_id          [INT]          REFERENCES shop (id),
    setting_id,
    lat              [DOUBLE],
    lng              [DOUBLE],
    city             [VARCHAR] (50),
    state            [VARCHAR] (50),
    country          [VARCHAR] (50),
    formattedAddress [VARCHAR] (50) 
);

INSERT INTO location (id, shop_id, setting_id, lat, lng, city, state, country, formattedAddress) VALUES (1, 1, 1, 40.757272, -74.089508, 'Kearny', 'NJ', 'United States', 'New Jersey Turnpike, Kearny, NJ, USA');

-- Table: logo
CREATE TABLE logo (
    id         INTEGER         PRIMARY KEY AUTOINCREMENT,
    shop_id    [INT]           REFERENCES shop (id),
    setting_id INT,
    original   [VARCHAR] (500),
    thumbnail  [VARCHAR] (500) 
);

INSERT INTO logo (id, shop_id, setting_id, original, thumbnail) VALUES (1, 1, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/PickBazar.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/conversions/PickBazar-thumbnail.jpg');
INSERT INTO logo (id, shop_id, setting_id, original, thumbnail) VALUES (2, 2, 2, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/PickBazar.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/conversions/PickBazar-thumbnail.jpg');
INSERT INTO logo (id, shop_id, setting_id, original, thumbnail) VALUES (3, 3, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/PickBazar.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/conversions/PickBazar-thumbnail.jpg');
INSERT INTO logo (id, shop_id, setting_id, original, thumbnail) VALUES (4, 4, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/PickBazar.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/conversions/PickBazar-thumbnail.jpg');
INSERT INTO logo (id, shop_id, setting_id, original, thumbnail) VALUES (5, 5, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/PickBazar.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/conversions/PickBazar-thumbnail.jpg');
INSERT INTO logo (id, shop_id, setting_id, original, thumbnail) VALUES (6, 6, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/PickBazar.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/860/conversions/PickBazar-thumbnail.jpg');

-- Table: migrations
CREATE TABLE migrations (
    name TEXT PRIMARY KEY
);

INSERT INTO migrations (name) VALUES ('migration/00000000.sql');

-- Table: order_children
CREATE TABLE order_children (
    id                 INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_order_Id  [INT],
    customer_id        [INT],
    shop_id            [INT],
    parent_id          [INT],
    payment_id         [INT],
    coupon_id          [INT],
    tracking_number    [VARCHAR] (50),
    customer_contact   [VARCHAR] (50),
    amount             [DOUBLE],
    sales_tax          [INT],
    paid_total         [DOUBLE],
    total              [DOUBLE],
    discount           [INT],
    payment_gateway    [VARCHAR] (50),
    delivery_fee       [INT],
    logistics_provider [VARCHAR] (50),
    delivery_time      TEXT,
    deleted_at         TEXT,
    created_at         TEXT,
    updated_at         TEXT
);

INSERT INTO order_children (id, customer_order_Id, customer_id, shop_id, parent_id, payment_id, coupon_id, tracking_number, customer_contact, amount, sales_tax, paid_total, total, discount, payment_gateway, delivery_fee, logistics_provider, delivery_time, deleted_at, created_at, updated_at) VALUES (1, 1, 1, 1, 1, 1, 1, 'bVH8G97r6wSC', '19365141641631', 21.6, 0, 21.6, 21.6, 0, 'cod', 0, NULL, '11.00 AM - 2.00 PM', '2021-08-26T11:24:26.000000Z', '2021-08-26T11:24:26.000000Z', '2021-08-26T11:24:26.000000Z');

-- Table: order_status
CREATE TABLE order_status (
    id                INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_order_Id [INT],
    name              [VARCHAR] (50),
    serial            [INT],
    color             [VARCHAR] (50),
    created_at        TEXT,
    updated_at        TEXT
);

INSERT INTO order_status (id, customer_order_Id, name, serial, color, created_at, updated_at) VALUES (1, 1, 'Order Received', 1, '#23b848', '2021-03-08T21:33:52.000000Z', '2021-03-08T21:34:04.000000Z');

-- Table: parent_child
CREATE TABLE parent_child (
    id        INTEGER PRIMARY KEY AUTOINCREMENT,
    parent_id [INT],
    child_id  [INT]
);


-- Table: pivot
CREATE TABLE pivot (
    id                      INTEGER  PRIMARY KEY AUTOINCREMENT,
    product_Id              [INT],
    customer_order_id       [INT],
    customer_order_quantity [DOUBLE],
    unit_price              [DOUBLE],
    subtotal                [DOUBLE],
    created_at              TEXT,
    updated_at              TEXT,
    category_Id             [INT]
);

INSERT INTO pivot (id, product_Id, customer_order_id, customer_order_quantity, unit_price, subtotal, created_at, updated_at, category_Id) VALUES (1, 1, NULL, NULL, NULL, NULL, NULL, NULL, 2);
INSERT INTO pivot (id, product_Id, customer_order_id, customer_order_quantity, unit_price, subtotal, created_at, updated_at, category_Id) VALUES (2, 2, NULL, NULL, NULL, NULL, NULL, NULL, 1);

-- Table: product
CREATE TABLE product (
    id                INTEGER                   PRIMARY KEY AUTOINCREMENT,
    shop_id           [INT],
    category_type_id  [INT],
    name              [VARCHAR] (100),
    slug              [VARCHAR] (50),
    description       [VARCHAR] (500),
    price             [DOUBLE],
    sale_price        [DOUBLE],
    sku               [VARCHAR] (50),
    quantity          [INT],
    in_stock          [INT],
    is_taxable        [INT],
    status            [VARCHAR] (50),
    shipping_class_id VARCHAR (50),
    product_type      [VARCHAR] (50),
    unit              [VARCHAR] (50),
    height            [DOUBLE],
    length            [DOUBLE],
    width             [DOUBLE],
    deleted_at        TEXT,
    created_at        TEXT,
    updated_at        "TEXT MAX_PRICE [DOUBLE]",
    max_price         [DOUBLE],
    min_price         [DOUBLE],
    video             [VARCHAR] (50) 
);

INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (1, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (2, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (3, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (4, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (5, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (6, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (7, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (8, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (9, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (10, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (11, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (12, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (13, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (14, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (15, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (16, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (17, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (18, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (19, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');
INSERT INTO product (id, shop_id, category_type_id, name, slug, description, price, sale_price, sku, quantity, in_stock, is_taxable, status, shipping_class_id, product_type, unit, height, length, width, deleted_at, created_at, updated_at, max_price, min_price, video) VALUES (20, 1, 1, 'Apples', 'apples', 'An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.', 2.0, 1.6, '1', 18, 1, 0, 'publish', NULL, 'simple', '1lb', 22.0, 2.0, 2.0, '2021-03-08T10:24:53.000000Z', '2021-03-08T10:24:53.000000Z', '2021-06-27T03:56:42.000000Z', 22.0, 22.0, 'video');

-- Table: product_gallery
CREATE TABLE product_gallery (
    id         INTEGER         PRIMARY KEY AUTOINCREMENT,
    product_Id [INT]           REFERENCES product (id),
    original   [VARCHAR] (200),
    thumbnail  [VARCHAR] (200) 
);

INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (1, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/apple-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/conversions/apple-1-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (2, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (3, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/apple.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/conversions/apple-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (4, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (5, 2, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/apple-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/conversions/apple-1-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (6, 2, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (7, 2, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/apple.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/conversions/apple-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (8, 2, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (9, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/apple-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/conversions/apple-1-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (10, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (11, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/apple.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/conversions/apple-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (12, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (13, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/apple-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/conversions/apple-1-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (14, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (15, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/apple.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/conversions/apple-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (16, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (17, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/apple-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/conversions/apple-1-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (18, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (19, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/apple.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/conversions/apple-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (20, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (21, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/apple-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/conversions/apple-1-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (22, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (23, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/apple.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/conversions/apple-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (24, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (25, 7, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/apple-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/conversions/apple-1-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (26, 7, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (27, 7, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/apple.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/conversions/apple-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (28, 7, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (29, 8, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/apple-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/conversions/apple-1-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (30, 8, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (31, 8, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/apple.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/conversions/apple-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (32, 8, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (33, 9, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/apple-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/conversions/apple-1-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (34, 9, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (35, 9, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/apple.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/conversions/apple-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (36, 9, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (37, 10, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/apple-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/572/conversions/apple-1-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (38, 10, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/573/conversions/apple-2-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (39, 10, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/apple.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/574/conversions/apple-thumbnail.jpg');
INSERT INTO product_gallery (id, product_Id, original, thumbnail) VALUES (40, 10, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/apple-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/736/conversions/apple-2-thumbnail.jpg');

-- Table: product_image
CREATE TABLE product_image (
    id         INTEGER         PRIMARY KEY AUTOINCREMENT,
    product_Id [INT]           REFERENCES product,
    original   [VARCHAR] (500),
    thumbnail  [VARCHAR] (500) 
);

INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (1, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (2, 2, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (3, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (4, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (5, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (6, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (7, 7, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (8, 8, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (9, 9, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (10, 10, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (11, 11, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (12, 12, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (13, 13, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (14, 14, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (15, 15, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (16, 16, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (17, 17, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (18, 18, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (19, 19, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');
INSERT INTO product_image (id, product_Id, original, thumbnail) VALUES (20, 20, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/Apples.jpg', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/1/conversions/Apples-thumbnail.jpg');

-- Table: profile
CREATE TABLE profile (
    id          INTEGER         PRIMARY KEY AUTOINCREMENT,
    user_id     [INT]           REFERENCES user (id),
    bio         [VARCHAR] (257),
    contact     [VARCHAR] (50),
    customer_id [INT],
    created_at  TEXT            NOT NULL,
    updated_at  TEXT            NOT NULL
);

INSERT INTO profile (id, user_id, bio, contact, customer_id, created_at, updated_at) VALUES (1, 1, 'This is the store owner and we have 6 shops under our banner. We are running all the shops to give our customers hassle-free service and quality products. Our goal is to provide best possible customer service and products for our clients', '12365141641631', 1, '2021-06-30T11:20:29.000000Z', '2021-06-30T14:13:53.000000Z');

-- Table: promotional_slider
CREATE TABLE promotional_slider (
    id               INTEGER         PRIMARY KEY AUTOINCREMENT,
    category_type_id [INT],
    original         [VARCHAR] (50),
    thumbnail        [VARCHAR] (114) 
);

INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (1, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/offer-5.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/conversions/offer-5-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (2, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/offer-4.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/conversions/offer-4-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (3, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/offer-3.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/conversions/offer-3-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (4, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/offer-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/conversions/offer-2-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (5, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/offer-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/conversions/offer-1-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (6, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/offer-5.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/conversions/offer-5-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (7, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/offer-4.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/conversions/offer-4-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (8, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/offer-3.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/conversions/offer-3-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (9, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/offer-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/conversions/offer-2-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (10, 3, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/offer-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/conversions/offer-1-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (11, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/offer-5.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/conversions/offer-5-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (12, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/offer-4.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/conversions/offer-4-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (13, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/offer-3.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/conversions/offer-3-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (14, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/offer-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/conversions/offer-2-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (15, 4, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/offer-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/conversions/offer-1-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (16, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/offer-5.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/conversions/offer-5-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (17, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/offer-4.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/conversions/offer-4-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (18, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/offer-3.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/conversions/offer-3-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (19, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/offer-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/conversions/offer-2-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (20, 5, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/offer-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/conversions/offer-1-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (21, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/offer-5.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/conversions/offer-5-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (22, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/offer-4.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/conversions/offer-4-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (23, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/offer-3.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/conversions/offer-3-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (24, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/offer-2.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/conversions/offer-2-thumbnail.jpg');
INSERT INTO promotional_slider (id, category_type_id, original, thumbnail) VALUES (25, 6, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/offer-1.png', 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/conversions/offer-1-thumbnail.jpg');

-- Table: seo
CREATE TABLE seo (
    id              INTEGER        PRIMARY KEY AUTOINCREMENT,
    setting_id      [INT],
    shop_id         INT,
    ogImage         [VARCHAR] (50),
    ogTitle         [VARCHAR] (50),
    metaTags        [VARCHAR] (50),
    metaTitle       [VARCHAR] (50),
    canonicalUrl    [VARCHAR] (50),
    ogDescription   [VARCHAR] (50),
    twitterHandle   [VARCHAR] (50),
    metaDescription [VARCHAR] (50),
    twitterCardType [VARCHAR] (50) 
);

INSERT INTO seo (id, setting_id, shop_id, ogImage, ogTitle, metaTags, metaTitle, canonicalUrl, ogDescription, twitterHandle, metaDescription, twitterCardType) VALUES (1, 1, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

-- Table: setting
CREATE TABLE setting (
    id               INTEGER        PRIMARY KEY AUTOINCREMENT,
    shop_id          [INT],
    category_id      [INT],
    category_type_id INT,
    isHome           INT,
    layoutType       [VARCHAR] (50),
    productCard      [VARCHAR] (50),
    created_at       TEXT,
    updated_at       TEXT
);

INSERT INTO setting (id, shop_id, category_id, category_type_id, isHome, layoutType, productCard, created_at, updated_at) VALUES (1, 1, 1, 1, 1, 'classic', 'neon', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO setting (id, shop_id, category_id, category_type_id, isHome, layoutType, productCard, created_at, updated_at) VALUES (2, 2, NULL, 2, 0, 'standard', 'argon', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO setting (id, shop_id, category_id, category_type_id, isHome, layoutType, productCard, created_at, updated_at) VALUES (3, 3, NULL, 3, 1, 'classic', 'helium', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO setting (id, shop_id, category_id, category_type_id, isHome, layoutType, productCard, created_at, updated_at) VALUES (4, 4, NULL, 4, 1, 'classic', 'helium', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO setting (id, shop_id, category_id, category_type_id, isHome, layoutType, productCard, created_at, updated_at) VALUES (5, 5, NULL, 5, 1, 'classic', 'xenon', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO setting (id, shop_id, category_id, category_type_id, isHome, layoutType, productCard, created_at, updated_at) VALUES (6, 6, NULL, 6, 1, 'modern', 'krypton', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');
INSERT INTO setting (id, shop_id, category_id, category_type_id, isHome, layoutType, productCard, created_at, updated_at) VALUES (7, 6, NULL, 7, 1, 'minimal', 'neon', '2021-03-08T07:18:25.000000Z', '2021-09-26T15:23:32.000000Z');

-- Table: setting_option
CREATE TABLE setting_option (
    int                INTEGER        PRIMARY KEY AUTOINCREMENT,
    setting_id         [INT],
    currency           [VARCHAR] (50),
    taxClass           [INT],
    siteTitle          [VARCHAR] (50),
    siteSubtitle       [VARCHAR] (50),
    shippingClass      [INT],
    minimumOrderAmount INT
);

INSERT INTO setting_option (int, setting_id, currency, taxClass, siteTitle, siteSubtitle, shippingClass, minimumOrderAmount) VALUES (1, 1, 'USD', 1, 'Pickbazar', 'Your next ecommerce', 1, 1);

-- Table: shipping
CREATE TABLE shipping (
    id         INTEGER        PRIMARY KEY AUTOINCREMENT,
    name       [VARCHAR] (50),
    amount     [INT],
    is_global  [INT],
    type       [VARCHAR] (50),
    created_at TEXT,
    updated_at TEXT
);

INSERT INTO shipping (id, name, amount, is_global, type, created_at, updated_at) VALUES (1, 'Global', 50, 1, 'fixed', '2021-03-25T13:27:49.000000Z', '2021-03-25T13:27:49.000000Z');

-- Table: shipping_address
CREATE TABLE shipping_address (
    id                INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_order_Id [INT],
    zip               [VARCHAR] (50),
    city              [VARCHAR] (50),
    state             [VARCHAR] (50),
    country           [VARCHAR] (50),
    street_address    [VARCHAR] (50) 
);

INSERT INTO shipping_address (id, customer_order_Id, zip, city, state, country, street_address) VALUES (1, 1, '40391', 'Winchester', 'KY', 'United States', '2148  Straford Park');

-- Table: shop
CREATE TABLE shop (
    id          INTEGER        PRIMARY KEY AUTOINCREMENT,
    owner_id    [INT],
    name        [VARCHAR] (50),
    slug        [VARCHAR] (50),
    description TEXT,
    is_active   [INT],
    created_at  TEXT           NOT NULL,
    updated_at  TEXT           NOT NULL
);

INSERT INTO shop (id, owner_id, name, slug, description, is_active, created_at, updated_at) VALUES (1, 1, 'Furniture Shop', 'furniture-shop', 'The furniture shop is the best shop around the city. This is being run under the store owner and our aim is to provide quality product and hassle free customer service.', 1, '2021-06-27T03:46:14.000000Z', '2021-07-08T09:27:14.000000Z');
INSERT INTO shop (id, owner_id, name, slug, description, is_active, created_at, updated_at) VALUES (2, 1, 'Clothing Shop', 'clothing-shop', 'The clothing shop is the best shop around the city. This is being run under the store owner and our aim is to provide quality product and hassle free customer service.', 1, '2021-06-27T03:47:10.000000Z', '2021-07-08T09:26:24.000000Z');
INSERT INTO shop (id, owner_id, name, slug, description, is_active, created_at, updated_at) VALUES (3, 1, 'Bags Shop', 'bags-shop', 'The Bag shop is the best shop around the city. This is being run under the store owner and our aim is to provide quality product and hassle free customer service.', 1, '2021-06-27T03:47:10.000000Z', '2021-07-08T09:26:24.000000Z');
INSERT INTO shop (id, owner_id, name, slug, description, is_active, created_at, updated_at) VALUES (4, 1, 'Makeup Shop', 'makeup-shop', 'The makeup shop is the best shop around the city. This is being run under the store owner and our aim is to provide quality product and hassle free customer service.', 1, '2021-06-27T03:47:10.000000Z', '2021-07-08T09:26:24.000000Z');
INSERT INTO shop (id, owner_id, name, slug, description, is_active, created_at, updated_at) VALUES (5, 1, 'Bakery Shop', 'bakery-shop', 'The bakery shop is the best shop around the city. This is being run under the store owner and our aim is to provide fresh and quality product and hassle free customer service.', 1, '2021-06-27T03:47:10.000000Z', '2021-07-08T09:26:24.000000Z');
INSERT INTO shop (id, owner_id, name, slug, description, is_active, created_at, updated_at) VALUES (6, 1, 'Grocery Shop', 'grocery-shop', 'The grocery shop is the best shop around the city. This is being run under the store owner and our aim is to provide fresh and quality product and hassle free customer service.', 1, '2021-06-27T03:47:10.000000Z', '2021-07-08T09:26:24.000000Z');

-- Table: social
CREATE TABLE social (
    id         INTEGER        PRIMARY KEY AUTOINCREMENT,
    shop_id    [INT]          REFERENCES shop (id),
    setting_id INT,
    url        [VARCHAR] (50),
    icon       [VARCHAR] (50) 
);

INSERT INTO social (id, shop_id, setting_id, url, icon) VALUES (2, 1, 1, 'https://www.facebook.com/', 'FacebookIcon');
INSERT INTO social (id, shop_id, setting_id, url, icon) VALUES (3, 1, 1, 'https://twitter.com/home', 'TwitterIcon');
INSERT INTO social (id, shop_id, setting_id, url, icon) VALUES (4, 1, 1, 'https://www.instagram.com/', 'InstagramIcon');

-- Table: tax
CREATE TABLE tax (
    id          INTEGER        PRIMARY KEY AUTOINCREMENT,
    country     [VARCHAR] (50),
    state       [VARCHAR] (50),
    zip         [VARCHAR] (50),
    city        [VARCHAR] (50),
    rate        [INT],
    name        [VARCHAR] (50),
    is_global   [BIT],
    priority    [VARCHAR] (50),
    on_shipping [BIT],
    created_at                 DEFAULT (CURRENT_TIMESTAMP),
    updated_at                 DEFAULT (CURRENT_TIMESTAMP) 
);

INSERT INTO tax (id, country, state, zip, city, rate, name, is_global, priority, on_shipping, created_at, updated_at) VALUES (1, NULL, NULL, NULL, NULL, 2, 'Global', 1, NULL, 1, '2021-03-25T13:26:57.000000Z', '2021-03-25T16:07:18.000000Z');

-- Table: user
CREATE TABLE user (
    Id                INTEGER        PRIMARY KEY AUTOINCREMENT,
    name              [VARCHAR] (50),
    email             [VARCHAR] (50),
    api_key           TEXT           NOT NULL
                                     UNIQUE,
    email_verified_at DATETIME,
    is_active         [INT],
    shop_Id           [INT],
    profile_id        [INT],
    created_at        TEXT           NOT NULL,
    updated_at        TEXT           NOT NULL
);

INSERT INTO user (Id, name, email, api_key, email_verified_at, is_active, shop_Id, profile_id, created_at, updated_at) VALUES (1, 'Mustapha', 'mustapha.manjoura@gmail.com', '2d92763a84fd760444793a99432735378fce7e3d7af73fc213a81d3a08a9e715', 'null', 1, 1, 1, '2022-01-07T12:37:57Z', '2022-01-10T13:38:33Z');
INSERT INTO user (Id, name, email, api_key, email_verified_at, is_active, shop_Id, profile_id, created_at, updated_at) VALUES (2, 'mohamed', 'mohamed.manjoura@gmail.com', '521155290e1f40a40713f6f183e508486774175ca2f6ba3916d55fbf93fdbb5c', 'null', 1, 1, 1, '2022-01-07T12:37:57Z', '2022-01-10T13:38:33Z');

-- Table: value
CREATE TABLE value (
    id         INTEGER      PRIMARY KEY AUTOINCREMENT,
    shop_id                 REFERENCES shop (id),
    value      VARCHAR (50) NOT NULL,
    meta       VARCHAR (50) NOT NULL,
    created_at TEXT,
    updated_at TEXT
);

INSERT INTO value (id, shop_id, value, meta, created_at, updated_at) VALUES (1, 1, 'Red', '#ce1f6a', '2021-05-09T16:10:56.000000Z', '2021-05-09T18:53:16.000000Z');
INSERT INTO value (id, shop_id, value, meta, created_at, updated_at) VALUES (2, 1, 'Blue', '#344fa1', '2021-05-09T16:11:20.000000Z', '2021-05-09T18:52:35.000000Z');

COMMIT TRANSACTION;
PRAGMA foreign_keys = on;
