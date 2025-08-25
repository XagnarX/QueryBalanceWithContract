-- Validation script for init_database.sql
-- This script helps verify that all necessary components are present

-- Check if all expected tables exist
SELECT 'Checking tables...' as status;

DO $$
DECLARE
    expected_tables text[] := ARRAY['users', 'chains', 'wallet_groups', 'wallet_addresses', 'balance_contracts', 'user_rpc_endpoints', 'user_tokens'];
    table_name text;
    table_exists boolean;
BEGIN
    FOREACH table_name IN ARRAY expected_tables
    LOOP
        SELECT EXISTS (
            SELECT FROM information_schema.tables 
            WHERE table_schema = 'public' 
            AND table_name = table_name
        ) INTO table_exists;
        
        IF table_exists THEN
            RAISE NOTICE 'Table % exists: OK', table_name;
        ELSE
            RAISE WARNING 'Table % missing: FAIL', table_name;
        END IF;
    END LOOP;
END $$;

-- Check if trigger function exists
SELECT 'Checking trigger function...' as status;
SELECT 
    CASE 
        WHEN EXISTS (
            SELECT 1 FROM information_schema.routines 
            WHERE routine_name = 'update_updated_at_column'
        ) 
        THEN 'Trigger function exists: OK'
        ELSE 'Trigger function missing: FAIL'
    END as function_check;

-- Check if initial data is inserted
SELECT 'Checking initial data...' as status;
SELECT 
    CASE 
        WHEN (SELECT COUNT(*) FROM chains WHERE chain_id IN (56, 97, 1, 137)) >= 4
        THEN 'Initial chain data exists: OK'
        ELSE 'Initial chain data missing: FAIL'
    END as chain_data_check;

SELECT 
    CASE 
        WHEN (SELECT COUNT(*) FROM balance_contracts WHERE chain_id IN (56, 97)) >= 2
        THEN 'Initial contract data exists: OK'
        ELSE 'Initial contract data missing: FAIL'
    END as contract_data_check;

-- Summary
SELECT 'Validation complete!' as status;