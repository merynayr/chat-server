-- +goose Up
-- +goose StatementBegin
CREATE TABLE chat (
    chat_id     SERIAL PRIMARY KEY,
    chat_name   VARCHAR(255)  NOT NULL,
    created_at  TIMESTAMP     NOT NULL DEFAULT NOW()
);

CREATE TABLE roster
(
    chat_id BIGINT REFERENCES chat (chat_id) ON DELETE CASCADE NOT NULL,
    user_id BIGINT                                             NOT NULL,
    UNIQUE (chat_id, user_id)
);
    
CREATE TABLE messages
(
    message_id      BIGSERIAL PRIMARY KEY,
    chat_id         BIGINT REFERENCES chat (chat_id) ON DELETE CASCADE NOT NULL,
    user_id         BIGINT                                         NOT NULL,
    content         VARCHAR(10000),
    created_at      TIMESTAMP                                      NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS messages;
DROP TABLE IF EXISTS roster;
DROP TABLE IF EXISTS chat;
-- +goose StatementEnd
