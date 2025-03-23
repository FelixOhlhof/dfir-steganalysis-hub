import io
import numpy as np
from typing import Optional
from .extractor import Extractor
from PIL import Image


class LsbExtractor(Extractor):
    def extract(self, input_file: bytes) -> Optional[bytes]:
        return self.lsb_extract(input_file).tobytes()
        
    def lsb_extract(self, input_file: bytes, bits=1, channels='RGB', endian='little', direction='row') -> np.ndarray:
        """Extract a message from an image using the LSBs method and returns it.

        :param input_file: the input image
        :param bits: the number of bits to extract
        :param channels: the channels to extract the message from. Options: 'R', 'G', 'B', 'A' or a combination of them. **Default**: 'RGB'
        :param endian: the endianness of the message. Options: 'little' or 'big'. **Default**: 'little'
        :param direction: the direction in which to traverse the image. Options: 'row' or 'col'. **Default**: 'row'
        :return: the extracted message
        """

        def _extract_bits_opt_little(data):
            div = 8 // bits
            message = np.zeros(len(data) // div, dtype=np.uint8)
            mask = (1 << bits) - 1
            for i in range(div):
                shift = bits * i
                message |= (data[i::div] & mask) << shift
            return message

        def _extract_bits_opt_big(data):
            div = 8 // bits
            message = np.zeros(len(data) // div, dtype=np.uint8)
            mask = (1 << bits) - 1
            for i in range(div):
                shift = 8 - bits - (bits * i)
                message |= (data[i::div] & mask) << shift
            return message

        def _extract_bits_little(data):
            msg_byte = 0
            shift = 0
            message = []
            mask = (1 << bits) - 1
            for byte in data:
                msg_byte |= (byte & mask) << shift
                shift += bits
                if shift >= 8:
                    tmp = msg_byte >> 8
                    message.append(msg_byte & 0xFF)
                    msg_byte = tmp
                    shift -= 8
            return np.array(message, dtype=np.uint8)

        def _extract_bits_big(data):
            msg_byte = 0
            shift = 8 - bits
            message = []
            mask = (1 << bits) - 1
            for byte in data:
                msg_byte |= (byte & mask) << shift
                shift += bits
                if shift <= 0:
                    tmp = msg_byte >> 8
                    message.append(msg_byte & 0xFF)
                    msg_byte = tmp
                    shift += 8
            return np.array(message, dtype=np.uint8)

        _COL_MAP = {'R': 0, 'G': 1, 'B': 2, 'A': 3}

        def _load_image(input_file: bytes, convert_mode='RGB', channels=None, direction='row'):
            if 'A' in channels:
                convert_mode = 'RGBA'

            with Image.open(io.BytesIO(input_file)) as img:
                arr = np.array(img.convert(convert_mode))

            if direction == 'col' or direction == 'column':
                arr = arr.transpose(1, 0, 2)

            channels = [*channels] if channels else None
            if (convert_mode == 'RGB' and 0 < len(channels) < 3) or (convert_mode == 'RGBA' and 0 < len(channels) < 4):
                arr = arr[:, :, [_COL_MAP[c] for c in channels]]
            return arr.reshape(-1)

        def _extract_message(input_file: bytes, convert_mode='RGB'):
            data = _load_image(input_file, convert_mode, channels, direction)
            if bits == 1 or bits.bit_count() == 1:
                if endian == 'big':
                    return _extract_bits_opt_big(data)
                else:
                    return _extract_bits_opt_little(data)
            else:
                if endian == 'big':
                    return _extract_bits_big(data)
                else:
                    return _extract_bits_little(data)

        return _extract_message(input_file)
