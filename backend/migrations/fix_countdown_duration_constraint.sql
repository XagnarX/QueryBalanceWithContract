-- Fix countdown_duration check constraint
-- Change minimum value from 10 to 1 second

-- Drop the old constraint
ALTER TABLE wallet_group_settings
DROP CONSTRAINT IF EXISTS wallet_group_settings_countdown_duration_check;

-- Add new constraint with minimum value of 1
ALTER TABLE wallet_group_settings
ADD CONSTRAINT wallet_group_settings_countdown_duration_check
CHECK (countdown_duration >= 1 AND countdown_duration <= 7200);

-- Update comment
COMMENT ON COLUMN wallet_group_settings.countdown_duration IS 'Countdown duration in seconds (1-7200)';
