import hashlib

file_signatures = {
        b"\x89PNG": "PNG",
        b"\xFF\xD8\xFF": "JPEG",
        b"%PDF": "PDF",
        b"PK\x03\x04": "ZIP",
        b"GIF87a": "GIF",
        b"GIF89a": "GIF",
        b"\x7FELF": "ELF Executable",
        b"BM": "BMP",
        b"\x1F\x8B": "GZIP",
        b"OggS": "OGG Audio",
        b"fLaC": "FLAC Audio",
        b"ID3": "MP3 Audio",
        b"\x00\x00\x01\x00": "ICO",
        b"\x00\x00\x02\x00": "CUR",
        b"ftyp": "MP4 Video",
        b"MThd": "MIDI Audio",
        b"WAVE": "WAV Audio",
        b"MZ": "PE Executable",
    }

def detect_file_type(file_data: bytes) -> str:
    """Detects the file type based on magic bytes."""
    for signature, file_type in file_signatures.items():
        if file_data.startswith(signature):
            return file_type

    return "Unknown"

def get_sha_256(data):
    return hashlib.sha256(data).hexdigest()