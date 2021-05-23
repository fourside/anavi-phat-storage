import sys
from parse_input import parseInput
from cli import parseArgs

def main() -> None:
	try:
		collectionName = parseArgs(sys.argv[1:])
		if collectionName is None:
			raise Exception('not passed collection with `--collection`')
		jsonData = parseInput(sys.stdin)
		print(jsonData, collectionName)
	except Exception as error:
		print("Error:", error)
		sys.exit(-1)

if __name__ == "__main__":
	main()
