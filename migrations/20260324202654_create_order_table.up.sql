CREATE TABLE orders (
    order_id   VARCHAR(36)    PRIMARY KEY,
    user_id    VARCHAR(36)    NOT NULL,
    items      JSON           NOT NULL,
    amount     DECIMAL(15, 2) NOT NULL,
    created_at TIMESTAMP      DEFAULT CURRENT_TIMESTAMP
);