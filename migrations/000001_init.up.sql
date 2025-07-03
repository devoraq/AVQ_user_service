CREATE TABLE IF NOT EXISTS
    users (
        id UUID PRIMARY KEY NOT NULL,
        nickname VARCHAR(50) NOT NULL,
        user_role VARCHAR(20) NOT NULL DEFAULT 'user',
        -- current_status VARCHAR(20) NOT NULL,
        system_status VARCHAR(15) NOT NULL DEFAULT 'active',
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        last_login_at TIMESTAMPTZ DEFAULT NOW(),
        last_activity_at TIMESTAMPTZ DEFAULT NOW(),
        password_changed_at TIMESTAMPTZ DEFAULT NOW(),
        deleted_at TIMESTAMPTZ,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS
    private_data (
        user_id UUID PRIMARY KEY REFERENCES users (id) ON DELETE RESTRICT,
        first_name VARCHAR(100) DEFAULT NULL,
        last_name VARCHAR(100) DEFAULT NULL,
        middle_name VARCHAR(100) DEFAULT NULL,
        date_of_birth TIMESTAMPTZ DEFAULT NULL,
        gender VARCHAR(20) DEFAULT NULL,
        deleted_at TIMESTAMPTZ DEFAULT NULL,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS
    account_data (
        user_id UUID PRIMARY KEY REFERENCES users (id) ON DELETE RESTRICT,
        avatar_url TEXT DEFAULT NULL,
        banner_url TEXT DEFAULT NULL,
        bio TEXT DEFAULT NULL,
        status VARCHAR(255) DEFAULT NULL,
        socials JSONB DEFAULT NULL,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
        deleted_at TIMESTAMPTZ
    );

CREATE TABLE IF NOT EXISTS
    contact_data (
        user_id UUID PRIMARY KEY REFERENCES users (id) ON DELETE RESTRICT,
        phone VARCHAR(20),
        country VARCHAR(100),
        city VARCHAR(100),
        street VARCHAR(150),
        building VARCHAR(50),
        apartment VARCHAR(50),
        postal_code VARCHAR(20),
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
        deleted_at TIMESTAMPTZ
    );

CREATE TABLE IF NOT EXISTS
    security_data (
        user_id UUID PRIMARY KEY REFERENCES users (id) ON DELETE RESTRICT,
        login VARCHAR(50) UNIQUE NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        lockout_until TIMESTAMPTZ,
        error_login_attempts SMALLINT NOT NULL DEFAULT 0,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
        deleted_at TIMESTAMPTZ
    );

CREATE TABLE IF NOT EXISTS
    user_settings (
        user_id UUID PRIMARY KEY REFERENCES users (id) ON DELETE RESTRICT,
        settings_language VARCHAR(2) NOT NULL,
        profile_visibility VARCHAR(20) NOT NULL,
        messages_permission VARCHAR(20) NOT NULL,
        email_notifications BOOLEAN NOT NULL DEFAULT TRUE,
        push_notifications BOOLEAN NOT NULL DEFAULT TRUE,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
        deleted_at TIMESTAMPTZ,
        two_factor_enabled BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS
    user_chat (
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
        chat_id UUID NOT NULL,
        pinned BOOLEAN NOT NULL DEFAULT FALSE,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
        deleted_at TIMESTAMPTZ,
        PRIMARY KEY (user_id, chat_id)
    );

CREATE TABLE IF NOT EXISTS
    user_notification (
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
        notification_id UUID NOT NULL,
        unread BOOLEAN NOT NULL DEFAULT TRUE,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
        deleted_at TIMESTAMPTZ,
        PRIMARY KEY (user_id, notification_id)
    );

CREATE TABLE IF NOT EXISTS
    user_post (
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
        post_id UUID NOT NULL,
        pinned BOOLEAN NOT NULL DEFAULT FALSE,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
        deleted_at TIMESTAMPTZ,
        PRIMARY KEY (user_id, post_id)
    );