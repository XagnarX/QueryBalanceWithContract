#!/usr/bin/env python3
"""
One-Click Database Import Script
Automatically fixes constraints, clears tables, and imports data
"""

import psycopg2
import json
import sys
from datetime import datetime

# Target database configuration
# IMPORTANT: Update these settings to match your TARGET database
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'wallet_manager',
    'user': 'postgres',
    'password': 'postgres'
}

# Tables to import (in order, respecting foreign key dependencies)
TABLES = [
    'users',
    'chains',
    'user_rpc_endpoints',      # Must be before wallet_group_settings
    'user_tokens',              # Must be before wallet_group_settings
    'wallet_groups',
    'wallet_addresses',
    'wallet_group_settings',    # References user_rpc_endpoints and user_tokens
    'balance_contracts'
]

def fix_constraints(cursor, conn):
    """Fix database constraints before import"""
    print("Step 1: Fixing database constraints...")

    # Fix countdown_duration constraint
    cursor.execute("""
        ALTER TABLE wallet_group_settings
        DROP CONSTRAINT IF EXISTS wallet_group_settings_countdown_duration_check;
    """)

    cursor.execute("""
        ALTER TABLE wallet_group_settings
        ADD CONSTRAINT wallet_group_settings_countdown_duration_check
        CHECK (countdown_duration >= 1 AND countdown_duration <= 7200);
    """)

    conn.commit()
    print("  ✓ Constraints fixed (countdown_duration: min=1, max=7200)")
    print()

def truncate_all_tables(cursor, conn):
    """Clear all tables"""
    print("Step 2: Clearing all existing data...")

    # Disable foreign key checks temporarily
    cursor.execute("SET session_replication_role = 'replica';")

    # Truncate all tables in reverse order
    for table_name in reversed(TABLES):
        cursor.execute(f"TRUNCATE TABLE {table_name} RESTART IDENTITY CASCADE")
        print(f"  ✓ Cleared {table_name}")

    # Re-enable foreign key checks
    cursor.execute("SET session_replication_role = 'origin';")

    conn.commit()
    print("  ✓ All tables cleared")
    print()

def import_table(cursor, table_name, data, conn):
    """Import data into a single table"""
    if not data:
        print(f"  {table_name}: No data to import")
        return 0

    print(f"  Importing {table_name}...", end='', flush=True)

    # Get column names from first record
    columns = list(data[0].keys())
    columns_str = ', '.join(columns)
    placeholders = ', '.join(['%s'] * len(columns))

    # Insert all records
    inserted = 0
    for record in data:
        values = []
        for col in columns:
            value = record[col]
            # Convert list to JSON string for JSONB columns
            if isinstance(value, list) and table_name == 'wallet_group_settings' and col == 'selected_token_ids':
                value = json.dumps(value)
            values.append(value)
        try:
            cursor.execute(
                f"INSERT INTO {table_name} ({columns_str}) VALUES ({placeholders})",
                values
            )
            inserted += 1
        except Exception as e:
            print(f"\n  ✗ Error inserting record: {e}")
            print(f"    Record: {record}")
            raise

    # Reset sequence for id column if it exists
    if 'id' in columns:
        cursor.execute(f"""
            SELECT setval(
                pg_get_serial_sequence('{table_name}', 'id'),
                COALESCE((SELECT MAX(id) FROM {table_name}), 1)
            )
        """)

    conn.commit()
    print(f" {inserted} records")
    return inserted

def import_data(cursor, export_data, conn):
    """Import all data from export file"""
    print("Step 3: Importing data...")

    imported_total = 0
    for table_name in TABLES:
        data = export_data.get(table_name, [])
        count = import_table(cursor, table_name, data, conn)
        imported_total += count

    print(f"  ✓ Total records imported: {imported_total}")
    print()
    return imported_total

def main():
    if len(sys.argv) < 2:
        print("Usage: python3 full_import.py <export_file.json>")
        print("Example: python3 full_import.py database_export_20241013_123456.json")
        print()
        print("This script will:")
        print("  1. Fix database constraints")
        print("  2. Clear all existing data")
        print("  3. Import data from the export file")
        sys.exit(1)

    export_file = sys.argv[1]

    print("=" * 70)
    print("ONE-CLICK DATABASE IMPORT")
    print("=" * 70)
    print()

    # Read export file
    try:
        print(f"Reading export file: {export_file}")
        with open(export_file, 'r', encoding='utf-8') as f:
            export_data = json.load(f)
        print("✓ Export file loaded successfully")
        print()
    except FileNotFoundError:
        print(f"✗ Error: File not found: {export_file}")
        sys.exit(1)
    except json.JSONDecodeError as e:
        print(f"✗ Error: Invalid JSON file: {e}")
        sys.exit(1)

    # Show data summary
    print("Data Summary:")
    total_records = 0
    for table_name in TABLES:
        count = len(export_data.get(table_name, []))
        print(f"  {table_name}: {count} records")
        total_records += count
    print(f"  Total: {total_records} records")
    print()

    # Show target database configuration
    print("Target Database:")
    print(f"  Host: {DB_CONFIG['host']}")
    print(f"  Port: {DB_CONFIG['port']}")
    print(f"  Database: {DB_CONFIG['database']}")
    print(f"  User: {DB_CONFIG['user']}")
    print()

    # WARNING
    print("⚠️  WARNING: This will DELETE ALL EXISTING DATA!")
    print()
    print("This script will:")
    print("  1. Fix database constraints (countdown_duration: 1-7200 seconds)")
    print("  2. TRUNCATE all tables (delete all existing data)")
    print("  3. Import data from the export file")
    print()
    confirmation = input("Type 'IMPORT' to proceed: ")
    if confirmation != 'IMPORT':
        print("Import cancelled.")
        sys.exit(0)
    print()

    # Connect to target database
    try:
        print("Connecting to database...")
        conn = psycopg2.connect(**DB_CONFIG)
        cursor = conn.cursor()
        print("✓ Connected successfully")
        print()
    except Exception as e:
        print(f"✗ Failed to connect to database: {e}")
        sys.exit(1)

    # Execute all steps
    try:
        # Step 1: Fix constraints
        fix_constraints(cursor, conn)

        # Step 2: Clear all tables
        truncate_all_tables(cursor, conn)

        # Step 3: Import data
        imported_total = import_data(cursor, export_data, conn)

        print("=" * 70)
        print("✓ IMPORT COMPLETED SUCCESSFULLY!")
        print("=" * 70)
        print(f"Total records imported: {imported_total}")
        print()
        print("Next steps:")
        print("  1. Restart your application")
        print("  2. Verify all data is working correctly")
        print("  3. Delete the export file for security")
        print()

    except Exception as e:
        print()
        print("=" * 70)
        print("✗ IMPORT FAILED")
        print("=" * 70)
        print(f"Error: {e}")
        print()
        print("Rolling back changes...")
        conn.rollback()
        print("Database state restored to before the import attempt.")
        sys.exit(1)
    finally:
        cursor.close()
        conn.close()

if __name__ == '__main__':
    main()
