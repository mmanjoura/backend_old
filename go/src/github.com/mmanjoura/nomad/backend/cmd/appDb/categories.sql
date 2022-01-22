select * from  category;
select * from  [category_type];
select * from  setting;
select * from  [category_image];

update category set parent_id = 0 where id = 1


select * from  promotional_slider;
select * from  parent_child;

select * from user
delete from sqlite_sequence where name='category';
  
update category set shop_id = 1
delete from sqlite_sequence where name='category_image';
delete from sqlite_sequence where name='category_type';
delete from sqlite_sequence where name='setting';
delete from sqlite_sequence where name='promotional_slider';
delete from sqlite_sequence where name='parent_child';
delete from sqlite_sequence where name='category';

INSERT INTO category(category_type_id, parent_id, name,slug,icon,details,created_at,updated_at,deleted_at) VALUES (1, 1, 'Fruits & Vegetables','fruits-vegetables','FruitsVegetable',NULL,'2021-03-08T07:21:31.000000Z','2021-03-08T07:21:31.000000Z','2021-03-08T07:21:31.000000Z');
INSERT INTO category(category_type_id, parent_id, name,slug,icon,details,created_at,updated_at,deleted_at) VALUES (1, 1, 'Fruits','fruits-vegetables',null,NULL,'2021-03-08T07:21:31.000000Z','2021-03-08T07:21:31.000000Z','2021-03-08T07:21:31.000000Z');
INSERT INTO category(category_type_id, parent_id, name,slug,icon,details,created_at,updated_at,deleted_at) VALUES (1, 1, 'Vegetables','vegetables',null, NULL,'2021-03-08T07:21:31.000000Z','2021-03-08T07:21:31.000000Z','2021-03-08T07:21:31.000000Z');

INSERT INTO category_type(category_id,name,slug,icon,created_at,updated_at) VALUES (1,'Grocery','grocery','FruitsVegetable','2021-03-08T07:18:25.000000Z','2021-09-26T15:23:32.000000Z');
INSERT INTO setting(shop_id,category_id,contact,website,isHome,layoutType,productCard,created_at,updated_at) VALUES (NULL,1,NULL,NULL,1,'classic','neon','2021-03-08T07:18:25.000000Z','2021-09-26T15:23:32.000000Z');
INSERT INTO setting(shop_id,category_id,contact,website,isHome,layoutType,productCard,created_at,updated_at) VALUES (NULL,1,NULL,NULL,1,'classic','neon','2021-03-08T07:18:25.000000Z','2021-09-26T15:23:32.000000Z');
INSERT INTO promotional_slider(category_type_id,original,thumbnail) VALUES (1,'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/offer-5.png','https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/902/conversions/offer-5-thumbnail.jpg');
INSERT INTO promotional_slider(category_type_id,original,thumbnail) VALUES (1,'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/offer-4.png','https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/903/conversions/offer-4-thumbnail.jpg');
INSERT INTO promotional_slider(category_type_id,original,thumbnail) VALUES (1,'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/offer-3.png','https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/904/conversions/offer-3-thumbnail.jpg');
INSERT INTO promotional_slider(category_type_id,original,thumbnail) VALUES (1,'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/offer-2.png','https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/905/conversions/offer-2-thumbnail.jpg');
INSERT INTO promotional_slider(category_type_id,original,thumbnail) VALUES (1,'https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/offer-1.png','https://pickbazarlaravel.s3.ap-southeast-1.amazonaws.com/906/conversions/offer-1-thumbnail.jpg');


SELECT category.id,
       category.name,
       category.slug,
       category.icon,
       category.category_type_id,
       category.created_at,
       category.updated_at,
       category.deleted_at,
       category.parent_Id
  FROM category INNER JOIN shop on shop.id = category.shop_id
  where category.parent_Id = 0