import unittest
from typing import TextIO
from io import TextIOWrapper, BytesIO
from json import JSONDecodeError

from parse_input import parseInput

class TestParseInput(unittest.TestCase):
	def test_parse_input_should_add_date_field(self):
		# arrange
		mockStdin = mockTextIo(b'{"hoge": 123}')

		# act
		actual = parseInput(mockStdin)

		# assert
		self.assertEquals(123, actual["hoge"])
		self.assertNotEqual(None, actual["date"]) # should be added date field

	def test_parse_input_should_raise_exception(self):
		# assert
		with self.assertRaises(JSONDecodeError):
			# arrange
			mockStdin = mockTextIo(b'"hoge": 123')

			# act
			parseInput(mockStdin)

def mockTextIo(input: bytes) -> TextIO:
	return TextIOWrapper(BytesIO(input))

if __name__ == "__main__":
	unittest.main()
