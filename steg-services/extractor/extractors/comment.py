import io
import sys
from pathlib import Path
from typing import Optional
from PIL import Image
from PIL.PngImagePlugin import PngImageFile
from PIL.JpegImagePlugin import JpegImageFile
from PIL.TiffImagePlugin import TiffImageFile
from .extractor import Extractor

sys.path.append(str(Path(__file__).resolve().parent.parent.parent))
from pyutils import file_utils


class CommentExtractor(Extractor):
    """Extractor for data embedded into the comment section."""

    def extract(self, input_file: bytes) -> Optional[bytes]:
        try:
            file_type = file_utils.detect_file_type(input_file)
            img = Image.open(io.BytesIO(input_file))

            if file_type.lower() == 'jpeg':
                # JPEG comments are stored in the info dictionary
                if isinstance(img, JpegImageFile):
                    return img.info.get("comment", None)
            elif file_type.lower() == 'png':
                # PNG comments are stored in the info dictionary as tEXt chunks
                if isinstance(img, PngImageFile):
                    return img.info.get("Comment", None) or img.info.get("Description", None)
            elif file_type.lower() == 'tiff':
                # TIFF comments may be stored in Exif tags such as 270 (ImageDescription)
                if isinstance(img, TiffImageFile):
                    description = img.tag_v2.get(270, None)  # ImageDescription
                    if description:
                        return description

                    user_comment = img.tag_v2.get(37510, None)  # UserComment
                    if user_comment:
                        return user_comment
            else:
                raise ValueError(f"Unsupported file type: {file_type}")
        except Exception as e:
            print(f"Error extracting comment: {e}")
        return None
