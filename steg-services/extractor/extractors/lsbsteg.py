import cv2
import numpy as np
from LSBSteganography import LSBSteg
from typing import Optional
from .extractor import Extractor


class LSBSteganographyExtractor(Extractor):
    """Extractor for data using LSB-Steganography."""

    def extract(self, input_file: bytes) -> Optional[bytes]:
        nparr = np.frombuffer(input_file, np.uint8)
        im = cv2.imdecode(nparr, cv2.IMREAD_COLOR)
        steg = LSBSteg.LSBSteg(im)
        return steg.decode_binary()
        
