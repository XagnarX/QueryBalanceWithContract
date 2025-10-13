#!/usr/bin/env python3
"""
Database Export Script for Cross-Machine Migration
Exports all data from the source database to a JSON file
"""

import psycopg2
import json
from datetime import datetime
import sys

# Source database configuration
# IMPORTANT: Update these settings to match your SOURCE database
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'wallet_manager',
    'user': 'postgres',
    'password': 'postgres'
}

# Tables to export (in order, respecting foreign key dependencies)
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

def export_table(cursor, table_name, order_by='id'):
    """Export a single table to dict"""
    print(f"  Exporting {table_name}...", end='', flush=True)

    cursor.execute(f"SELECT * FROM {table_name} ORDER BY {order_by}")
    rows = cursor.fetchall()
    columns = [desc[0] for desc in cursor.description]

    # Convert rows to list of dicts
    data = []
    for row in rows:
        row_dict = {}
        for i, col in enumerate(columns):
            value = row[i]
            # Handle datetime objects
            if hasattr(value, 'isoformat'):
                value = value.isoformat()
            row_dict[col] = value
        data.append(row_dict)

    print(f" {len(data)} records")
    return data

def main():
    print("=" * 60)
    print("Database Export Tool - Source Machine")
    print("=" * 60)
    print()

    # Show configuration
    print("Source Database Configuration:")
    print(f"  Host: {DB_CONFIG['host']}")
    print(f"  Port: {DB_CONFIG['port']}")
    print(f"  Database: {DB_CONFIG['database']}")
    print(f"  User: {DB_CONFIG['user']}")
    print()

    # Connect to source database
    try:
        print("Connecting to source database...")
        conn = psycopg2.connect(**DB_CONFIG)
        cursor = conn.cursor()
        print("✓ Connected successfully")
        print()
    except Exception as e:
        print(f"✗ Failed to connect to database: {e}")
        sys.exit(1)

    # Export all tables
    export_data = {}
    total_records = 0

    try:
        print("Exporting tables:")
        for table_name in TABLES:
            data = export_table(cursor, table_name)
            export_data[table_name] = data
            total_records += len(data)

        print()
        print(f"Total records exported: {total_records}")
        print()

        # Generate output filename with timestamp
        timestamp = datetime.now().strftime('%Y%m%d_%H%M%S')
        output_file = f"database_export_{timestamp}.json"

        # Write to JSON file
        print(f"Writing to {output_file}...")
        with open(output_file, 'w', encoding='utf-8') as f:
            json.dump(export_data, f, indent=2, ensure_ascii=False)

        print(f"✓ Export completed successfully!")
        print()
        print("=" * 60)
        print("Next Steps:")
        print("=" * 60)
        print(f"1. Transfer {output_file} to the target machine")
        print("2. On the target machine, run:")
        print(f"   python3 import_database.py {output_file}")
        print()

    except Exception as e:
        print(f"✗ Export failed: {e}")
        sys.exit(1)
    finally:
        cursor.close()
        conn.close()

if __name__ == '__main__':
    main()
