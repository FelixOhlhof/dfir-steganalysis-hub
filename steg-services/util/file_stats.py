
class FileStats:
    def compare_bytes_size(bytes_a, bytes_b):
        """
        Compares the size of two byte variables and returns the difference in size in percent and bytes.

        Args:
            bytes_a (bytes): First bytes variable.
            bytes_b (bytes): Second bytes variable.

        Returns:
            dict: A dictionary with the sizes of `bytes_a` and `bytes_b`, the difference in size in bytes and the percentage difference relative to `bytes_a`.
        """
        size_a = len(bytes_a)
        size_b = len(bytes_b)

        difference_in_bytes = abs(size_a - size_b)

        if size_a == 0:
            percentage_difference = float('inf') if size_b > 0 else 0
        else:
            percentage_difference = (difference_in_bytes / size_a) * 100

        return {
            "size_a": size_a,
            "size_b": size_b,
            "diff_in_bytes": difference_in_bytes,
            "diff_in_percentage": round(percentage_difference, 2)
        }
