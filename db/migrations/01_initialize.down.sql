DROP TABLE message_reactions;

COMMENT ON TABLE reactions IS NULL;

DROP TABLE reactions;

DROP TABLE stamps;

COMMENT ON COLUMN messages.type IS NULL;
COMMENT ON COLUMN messages.content IS NULL;

DROP TABLE messages;

DROP TABLE message_types;

COMMENT ON COLUMN room_users.last_read_at IS NULL;

DROP TABLE room_users;

DROP TABLE chat_rooms;

COMMENT ON COLUMN follows.display_name IS NULL;

DROP TABLE follows;

COMMENT ON COLUMN users.email IS NULL;
COMMENT ON COLUMN users.profile_picture_url IS NULL;

DROP TABLE users;
