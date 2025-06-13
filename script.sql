CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    date_of_birth DATE NOT NULL,
    place_of_birth VARCHAR(100),
    address TEXT
);

CREATE TABLE accounts (
    account_id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID NOT NULL,
    bic VARCHAR(11) NOT NULL,
    iban VARCHAR(34) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE cards (
    card_id UUID PRIMARY KEY DEFAULT gen_random_uuid (),
    user_id UUID NOT NULL,
    account_id UUID NOT NULL,
    card_number VARCHAR(19) NOT NULL,
    expiration DATE NOT NULL,
    cvv VARCHAR(4) NOT NULL,
    "limit" int not null,
    FOREIGN KEY (user_id) REFERENCES users (user_id),
    FOREIGN KEY (account_id) REFERENCES accounts (account_id)
);

INSERT INTO
    users (
        first_name,
        last_name,
        date_of_birth,
        place_of_birth,
        address
    )
values (
        'John',
        'Doe',
        '1990-01-01',
        'New York',
        '123 Main St, New York, NY 10001'
    ),
    (
        'Jane',
        'Smith',
        '1985-05-15',
        'Los Angeles',
        '456 Elm St, Los Angeles, CA 90001'
    ),
    (
        'Alice',
        'Johnson',
        '1992-07-20',
        'Chicago',
        '789 Oak St, Chicago, IL 60601'
    ),
    (
        'Bob',
        'Brown',
        '1988-03-30',
        'Houston',
        '321 Pine St, Houston, TX 77001'
    );

INSERT INTO accounts (
    user_id,
    bic,
    iban
)
VALUES
    (
        (SELECT user_id FROM users WHERE first_name = 'John' AND last_name = 'Doe'),
        'BNPPDEFF',
        'DE89370400440532013000'
    ),
    (
        (SELECT user_id FROM users WHERE first_name = 'Jane' AND last_name = 'Smith'),
        'BNPPDEFF',
        'DE89370400440532013001'
    ),
    (
        (SELECT user_id FROM users WHERE first_name = 'Alice' AND last_name = 'Johnson'),
        'BNPPDEFF',
        'DE89370400440532013002'
    ),
    (
        (SELECT user_id FROM users WHERE first_name = 'Bob' AND last_name = 'Brown'),
        'BNPPDEFF',
        'DE89370400440532013003'
    );


INSERT INTO cards (
    user_id,
    account_id,
    card_number,
    expiration,
    cvv,
    "limit"
)
VALUES
    (
        (SELECT user_id FROM users WHERE first_name = 'John' AND last_name = 'Doe'),
        (SELECT account_id FROM accounts WHERE user_id = (SELECT user_id FROM users WHERE first_name = 'John' AND last_name = 'Doe')),
        '1234567890123456',
        '2025-12-31',
        '123',
        5000
    ),
    (
        (SELECT user_id FROM users WHERE first_name = 'Jane' AND last_name = 'Smith'),
        (SELECT account_id FROM accounts WHERE user_id = (SELECT user_id FROM users WHERE first_name = 'Jane' AND last_name = 'Smith')),
        '2345678901234567',
        '2026-11-30',
        '456',
        6000
    ),
    (
        (SELECT user_id FROM users WHERE first_name = 'Alice' AND last_name = 'Johnson'),
        (SELECT account_id FROM accounts WHERE user_id = (SELECT user_id FROM users WHERE first_name = 'Alice' AND last_name = 'Johnson')),
        '3456789012345678',
        '2024-10-31',
        '789',
        7000
    ),
    (
        (SELECT user_id FROM users WHERE first_name = 'Bob' AND last_name = 'Brown'),
        (SELECT account_id FROM accounts WHERE user_id = (SELECT user_id FROM users WHERE first_name = 'Bob' AND last_name = 'Brown')),
        '4567890123456789',
        '2023-09-30',
        '012',
        8000
    );

