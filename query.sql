select brands.name,
       outlets.name                                                               AS outlet_name,
       outlets.picture,
       outlets.address,
       outlets.latitude,
       outlets.longitude,
       (select count(products.id) FROM products WHERE products.brand_id = outlets.brand_id) AS total_product,
       SQRT(POW(111.12 * (-6.1753924::float -  outlets.latitude::float), 2) +
            POW(111.12 * (outlets.longitude::float - 106.824964::float) * COS(-6.1753924::float / 92.215), 2)
           )                                                                      as distance_outlet_from_monas
from outlets
         join brands on outlets.brand_id = brands.id;