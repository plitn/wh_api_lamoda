CREATE TABLE "products"(
                           "product_id" BIGINT NOT NULL UNIQUE,
                           "width" DOUBLE PRECISION NOT NULL,
                           "height" DOUBLE PRECISION NOT NULL,
                           "depth" DOUBLE PRECISION NOT NULL,
                           "volume" DOUBLE PRECISION NOT NULL,
                           "quantity_total" BIGINT NOT NULL
);
ALTER TABLE
    "products" ADD PRIMARY KEY("product_id");
CREATE TABLE "warehouses"(
                             "wh_id" BIGINT NOT NULL UNIQUE,
                             "wh_name" VARCHAR(255) NOT NULL,
                             "is_active" BOOLEAN NOT NULL,
                             "is_active_dt" TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL
);
ALTER TABLE
    "warehouses" ADD PRIMARY KEY("wh_id");
CREATE TABLE "wh_product"(
                             "id" SERIAL NOT NULL UNIQUE,
                             "wh_id" BIGINT NOT NULL,
                             "product_id" BIGINT NOT NULL,
                             "product_qty" BIGINT NOT NULL,
                             "reserved_qty" BIGINT NOT NULL
);
ALTER TABLE
    "wh_product" ADD PRIMARY KEY("id");
ALTER TABLE
    "wh_product" ADD CONSTRAINT "wh_product_product_id_foreign" FOREIGN KEY("product_id") REFERENCES "products"("product_id");
ALTER TABLE
    "wh_product" ADD CONSTRAINT "wh_product_wh_id_foreign" FOREIGN KEY("wh_id") REFERENCES "warehouses"("wh_id");

INSERT INTO products (product_id, width, height, depth, volume, quantity_total) VALUES
                                                                                    (35650, 4.88, 54.85, 56.61, 15152.69, 613), --+
                                                                                    (39741, 22.26, 86.55, 71.23, 137231.93, 823), --+
                                                                                    (93098, 60.55, 24.53, 82.74, 122893.02, 140), --+
                                                                                    (79753, 92.47, 58.21, 81.91, 440895.21, 102), --+
                                                                                   -- (11891, 76.36, 64.26, 57.9, 284109.14, 911),
                                                                                    (84979, 47.13, 45.95, 44.9, 97236.5, 285); -- +
                                                                                  --  (2385, 87.51, 43.06, 46.2, 174089.94, 801),
                                                                                  --  (21931, 11.91, 40.03, 72.97, 34788.98, 962),
                                                                                  --  (67478, 30.76, 13.72, 45.43, 19172.7, 753),
                                                                                    --(9335, 5.84, 6.24, 88.21, 3214.51, 488);
INSERT INTO warehouses (wh_id, wh_name, is_active, is_active_dt)
VALUES
    (751, 'Hubbard, Brown and Smith', True, '2024-02-01 23:18:35'),
    (789, 'Burton, Wilson and Rush', False, '2023-06-22 13:00:10'),
    (871, 'Blevins Inc', True, '2023-06-08 08:53:37'),
    (661, 'Dunn, Robbins and Hobbs', False, '2023-08-06 02:41:15'),
    (112, 'Davis-Harris', True, '2023-05-09 07:10:27');

INSERT INTO wh_product (wh_id, product_id, product_qty, reserved_qty) VALUES
                                                                          (751, 84979, 150, 85),
                                                                          (661, 84979, 10, 40),
                                                                          (112, 35650, 300, 0),
                                                                          (871, 35650, 0, 313),
                                                                          (789, 39741, 823, 0),
                                                                          (871, 93098, 70, 70),
                                                                          (112, 79753, 64, 18),
                                                                          (661, 79753, 15, 5);


