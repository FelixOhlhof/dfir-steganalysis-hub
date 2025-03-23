from abc import ABC, abstractmethod
from typing import Optional


class Extractor(ABC):
    """An abstract extractor for a specific attribute of image files."""

    @abstractmethod
    def extract(self, input_file: bytes) -> Optional[bytes]:
        """Tries to extract a specific attribute in the input file.
        :param input_file: Path or file-like object of the input image file
        :param format: The format of the image
        """
        pass
