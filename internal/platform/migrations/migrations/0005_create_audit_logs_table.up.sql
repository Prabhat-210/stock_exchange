CREATE TABLE IF NOT EXISTS auth.audit_logs (
    id UUID PRIMARY KEY,
    user_id UUID,
    event_type VARCHAR(100) NOT NULL,
    event_status VARCHAR(50) NOT NULL,
    ip_address INET,
    user_agent TEXT,
    metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_audit_logs_user
        FOREIGN KEY (user_id)
        REFERENCES auth.users(id)
        ON DELETE SET NULL
);