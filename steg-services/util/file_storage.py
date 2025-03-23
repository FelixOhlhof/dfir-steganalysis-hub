import threading
import os
import time
import sqlite3

import sys
from pathlib import Path
sys.path.append(str(Path(__file__).resolve().parent.parent))
from pyutils import file_utils

class FileStorage:
    def __init__(self, temp_dir=None, db_path="file_storage.db"):
        self.temp_dir = temp_dir or os.path.join(os.getcwd(), 'files')
        self.lock = threading.Lock()  # Mutex
        self.db_path = db_path

        # Initialize database and create the table
        self._initialize_database()

        # Clean up expired files
        self.clean_expired_files()

    def _initialize_database(self):
        """Initializes the SQLite database and creates the necessary table."""
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("""
                CREATE TABLE IF NOT EXISTS saved_files (
                    file_name TEXT PRIMARY KEY,
                    file_path TEXT NOT NULL,
                    file_type TEXT,
                    sha256 TEXT,
                    expires_at INTEGER
                )
            """)
            conn.commit()

    def save_file(self, file_data: bytes, file_name: str, life_span: int = 10) -> dict:
        """Saves a file with a lifespan and stores its metadata in the database."""
        expires_at = int(time.time() + (life_span * 60)) if life_span != -1 else -1
        file_path = os.path.join(self.temp_dir, file_name)
        file_type = file_utils.detect_file_type(file_data)
        sha256 = file_utils.get_sha_256(file_data)

        # Save the file
        with open(file_path, 'wb') as f:
            f.write(file_data)

        # Store file metadata in the database
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("""
                INSERT OR REPLACE INTO saved_files (file_name, file_path, expires_at, sha256, file_type)
                VALUES (?, ?, ?, ?, ?)
            """, (file_name, file_path, expires_at, sha256, file_type))
            conn.commit()

        return {
            "file_name": file_name,
            "file_count": self._get_file_count(),
            "expires_at": "never" if expires_at == -1 else time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(expires_at)),
            "file_type": file_type,
            "sha256" : sha256
        }

    def get_file(self, file_name: str, delete_after: bool = False) -> dict:
        """Retrieves the requested file and optionally deletes it."""
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("""
                SELECT file_path, file_type, expires_at, sha256 FROM saved_files WHERE file_name = ?
            """, (file_name,))
            result = cursor.fetchone()

        if not result:
            return None

        file_path, file_type, expires_at, sha256 = result
        with open(file_path, "rb") as f:
            file_data = f.read()

        if delete_after:
            self.delete_file(file_name)

        return {"file_name": file_name, "file_data": file_data, "file_type": file_type, "expires_at": "never" if expires_at == -1 else time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(expires_at)), "sha256": sha256}

    def get_last_file(self, delete_after: bool = False) -> dict:
        """Retrieves the most recently added file."""
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("""
                SELECT file_name, file_path, file_type, expires_at, sha256 FROM saved_files ORDER BY ROWID DESC LIMIT 1
            """)
            result = cursor.fetchone()

        if not result:
            return None

        file_name, file_path, file_type, expires_at, sha256 = result
        with open(file_path, "rb") as f:
            file_data = f.read()

        if delete_after:
            self.delete_file(file_name)

        return {"file_name": file_name, "file_data": file_data, "file_type": file_type, "expires_at": "never" if expires_at == -1 else time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(expires_at)), "sha256": sha256}

    def delete_file(self, file_name: str):
        """Deletes a file from the filesystem and removes its metadata from the database."""
        with self.lock:
            with sqlite3.connect(self.db_path) as conn:
                cursor = conn.cursor()
                cursor.execute("""
                    SELECT file_path FROM saved_files WHERE file_name = ?
                """, (file_name,))
                result = cursor.fetchone()

                if result:
                    file_path = result[0]
                    if os.path.exists(file_path):
                        os.remove(file_path)

                    cursor.execute("""
                        DELETE FROM saved_files WHERE file_name = ?
                    """, (file_name,))
                    conn.commit()

    def delete_all_files(self):
        """Removes all files."""
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("""SELECT file_name FROM saved_files""")
            expired_files = [row[0] for row in cursor.fetchall()]

        for file_name in expired_files:
            self.delete_file(file_name)


    def clean_expired_files(self):
        """Removes all expired files based on their expiration time."""
        current_time = int(time.time())
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("""
                SELECT file_name FROM saved_files WHERE expires_at != -1 AND expires_at < ?
            """, (current_time,))
            expired_files = [row[0] for row in cursor.fetchall()]

        for file_name in expired_files:
            self.delete_file(file_name)

        threading.Timer(30, self.clean_expired_files).start()

    def _get_file_count(self) -> int:
        """Returns the number of files currently stored."""
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("SELECT COUNT(*) FROM saved_files")
            return cursor.fetchone()[0]
        
    def get_stats(self, file_limit = 100) -> dict:
        """Returns statistics about the currently stored files."""
        stats = {}
        with sqlite3.connect(self.db_path) as conn:
            cursor = conn.cursor()
            cursor.execute("SELECT COUNT(*) FROM saved_files")
            stats['file_count'] = cursor.fetchone()[0]

            cursor.execute(f"SELECT file_name, expires_at, file_type, sha256 FROM saved_files LIMIT {file_limit}")
            files = cursor.fetchall()

            stats['files'] = [
                {
                    "file_name": file[0],
                    "expires_at": "never" if file[1] == -1 else time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(file[1])),
                    "file_type": file[2],
                    "sha256": file[3]
                }
                for file in files
            ]

        return stats
        