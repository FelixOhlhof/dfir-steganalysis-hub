import sys
from pathlib import Path
from typing import Optional
from .extractor import Extractor

sys.path.append(str(Path(__file__).resolve().parent.parent.parent))
from pyutils import file_utils


class AppendToEoiExtractor(Extractor):
    """Extractor for data appended to the end-of-image marker."""

    def extract(self, input_file: bytes) -> Optional[bytes]:
        type = file_utils.detect_file_type(input_file)

        eof = None
        if type.lower() in ["jpeg", "jpg"]:
            eof = b"\xFF\xD9"
        elif type.lower() in ["gif"]:
            eof = b"\x00\x3B"
        elif type.lower() in ["png"]:
            eof = b"\x49\x45\x4E\x44\xAE\x42\x60\x82"
        else:
            raise Exception(f"Extension {type} not supported!")

        # Looking for the EOI marker
        eof_data = input_file.rsplit(eof, 1)

        if len(eof_data) > 1:
            return eof_data[1]
        else:
            return None
