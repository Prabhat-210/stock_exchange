CREATE TABLE IF NOT EXISTS auth.refresh_tokens (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    session_id UUID NOT NULL,
    token_hash TEXT NOT NULL UNIQUE,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_refresh_tokens_user
        FOREIGN KEY (user_id)
        REFERENCES auth.users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_refresh_tokens_session
        FOREIGN KEY (session_id)
        REFERENCES auth.sessions(id)
        ON DELETE CASCADE
);