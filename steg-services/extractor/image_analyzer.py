import io
import re

from enum import Enum
from PIL import Image, ExifTags


class ImageType(Enum):
    PNG = 1
    JPG = 2
    BMP = 3
    GIF = 4
    TIF = 5
    SVG = 6
    ICO = 7
    WEBP = 8
    UNKNOWN = 99


class ImageAnalyzer:
    def get_image_type(self, data: bytes) -> ImageType:
        # PNG: 89 50 4E 47
        if len(data) >= 4 and data[:4] == b'\x89PNG':
            return ImageType.PNG

        # JPG: FF D8
        elif len(data) >= 2 and data[:2] == b'\xFF\xD8':
            return ImageType.JPG

        # BMP: 42 4D
        elif len(data) >= 2 and data[:2] == b'\x42\x4D':
            return ImageType.BMP

        # GIF: 47 49 46 38
        elif len(data) >= 4 and data[:4] == b'GIF8':
            return ImageType.GIF

        # TIF: 49 49 2A 00 or 4D 4D 00 2A
        elif len(data) >= 4 and (
            data[:4] == b'\x49\x49\x2A\x00' or data[:4] == b'\x4D\x4D\x00\x2A'
        ):
            return ImageType.TIF

        # SVG: XML header < ? x m l
        elif len(data) >= 5 and data[:5] == b'\x3C\x3F\x78\x6D\x6C':
            return ImageType.SVG

        # ICO: 00 00 01 00
        elif len(data) >= 4 and data[:4] == b'\x00\x00\x01\x00':
            return ImageType.ICO

        # WEBP: 52 49 46 46 with WEBP marker
        elif len(data) >= 12 and data[:4] == b'RIFF' and data[8:12] == b'WEBP':
            return ImageType.WEBP

        # Default: Unknown type
        else:
            return ImageType.UNKNOWN

    def extract_exif(self, data: bytes) -> dict:
        """Extracts EXIF metadata from an image file."""
        try:
            with Image.open(io.BytesIO(data)) as img:
                exif_data = img._getexif()
                if not exif_data:
                    return {}

                # Convert EXIF tag IDs to human-readable names
                exif = {
                    ExifTags.TAGS.get(tag_id, tag_id): value
                    for tag_id, value in exif_data.items()
                    if tag_id in ExifTags.TAGS
                }
                return exif

        except Exception as e:
            raise Exception(f"Error extracting EXIF data: {e}")

    def extract_strings(self, data: bytes, min_length: int = 4) -> list[str]:
        """
        Extracts human-readable strings from binary data, similar to the Linux 'strings' command.
        :param data: Binary data to search for strings.
        :param min_length: Minimum length of strings to include.
        :return: List of extracted strings.
        """
        # Regex pattern for printable ASCII strings
        pattern = rb'[ -~]{' + str(min_length).encode() + rb',}'
        strings = re.findall(pattern, data)
        return [s.decode('utf-8', errors='ignore') for s in strings]
    
    def compare_exif(self, image_a: bytes, image_b: bytes):
        """
        Compares the EXIF data of two images and returns a dictionary with the values.

        Args:
            image_a (bytes): The first image (JPG or PNG).
            image_b (bytes): The second image (JPG or PNG).

        Returns:
            dict: A dictionary with the EXIF field name as the key and the different values in the form “ExifData_a:ExifData_b” as the value.
        """

        exif_a = self.extract_exif(image_a)
        exif_b = self.extract_exif(image_b)

        differences = {}
        all_keys = set(exif_a.keys()).union(set(exif_b.keys()))

        for key in all_keys:
            value_a = exif_a.get(key, None)
            value_b = exif_b.get(key, None)
            if value_a != value_b:
                differences[key] = f"{value_a}:{value_b}"

        return differences