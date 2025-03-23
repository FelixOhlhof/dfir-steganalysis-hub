from PIL import Image
import io

class ImageConverter:
    @staticmethod
    def convert_to_jpg(file_data: bytes) -> bytes:
        """Converts a PNG image to JPG."""
        with Image.open(io.BytesIO(file_data)) as img:
            rgb_img = img.convert('RGB')
            output = io.BytesIO()
            rgb_img.save(output, format='JPEG')
            return output.getvalue()

    @staticmethod
    def convert_to_png(file_data: bytes) -> bytes:
        """Converts a JPG image to PNG."""
        with Image.open(io.BytesIO(file_data)) as img:
            output = io.BytesIO()
            img.save(output, format='PNG')
            return output.getvalue()
