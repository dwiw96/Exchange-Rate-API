BEGIN;
CREATE TABLE currencies(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_currencies_id PRIMARY KEY,
    currency_code VARCHAR(3) NOT NULL,
        CONSTRAINT uq_currencies_currencyCode UNIQUE(currency_code),
            CONSTRAINT ck_currencies_currencyCode_length CHECK (LENGTH(currency_code) = 3),
    currency_name TEXT NOT NULL,
        CONSTRAINT ck_currencies_currencyName_length CHECK (LENGTH(currency_name) > 0)
);

CREATE TABLE exchange_rates(
    id INT GENERATED ALWAYS AS IDENTITY
        CONSTRAINT pk_exchangeRates_id PRIMARY KEY,
    currency_code_from VARCHAR(3) NOT NULL,
        CONSTRAINT fk_exchangeRates_currencyCodeFrom FOREIGN KEY(currency_code_from) REFERENCES currencies(currency_code),
    currency_code_to VARCHAR(3) NOT NULL,
        CONSTRAINT ck_exchangeRates_currencyCodeTo_length CHECK (LENGTH(currency_code_to) > 0),
    buy DOUBLE PRECISION NOT NULL,
        CONSTRAINT ck_exchangeRates_buy_zero CHECK (buy > 0),
    sell DOUBLE PRECISION NOT NULL,
        CONSTRAINT ck_exchangeRates_sell_zero CHECK (sell > 0),
    validate_date DATE DEFAULT NOW() NOT NULL
);

CREATE UNIQUE INDEX uq_currencies_currencyName ON currencies(currency_name) WHERE NOT (currency_name='CHINA YUAN');

CREATE INDEX ix_currencies_currency_code ON currencies(currency_code);
CREATE INDEX ix_currencies_currency_name ON currencies(currency_name);
CREATE INDEX ix_exchangeRates_currency_code_from ON exchange_rates(currency_code_from);
CREATE INDEX ix_exchangeRates_currency_code_to ON exchange_rates(currency_code_to);
CREATE INDEX ix_exchangeRates_buy ON exchange_rates(buy);
CREATE INDEX ix_exchangeRates_sell ON exchange_rates(sell);
COMMIT;