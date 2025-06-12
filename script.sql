CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
                       user_id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       first_name     VARCHAR(100) NOT NULL,
                       last_name      VARCHAR(100) NOT NULL,
                       date_of_birth  DATE NOT NULL,
                       place_of_birth VARCHAR(100),
                       address        TEXT
);

CREATE TABLE accounts (
                          account_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                          user_id    UUID NOT NULL,
                          bic        VARCHAR(11) NOT NULL,
                          iban       VARCHAR(34) NOT NULL,
                          FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE cards (
                       card_id     UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                       user_id     UUID NOT NULL,
                       account_id  UUID NOT NULL,
                       card_number VARCHAR(19) NOT NULL,
                       expiration  DATE NOT NULL,
                       cvv         VARCHAR(4) NOT NULL,
                       "limit"     int not null,
                       FOREIGN KEY (user_id) REFERENCES users(user_id),
                       FOREIGN KEY (account_id) REFERENCES accounts(account_id)
);
