import argparse
from typing import List
import sys

def parseArgs(args: List[str]) -> str:
	parser = argparse.ArgumentParser()
	parser.add_argument('--collection')
	parsed = parser.parse_args(args)
	return parsed.collection
