CREATE TABLE stocks (
    id SERIAL PRIMARY KEY,
    company_name VARCHAR(100) NOT NULL,
    company_symbol VARCHAR(10) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    last_div FLOAT NOT NULL
);
