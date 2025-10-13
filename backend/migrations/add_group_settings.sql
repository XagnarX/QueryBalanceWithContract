-- Create wallet_group_settings table
-- This table stores per-group configuration like countdown, RPC selection, and token selection

CREATE TABLE IF NOT EXISTS wallet_group_settings (
    id SERIAL PRIMARY KEY,
    group_id INTEGER NOT NULL,
    user_id INTEGER NOT NULL,

    -- Countdown configuration
    countdown_enabled BOOLEAN DEFAULT false,
    countdown_duration INTEGER DEFAULT 600 CHECK (countdown_duration >= 10 AND countdown_duration <= 7200),

    -- RPC configuration
    selected_rpc_id INTEGER,

    -- Token configuration (stored as JSONB array of token IDs)
    selected_token_ids JSONB DEFAULT '[]'::jsonb,

    -- Timestamps
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    -- Foreign key constraints
    CONSTRAINT fk_group_settings_group FOREIGN KEY (group_id) REFERENCES wallet_groups(id) ON DELETE CASCADE,
    CONSTRAINT fk_group_settings_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT fk_group_settings_rpc FOREIGN KEY (selected_rpc_id) REFERENCES user_rpc_endpoints(id) ON DELETE SET NULL,

    -- Unique constraint: one setting per group per user
    CONSTRAINT uq_group_settings_user_group UNIQUE (user_id, group_id)
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_group_settings_user_id ON wallet_group_settings(user_id);
CREATE INDEX IF NOT EXISTS idx_group_settings_group_id ON wallet_group_settings(group_id);
CREATE INDEX IF NOT EXISTS idx_group_settings_user_group ON wallet_group_settings(user_id, group_id);

-- Add comment for documentation
COMMENT ON TABLE wallet_group_settings IS 'Stores user-specific configuration for each wallet group including countdown, RPC, and token selections';
COMMENT ON COLUMN wallet_group_settings.countdown_enabled IS 'Whether countdown timer is enabled for this group';
COMMENT ON COLUMN wallet_group_settings.countdown_duration IS 'Countdown duration in seconds (10-7200)';
COMMENT ON COLUMN wallet_group_settings.selected_rpc_id IS 'User-selected RPC endpoint for this group (NULL = use default)';
COMMENT ON COLUMN wallet_group_settings.selected_token_ids IS 'JSON array of selected token IDs for balance queries';
