CREATE TABLE IF NOT EXISTS
    USER(
        id UUID PRIMARY KEY,
        user_role VARCHAR(20) NOT NULL,
        current_status VARCHAR(20) NOT NULL,
        system_status VARCHAR(15) NOT NULL,
        created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
        last_login_at TIMESTAMPTZ,
        last_activity_at TIMESTAMPTZ,
        password_changed_at TIMESTAMPTZ,
        deleted_at TIMESTAMPTZ,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS
    private_data (
        id UUID PRIMARY KEY,
        user_id UUID NOT NULL REFERENCES users (id) ON DELETE RESTRICT,
        first_name VARCHAR(100) NOT NULL,
        last_name VARCHAR(100) NOT NULL,
        middle_name VARCHAR(100),
        date_of_birth DATE,
        gender VARCHAR(20),
        deleted_at TIMESTAMPTZ,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE
    );

CREATE TABLE IF NOT EXISTS
    account_data (
        user_id UUID PRIMARY KEY REFERENCES users (id) ON DELETE RESTRICT,
        nick_name VARCHAR(50),
        avatar_url TEXT,
        banner_url TEXT,
        bio TEXT,
        status VARCHAR(255),
        socials JSONB,
        is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
        deleted_at TIMESTAMPTZ
    );

CREATE TABLE IF NOT EXISTS
    contact_data (
        user_id UUID PRIMARY KEY REFERENCES users (id) ON DELETE RESTRICT,
        phone VARCHAR(20),
        email VARCHAR(255),
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