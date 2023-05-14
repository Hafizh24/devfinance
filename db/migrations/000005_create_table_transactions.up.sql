CREATE TABLE IF NOT EXISTS transactions (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    type VARCHAR(10) NOT NULL,
    note text,
    amount int,
    category_id bigint not null,
    currency_id bigint not null ,
    user_id bigint not null,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    FOREIGN key (category_id) REFERENCES categories (id) ON DELETE CASCADE,
    FOREIGN key (currency_id) REFERENCES currencies (id) ON DELETE CASCADE,
    FOREIGN key (user_id) REFERENCES users (id) ON DELETE CASCADE
);