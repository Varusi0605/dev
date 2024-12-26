CREATE OR REPLACE FUNCTION getProducts_fn(userId uuid,brandName text,categoryName text)
RETURNS TABLE (product_id uuid,category_id uuid,brand_id uuid,user_id uuid,product_name text,price numeric,rating numeric) AS
$$
BEGIN
	
	IF brandName != '' AND categoryName != '' THEN
		RETURN QUERY	
		SELECT p.product_id,p.category_id,p.brand_id,p.user_id,p.product_name,p.price,p.rating FROM products AS p 
	    INNER JOIN categories USING(category_id) 
		INNER JOIN brands USING(Brand_id)
		WHERE category_name = categoryName AND brand_name = brandName AND p.user_id = userId;
	ELSEIF brandName = '' AND categoryName != '' THEN
		RETURN QUERY	
	    SELECT p.product_id,p.category_id,p.brand_id,p.user_id,p.product_name,p.price,p.rating FROM products AS p 
	    INNER JOIN categories USING(category_id) 
		INNER JOIN brands USING(Brand_id)
		WHERE category_name = categoryName AND p.user_id = userId ;
	ELSEIF brandName != '' AND categoryName = '' THEN
		RETURN QUERY	
	    SELECT p.product_id,p.category_id,p.brand_id,p.user_id,p.product_name,p.price,p.rating FROM products AS p 
	    INNER JOIN categories USING(category_id) 
		INNER JOIN brands USING(Brand_id)
		WHERE brand_Name = brandName AND p.user_id = userId ;
	ELSE 
		RETURN QUERY	
		SELECT p.product_id,p.category_id,p.brand_id,p.user_id,p.product_name,p.price,p.rating FROM products AS p 
	    INNER JOIN categories USING(category_id) 
		INNER JOIN brands USING(Brand_id)
		WHERE p.user_id = userId;
	END IF;	
END;
$$
LANGUAGE plpgsql;

SELECT * FROM getProducts_fn('1ba51d85-bde6-4f12-b153-92e9233a0492','','phone')

