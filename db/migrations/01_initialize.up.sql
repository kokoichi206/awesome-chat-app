-- uuid_generate_v4 を使うための拡張機能。
-- uuid-ossp は PostgreSQL に標準で含まれている。
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    profile_info TEXT,
    profile_picture_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);
COMMENT ON COLUMN users.email IS 'Firebase Auth などで取得した email アドレス。';
COMMENT ON COLUMN users.profile_picture_url IS 'S3 などの外部ストレージに格納された写真の URL。レスポンスに含めるときは、URL のまま返す想定。';

CREATE TABLE follows (
    user_id uuid REFERENCES users(id),
    follow_id uuid REFERENCES users(id),
    display_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (user_id, follow_id)
);
COMMENT ON COLUMN follows.display_name IS 'フォローしたユーザーの表示名。';

CREATE TABLE room_types (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE chat_rooms (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    type uuid REFERENCES room_types(id),
    room_name VARCHAR(255) NOT NULL,
    created_by uuid REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ
);

CREATE TABLE room_users (
    room_id uuid REFERENCES chat_rooms(id),
    user_id uuid REFERENCES users(id),
    last_read_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    deleted_at TIMESTAMPTZ,
    PRIMARY KEY (room_id, user_id, created_at)
);
COMMENT ON COLUMN room_users.last_read_at IS 'ユーザーの該当ルームに対する既読確認用';

CREATE TABLE message_types (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE messages (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    room_id uuid REFERENCES chat_rooms(id),
    user_id uuid REFERENCES users(id),
    type uuid REFERENCES message_types(id),
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);
COMMENT ON COLUMN messages.type IS 'メッセージの種類。テキスト、画像、動画、音声など。';
COMMENT ON COLUMN messages.content IS '画像などのバイナリデータに関しては外部 URL のパスとする。';

CREATE TABLE stamps (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    image_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL
);
COMMENT ON TABLE stamps IS '利用可能なスタンプ情報。画像は外部 URL のパスとする。';

CREATE TABLE reactions (
    id uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);
COMMENT ON TABLE reactions IS 'メッセージへの反応。具体的な表示はフロントに依存。';

CREATE TABLE message_reactions (
    message_id uuid REFERENCES messages(id),
    user_id uuid REFERENCES users(id),
    reaction_id uuid REFERENCES reactions(id),
    created_at TIMESTAMPTZ NOT NULL,
    PRIMARY KEY (message_id, user_id)
);
