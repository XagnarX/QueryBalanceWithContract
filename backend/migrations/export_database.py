#!/usr/bin/env python3
"""
Export all wallet manager database data to JSON file
Use this to backup data from one machine and transfer to another
"""

import json
import sys
import psycopg2
from datetime import datetime
from decimal import Decimal

# Source database configuration
# Modify these settings to match your source database
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'wallet_manager',
    'user': 'postgres',
    'password': 'postgres'
}

def json_serial(obj):
    """JSON serializer for objects not serializable by default"""
    if isinstance(obj, datetime):
        return obj.isoformat()
    if isinstance(obj, Decimal):
        return float(obj)
    raise TypeError(f"Type {type(obj)} not serializable")


def export_table(cursor, table_name, order_by='id'):
    """Export a single table to dict"""
    print(f"  📦 Exporting {table_name}...")

    cursor.execute(f"SELECT * FROM {table_name} ORDER BY {order_by}")
    rows = cursor.fetchall()
    columns = [desc[0] for desc in cursor.description]

    data = []
    for row in rows:
        record = {}
        for i, col in enumerate(columns):
            value = row[i]
            # Handle JSONB fields
            if isinstance(value, str) and col.endswith('_ids'):
                try:
                    record[col] = json.loads(value)
                except:
                    record[col] = value
            else:
                record[col] = value
        data.append(record)

    print(f"     ✅ {len(data)} records")
    return data


def export_database():
    """Export entire database to JSON file"""
    print("\n" + "=" * 70)
    print("DATABASE EXPORT TOOL - Cross-Machine Migration")
    print("=" * 70)
    print(f"Database: {DB_CONFIG['database']}@{DB_CONFIG['host']}:{DB_CONFIG['port']}")
    print("=" * 70 + "\n")

    try:
        # Connect to database
        print("🔌 Connecting to database...")
        conn = psycopg2.connect(**DB_CONFIG)
        cursor = conn.cursor()
        print("✅ Connected successfully\n")

        print("📦 Exporting tables...\n")

        # Export all tables
        export_data = {
            'export_info': {
                'timestamp': datetime.now().isoformat(),
                'source_database': DB_CONFIG['database'],
                'source_host': DB_CONFIG['host'],
                'exporter_version': '1.0'
            },
            'tables': {}
        }

        # Core tables
        tables_to_export = [
            ('users', 'id'),
            ('chains', 'chain_id'),
            ('wallet_groups', 'id'),
            ('wallet_addresses', 'id'),
            ('wallet_group_settings', 'id'),
            ('user_rpc_endpoints', 'id'),
            ('user_tokens', 'id'),
            ('balance_contracts', 'id'),
        ]

        for table_name, order_col in tables_to_export:
            try:
                export_data['tables'][table_name] = export_table(cursor, table_name, order_col)
            except psycopg2.Error as e:
                print(f"     ⚠️  Warning: Could not export {table_name}: {e}")
                export_data['tables'][table_name] = []

        # Calculate statistics
        total_records = sum(len(data) for data in export_data['tables'].values())

        print(f"\n📊 Export Summary:")
        print(f"{'Table':<30} {'Records':<10}")
        print("-" * 40)
        for table, data in export_data['tables'].items():
            print(f"{table:<30} {len(data):<10}")
        print("-" * 40)
        print(f"{'TOTAL':<30} {total_records:<10}")

        # Save to file
        filename = f"wallet_manager_export_{datetime.now().strftime('%Y%m%d_%H%M%S')}.json"
        print(f"\n💾 Saving to file: {filename}")

        with open(filename, 'w', encoding='utf-8') as f:
            json.dump(export_data, f, indent=2, default=json_serial, ensure_ascii=False)

        print(f"✅ Export completed\n")

        print("=" * 70)
        print("🎉 DATABASE EXPORT SUCCESSFUL!")
        print("=" * 70)
        print(f"\n📄 Export file: {filename}")
        print(f"📊 Total records: {total_records}")
        print(f"🕐 Export time: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        print("\n📦 Next steps:")
        print(f"  1. Transfer {filename} to target machine")
        print(f"  2. Run: python3 import_database.py {filename}")
        print("\n")

        cursor.close()
        conn.close()

    except psycopg2.Error as e:
        print(f"\n❌ Database error: {e}")
        sys.exit(1)
    except Exception as e:
        print(f"\n❌ Export failed: {e}")
        sys.exit(1)


if __name__ == "__main__":
    print("\n💡 This will export ALL data from your database to a JSON file")
    print("   You can then transfer this file to another machine\n")

    response = input("Do you want to proceed? (yes/no): ").strip().lower()
    if response in ['yes', 'y']:
        export_database()
    else:
        print("Export cancelled.")
