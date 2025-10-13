#!/usr/bin/env python3
"""
Import database data from JSON file
Use this to restore data on a new machine
"""

import json
import sys
import psycopg2
from datetime import datetime

# Target database configuration
# Modify these settings to match your target database
DB_CONFIG = {
    'host': 'localhost',
    'port': 5432,
    'database': 'wallet_manager',
    'user': 'postgres',
    'password': 'password123'
}


def import_table(cursor, table_name, data, conn):
    """Import data into a single table"""
    if not data:
        print(f"  ⏭️  Skipping {table_name} (no data)")
        return 0

    print(f"  📥 Importing {table_name}...")

    # Get column names from first record
    columns = list(data[0].keys())
    placeholders = ', '.join(['%s'] * len(columns))
    columns_str = ', '.join(columns)

    # Clear existing data
    cursor.execute(f"TRUNCATE TABLE {table_name} RESTART IDENTITY CASCADE")

    # Insert records
    inserted = 0
    for record in data:
        # Handle JSONB fields
        values = []
        for col in columns:
            value = record[col]
            # Convert list/dict to JSON string for JSONB fields
            if isinstance(value, (list, dict)) and col.endswith('_ids'):
                value = json.dumps(value)
            values.append(value)

        cursor.execute(
            f"INSERT INTO {table_name} ({columns_str}) VALUES ({placeholders})",
            values
        )
        inserted += 1

    conn.commit()
    print(f"     ✅ {inserted} records imported")
    return inserted


def import_database(filename):
    """Import database from JSON file"""
    print("\n" + "=" * 70)
    print("DATABASE IMPORT TOOL - Cross-Machine Migration")
    print("=" * 70)
    print(f"Target DB: {DB_CONFIG['database']}@{DB_CONFIG['host']}:{DB_CONFIG['port']}")
    print("=" * 70 + "\n")

    # Load JSON file
    try:
        print(f"📖 Reading file: {filename}")
        with open(filename, 'r', encoding='utf-8') as f:
            import_data = json.load(f)
        print("✅ File loaded successfully\n")
    except FileNotFoundError:
        print(f"❌ File not found: {filename}")
        sys.exit(1)
    except json.JSONDecodeError as e:
        print(f"❌ Invalid JSON file: {e}")
        sys.exit(1)

    # Show export info
    export_info = import_data.get('export_info', {})
    print("📋 Export Information:")
    print(f"  Source: {export_info.get('source_database', 'unknown')}@{export_info.get('source_host', 'unknown')}")
    print(f"  Export time: {export_info.get('timestamp', 'unknown')}")
    print(f"  Version: {export_info.get('exporter_version', 'unknown')}")

    # Show tables summary
    tables = import_data.get('tables', {})
    total_records = sum(len(data) for data in tables.values())

    print(f"\n📊 Data Summary:")
    print(f"{'Table':<30} {'Records':<10}")
    print("-" * 40)
    for table, data in tables.items():
        print(f"{table:<30} {len(data):<10}")
    print("-" * 40)
    print(f"{'TOTAL':<30} {total_records:<10}")

    print("\n⚠️  WARNING: This will DELETE all existing data in the target database!")
    print("⚠️  All tables will be truncated and replaced with imported data\n")

    response = input("Type 'IMPORT' to confirm: ").strip()
    if response != 'IMPORT':
        print("Import cancelled.")
        sys.exit(0)

    try:
        # Connect to database
        print("\n🔌 Connecting to target database...")
        conn = psycopg2.connect(**DB_CONFIG)
        cursor = conn.cursor()
        print("✅ Connected successfully\n")

        print("📥 Importing tables...\n")

        # Import tables in correct order (respect foreign keys)
        import_order = [
            'users',
            'chains',
            'wallet_groups',
            'wallet_addresses',
            'balance_contracts',
            'user_rpc_endpoints',
            'user_tokens',
            'wallet_group_settings',
        ]

        imported_counts = {}
        for table_name in import_order:
            if table_name in tables:
                count = import_table(cursor, table_name, tables[table_name], conn)
                imported_counts[table_name] = count

        # Reset sequences
        print("\n🔄 Resetting ID sequences...")
        for table_name in imported_counts.keys():
            try:
                cursor.execute(f"""
                    SELECT setval(pg_get_serial_sequence('{table_name}', 'id'),
                    (SELECT MAX(id) FROM {table_name}))
                """)
            except:
                # Some tables might not have id column
                pass
        conn.commit()
        print("✅ Sequences reset\n")

        # Verify import
        print("🔍 Verifying import...")
        total_imported = sum(imported_counts.values())

        print(f"\n📊 Import Summary:")
        print(f"{'Table':<30} {'Imported':<10}")
        print("-" * 40)
        for table, count in imported_counts.items():
            print(f"{table:<30} {count:<10}")
        print("-" * 40)
        print(f"{'TOTAL':<30} {total_imported:<10}")

        print("\n" + "=" * 70)
        print("🎉 DATABASE IMPORT SUCCESSFUL!")
        print("=" * 70)
        print(f"\n✅ Imported {total_imported} records")
        print(f"📄 Source file: {filename}")
        print(f"🕐 Import time: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
        print("\n💡 Next steps:")
        print("  1. Verify data in target database")
        print("  2. Update backend database connection if needed")
        print("  3. Restart backend service")
        print("  4. Test application functionality\n")

        cursor.close()
        conn.close()

    except psycopg2.Error as e:
        print(f"\n❌ Database error: {e}")
        if conn:
            conn.rollback()
            conn.close()
        sys.exit(1)
    except Exception as e:
        print(f"\n❌ Import failed: {e}")
        if conn:
            conn.rollback()
            conn.close()
        sys.exit(1)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        print("\n❌ Usage: python3 import_database.py <json_file>\n")
        print("Example:")
        print("  python3 import_database.py wallet_manager_export_20250113_143022.json\n")
        sys.exit(1)

    filename = sys.argv[1]
    import_database(filename)
