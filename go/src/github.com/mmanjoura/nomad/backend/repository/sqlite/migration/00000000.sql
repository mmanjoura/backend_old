CREATE TABLE user (
    Id         INTEGER PRIMARY KEY AUTOINCREMENT,
    [name]              [VARCHAR] (50),
    email             [VARCHAR] (50),
	api_key    TEXT NOT NULL UNIQUE,
    email_verified_at DATETIME,
    is_active         [INT],
    shop_Id           [INT],
	profile_id        [INT],
    created_at                        TEXT NOT NULL,
    updated_at                        TEXT NOT NULL  
);


CREATE TABLE auth (
	Id            INTEGER PRIMARY KEY AUTOINCREMENT,
	user_Id       INTEGER NOT NULL REFERENCES user (Id) ON DELETE CASCADE,
	source        TEXT NOT NULL,
	source_Id     TEXT NOT NULL,
	access_token  TEXT NOT NULL,
	refresh_token TEXT NOT NULL,
	expiry        TEXT,
	created_at    TEXT NOT NULL,
	updated_at    TEXT NOT NULL,

	UNIQUE(user_Id, source),  -- one source per user
	UNIQUE(source, source_Id) -- one auth per source user
);

----------------------------PickBazar attributes--------------------------------------------
CREATE TABLE shop (
    id             INTEGER        PRIMARY KEY AUTOINCREMENT,
    owner_id       [INT]          REFERENCES user (id),
    name           [VARCHAR] (50),
    slug           [VARCHAR] (50),
    description    TEXT,
    is_active      [INT],
	created_at     TEXT           NOT NULL,
    updated_at     TEXT           NOT NULL
);

CREATE TABLE location (
    id               INTEGER        PRIMARY KEY AUTOINCREMENT,
	shop_id  [INT] REFERENCES shop (id),
    lat              [DOUBLE],
    lng              [DOUBLE],
    city             [VARCHAR] (50),
    state            [VARCHAR] (50),
    country          [VARCHAR] (50),
    formattedAddress [VARCHAR] (50) 
);

CREATE TABLE avatar (
    id        INTEGER         PRIMARY KEY AUTOINCREMENT,
	user_id  [INT] REFERENCES user (id),
    original  [VARCHAR] (124),
    thumbnail [VARCHAR] (146) 
);

CREATE TABLE profile (
    id          INTEGER         PRIMARY KEY AUTOINCREMENT,
    user_id  [INT] REFERENCES user (id),
    bio         [VARCHAR] (257),
    contact     [VARCHAR] (50),
    customer_id [INT],
    created_at                  TEXT           NOT NULL,
    updated_at                  TEXT           NOT NULL 
);
CREATE TABLE social (
    id            INTEGER        PRIMARY KEY AUTOINCREMENT,
	shop_id  [INT] REFERENCES shop (id),
    url                  [VARCHAR] (50),
    icon                 [VARCHAR] (50) 
);

CREATE TABLE setting (
    id                 INTEGER        PRIMARY KEY AUTOINCREMENT,
    shop_id            [INT]  ,
    category_id        [INT],
    contact            [VARCHAR] (50),
    website            [VARCHAR] (50),
    isHome             INT,
    layoutType         [VARCHAR] (50),
    productCard        [VARCHAR] (50),
        created_at TEXT,
    updated_at TEXT
);

CREATE TABLE seo (
    id              INTEGER        PRIMARY KEY AUTOINCREMENT,
    shop_id         [INT],
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

CREATE TABLE shop_option (
    int       INTEGER        PRIMARY KEY AUTOINCREMENT,
    shop_id         [INT],
    currency             [VARCHAR] (50),
    taxClass             [INT],
    siteTitle            [VARCHAR] (50),
    siteSubtitle         [VARCHAR] (50),
    shippingClass        [INT],
    minimumOrderAmount   [INT]
);



CREATE TABLE cover_image (
    id        INTEGER         PRIMARY KEY AUTOINCREMENT,
	shop_id  [INT] REFERENCES shop (id),
    original  [VARCHAR] (500),
    thumbnail [VARCHAR] (500) 
);

CREATE TABLE logo (
    id        INTEGER         PRIMARY KEY AUTOINCREMENT,
	shop_id  [INT] REFERENCES shop (id),
    original  [VARCHAR] (500),
    thumbnail [VARCHAR] (500) 
);

CREATE TABLE address (
    id             INTEGER        PRIMARY KEY AUTOINCREMENT,
	Shop_id    [INT]       REFERENCES shop (id),
    zip            [VARCHAR] (50),
    city           [VARCHAR] (50),
    state          [VARCHAR] (50),
    country        [VARCHAR] (50),
    street_address [VARCHAR] (50),
    type           [VARCHAR] (50),
    title          [NCHAR] (50),
    [default]      [INT]
);

CREATE TABLE attribute (
    id         INTEGER        PRIMARY KEY AUTOINCREMENT,
    Shop_id    [INT]       REFERENCES shop (id),
    slug       [VARCHAR] (50),
    name       [VARCHAR] (50),
    created_at TEXT,
    updated_at TEXT
);

CREATE TABLE value (
    id           INTEGER      PRIMARY KEY AUTOINCREMENT,
    shop_id REFERENCES shop (id),
    value        VARCHAR (50) NOT NULL,
    meta         VARCHAR (50) NOT NULL,
    created_at   TEXT,
    updated_at   TEXT
);

----------------------------PickBazar Categories--------------------------------------------
CREATE TABLE category (
    id         INTEGER        PRIMARY KEY AUTOINCREMENT,
    type_id    [INT],
    parent_Id  INTEGER,
    product_id  [INT],
    name       [VARCHAR] (50),
    slug       [VARCHAR] (50),
    icon       [VARCHAR] (50),
    deleted_at TEXT,
    created_at TEXT,
    updated_at TEXT
);
CREATE TABLE category_type (
    id     INTEGER        PRIMARY KEY AUTOINCREMENT,
    category_Id [INT]          REFERENCES category (id),
    name        [VARCHAR] (50),
    slug        [VARCHAR] (50),
    icon        [VARCHAR] (50),
    created_at  TEXT,
    updated_at  TEXT
);

CREATE TABLE promotional_slider (
    id               INTEGER         PRIMARY KEY AUTOINCREMENT,
    category_type_id [INT],
    shop_id       [INT],
     product_id       [INT],
    original         [VARCHAR] (50),
    thumbnail        [VARCHAR] (114) 
);


CREATE TABLE category_children (
    id               INTEGER        PRIMARY KEY AUTOINCREMENT,
    category_Id      [INT]          REFERENCES category (id),
    parent_id        [INT],
    name             [VARCHAR] (50),
    slug             [VARCHAR] (50),
    icon             [VARCHAR] (50),
    created_at       TEXT,
    updated_at       TEXT,
    deleted_at       TEXT
);

CREATE TABLE image (
    id                   INTEGER        PRIMARY KEY AUTOINCREMENT,
    category_Id          [INT],
    coupon_Id            [INT],
    product_Id           [INT],
    is_cover             [INT],
    is_main              [INT],
    is_logo              [INT],
    is_gallery           [INT],
    category_children_id [INT],
    shop_id              [INT],
    original             [VARCHAR] (50),
    thumbnail            [VARCHAR] (50) 
);


---------------------------- Start coupon----------------------------------------
CREATE TABLE coupon (
    id   INTEGER        PRIMARY KEY AUTOINCREMENT,
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

---------------------------- Start product----------------------------------------
CREATE TABLE product (
    id                INTEGER                   PRIMARY KEY AUTOINCREMENT,
    shop_id           [INT],
    shipping_class_id [INT],
    type_id           [INT],
    name              [VARCHAR] (50),
    slug              [VARCHAR] (50),
    description       [VARCHAR] (311),
    price             [DOUBLE],
    sale_price        [DOUBLE],
    sku               [VARCHAR] (50),
    quantity          [INT],
    in_stock          [INT],
    is_taxable        [INT],
    status            [VARCHAR] (50),
    product_type      [VARCHAR] (50),
    unit              [VARCHAR] (50),
    height            [DOUBLE],
    length            [DOUBLE],
    width             [DOUBLE],
    deleted_at        TEXT,
    created_at        TEXT,
    updated_at        TEXT,
    max_price         [DOUBLE],
    min_price         [DOUBLE],
    video             [VARCHAR] (50) 
)

CREATE TABLE gallery (
    id         INTEGER         PRIMARY KEY AUTOINCREMENT,
    product_Id [INT]           REFERENCES product (id),
    original   [VARCHAR] (200),
    thumbnail  [VARCHAR] (200) 
);


---------------------------- Orders ----------------------------------------

CREATE TABLE customer_order (
    id                 INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_id        [INT],
    coupon_id          [INT],
    parent_id          [INT],
    shop_id            [INT],
    payment_id         [INT],
    coupon_id         [INT],
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

CREATE TABLE shipping_address (
    id                INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_order_Id [INT],
    zip               [VARCHAR] (50),
    city              [VARCHAR] (50),
    state             [VARCHAR] (50),
    country           [VARCHAR] (50),
    street_address    [VARCHAR] (50) 
);

CREATE TABLE billing_address (
    id                INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_order_Id [INT],
    zip               [VARCHAR] (50),
    city              [VARCHAR] (50),
    state             [VARCHAR] (50),
    country           [VARCHAR] (50),
    street_address    [VARCHAR] (50) 
);

CREATE TABLE order_status (
    id                INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_order_Id [INT],
    name              [VARCHAR] (50),
    serial            [INT],
    color             [VARCHAR] (50),
    created_at        TEXT,
    updated_at        TEXT
);

CREATE TABLE order_children (
    id                INTEGER        PRIMARY KEY AUTOINCREMENT,
    customer_order_Id [INT],
    customer_id       [INT],
    shop_id           [INT],
    parent_id         [INT],
    payment_id        [INT],
    coupon_id        [INT],
    tracking_number   [VARCHAR] (50),
    customer_contact  [VARCHAR] (50),
    amount            [DOUBLE],
    sales_tax         [INT],
    paid_total        [DOUBLE],
    total             [DOUBLE],
    discount          [INT],
    payment_gateway   [VARCHAR] (50),
    delivery_fee      [INT],
    delivery_time     TEXT,
    created_at        TEXT,
    updated_at        TEXT
);

CREATE TABLE pivot (
    id                INTEGER  PRIMARY KEY AUTOINCREMENT,
    product_Id              [INT]   ,
    customer_order_id       [INT],
    categorie_Id            [INT]
    customer_order_quantity [DOUBLE],
    unit_price              [DOUBLE],
    subtotal                [DOUBLE],
    created_at               TEXT,
    updated_at               TEXT
    
);

CREATE TABLE shipping (
    id INTEGER        PRIMARY KEY AUTOINCREMENT,
    name        [VARCHAR] (50),
    amount      [INT],
    is_global   [INT],
    type        [VARCHAR] (50),
    created_at                 TEXT,
    updated_at                 TEXT 
);


CREATE TABLE tax (
    id      INTEGER        PRIMARY KEY AUTOINCREMENT,
    country     [VARCHAR] (50),
    state       [VARCHAR] (50),
    zip         [VARCHAR] (50),
    city        [VARCHAR] (50),
    rate        [INT],
    name        [VARCHAR] (50),
    is_global   [INT],
    priority    [VARCHAR] (50),
    on_shipping [INT],
    created_at                 DEFAULT (CURRENT_TIMESTAMP),
    updated_at                 DEFAULT (CURRENT_TIMESTAMP) 
);


---------------------------- Start Inserts User/Auth----------------------------------------

INSERT INTO user([name], email, api_key, email_verified_at, is_active, shop_Id , profile_id, created_at, updated_at ) VALUES ('Mustapha','mustapha.manjoura@gmail.com', '2d92763a84fd760444793a99432735378fce7e3d7af73fc213a81d3a08a9e715',  'null', 1, 1, 1, '2022-01-07T12:37:57Z', '2022-01-10T13:38:33Z');
INSERT INTO user([name], email, api_key, email_verified_at, is_active, shop_Id , profile_id, created_at, updated_at ) VALUES ('mohamed','mohamed.manjoura@gmail.com', '521155290e1f40a40713f6f183e508486774175ca2f6ba3916d55fbf93fdbb5c',  'null', 1, 1, 1, '2022-01-07T12:37:57Z', '2022-01-10T13:38:33Z');
INSERT INTO auth ( user_id, source, source_id, access_token, refresh_token, created_at, updated_at ) VALUES ( 1, 'github', 5604914, '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', '', '2022-01-07T12:37:57Z', '2022-01-10T13:38:33Z' );

----------------------------PickBazar attributes--------------------------------------------
INSERT INTO shop(owner_id,name,slug,description, is_active, created_at,updated_at,orders_count,products_count) VALUES (1,'Furniture Shop','furniture-shop','The furniture shop is the best shop around the city. This is being run under the store owner and our aim is to provide quality product and hassle free customer service.', 1, '2021-06-27T03:46:14.000000Z','2021-07-08T09:27:14.000000Z',0,55);

INSERT INTO location(shop_id, lat,lng,city,state,country,formattedAddress) VALUES (1, 40.757272,-74.089508,'Kearny','NJ','United States','New Jersey Turnpike, Kearny, NJ, USA');
INSERT INTO avatar(user_id, original,thumbnail) VALUES (1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/881/aatik-tasneem-7omHUGhhmZ0-unsplash%402x.png','https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/881/conversions/aatik-tasneem-7omHUGhhmZ0-unsplash%402x-thumbnail.jpg');
INSERT INTO profile(user_id, bio,contact,customer_id,created_at,updated_at) VALUES (1, 'This is the store owner and we have 6 shops under our banner. We are running all the shops to give our customers hassle-free service and quality products. Our goal is to provide best possible customer service and products for our clients','12365141641631',1,'2021-06-30T11:20:29.000000Z','2021-06-30T14:13:53.000000Z');
INSERT INTO social(shop_id, url,icon) VALUES ( 1, 'https://www.instagram.com/','InstagramIcon');
INSERT INTO setting(shop_id, category_id, contact, website, isHome, layoutType, productCard) VALUES ( 1, 1, '212901921221', 'https://www.instagram.com/', 1, "layout", "productcard");
INSERT INTO cover_image(shop_id, original,thumbnail) VALUES (1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/Untitled-6.jpg','https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/883/conversions/Untitled-6-thumbnail.jpg');
INSERT INTO logo(shop_id, original,thumbnail) VALUES (1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/882/Furniture.png','https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/882/conversions/Furniture-thumbnail.jpg');
INSERT INTO address(shop_id, zip,city,state,country,street_address) VALUES (1, '08753','East Dover','New Jersey','USA','588  Finwood Road');

INSERT INTO attribute(shop_id,  slug,name, created_at,updated_at) VALUES (1, 'color','Color','2021-05-09T16:10:31.000000Z','2021-05-09T16:10:31.000000Z');
INSERT INTO value(shop_id,value,meta,created_at,updated_at) VALUES (1,'Red','#ce1f6a','2021-05-09T16:10:56.000000Z','2021-05-09T18:53:16.000000Z');
INSERT INTO value(shop_id,value,meta,created_at,updated_at) VALUES (1,'Blue','#344fa1','2021-05-09T16:11:20.000000Z','2021-05-09T18:52:35.000000Z');

----------------------------PickBazar categories--------------------------------------------

INSERT INTO category(type_id,parent_id, product_id, name, slug,icon,created_at,updated_at,deleted_at) VALUES (1, 1, 1, 'Fruits & Vegetables','fruits-vegetables','FruitsVegetable','2021-03-08T07:21:31.000000Z','2021-03-08T07:21:31.000000Z', '2021-03-08T07:21:31.000000Z');
INSERT INTO category_type(category_Id, name,slug,icon,created_at,updated_at) VALUES (1, 'Grocery','grocery','FruitsVegetable','2021-03-08T07:18:25.000000Z','2021-09-26T15:23:32.000000Z');
INSERT INTO promotional_slider(category_type_id, shop_id, product_id, original,thumbnail) VALUES (1, 1, 1, 'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/offer-5.png','https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/conversions/offer-5-thumbnail.jpg');
INSERT INTO category_children(category_Id, parent_id, name,slug,icon,created_at,updated_at, deleted_at) VALUES (1, 1,'Fruits','fruits', 'icon','2021-03-08T07:22:04.000000Z','2021-03-08T07:22:04.000000Z','2021-03-08T07:22:04.000000Z');
INSERT INTO image(category_Id, coupon_Id, product_Id, category_children_id, shop_id, is_cover, is_main, is_logo, is_gallery, original, thumbnail) VALUES (1, 1, 1, 1, 1, 1,1,1,1, 'dummy image','dummy thumbnail');

---------------------------- Start Inserts coupon----------------------------------------

INSERT INTO coupon(code,description,type,amount,active_from,expire_at,created_at,updated_at,deleted_at,is_valid) VALUES ('EID2','long sescription','fixed',2,'2021-03-10T21:48:15.000Z','2024-07-27T21:48:15.000Z','2021-03-10T21:49:46.000000Z','2021-08-19T04:00:12.000000Z','2021-08-19T04:00:12.000000Z',1);

---------------------------- Start Inserts product----------------------------------------

INSERT INTO product(shop_id,shipping_class_id,type_id,name,slug,description,price,sale_price,sku,quantity,in_stock,is_taxable,status,product_type,unit,height,width,length,deleted_at,created_at,updated_at,max_price,min_price,video) VALUES (6,0,1,'Apples','apples','An apple is a sweet, edible fruit produced by an apple tree (Malus domestica). Apple trees are ... The skin of ripe apples is generally red, yellow, green, pink, or russetted, though many bi- or tri-colored cultivars may be found.',2,1.6,1,18,1,0,'publish','simple','1lb',22,2,2,'2021-03-08T10:24:53.000000Z','2021-03-08T10:24:53.000000Z','2021-06-27T03:56:42.000000Z',22,22,'video');

---------------------------- Start Inserts customer order----------------------------------------
INSERT INTO order_status(customer_order_Id,name,serial,color,created_at,updated_at) VALUES (1,'Order Received',1,'#23b848','2021-03-08T21:33:52.000000Z','2021-03-08T21:34:04.000000Z');
INSERT INTO shipping_address(customer_order_Id,zip,city,state,country,street_address) VALUES (1,40391,'Winchester','KY','United States','2148  Straford Park');
INSERT INTO billing_address(customer_order_Id,zip,city,state,country,street_address) VALUES (1,40391,'Winchester','KY','United States','2148  Straford Park');
INSERT INTO customer(customer_order_Id,shop_id,name,email,email_verified_at,created_at,updated_at,is_active) VALUES (1,1,'Customer','customer@demo.com','2021-08-18T10:30:29.000000Z','2021-08-18T10:30:29.000000Z','2021-08-18T13:17:53.000000Z',1);
INSERT INTO order_children(customer_order_Id,coupon_id,parent_id,shop_id,payment_id,customer_id,tracking_number,customer_contact,amount,sales_tax,paid_total,total,discount,payment_gateway,logistics_provider,delivery_fee,delivery_time,deleted_at,created_at,updated_at) VALUES (1,1,1,1,1,1,'bVH8G97r6wSC',19365141641631,21.6,0,21.6,21.6,0,'cod',NULL,0,'11.00 AM - 2.00 PM','2021-08-26T11:24:26.000000Z','2021-08-26T11:24:26.000000Z','2021-08-26T11:24:26.000000Z');
INSERT INTO customer_order(id,customer_id,tracking_number,coupon_id,parent_id,shop_id,payment_id,customer_contact,amount,sales_tax,paid_total,total,discount,payment_gateway,logistics_provider,logistics_provider, delivery_fee,delivery_time,deleted_at,created_at,updated_at) VALUES (1,1,'KN72GQqD4jJ0',1,1,1,1,19365141641631,21.6,0.432,72.032,72.032,0,'cod',"NULL",50,'11.00 AM - 2.00 PM','2021-08-26T11:24:26.000000Z','2021-08-26T11:24:26.000000Z','2021-08-26T11:24:26.000000Z');

INSERT INTO tax(id,country,state,zip,city,rate,name,is_global,priority,on_shipping,created_at,updated_at) VALUES (1,NULL,NULL,NULL,NULL,2,'Global',1,NULL,1,'2021-03-25T13:26:57.000000Z','2021-03-25T16:07:18.000000Z');




