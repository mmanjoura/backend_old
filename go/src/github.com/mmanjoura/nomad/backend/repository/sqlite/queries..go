package sqlite

var GetAttributes = `
SELECT attribute.id,
       attribute.slug,
       attribute.name,
       attribute.Shop_id,
       attribute.created_at,
       attribute.updated_at,
       shop.id,
       shop.owner_id,
       shop.name,
       shop.slug,
       shop.description,
       shop.is_active,
       shop.created_at,
       shop.updated_at,
       cover_image.id,
       cover_image.original,
       cover_image.thumbnail,
       logo.id,
       logo.original,
       logo.thumbnail,
       address.zip,
       address.city,
       address.state,
       address.country,
       address.street_address,
       contactdetail.contact,
       contactdetail.website,
       location.lat,
       location.lng,
       location.city,
       location.state,
       location.country,
       location.formattedAddress
  FROM attribute
       INNER JOIN
       shop ON shop.id = attribute.Shop_id
       INNER JOIN
       cover_image ON cover_image.Shop_id = shop.id
       INNER JOIN
       logo ON logo.Shop_id = shop.id
       INNER JOIN
       address ON address.Shop_id = shop.id
       INNER JOIN
       contactdetail ON contactdetail.shop_id = shop.id
       INNER JOIN
       location ON location.Shop_id = shop.id
	WHERE `

var GetAttributeValues = `
SELECT value.id,
       attribute.id,
       value.value,
       value.meta,
       value.created_at,
       value.updated_at
  FROM value
       INNER JOIN
       attribute ON value.shop_id = attribute.Shop_id
 WHERE `

var GetSocials = `
SELECT
    url,
    icon
  FROM social
 WHERE `

var GetCategories = `
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
      WHERE `

var GetCategoryChildren = `
SELECT category.id,
       category.name,
       category.slug,
       category.icon,
       category.parent_id,
       category.category_type_id,
       category.created_at,
       category.updated_at,
       category.deleted_at
  FROM category
       INNER JOIN
       category_type ON category_type.id = category.category_type_id 
      WHERE `

var GetPromotionalSlider = `
SELECT promotional_slider.id,
       promotional_slider.original,
       promotional_slider.thumbnail
  FROM promotional_slider
 WHERE `

var GetImages = `
 SELECT category_image.original,
      category_image.thumbnail
FROM category_image
WHERE `

var GetBanners = `
SELECT id,
       category_type_id,
       title,
       description,
       created_at,
       updated_at
  FROM banner
 WHERE `

var GetCategoryType = `
SELECT category_type.id,
       category_type.name,
       category_type.slug,
       category_type.icon,
       category_type.created_at,
       category_type.updated_at
  FROM category_type
  WHERE `

var GetCategoryTypeSetting = `
  SELECT setting.isHome,
  setting.layoutType,
  setting.productCard
FROM setting
WHERE `

var GetCoupons = `
SELECT id,
       code,
       description,
       type,
       amount,
       active_from,
       expire_at,
       created_at,
       updated_at,
       deleted_at,
       is_valid
  FROM coupon
  WHERE `

var GetProducts = `
SELECT product.id,
       product.name,
       product.slug,
       product.description,
       product.category_type_id,
       product.price,
       product.shop_id,
       product.sale_price,
       product.sku,
       product.is_taxable,
       product.shipping_class_id,
       product.status,
       product.product_type,
       product.unit,
       product.height,
       product.width,
       product.length,
       product.deleted_at,
       product.created_at,
       product.updated_at,
       product.video,
       category_type.id,
       category_type.name,
       category_type.slug,
       category_type.icon,
       category_type.created_at,
       category_type.updated_at,
       setting.isHome,
       setting.layoutType,
       setting.productCard,
       shop.id,
       shop.owner_id,
       shop.name,
       shop.slug,
       shop.description,
       shop.is_active,
       shop.created_at,
       shop.updated_at,
       product_image.id,
       product_image.original,
       product_image.thumbnail,
       logo.id,
       logo.original,
       logo.thumbnail
  FROM product
       INNER JOIN
       product_image ON product_image.product_Id = product.id
       INNER JOIN
       category_type ON category_type.category_Id = product.category_type_id
       INNER JOIN
       setting ON setting.shop_id = product.shop_id
       INNER JOIN
       shop ON shop.id = product.shop_id
       INNER JOIN
       logo ON logo.shop_id = shop.id
       INNER JOIN
       address ON address.Shop_id = shop.id
       WHERE `

var GetOrders = `
select customer_order.id
    ,customer_order.tracking_number
    ,customer_order.customer_id
    ,customer_order.customer_contact
    ,customer_order.amount
    ,customer_order.sales_tax
    ,customer_order.paid_total
    ,customer_order.total
    ,customer_order.coupon_id
    ,customer_order.parent_id
    ,customer_order.shop_id
    ,customer_order.discount
    ,customer_order.payment_id
    ,customer_order.payment_gateway
    ,customer_order.logistics_provider
    ,customer_order.delivery_fee
    ,customer_order.delivery_time
    ,customer_order.deleted_at
    ,customer_order.created_at
    ,customer_order.updated_at
    ,shipping_address.zip
    ,shipping_address.city
    ,shipping_address.state
    ,shipping_address.country
    ,shipping_address.street_address
    ,billing_address.zip
    ,billing_address.city
    ,billing_address.state
    ,billing_address.country
    ,billing_address.street_address
    ,customer.id
    ,customer.name
    ,customer.email
    ,customer.email_verified_at
    ,customer.created_at
    ,customer.updated_at
    ,customer.is_active
    ,customer.shop_id
    ,order_status.id
    ,order_status.name
    ,order_status.serial
    ,order_status.color
    ,order_status.created_at
    ,order_status.updated_at
from customer_order 
    INNER JOIN shipping_address on shipping_address.customer_order_Id = customer_order.id
    INNER JOIN billing_address on billing_address.customer_order_Id = customer_order.id
    INNER JOIN customer on customer.customer_order_Id = customer_order.id
    INNER jOIN order_status on order_status.customer_order_Id = customer_order.id
       WHERE `

// var GetShopSettings = `
// SELECT setting.id,
//     setting.created_at,
//     setting.updated_at,
//     shop_option.currency,
//     shop_option.taxClass,
//     shop_option.siteTitle,
//     shop_option.siteSubtitle,
//     shop_option.shippingClass,
//     shop_option.minimumOrderAmount
// FROM setting
// INNER JOIN
//   shop_option ON shop_option.shop_id = setting.shop_id
//        WHERE `

var GetShippings = `
SELECT id,
       name,
       amount,
       is_global,
       type,
       created_at,
       updated_at
  FROM shipping
  WHERE `

var GetTaxes = `
SELECT id,
       country,
       state,
       zip,
       city,
       rate,
       name,
       is_global,
       priority,
       on_shipping,
       created_at,
       updated_at
  FROM tax
  WHERE `

var GetCategory_types = `
SELECT category_type.id,
  category_type.name,
  category_type.slug,
  category_type.icon,
  category_type.created_at,
  category_type.updated_at
  FROM category_type
       WHERE `

var GetBannerImge = `
  SELECT id,
       original,
       thumbnail
  FROM banner_image
WHERE `

var GetSettingOptions = `
SELECT currency,
       taxclass,
       sitetitle,
       sitesubtitle,
       shippingclass,
       minimumOrderAmount
  FROM setting_option
  WHEREE `
var GetSEOs = `
SELECT ogImage,
       ogTitle,
       metaTags,
       metaTitle,
       canonicalUrl,
       ogDescription,
       twitterHandle,
       metaDescription,
       twitterCardType
  FROM seo
  WHERE `

var GetLogo = `
  SELECT id,
       original,
       thumbnail
  FROM logo
  WHERE `

var GetLocation = `
  SELECT lat,
       lng,
       city,
       state,
       country,
       formattedAddress
  FROM location
  WHERE `

var GetDeliveryTime = `
  SELECT title,
       description
  FROM deliverytime
  WHERE `

var GetContactDetils = `
  SELECT contact,
       website
  FROM contactdetail
  WHERE `

var GetSettings = `
  SELECT setting.id,
       setting.created_at,
       setting.updated_at,
       setting_option.currency,
       setting_option.taxClass,
       setting_option.siteTitle,
       setting_option.siteSubtitle,
       setting_option.shippingClass,
       setting_option.minimumOrderAmount,
       seo.ogImage,
       seo.ogTitle,
       seo.metaTags,
       seo.metaTitle,
       seo.canonicalUrl,
       seo.ogDescription,
       seo.twitterCardType,
       seo.twitterCardType,
       logo.id,
       logo.original,
       logo.thumbnail,
       contactDetail.contact,
       contactDetail.website,
       location.lat,
       location.lng,
       location.state,
       location.country,
       location.formattedAddress
  FROM setting
       INNER JOIN
       setting_option ON setting_option.setting_id = setting.id
       INNER JOIN
       seo ON seo.setting_id = setting.id
       INNER JOIN
       logo ON logo.setting_id = setting.id
       INNER JOIN
       contactDetail ON contactDetail.setting_id = setting.id
       INNER JOIN
       location ON location.setting_id = setting.id
       WHERE `
